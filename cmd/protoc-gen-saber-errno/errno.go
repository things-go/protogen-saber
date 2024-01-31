package main

import (
	"fmt"
	"log"
	"os"
	"slices"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/things-go/protogen-saber/internal/protoerrno"
	"github.com/things-go/protogen-saber/internal/protoutil"
)

func runProtoGen(gen *protogen.Plugin) error {
	if slices.Contains([]string{"builtin-eno", "builtin-est"}, args.CustomTemplate) && args.ErrorsPackage == "" {
		log.Fatal("errors package import path must be give with '--saber-errno_out=epk=xxx'")
	}
	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	usedTemplate, err := GetUsedTemplate()
	if err != nil {
		return err
	}
	for _, file := range gen.Files {
		if !file.Generate || len(file.Enums) == 0 {
			continue
		}
		enums := protoerrno.IntoEnums("", file.Enums)
		enums = append(enums, protoerrno.IntoEnumsFromMessage("", file.Messages)...)
		if len(enums) == 0 {
			continue
		}

		filename := file.GeneratedFilenamePrefix + ".errno.pb.go"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)
		mt := &ErrnoFile{
			Version:       version,
			ProtocVersion: protoutil.ProtocVersion(gen),
			IsDeprecated:  file.Proto.GetOptions().GetDeprecated(),
			Source:        file.Desc.Path(),
			Package:       string(file.GoPackageName),
			Epk:           args.ErrorsPackage,
			Errors:        enums,
		}
		err := mt.execute(usedTemplate, g)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr,
				"\u001B[31mWARN\u001B[m: execute template failed. %v\n", err)
		}
	}
	return nil
}
