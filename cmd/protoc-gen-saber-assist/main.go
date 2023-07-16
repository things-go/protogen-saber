package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

var args = &Args{
	ShowVersion:      false,
	TrimPrefix:       false,
	DisableOrComment: false,
	Package:          "",
	ModelImportPath:  "",
}

func init() {
	flag.BoolVar(&args.ShowVersion, "version", false, "print the version and exit")
	flag.BoolVar(&args.TrimPrefix, "trim_prefix", false, "trim filename prefix")
	flag.BoolVar(&args.DisableOrComment, "disable_or_comment", false, "disable use comment if mapping value not exist. just use empty string ")
	flag.StringVar(&args.Package, "package", "", "override default package name")
	flag.StringVar(&args.ModelImportPath, "model_import_path", "", "model import path")
}

func main() {
	flag.Parse()
	if args.ShowVersion {
		fmt.Printf("protoc-gen-saber-assist %v\n", version)
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
