package main

const version = "v0.0.3"

// Args flag 参数
type Args struct {
	ShowVersion      bool   // 显示版本
	DisableOrComment bool   // 不使用注释作为mapping
	CustomTemplate   string // 自定义模板
	Suffix           string // 自定义文件名后缀, 默认 .mapping.pb.go
	Merge            bool   // 合并到一个文件
	Filename         string // 合并文件名
	Package          string // 合并包名
	GoPackage        string // 合并go包名
}
