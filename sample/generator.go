package sample

import (
	"github.com/manh737/go-grpc/protos"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewKeyboards returns a sample keyboard
func NewKeyboard() *protos.Keyboard {
	keyboard := &protos.Keyboard{
		Layout:   randomKeyboardLayout(),
		Backlist: randomBoolean(),
	}
	return keyboard
}

// NewMemory returns a sample memory
func NewMemory() *protos.Memory {
	memory := &protos.Memory{
		Unit:  randomMemoryUnit(),
		Value: uint32(randomInt(2, 8)),
	}
	return memory
}

// NewCPU returns a sample cpu
func NewCPU() *protos.CPU {
	number_cores := randomInt(2, 8)
	min_ghz := randomFloat64(2.0, 3.5)
	cpu := &protos.CPU{
		Brand:         randomCPUBrand(),
		Name:          randomCPUName(),
		NumberCores:   uint32(number_cores),
		NumberThreads: uint32(number_cores * 2),
		MinGhz:        min_ghz,
		MaxGhz:        randomFloat64(min_ghz, 5.0),
	}
	return cpu
}

// NewGPU returns a sample gpu
func NewGPU() *protos.GPU {
	number_cores := randomInt(2, 8)
	min_ghz := randomFloat64(2.0, 3.5)
	gpu := &protos.GPU{
		Brand:         randomGPUBrand(),
		Name:          randomGPUName(),
		NumberCores:   uint32(number_cores),
		NumberThreads: uint32(number_cores * 2),
		MinGhz:        min_ghz,
		MaxGhz:        randomFloat64(min_ghz, 5.0),
		Memory:        NewMemory(),
	}
	return gpu
}

// NewScreen returns a sample screen
func NewScreen() *protos.Screen {
	screen := &protos.Screen{
		Resolution: &protos.Screen_Resolution{
			Width:  uint32(randomInt(1920, 3840)),
			Height: uint32(randomInt(1080, 2160)),
		},
		Panel:      randomScreenPanel(),
		SizeInch:   float32(randomFloat64(16.0, 34.0)),
		Multitouch: randomBoolean(),
	}
	return screen
}

// NewStorage returns a sample storage
func NewStorage() *protos.Storage {
	storage := &protos.Storage{
		Driver: randomStorageDriver(),
		Memory: NewMemory(),
	}
	return storage
}

// NewLaptop returns a sample laptop
func NewLaptop() *protos.Laptop {
	laptop := &protos.Laptop{
		Id:       randomID(),
		Brand:    randomLaptopBrand(),
		Name:     randomLaptopName(),
		Cpu:      NewCPU(),
		Gpus:     []*protos.GPU{NewGPU()},
		Storages: []*protos.Storage{NewStorage()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Memory:   NewMemory(),
		Weight: &protos.Laptop_KgWeight{
			KgWeight: randomFloat64(1.5, 3.0),
		},
		PriceUsd:    randomFloat64(1500.0, 2000.0),
		ReleaseYear: uint32(randomInt(2015, 2022)),
		UpdatedAt:   timestamppb.Now(),
	}
	return laptop
}
