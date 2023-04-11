package asynq_auxiliary

import (
	"encoding/json"

	"google.golang.org/protobuf/proto"
)

// HandlerSettings 处理器配置
type HandlerSettings struct {
	// UnmarshalBinary parses the binary data and stores the result
	// in the value pointed to by v.
	UnmarshalBinary func([]byte, any) error
}

// HandlerOption 处理器选项
type HandlerOption func(*HandlerSettings)

// NewHandlerSettings 新建处理器选项, 默认使用 proto.Unmarshal
func NewHandlerSettings() *HandlerSettings {
	return &HandlerSettings{
		UnmarshalBinary: func(b []byte, v any) error {
			return proto.Unmarshal(b, v.(proto.Message))
		},
	}
}

// WithHandlerUnmarshalBinary 指定反序列化
func WithHandlerUnmarshalBinary(f func([]byte, any) error) HandlerOption {
	return func(cs *HandlerSettings) {
		if f != nil {
			cs.UnmarshalBinary = f
		}
	}
}

// WithHandlerUnmarshalBinary 使用 json.Unmarshal
func WithHandlerJsonUnmarshalBinary() HandlerOption {
	return func(cs *HandlerSettings) {
		cs.UnmarshalBinary = json.Unmarshal
	}
}
