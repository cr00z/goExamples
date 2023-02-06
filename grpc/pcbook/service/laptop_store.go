package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"log"
	"pcbook/pb"
	"sync"
)

var ErrorAlreadyExists = errors.New("record already exists")

type LaptopStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (*pb.Laptop, error)
	Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error
}

type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (s *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.data[laptop.Id] != nil {
		return ErrorAlreadyExists
	}

	laptopCopy, err := DeepCopyLaptop(laptop)
	if err != nil {
		return err
	}
	s.data[laptopCopy.Id] = laptopCopy

	return nil
}

func (s *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	laptop := s.data[id]
	if laptop == nil {
		return nil, nil
	}

	return DeepCopyLaptop(laptop)
}

func (s *InMemoryLaptopStore) Search(
	ctx context.Context,
	filter *pb.Filter,
	found func(laptop *pb.Laptop) error,
) error {
	for _, laptop := range s.data {
		//time.Sleep(time.Second)

		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Print("context is cancelled")
			return errors.New("context is cancelled")
		}

		if isQualified(filter, laptop) {
			other, err := DeepCopyLaptop(laptop)
			if err != nil {
				return err
			}

			err = found(other)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// service

func DeepCopyLaptop(laptop *pb.Laptop) (*pb.Laptop, error) {
	laptopCopy := &pb.Laptop{}
	err := copier.Copy(laptopCopy, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %v", err)
	}
	return laptopCopy, err
}

func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool {
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd() {
		return false
	}
	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores() {
		return false
	}
	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz() {
		return false
	}
	if toBits(laptop.GetRam()) < toBits(laptop.GetRam()) {
		return false
	}
	return true
}

func toBits(memory *pb.Memory) uint64 {
	value := memory.GetValue()
	switch memory.GetUnit() {
	case pb.Memory_BYTE:
		value <<= 3
	case pb.Memory_KILOBYTE:
		value <<= 13
	case pb.Memory_MEGABYTE:
		value <<= 23
	case pb.Memory_GIGABYTE:
		value <<= 33
	case pb.Memory_TERABYTE:
		value <<= 43
	}
	return value
}
