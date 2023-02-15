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

func createUser(userStore service.UserStore, username string, password string, role string) {
	user, _ := service.NewUser(username, password, role)
	userStore.Save(user)
}

func seedUsers(userStore service.UserStore) {
	createUser(userStore, "admin", "admin", "admin")
	createUser(userStore, "user", "user", "user")
}

func main() {
	port := flag.Int("port", 8080, "the server port")
	flag.Parse()

	serverAddr := fmt.Sprintf("0.0.0.0:%d", *port)
	fmt.Printf("start server on %s\n", serverAddr)

	// laptop service

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()

	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	// auth service

	userStore := service.NewInMemoryUserStore()
	jwtManager := service.NewJWTManager(secretKey, tokenDuration)

	authServer := service.NewAuthServer(userStore, jwtManager)

	pb.RegisterAuthServiceServer(grpcServer, authServer)

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
