package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

var args = &Args{
	ShowVersion:      false,
	DisableOrComment: false,
	ModelSource:      "",
}

func init() {
	flag.BoolVar(&args.ShowVersion, "version", false, "print the version and exit")
	flag.BoolVar(&args.DisableOrComment, "disable_or_comment", false, "disable use comment if mapping value not exist. just use empty string ")
	flag.StringVar(&args.ModelSource, "model_source", "", "model package")
}

func main() {
	flag.Parse()
	if args.ShowVersion {
		fmt.Printf("protoc-gen-saber-assist %v\n", version)
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
