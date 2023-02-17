package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"pcbook/client"
	"pcbook/pb"
	"pcbook/sample"
	"strings"
	"time"
)

const (
	tokenRefreshDuration = 30 * time.Second
	username             = "admin"
	password             = "admin"
)

func authMethods() map[string]bool {
	const laptopServicePath = "/techschool_pcbook.LaptopService/"

	return map[string]bool{
		laptopServicePath + "CreateLaptop": true,
		laptopServicePath + "UploadImage":  true,
		laptopServicePath + "RateLaptop":   true,
	}
}

func testCreateLaptop(laptopClient *client.LaptopClient) {
	laptopClient.CreateLaptop(sample.NewLaptop())
}

func testSearchLaptop(laptopClient *client.LaptopClient) {
	for i := 0; i < 10; i++ {
		laptopClient.CreateLaptop(sample.NewLaptop())
	}

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   1.5,
		MinRam: &pb.Memory{
			Value: 8,
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	laptopClient.SearchLaptop(filter)
}

func testUploadImage(laptopClient *client.LaptopClient) {
	laptop := sample.NewLaptop()
	laptopClient.CreateLaptop(laptop)
	laptopClient.UploadImage(laptop.GetId(), "tmp/laptop.jpg")
}

func testRateLaptop(laptopClient *client.LaptopClient) {
	nLaptops := 3
	laptopIDs := make([]string, nLaptops)
	for i := 0; i < nLaptops; i++ {
		laptop := sample.NewLaptop()
		laptopIDs[i] = laptop.GetId()
		laptopClient.CreateLaptop(laptop)
	}

	scores := make([]float64, nLaptops)
	for {
		fmt.Print("rate laptops (y/n)? ")
		var answer string
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "y" {
			break
		}

		for i := 0; i < nLaptops; i++ {
			scores[i] = sample.RandomLaptopScore()
		}

		err := laptopClient.RateLaptop(laptopIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	address := flag.String("a", "127.0.0.1:8080", "server address")
	flag.Parse()

	authConn, err := grpc.Dial(
		*address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cannot connect to auth server: %v", err)
	}

	authClient := client.NewAuthClient(authConn, username, password)
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), tokenRefreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}

	laptopConn, err := grpc.Dial(
		*address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatalf("cannot connect to laptop server: %v", err)
	}

	laptopClient := client.NewLaptopClient(laptopConn)

	// testCreateLaptop(laptopClient)
	// testSearchLaptop(client)
	// testUploadImage(client)
	testRateLaptop(laptopClient)
}
