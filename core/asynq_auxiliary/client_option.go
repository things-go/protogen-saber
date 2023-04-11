package asynq_auxiliary

import "google.golang.org/protobuf/proto"

type ClientSettings struct {
	MarshalBinary func(any) ([]byte, error)
}

type ClientOption func(*ClientSettings)

func NewClientSettings() *ClientSettings {
	return &ClientSettings{
		MarshalBinary: func(v any) ([]byte, error) {
			return proto.Marshal(v.(proto.Message))
		},
	}
}

func WithClientMarshalBinary(f func(any) ([]byte, error)) ClientOption {
	return func(cs *ClientSettings) {
		if f != nil {
			cs.MarshalBinary = f
		}
	}
}
