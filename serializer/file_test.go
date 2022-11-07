package serializer

import (
	"testing"

	"github.com/manh737/go-grpc/protos"
	"github.com/manh737/go-grpc/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/test_laptop.bin"

	laptop1 := sample.NewLaptop()

	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &protos.Laptop{}

	err = ReadProtobufFromBinaryFile(laptop2, binaryFile)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))

	laptop3 := sample.NewLaptop()

	jsonFile := "../tmp/test_laptop.json"

	err = WriteProtobufToJSONFile(laptop3, jsonFile)
	require.NoError(t, err)

	laptop4 := &protos.Laptop{}

	err = ReadProtobufFromJSONFile(laptop4, jsonFile)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop3, laptop4))

}
