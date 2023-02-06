package service

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"pcbook/pb"
	"pcbook/sample"
	"testing"
)

func Test_LaptopServer_CreateLaptop(t *testing.T) {
	t.Parallel()

	// Arrange
	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	laptopDuplicateID := sample.NewLaptop()
	storeDuplicateID := NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID)
	require.NoError(t, err)

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "success_without_id",
			laptop: laptopNoID,
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "failure_invalid_uuid",
			laptop: laptopInvalidID,
			store:  NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			name:   "failure_duplicate_uuid",
			laptop: laptopDuplicateID,
			store:  storeDuplicateID,
			code:   codes.AlreadyExists,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Act
			server := NewLaptopServer(tc.store, nil)
			response, err := server.CreateLaptop(context.Background(), &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			})

			// Assert
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.NotEmpty(t, response)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, response.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, response)
				stat, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, stat.Code())
			}
		})
	}
}
