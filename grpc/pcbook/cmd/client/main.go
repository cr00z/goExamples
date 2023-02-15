package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"
	"path"
	"pcbook/pb"
	"pcbook/sample"
	"time"
)

func createLaptop(client pb.LaptopServiceClient, laptop *pb.Laptop) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// laptop.Id = ""
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	resp, err := client.CreateLaptop(ctx, req)
	if err != nil {
		state, ok := status.FromError(err)
		if ok && state.Code() == codes.AlreadyExists {
			log.Printf("laptop already exists")
		} else {
			log.Fatalf("cannot create laptop request: %v", err)
		}
		return
	}

	log.Printf("laptop created with id: %s", resp.Id)
}

func searchLaptop(client pb.LaptopServiceClient, filter *pb.Filter) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.SearchLaptopRequest{
		Filter: filter,
	}

	stream, err := client.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatalf("cannot create search request: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatalf("cannot receive response: %v", err)
		}
		laptop := res.GetLaptop()
		fmt.Println("found: ", laptop.GetId())
		fmt.Printf("\t%s %s\n", laptop.GetBrand(), laptop.GetName())
		fmt.Printf("\tcpu %d cores, min %.2f gHz, max %.2f ghz\n",
			laptop.GetCpu().GetNumberCores(),
			laptop.GetCpu().GetMinGhz(),
			laptop.GetCpu().GetMaxGhz(),
		)
		fmt.Printf("\tram %d %s\n",
			laptop.GetRam().GetValue(),
			laptop.GetRam().GetUnit(),
		)
	}
}

func testCreateLaptop(laptopClient pb.LaptopServiceClient) {
	createLaptop(laptopClient, sample.NewLaptop())
}

func testSearchLaptop(laptopClient pb.LaptopServiceClient) {
	for i := 0; i < 10; i++ {
		createLaptop(laptopClient, sample.NewLaptop())
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
	searchLaptop(laptopClient, filter)
}

func testUploadImage(laptopClient pb.LaptopServiceClient) {
	laptop := sample.NewLaptop()
	createLaptop(laptopClient, laptop)
	uploadImage(laptopClient, laptop.GetId(), "tmp/laptop.jpg")
}

func uploadImage(laptopClient pb.LaptopServiceClient, laptopID string, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := laptopClient.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId:  laptopID,
				ImageType: path.Ext(imagePath),
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info: ", err)
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunks to buffer: ", err)
		}

		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			err2 := stream.RecvMsg(nil)
			log.Fatal("cannot send chunk to server: ", err, err2)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("image uploaded with id %s, size %d bytes", res.GetId(), res.GetSize())
}

func main() {
	address := flag.String("addr", "127.0.0.1:8080", "server address")
	flag.Parse()

	conn, err := grpc.Dial(
		*address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("cannot connect to server: %v", err)
	}

	client := pb.NewLaptopServiceClient(conn)

	// testCreateLaptop(client)
	// testSearchLaptop(client)
	testUploadImage(client)
}
