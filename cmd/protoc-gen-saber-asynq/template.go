package main

import (
	"embed"
	"html/template"
	"io"
)

//go:embed asynq.tpl
var Static embed.FS

var ginHttpTemplate = template.Must(template.New("components").ParseFS(Static, "asynq.tpl")).
	Lookup("asynq.tpl")

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/v1/helloworld.proto
	Methods     []*methodDesc
}

type methodDesc struct {
	// method
	Name    string // 方法名
	Num     int    // 方法号
	Request string // 请求结构
	Reply   string // 回复结构
	Comment string // 方法注释
	// asynq rule
	Pattern string // 匹配器
}

func (s *serviceDesc) execute(w io.Writer) error {
	return ginHttpTemplate.Execute(w, s)
}
