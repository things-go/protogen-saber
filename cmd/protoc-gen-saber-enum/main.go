package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

var args = &Args{
	ShowVersion:      false,
	DisableOrComment: false,
	CustomTemplate:   "",
	Suffix:           "",
	Merge:            false,
	Filename:         "",
	Package:          "",
	GoPackage:        "",
}

func init() {
	flag.BoolVar(&args.ShowVersion, "version", false, "print the version and exit")
	flag.BoolVar(&args.DisableOrComment, "disable_or_comment", false, "disable use comment if mapping value not exist. just use empty string ")
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
