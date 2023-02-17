package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"pcbook/pb"
	"pcbook/service"
	"time"
)

const (
	secretKey     = "12345"
	tokenDuration = time.Minute * 15
)

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	log.Print("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {

	log.Print("--> stream interceptor: ", info.FullMethod)
	return handler(srv, ss)
}

func createUser(userStore service.UserStore, username string, password string, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func seedUsers(userStore service.UserStore) error {
	err := createUser(userStore, "admin", "admin", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user", "user", "user")
}

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/techschool_pcbook.LaptopService/"

	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
		laptopServicePath + "UploadImage":  {"admin"},
		laptopServicePath + "RateLaptop":   {"admin", "user"},
	}
}

func main() {
	port := flag.Int("port", 8080, "the server port")
	flag.Parse()

	serverAddr := fmt.Sprintf("0.0.0.0:%d", *port)
	fmt.Printf("start server on %s\n", serverAddr)

	// jwt manager

	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authInterceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
		grpc.StreamInterceptor(authInterceptor.Stream()),
	)

	// auth service

	userStore := service.NewInMemoryUserStore()
	_ = seedUsers(userStore)

	authServer := service.NewAuthServer(userStore, jwtManager)

	pb.RegisterAuthServiceServer(grpcServer, authServer)

	// laptop service

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()

	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	// ...

	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
