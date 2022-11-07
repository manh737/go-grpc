package services

import (
	"context"
	"testing"

	"github.com/manh737/go-grpc/protos"
	"github.com/manh737/go-grpc/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLaptopServer(t *testing.T) {
	t.Parallel()
	laptopWithoutId := sample.NewLaptop()
	laptopWithoutId.Id = ""
	laptopInvalidId := sample.NewLaptop()
	laptopInvalidId.Id = "invalid-uuid"
	laptopDuplicateId := sample.NewLaptop()
	storeDublicateId := NewInMemoryLaptopStore()
	err := storeDublicateId.Save(laptopDuplicateId)
	require.NoError(t, err)
	testCases := []struct {
		name   string
		laptop *protos.Laptop
		store  LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		}, {
			name:   "success_without_id",
			laptop: laptopWithoutId,
			store:  NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			name:   "laptop_invalid_id",
			laptop: laptopInvalidId,
			store:  NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		}, {
			name:   "laptop_already_exists",
			laptop: laptopDuplicateId,
			store:  storeDublicateId,
			code:   codes.AlreadyExists,
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			server := NewLaptopServer(testCase.store)
			req := &protos.CreateLaptopRequest{
				Laptop: testCase.laptop,
			}
			res, err := server.CreateLaptop(context.Background(), req)
			if testCase.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotNil(t, res.Laptop)
				require.NotEmpty(t, res.Laptop.Id)
				if len(testCase.laptop.Id) > 0 {
					require.Equal(t, testCase.laptop.Id, res.Laptop.Id)
				}
				requireEqualLaptop(t, testCase.laptop, res.Laptop)
			} else {
				require.Error(t, err)
				grpcError, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, testCase.code, grpcError.Code())
			}
		})
	}
}
func requireEqualLaptop(t *testing.T, expected, actual *protos.Laptop) {
	require.NotNil(t, expected)
	require.NotNil(t, actual)
	require.Equal(t, expected.Id, actual.Id)
	require.Equal(t, expected.Brand, actual.Brand)
	require.Equal(t, expected.Name, actual.Name)
	require.Equal(t, expected.Cpu.NumberCores, actual.Cpu.NumberCores)
	require.Equal(t, expected.Cpu.NumberThreads, actual.Cpu.NumberThreads)
	require.Equal(t, expected.Cpu.MinGhz, actual.Cpu.MinGhz)
	require.Equal(t, expected.Cpu.MaxGhz, actual.Cpu.MaxGhz)
	require.Equal(t, expected.Cpu.Brand, actual.Cpu.Brand)
	require.Equal(t, expected.Memory, actual.Memory)
	require.Equal(t, expected.Gpus[0], actual.Gpus[0])
	require.Equal(t, expected.Storages[0], actual.Storages[0])
	require.Equal(t, expected.PriceUsd, actual.PriceUsd)
}
