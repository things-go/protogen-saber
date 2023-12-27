package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

const version = "v0.5.0"

var args = &struct {
	ShowVersion    bool   // 显示版本
	CustomTemplate string // 自定义模板
	Suffix         string // 自定义文件名后缀, 默认 .mapping.pb.go
	Merge          bool   // 合并到一个文件
	Filename       string // 合并文件名
	Package        string // 合并包名
	GoPackage      string // 合并go包名
}{
	ShowVersion:    false,
	CustomTemplate: "",
	Suffix:         "",
	Merge:          false,
	Filename:       "",
	Package:        "",
	GoPackage:      "",
}

func init() {
	flag.BoolVar(&args.ShowVersion, "version", false, "print the version and exit")
	flag.StringVar(&args.CustomTemplate, "template", "", "use custom template")
	flag.StringVar(&args.Suffix, "suffix", ".mapping.pb.go", "use custom file suffix")

	flag.BoolVar(&args.Merge, "merge", false, "merge in a file")
	flag.StringVar(&args.Filename, "filename", "", "filename when merge enabled")
	flag.StringVar(&args.Package, "package", "", "package name when merge enabled")
	flag.StringVar(&args.GoPackage, "go_package", "", "go package when merge enabled")
}
func main() {
	flag.Parse()
	if args.ShowVersion {
		fmt.Printf("protoc-gen-saber-enum %v\n", version)
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
