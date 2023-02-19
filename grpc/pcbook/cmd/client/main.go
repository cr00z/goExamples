package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
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

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemServerCA, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

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
	address := flag.String("address", "0.0.0.0:8080", "server address")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()

	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	authConn, err := grpc.Dial(
		*address,
		transportOption,
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
		transportOption,
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
