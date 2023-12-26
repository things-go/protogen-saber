package protoseaql

import (
	"fmt"
	"io"
	"strings"

	"github.com/things-go/protogen-saber/internal/infra"
	"github.com/things-go/protogen-saber/internal/protoenum"
	"github.com/things-go/protogen-saber/internal/protoutil"
	"github.com/things-go/protogen-saber/protosaber/seaql"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// annotation const value
const (
	annotation_Path           = "seaql"
	annotation_Key_Name       = "table_name"
	annotation_Key_Engine     = "engine"
	annotation_Key_Charset    = "charset"
	annotation_Key_Collate    = "collate"
	annotation_Key_Index      = "index"
	annotation_Key_ForeignKey = "foreign_key"
	annotation_Key_Type       = "type"
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
		rawTableName := string(pe.Desc.Name())
		var tableName = rawTableName
		var engine = "InnoDB"
		var charset = "utf8mb4"
		var collate = "utf8mb4_general_ci"
		var indexes []string
		var foreignKey []string

		// 先判断注解, 再判断扩展
		annotates, remainComments := protoutil.NewComments(pe.Comments.Leading).FindAnnotation2(annotation_Path)
		if len(annotates) > 0 {
			for _, v := range annotates {
				switch v.Key {
				case annotation_Key_Name:
					tableName = v.Value
				case annotation_Key_Engine:
					engine = v.Value
				case annotation_Key_Charset:
					charset = v.Value
				case annotation_Key_Collate:
					collate = v.Value
				case annotation_Key_Index:
					indexes = append(indexes, v.Value)
				case annotation_Key_ForeignKey:
					foreignKey = append(foreignKey, v.Value)
				}
			}
		} else {
			messageOptions := proto.GetExtension(pe.Desc.Options(), seaql.E_Options)
			seaOptions, ok := messageOptions.(*seaql.Options)
			if !ok || seaOptions == nil {
				continue
			}
			if seaOptions.TableName != "" {
				tableName = seaOptions.TableName
			}
			if seaOptions.Engine != "" {
				engine = seaOptions.Engine
			}
			if seaOptions.Charset != "" {
				charset = seaOptions.Charset
			}
			if seaOptions.Collate != "" {
				collate = seaOptions.Collate
			}
			indexes = seaOptions.Index
			foreignKey = seaOptions.ForeignKey
		}
		comment := strings.TrimSpace(strings.TrimPrefix(remainComments.LineString(), rawTableName))

		columns := make([]Column, 0, len(pe.Fields))
		for _, v := range pe.Fields {
			ty := ""
			// 先判断注解, 再判断扩展
			annotateValues, remainComments := protoutil.NewComments(v.Comments.Leading).FindAnnotationValues2(annotation_Path, annotation_Key_Type)
			if len(annotates) > 0 && annotateValues[0] != "" {
				ty = annotateValues[0]
			} else {
				messageFieldOptions := proto.GetExtension(v.Desc.Options(), seaql.E_Field)
				seaFieldOptions := messageFieldOptions.(*seaql.Field)
				if seaFieldOptions == nil {
					return nil, fmt.Errorf("seaql: message(%s) - field(%s) is not set seaql type", pe.Desc.Name(), string(v.Desc.Name()))
				}
				seaFieldOptions.Type = strings.TrimSpace(seaFieldOptions.Type)
				if seaFieldOptions.Type == "" {
					return nil, fmt.Errorf("seaql: message(%s) - field(%s) should be not empty", pe.Desc.Name(), string(v.Desc.Name()))
				}
				ty = seaFieldOptions.Type
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
			Name:        infra.SnakeCase(tableName),
			Comment:     comment,
			Engine:      engine,
			Charset:     charset,
			Collate:     collate,
			Columns:     columns,
			Indexes:     indexes,
			ForeignKeys: foreignKey,
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
