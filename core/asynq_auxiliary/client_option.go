package asynq_auxiliary

// ClientSettings 客户端配置
type ClientSettings struct {
	Marshaler BinaryMarshaler
}

// ClientOption 客户端选项
type ClientOption func(*ClientSettings)

// NewClientSettings 新建配置, 默认使用 json.Marshal
func NewClientSettings(opts ...ClientOption) *ClientSettings {
	cs := &ClientSettings{
		Marshaler: BinaryProtobuf,
	}
	for _, opt := range opts {
		opt(cs)
	}
	return cs
}

// WithClientMarshaler 使用指定序列化
func WithClientMarshaler(marshaler BinaryMarshaler) ClientOption {
	return func(cs *ClientSettings) {
		if marshaler != nil {
			cs.Marshaler = marshaler
		}
	}
}

// WithClientMarshalerJSON 使用 json.Marshal Marshaler
func WithClientMarshalerJSON() ClientOption {
	return func(cs *ClientSettings) {
		cs.Marshaler = BinaryJSON
	}
}

// WithClientMarshalerProtobuf 使用 proto.Marshal Marshaler
func WithClientMarshalerProtobuf() ClientOption {
	return func(cs *ClientSettings) {
		cs.Marshaler = BinaryProtobuf
	}
}
