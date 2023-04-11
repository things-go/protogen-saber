package main

import (
	"errors"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/things-go/protogen-saber/internal/protoenum"
	"github.com/things-go/protogen-saber/internal/protoutil"
)

func runProtoGen(gen *protogen.Plugin) error {
	var mergeEnums []*protoenum.Enum
	var source []string
	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	isMerge := *merge
	if *merge {
		if *_package == "" ||
			*filename == "" ||
			*goPackage == "" {
			return errors.New("when enable merge, filename,package,go_package must be set")
		}
		mergeEnums = make([]*protoenum.Enum, 0, len(gen.Files)*4)
		source = make([]string, 0, len(gen.Files))
	}
	usedTemplate := enumTemplate
	if *customTemplate != "" {
		t, err := ParseTemplateFromFile(*customTemplate)
		if err != nil {
			return err
		}
		usedTemplate = t
	}

	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		enums := protoenum.IntoEnums("", f.Enums, *disableOrComment)
		enums = append(enums, protoenum.IntoEnumsFromMessage("", f.Messages, *disableOrComment)...)
		if len(enums) == 0 {
			continue
		}
		if isMerge {
			source = append(source, f.Desc.Path())
			mergeEnums = append(mergeEnums, enums...)
			continue
		}
		g := gen.NewGeneratedFile(f.GeneratedFilenamePrefix+*suffix, f.GoImportPath)
		e := &EnumFile{
			Version:       version,
			ProtocVersion: protoutil.ProtocVersion(gen),
			IsDeprecated:  f.Proto.GetOptions().GetDeprecated(),
			Source:        f.Desc.Path(),
			Package:       string(f.GoPackageName),
			Enums:         enums,
		}
		_ = e.execute(usedTemplate, g)
	}
	if isMerge {
		g := gen.NewGeneratedFile(*filename+*suffix, protogen.GoImportPath(*goPackage))
		mergeFile := &EnumFile{
			Version:       version,
			ProtocVersion: protoutil.ProtocVersion(gen),
			IsDeprecated:  false,
			Source:        strings.Join(source, ","),
			Package:       *_package,
			Enums:         mergeEnums,
		}
		return mergeFile.execute(usedTemplate, g)
	}
	return nil
}
