package services

import (
	"errors"
	"sync"

	"github.com/manh737/go-grpc/protos"

	"github.com/jinzhu/copier"
)

var ErrAllreadyExists = errors.New("laptop with given id already exists")

type LaptopStore interface {
	Save(laptop *protos.Laptop) error
	Find(laptopId string) (*protos.Laptop, error)
	Size() int
}

type InMemoryLaptopStore struct {
	data  map[string]*protos.Laptop
	mutex sync.RWMutex
}

// NewInMemoryLaptopStore creates a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*protos.Laptop),
	}
}

func (s *InMemoryLaptopStore) Save(laptop *protos.Laptop) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.data[laptop.Id]; ok {
		return ErrAllreadyExists
	}

	//deep copy
	other := &protos.Laptop{}
	err := copier.Copy(other, laptop)

	if err != nil {
		return errors.New("cannot copy laptop")
	}
	s.data[other.Id] = other
	return nil
}

// CreateLaptop is a unary RPC to create a new laptop
func (s *InMemoryLaptopStore) Find(laptopId string) (*protos.Laptop, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	laptop, ok := s.data[laptopId]
	if !ok {
		return nil, nil
	}

	// deep copy
	result := &protos.Laptop{}
	err := copier.Copy(result, laptop)
	if err != nil {
		return nil, errors.New("cannot copy laptop")
	}
	return result, nil
}

// CreateLaptop is a unary RPC to create a new laptop
func (s *InMemoryLaptopStore) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.data)
}
