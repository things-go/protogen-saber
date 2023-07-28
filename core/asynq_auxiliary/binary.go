package asynq_auxiliary

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
)


var BinaryProtobuf = Protobuf{}
var BinaryJSON = JSON{}

type BinaryMarshaler interface {
	MarshalBinary(v any) (data []byte, err error)
}

type BinaryUnmarshaler interface {
	UnmarshalBinary(data []byte, v any) error
}

type Protobuf struct{}

func (Protobuf) MarshalBinary(v any) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (Protobuf) UnmarshalBinary(b []byte, v any) error {
	return proto.Unmarshal(b, v.(proto.Message))
}

type JSON struct{}

func (JSON) MarshalBinary(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (JSON) UnmarshalBinary(b []byte, v any) error {
	return json.Unmarshal(b, v)
}
