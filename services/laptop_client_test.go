package services

import (
	"context"
	"net"
	"testing"

	"github.com/manh737/go-grpc/protos"
	"github.com/manh737/go-grpc/sample"
	"github.com/manh737/go-grpc/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestLaptopClient(t *testing.T) {
	t.Parallel()
	laptopServer, serverAddress := startTestLaptopServer(t)
	laptopClient := newTestLaptopClient(serverAddress)
	laptop := sample.NewLaptop()
	expectedId := laptop.Id
	req := &protos.CreateLaptopRequest{
		Laptop: laptop,
	}
	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedId, res.Laptop.Id)
	otherLaptop := sample.NewLaptop()
	req.Laptop = otherLaptop
	res, err = laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, otherLaptop.Id, res.Laptop.Id)
	// check if the laptop is saved to the store
	otherLaptop, err = laptopServer.Store.Find(otherLaptop.Id)
	require.NoError(t, err)
	require.NotNil(t, otherLaptop)
	compareLaptop(t, otherLaptop, otherLaptop)
}

func startTestLaptopServer(t *testing.T) (*LaptopServer, string) {
	laptopServer := NewLaptopServer(NewInMemoryLaptopStore())
	grpcServer := grpc.NewServer()
	protos.RegisterLaptopServiceServer(grpcServer, laptopServer)
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	go grpcServer.Serve(listener)
	return laptopServer, listener.Addr().String()
}
func newTestLaptopClient(serverAddress string) protos.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return protos.NewLaptopServiceClient(conn)
}

func compareLaptop(t *testing.T, laptop1, laptop2 *protos.Laptop) {
	laptop1JSON, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)
	laptop2JSON, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)
	require.Equal(t, laptop1JSON, laptop2JSON)
}
