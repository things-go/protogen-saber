package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

const version = "v0.0.2"

var showVersion = flag.Bool("version", false, "print the version and exit")
var disableOrComment = flag.Bool("disable_or_comment", false, "disable use comment if mapping value not exist. just use empty string ")
var trimPrefix = flag.Bool("trim_prefix", false, "trim filename prefix")

var merge = flag.Bool("merge", false, "merge in a file")
var filename = flag.String("filename", "create_table", "filename when merge enabled")

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-saber-seaql %v\n", version)
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
