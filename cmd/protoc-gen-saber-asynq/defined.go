package main

const version = "v0.2.0"

// Args flag参数
type Args struct {
	ShowVersion bool // 显示版本
}

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/v1/helloworld.proto
	Methods     []*methodDesc
}

type methodDesc struct {
	Name    string // 方法名
	Num     int    // 方法号, not used
	Request string // 请求结构
	Reply   string // 回复结构
	Comment string // 方法注释
	// asynq rule
	Pattern  string // 匹配器
	CronSpec string // cron specification
}
