package main

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/things-go/protogen-saber/internal/infra"
	"github.com/things-go/protogen-saber/protosaber/enumerate"
	"github.com/things-go/protogen-saber/protosaber/seaql"
)

func runProtoGen(gen *protogen.Plugin) error {
	gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, f := range gen.Files {
		if !f.Generate {
			continue
		}
		tables := intoTable(f.Messages)
		if len(tables) == 0 {
			continue
		}

		dir := filepath.Dir(f.GeneratedFilenamePrefix)
		for _, tb := range tables {
			g := gen.NewGeneratedFile(filepath.Join(dir, tb.Name)+".sql", f.GoImportPath)
			e := &File{
				Version:       version,
				ProtocVersion: infra.ProtocVersion(gen),
				IsDeprecated:  f.Proto.GetOptions().GetDeprecated(),
				Source:        f.Desc.Path(),
				Table:         tb,
			}

			_ = e.execute(enumTemplate, g)
		}
	}
	return nil
}

// intoTable generates the errors definitions, excluding the package statement.
func intoTable(protoMessages []*protogen.Message) []Table {
	tables := make([]Table, 0, len(protoMessages))
	for _, pe := range protoMessages {
		if len(pe.Fields) == 0 {
			continue
		}
		messageOptions := proto.GetExtension(pe.Desc.Options(), seaql.E_Options)
		seaOptions, ok := messageOptions.(*seaql.Options)
		if !ok || seaOptions == nil {
			continue
		}

		columns := make([]Column, 0, len(pe.Fields))
		for _, v := range pe.Fields {
			messageFieldOptions := proto.GetExtension(v.Desc.Options(), seaql.E_Field)
			seaFieldOptions := messageFieldOptions.(*seaql.Field)

			comment := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSuffix(string(v.Comments.Leading), "\n"), "\n", ","), " ", "")
			if enumComment := intoEnumComment(v.Enum); enumComment != "" {
				comment += "," + enumComment
			}
			columns = append(columns, Column{
				Name:    string(v.Desc.Name()),
				Type:    seaFieldOptions.Type,
				Comment: comment,
			})
		}
		tableName := string(pe.Desc.Name())
		if seaOptions.TableName != "" {
			tableName = seaOptions.TableName
		}
		engine := "InnoDB"
		if seaOptions.Engine != "" {
			engine = seaOptions.Engine
		}
		charset := "utf8mb4"
		if seaOptions.Charset != "" {
			charset = seaOptions.Charset
		}
		collate := "utf8mb4_general_ci"
		if seaOptions.Collate != "" {
			collate = seaOptions.Collate
		}
		tables = append(tables, Table{
			Name:    infra.SnakeCase(tableName, false),
			Comment: strings.TrimSpace(strings.ReplaceAll(string(pe.Comments.Leading), "\n", "")),
			Engine:  engine,
			Charset: charset,
			Collate: collate,
			Columns: columns,
			Indexes: seaOptions.Index,
		})
		if len(pe.Messages) > 0 {
			tables = append(tables, intoTable(pe.Messages)...)
		}
	}

	return tables
}

// intoEnumComment generates enum comment if it exists
func intoEnumComment(pe *protogen.Enum) string {
	if pe == nil || len(pe.Values) == 0 {
		return ""
	}
	isEnabled := proto.GetExtension(pe.Desc.Options(), enumerate.E_Enabled)
	ok := isEnabled.(bool)
	if !ok {
		return ""
	}

	eValueMp := make(map[int]string, len(pe.Values))
	for _, v := range pe.Values {
		mpv := proto.GetExtension(v.Desc.Options(), enumerate.E_Mapping)
		mappingValue, _ := mpv.(string)

		eValueMp[v.Desc.Index()] = mappingValue
	}
	b, _ := json.Marshal(eValueMp)
	bb := strings.ReplaceAll(string(b), `"`, "")
	bb = strings.Replace(bb, "{", "[", 1)
	bb = strings.Replace(bb, "}", "]", 1)
	return bb
}
