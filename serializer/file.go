package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

// WriteProtobufToBinaryFile writes a protobuf message to a binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message: %w", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write to file: %w", err)
	}
	return nil
}

// WriteProtobufToJSONFile writes a protobuf message to a json file
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to json: %w", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write to file: %w", err)
	}
	return nil
}

// ReadProtobufFromBinaryFile reads a protobuf message from a binary file
func ReadProtobufFromBinaryFile(message proto.Message, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal message: %w", err)
	}
	return nil

}

// ReadProtobufFromJSONFile reads a protobuf message from a json file
func ReadProtobufFromJSONFile(message proto.Message, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}
	err = JSONToProtobuf(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal message: %w", err)
	}
	return nil

}
