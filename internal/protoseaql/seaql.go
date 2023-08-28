package protoseaql

import (
	"fmt"
	"io"
	"strings"

	"github.com/things-go/protogen-saber/internal/infra"
	"github.com/things-go/protogen-saber/internal/protoenum"
	"github.com/things-go/protogen-saber/protosaber/seaql"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

type Schema struct {
	Tables []Table // 表
}

type Column struct {
	Name    string // 列名
	Type    string // sql 定义
	Comment string // 注释
}

type Table struct {
	Name        string   // 表名
	Comment     string   // 注释
	Engine      string   // 引擎
	Charset     string   // 字符集
	Collate     string   // 排序规则
	Indexes     []string // 索引
	ForeignKeys []string // 外键
	Columns     []Column // 列项
}

func Execute(w io.Writer, data *Schema) error {
	for i, tb := range data.Tables {
		fmt.Fprintf(w, "-- %s\n", tb.Comment)
		fmt.Fprintf(w, "CREATE TABLE \n")
		fmt.Fprintf(w, "\t`%s` (\n", tb.Name)
		extLen := len(tb.Indexes) + len(tb.ForeignKeys)
		for idx, col := range tb.Columns {
			suffix := ","
			if len(tb.Columns) == idx+1 && extLen == 0 {
				suffix = ""
			}
			fmt.Fprintf(w, "\t\t`%s` %s COMMENT '%s'%s\n", col.Name, col.Type, col.Comment, suffix)
		}
		for idx, val := range append(tb.Indexes, tb.ForeignKeys...) {
			suffix := ","
			if idx == extLen-1 {
				suffix = ""
			}
			fmt.Fprintf(w, "\t\t%s%s\n", val, suffix)
		}
		fmt.Fprintf(w, "\t) ENGINE = %s DEFAULT CHARSET = %s COLLATE = %s COMMENT = '%s';\n", tb.Engine, tb.Charset, tb.Collate, tb.Comment)
		if len(data.Tables) != i+1 {
			fmt.Fprintf(w, "\n")
		}
	}
	return nil
}

// IntoTable generates the errors definitions, excluding the package statement.
func IntoTable(protoMessages []*protogen.Message) ([]Table, error) {
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
			if enumComment := protoenum.IntoEnumComment(v.Enum); enumComment != "" {
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
			Name:        infra.SnakeCase(tableName),
			Comment:     strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(strings.ReplaceAll(string(pe.Comments.Leading), "\n", "")), rawTableName)),
			Engine:      engine,
			Charset:     charset,
			Collate:     collate,
			Columns:     columns,
			Indexes:     seaOptions.Index,
			ForeignKeys: seaOptions.ForeignKey,
		})
		if len(pe.Messages) > 0 {
			tbs, err := IntoTable(pe.Messages)
			if err != nil {
				return nil, err
			}
			tables = append(tables, tbs...)
		}
	}
	return tables, nil
}
