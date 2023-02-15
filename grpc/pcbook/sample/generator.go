package sample

import (
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math"
	"math/rand"
	"pcbook/pb"
)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	cores := uint32(randomFromSet(1, 2, 4, 8))
	threads := cores * uint32(randomFromSet(1, 2, 4))
	minGhz := randomFloat(1.0, 2.0)
	maxGhz := randomFloat(minGhz, 5.0)

	return &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   cores,
		NumberThreads: threads,
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat(1.0, 1.5)
	maxGhz := randomFloat(minGhz, 2.0)
	memory := NewRAMinGigabytes(1, 2, 4, 8, 16)

	return &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
}

func NewRam() *pb.Memory {
	return NewRAMinGigabytes(4, 8, 16, 32, 64)
}

func NewSSD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: NewRAMinGigabytes(128, 256, 512, 1024),
	}
}

func NewHDD() *pb.Storage {
	return &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomFromSet(1, 2, 4, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
}

func NewScreen() *pb.Screen {
	size := math.Round(randomFloat(13, 17)*10) / 10

	return &pb.Screen{
		SizeInch:   float32(size),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	price := math.Round(randomFloat(1500, 3000)*100) / 100
	year := randomFloat(2015, 2023)
	return &pb.Laptop{
		Id:       uuid.New().String(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRam(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat(1.0, 3.0),
		},
		PriceUsd:    price,
		ReleaseYear: uint32(year),
		UpdatedAt:   timestamppb.Now(),
	}
}

func RandomLaptopScore() float64 {
	return float64(1 + rand.Intn(10-1))
}
