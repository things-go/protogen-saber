package asynq_auxiliary

// HandlerSettings 处理器配置
type HandlerSettings struct {
	Unmarshaler BinaryUnmarshaler
}

// HandlerOption 处理器选项
type HandlerOption func(*HandlerSettings)

// NewHandlerSettings 新建处理器选项, 默认使用 proto.Unmarshal
func NewHandlerSettings(opts ...HandlerOption) *HandlerSettings {
	hs := &HandlerSettings{
		Unmarshaler: BinaryProtobuf,
	}
	for _, opt := range opts {
		opt(hs)
	}
	return hs
}

// WithHandlerUnmarshaler 指定反序列化
func WithHandlerUnmarshaler(unmarshaler BinaryUnmarshaler) HandlerOption {
	return func(hs *HandlerSettings) {
		if unmarshaler != nil {
			hs.Unmarshaler = unmarshaler
		}
	}
}

// WithClientUnmarshalerJSON 使用 json.Unmarshal Unmarshaler
func WithClientUnmarshalerJSON() HandlerOption {
	return func(hs *HandlerSettings) {
		hs.Unmarshaler = BinaryJSON
	}
}

// WithClientUnmarshalerProtobuf 使用 proto.Unmarshal Unmarshaler
func WithClientUnmarshalerProtobuf() HandlerOption {
	return func(hs *HandlerSettings) {
		hs.Unmarshaler = BinaryProtobuf
	}
}
