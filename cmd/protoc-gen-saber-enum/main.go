package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

const version = "v0.0.2"

var showVersion = flag.Bool("version", false, "print the version and exit")
var disableOrComment = flag.Bool("disable_or_comment", false, "disable use comment if mapping value not exist. just use empty string ")
var customTemplate = flag.String("template", "", "use custom template")
var suffix = flag.String("suffix", ".mapping.pb.go", "use custom file suffix")

var merge = flag.Bool("merge", false, "merge in a file")
var filename = flag.String("filename", "", "filename when merge enabled")
var _package = flag.String("package", "", "package name when merge enabled")
var goPackage = flag.String("go_package", "", "go package when merge enabled")

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-saber-enum %v\n", version)
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
