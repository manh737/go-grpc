package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/manh737/go-grpc/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	Store LaptopStore
	protos.UnimplementedLaptopServiceServer
}

// NewLaptopServer returns a new LaptopServer
func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

// CreateLaptop is a unary RPC to create a new laptop
func (s *LaptopServer) CreateLaptop(ctx context.Context, req *protos.CreateLaptopRequest) (*protos.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("laptop id is not a valid UUID error: %s", err))
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot generate uuid error: %s", err))
		}
		laptop.Id = id.String()
	}
	switch ctx.Err() {
	case context.Canceled:
		log.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		log.Print("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	// save the laptop to the in-memory store
	err := s.Store.Save(laptop)
	if err != nil {
		if errors.Is(err, ErrAllreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, fmt.Sprintf("laptop with id %s already exists error: %s", laptop.Id, err))
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot save laptop error: %s", err))
	}
	res := &protos.CreateLaptopResponse{
		Laptop: laptop,
	}
	log.Default().Printf("saved laptop with id: %s", laptop.Id)
	log.Default().Printf("total laptop %d", s.Store.Size())
	return res, nil
}
