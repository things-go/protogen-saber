package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

const version = "v0.0.4"

var args = &struct {
	ShowVersion bool   // 显示当前版本
	TrimPrefix  bool   // 去掉文件前缀
	Package     string // 覆盖默认包名
	Schema      string // file+mysql, file+tidb
}{
	ShowVersion: false,
	TrimPrefix:  false,
	Package:     "",
	Schema:      "file+mysql",
}

func init() {
	flag.BoolVar(&args.ShowVersion, "version", false, "print the version and exit")
	flag.BoolVar(&args.TrimPrefix, "trim_prefix", false, "trim filename prefix")
	flag.StringVar(&args.Package, "package", "", "override default package name")
	flag.StringVar(&args.Schema, "schema", "file+mysql", "ens driver, [file+mysql, file+tidy]")
}

func main() {
	flag.Parse()
	if args.ShowVersion {
		fmt.Printf("protoc-gen-saber-model %v\n", version)
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
