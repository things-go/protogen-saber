package asynq_auxiliary

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
)

// ClientSettings 客户端配置
type ClientSettings struct {
	MarshalBinary func(any) ([]byte, error)
}

// ClientOption 客户端选项
type ClientOption func(*ClientSettings)

// NewClientSettings 新建配置, 默认使用 proto.Marshal
func NewClientSettings() *ClientSettings {
	return &ClientSettings{
		MarshalBinary: func(v any) ([]byte, error) {
			return proto.Marshal(v.(proto.Message))
		},
	}
}

// WithClientMarshalBinary 使用指定序列化
func WithClientMarshalBinary(f func(any) ([]byte, error)) ClientOption {
	return func(cs *ClientSettings) {
		if f != nil {
			cs.MarshalBinary = f
		}
	}
}

// WithClientJsonMarshalBinary 使用 json.Marshal
func WithClientJsonMarshalBinary() ClientOption {
	return func(cs *ClientSettings) {
		cs.MarshalBinary = json.Marshal
	}
}
