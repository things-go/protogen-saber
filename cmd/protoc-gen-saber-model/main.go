package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

const version = "v0.0.5"

var args = &struct {
	ShowVersion bool   // 显示当前版本
	TrimPrefix  bool   // 去掉文件前缀
	Package     string // 覆盖默认包名
	Schema      string // file+mysql, file+tidb
	// ens options
	EnableInt          bool
	EnableIntegerInt   bool
	EnableBoolInt      bool
	DisableNullToPoint bool
	EnableForeignKey   bool
}{
	ShowVersion:        false,
	TrimPrefix:         false,
	Package:            "",
	Schema:             "file+mysql",
	EnableInt:          false,
	EnableIntegerInt:   false,
	EnableBoolInt:      false,
	DisableNullToPoint: false,
	EnableForeignKey:   false,
}

func init() {
	flag.BoolVar(&args.ShowVersion, "version", false, "print the version and exit")
	flag.BoolVar(&args.TrimPrefix, "trim_prefix", false, "trim filename prefix")
	flag.StringVar(&args.Package, "package", "", "override default package name")
	flag.StringVar(&args.Schema, "schema", "file+mysql", "ens driver, [file+mysql, file+tidy]")
	flag.BoolVar(&args.EnableInt, "enable_int", false, "enable [int8,uint8,int16,uint16,int32,uint32] output [int,uint]")
	flag.BoolVar(&args.EnableIntegerInt, "enable_integer_int", false, "enable [int32,uint32] output [int,uint]")
	flag.BoolVar(&args.EnableBoolInt, "enable_bool_int", false, "enable [bool] output [int]")
	flag.BoolVar(&args.DisableNullToPoint, "disable_null_to_point", false, "disable null out point type, use sql.Nullxx")
	flag.BoolVar(&args.EnableForeignKey, "enable_foreign_key", false, "out foreign key")
}

func main() {
	flag.Parse()
	if args.ShowVersion {
		fmt.Printf("protoc-gen-saber-model %v\n", version)
		return
	}

	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
