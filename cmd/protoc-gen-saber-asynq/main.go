package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

var args = &struct {
	ShowVersion bool // 显示版本
}{
	ShowVersion: false,
}

func init() {
	flag.BoolVar(&args.ShowVersion, "version", false, "print the version and exit")
}

func main() {
	flag.Parse()
	if args.ShowVersion {
		fmt.Printf("protoc-gen-saber-asynq %v\n", version)
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
