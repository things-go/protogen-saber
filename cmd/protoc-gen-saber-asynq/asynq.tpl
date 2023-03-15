

{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

{{- range .MethodSets}}
const Pattern_{{$svrType}}_{{.Name}} = "{{.Pattern}}"
{{- end}}

type {{.ServiceType}}TaskHandler interface {
{{- range .MethodSets}}
	{{.Comment}}
	{{.Name}}(context.Context, *{{.Request}}) (error)
{{- end}}
	// UnmarshalBinary parses the binary data and stores the result
	// in the value pointed to by v.
	UnmarshalBinary([]byte, any) error
}

type Unimplemented{{.ServiceType}}TaskHandlerImpl struct {}

func (*Unimplemented{{.ServiceType}}TaskHandlerImpl) UnmarshalBinary(b []byte, v any) error {
	return proto.Unmarshal(b, v.(proto.Message))
}

func Register{{.ServiceType}}TaskHandler(mux *asynq.ServeMux, srv {{.ServiceType}}TaskHandler) {
	{{- range .Methods}}
	mux.HandleFunc(Pattern_{{$svrType}}_{{.Name}}, _{{$svrType}}_{{.Name}}_Task_Handler(srv))
	{{- end}}
}
{{range .Methods}}
func _{{$svrType}}_{{.Name}}_Task_Handler(srv {{$svrType}}TaskHandler) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, task *asynq.Task) error {
		var in {{.Request}}
		if err := srv.UnmarshalBinary(task.Payload(), &in); err != nil {
			return err
		}
		return srv.{{.Name}}(ctx, &in)
	}
}
{{end}}

type {{.ServiceType}}TaskClient interface {
	// SetMarshaler set marshal the binary encoding of v function.
	SetMarshaler(func(any) ([]byte, error)) {{.ServiceType}}TaskClient
{{- range .MethodSets}}
    {{.Comment}}
	{{.Name}}(ctx context.Context, req *{{.Request}}, opts ...asynq.Option) (info *asynq.TaskInfo, err error) 
{{- end}}
}

type {{.ServiceType}}TaskClientImpl struct {
	cc *asynq.Client
	marshaler func(any) ([]byte, error)
}

// New{{.ServiceType}}TaskClient new client. use default proto.Marhsal.
func New{{.ServiceType}}TaskClient (client *asynq.Client) {{.ServiceType}}TaskClient {
	return &{{.ServiceType}}TaskClientImpl{
		cc: client,
		marshaler: func(v any) ([]byte, error) {
			return proto.Marshal(v.(proto.Message))
		},
	}
}

// SetMarshaler set marshal the binary encoding of v function.
func (c *{{$svrType}}TaskClientImpl) SetMarshaler(marshaler func(any) ([]byte, error)) {{.ServiceType}}TaskClient {
	if marshaler != nil {
		c.marshaler = marshaler
	}
	return c
}

{{range .MethodSets}}
{{.Comment}}
func (c *{{$svrType}}TaskClientImpl) {{.Name}}(ctx context.Context, in *{{.Request}}, opts ...asynq.Option) (*asynq.TaskInfo, error) {
	payload, err := c.marshaler(in)
	if err != nil {
		return nil, err
	}
	task := asynq.NewTask(Pattern_{{$svrType}}_{{.Name}}, payload, opts...)
	taskInfo, err := c.cc.Enqueue(task)
	if err != nil {
		return nil, err
	}
	return taskInfo, nil
}
{{end}}