package main

import (
	"flag"
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

var args = &struct {
	ShowVersion  bool   // 显示版本
	DisableSaber bool   // 禁用saber依赖
	Codec        string // 编解码器(DisableSaber=true), 默认 proto
	CodecPackage string // 编解码器包路径, 需要实现函数 `Marshal(v any) ([]byte, error)` 和 `Unmarshal(data []byte, v any) error`
}{
	ShowVersion:  false,
	DisableSaber: false,
	Codec:        "proto",
	CodecPackage: "",
}

func init() {
	flag.BoolVar(&args.ShowVersion, "version", false, "print the version and exit")
	flag.BoolVar(&args.DisableSaber, "disable_saber", false, "disable saber dependence")
	flag.StringVar(&args.Codec, "codec", "proto", "codec(enabled when disable_saber=true)")
	flag.StringVar(&args.CodecPackage, "codec_package", "", "enabled when disable_saber=true and codec=custom, `Marshal(v any) ([]byte, error)` and `Unmarshal(data []byte, v any) error` should be implement)")
}

func main() {
	flag.Parse()
	if args.ShowVersion {
		fmt.Printf("protoc-gen-saber-asynq %v\n", version)
		return
	}
	protogen.Options{ParamFunc: flag.CommandLine.Set}.Run(runProtoGen)
}
