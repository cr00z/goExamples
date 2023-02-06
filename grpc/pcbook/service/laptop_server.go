package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"pcbook/pb"
)

type LaptopServer struct {
	laptopStore LaptopStore
	imageStore  ImageStore
	pb.UnimplementedLaptopServiceServer
}

func NewLaptopServer(laptopStore LaptopStore, imageStore ImageStore) *LaptopServer {
	return &LaptopServer{
		laptopStore: laptopStore,
		imageStore:  imageStore,
	}
}

func (s *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {

	//time.Sleep(time.Second * 5)
	if ctx.Err() == context.Canceled {
		fmt.Println("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	laptop := req.GetLaptop()
	log.Printf("receive create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}

	err := s.laptopStore.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrorAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop to the store: %v", err)
	}

	log.Printf("saved laptop to the store with id: %s", laptop.Id)
	return &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}, nil
}

func (s *LaptopServer) SearchLaptop(
	req *pb.SearchLaptopRequest,
	stream pb.LaptopService_SearchLaptopServer,
) error {

	filter := req.GetFilter()
	log.Printf("receive a search-laptop request with filter: %v", filter)

	err := s.laptopStore.Search(
		stream.Context(),
		filter,
		func(laptop *pb.Laptop) error {
			res := &pb.SearchLaptopResponse{
				Laptop: laptop,
			}
			err := stream.Send(res)
			if err != nil {
				return err
			}
			log.Printf("send laptop with id: %s", laptop.GetId())
			return nil
		},
	)
	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}

	return nil
}

func (s *LaptopServer) UploadImage(
	stream pb.LaptopService_UploadImageServer,
) error {
	req, err := stream.Recv()
	if err
}
