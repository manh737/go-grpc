package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) ([]byte, error) {
	return protojson.Marshal(message)
}

func JSONToProtobuf(data []byte, message proto.Message) error {
	return protojson.Unmarshal(data, message)
}
