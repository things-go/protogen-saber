package protoseaql

import (
	"embed"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/things-go/protogen-saber/internal/infra"
	"github.com/things-go/protogen-saber/internal/protoenum"
	"github.com/things-go/protogen-saber/protosaber/seaql"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

//go:embed ddl.tpl
var ddl embed.FS

var ddlTemplate = template.Must(template.New("components").
	Funcs(template.FuncMap{
		"add":            func(a, b int) int { return a + b },
		"snakecase":      func(s string) string { return infra.SnakeCase(s) },
		"kebabcase":      func(s string) string { return infra.Kebab(s) },
		"camelcase":      func(s string) string { return infra.CamelCase(s) },
		"smallcamelcase": func(s string) string { return infra.SmallCamelCase(s) },
	}).
	ParseFS(ddl, "ddl.tpl")).
	Lookup("ddl.tpl")

type Schema struct {
	Tables []Table
}

type Column struct {
	Name    string
	Type    string
	Comment string
}

type Table struct {
	Name    string   // 表名
	Comment string   // 注释
	Engine  string   // 引擎
	Charset string   // 字符集
	Collate string   // 排序规则
	Indexes []string // 索引
	Columns []Column // 列项
}

func Execute(w io.Writer, data *Schema) error {
	return ddlTemplate.Execute(w, data)
}

// IntoTable generates the errors definitions, excluding the package statement.
func IntoTable(protoMessages []*protogen.Message, disableOrComment bool) ([]Table, error) {
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
			if seaFieldOptions == nil {
				return nil, fmt.Errorf("seaql: message(%s) - field(%s) is not set seaql type", pe.Desc.Name(), string(v.Desc.Name()))
			}
			seaFieldOptions.Type = strings.TrimSpace(seaFieldOptions.Type)
			if seaFieldOptions.Type == "" {
				return nil, fmt.Errorf("seaql: message(%s) - field(%s) should be not empty", pe.Desc.Name(), string(v.Desc.Name()))
			}

			comment := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSuffix(string(v.Comments.Leading), "\n"), "\n", ","), " ", "")
			if enumComment := protoenum.IntoEnumComment(v.Enum, disableOrComment); enumComment != "" {
				comment += "," + enumComment
			}
			columns = append(columns, Column{
				Name:    string(v.Desc.Name()),
				Type:    seaFieldOptions.Type,
				Comment: comment,
			})
		}
		rawTableName := string(pe.Desc.Name())
		tableName := rawTableName
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
			Name:    infra.SnakeCase(tableName),
			Comment: strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(strings.ReplaceAll(string(pe.Comments.Leading), "\n", "")), rawTableName)),
			Engine:  engine,
			Charset: charset,
			Collate: collate,
			Columns: columns,
			Indexes: seaOptions.Index,
		})
		if len(pe.Messages) > 0 {
			tmpTables, err := IntoTable(pe.Messages, disableOrComment)
			if err != nil {
				return nil, err
			}
			tables = append(tables, tmpTables...)
		}
	}
	return tables, nil
}
