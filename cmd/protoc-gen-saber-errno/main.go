package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

const version = "v0.0.1"

var args = &struct {
	ShowVersion    bool   // 显示版本
	CustomTemplate string // 自定义模板
	ErrorsPackage  string // error package
}{

	ShowVersion:    false,
	CustomTemplate: "",
	ErrorsPackage:  "",
}

func init() {
	flag.BoolVar(&args.ShowVersion, "version", false, "print the version and exit")
	flag.StringVar(&args.CustomTemplate, "template", "builtin-eno", "use custom template except [builtin-eno, builtin-est]")
	flag.StringVar(&args.ErrorsPackage, "epk", "github.com/things-go/dyn/genproto/errors", "errors core package in your project")
}

func main() {
	flag.Parse()
	if args.ShowVersion {
		fmt.Printf("protoc-gen-saber-errno %v\n", version)
		return
	}
	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
