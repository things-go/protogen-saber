package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

var args = &Args{
	ShowVersion:      false,
	DisableOrComment: false,
	TrimPrefix:       false,
	Merge:            false,
	Filename:         "",
}

func init() {
	flag.BoolVar(&args.ShowVersion, "version", false, "print the version and exit")
	flag.BoolVar(&args.DisableOrComment, "disable_or_comment", false, "disable use comment if mapping value not exist. just use empty string ")
	flag.BoolVar(&args.TrimPrefix, "trim_prefix", false, "trim filename prefix")
	flag.BoolVar(&args.Merge, "merge", false, "merge in a file")
	flag.StringVar(&args.Filename, "filename", "create_table", "filename when merge enabled")
}

func main() {
	flag.Parse()
	if args.ShowVersion {
		fmt.Printf("protoc-gen-saber-seaql %v\n", version)
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
