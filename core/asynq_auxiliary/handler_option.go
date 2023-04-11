package asynq_auxiliary

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
)

type HandlerSettings struct {
	// UnmarshalBinary parses the binary data and stores the result
	// in the value pointed to by v.
	UnmarshalBinary func([]byte, any) error
}

type HandlerOption func(*HandlerSettings)

func NewHandlerSettings() *HandlerSettings {
	return &HandlerSettings{
		UnmarshalBinary: func(b []byte, v any) error {
			return proto.Unmarshal(b, v.(proto.Message))
		},
	}
}

func WithHandlerUnmarshalBinary(f func([]byte, any) error) HandlerOption {
	return func(cs *HandlerSettings) {
		if f != nil {
			cs.UnmarshalBinary = f
		}
	}
}

func WithHandlerJsonUnmarshalBinary() HandlerOption {
	return func(cs *HandlerSettings) {
		cs.UnmarshalBinary = json.Unmarshal
	}
}
