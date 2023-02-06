package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"pcbook/pb"
	"pcbook/service"
)

func main() {
	port := flag.Int("port", 8080, "the server port")
	flag.Parse()

	serverAddr := fmt.Sprintf("0.0.0.0:%d", *port)
	fmt.Printf("start server on %s\n", serverAddr)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")

	laptopServer := service.NewLaptopServer(laptopStore, imageStore)
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
