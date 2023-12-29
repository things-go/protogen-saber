package protoseaql

import (
	"fmt"
	"io"
	"strings"

	"github.com/things-go/protogen-saber/internal/infra"
	"github.com/things-go/protogen-saber/internal/protoenum"
	"google.golang.org/protobuf/compiler/protogen"
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
		seaqlAnnotation := ParseSeaqlDerive(string(pe.Desc.Name()), pe.Comments.Leading)
		if !seaqlAnnotation.Enabled {
			continue
		}

		columns := make([]Column, 0, len(pe.Fields))
		for _, v := range pe.Fields {
			seaqlValueAnnotate, remainComments := ParseSeaqlValueDerive(v.Comments.Leading)
			ty := seaqlValueAnnotate.Type
			if ty == "" {
				return nil, fmt.Errorf("seaql: message(%s) - field(%s) type should be not empty", pe.Desc.Name(), string(v.Desc.Name()))
			}

			comment := strings.ReplaceAll(remainComments.LineString(), " ", "")
			if enumComment := protoenum.IntoEnumComment(v.Enum); enumComment != "" {
				comment += "," + enumComment
			}

			columns = append(columns, Column{
				Name:    string(v.Desc.Name()),
				Type:    ty,
				Comment: comment,
			})
		}

		tables = append(tables, Table{
			Name:        infra.SnakeCase(seaqlAnnotation.Name),
			Comment:     seaqlAnnotation.Comment,
			Engine:      seaqlAnnotation.Engine,
			Charset:     seaqlAnnotation.Charset,
			Collate:     seaqlAnnotation.Collate,
			Indexes:     seaqlAnnotation.Indexes,
			ForeignKeys: seaqlAnnotation.ForeignKeys,
			Columns:     columns,
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
