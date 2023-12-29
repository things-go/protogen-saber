package protoseaql

import (
	"strings"

	"github.com/things-go/protogen-saber/internal/annotation"
	"github.com/things-go/protogen-saber/internal/protoutil"
	"google.golang.org/protobuf/compiler/protogen"
)

// annotation const value
const (
	Identity                  = "seaql"
	Attribute_Name_Name       = "name"
	Attribute_Name_Engine     = "engine"
	Attribute_Name_Charset    = "charset"
	Attribute_Name_Collate    = "collate"
	Attribute_Name_Index      = "index"
	Attribute_Name_ForeignKey = "foreign_key"
	Attribute_Name_Type       = "type"
)

type SeaqlDerive struct {
	Enabled     bool
	Name        string   // 表名
	Comment     string   // 注释
	Engine      string   // 引擎
	Charset     string   // 字符集
	Collate     string   // 排序规则
	Indexes     []string // 索引
	ForeignKeys []string // 外键
}

func ParseSeaqlDerive(rawTableName string, s protogen.Comments) *SeaqlDerive {
	ret := &SeaqlDerive{
		Enabled:     false,
		Name:        rawTableName,
		Comment:     "",
		Engine:      "InnoDB",
		Charset:     "utf8mb4",
		Collate:     "utf8mb4_general_ci",
		Indexes:     []string{},
		ForeignKeys: []string{},
	}
	derives, remainComments := protoutil.NewCommentLines(s).FindDerives(Identity)
	ret.Comment = strings.TrimSpace(strings.TrimPrefix(remainComments.LineString(), rawTableName))
	for _, d := range derives {
		if len(d.Attrs) == 0 {
			ret.Enabled = true
			continue
		}
		for _, v := range d.Attrs {
			switch v.Name {
			case Attribute_Name_Name:
				if vv, ok := v.Value.(annotation.String); ok && vv.Value != "" {
					ret.Name = vv.Value
				}
			case Attribute_Name_Engine:
				if vv, ok := v.Value.(annotation.String); ok && vv.Value != "" {
					ret.Engine = vv.Value
				}
			case Attribute_Name_Charset:
				if vv, ok := v.Value.(annotation.String); ok && vv.Value != "" {
					ret.Charset = vv.Value
				}
			case Attribute_Name_Collate:
				if vv, ok := v.Value.(annotation.String); ok && vv.Value != "" {
					ret.Collate = vv.Value
				}
			case Attribute_Name_Index:
				switch vv := v.Value.(type) {
				case annotation.String:
					if vv.Value != "" {
						ret.Indexes = append(ret.Indexes, vv.Value)
					}
				case annotation.StringList:
					for _, s := range vv.Value {
						if s != "" {
							ret.Indexes = append(ret.Indexes, s)
						}
					}
				}
			case Attribute_Name_ForeignKey:
				switch vv := v.Value.(type) {
				case annotation.String:
					if vv.Value != "" {
						ret.ForeignKeys = append(ret.ForeignKeys, vv.Value)
					}
				case annotation.StringList:
					for _, s := range vv.Value {
						if s != "" {
							ret.ForeignKeys = append(ret.ForeignKeys, s)
						}
					}
				}
			}
		}
	}
	return ret
}

type SeaqlValueDerive struct {
	Type string
}

func ParseSeaqlValueDerive(s protogen.Comments) (*SeaqlValueDerive, protoutil.CommentLines) {
	ret := &SeaqlValueDerive{Type: ""}
	derives, remainComments := protoutil.NewCommentLines(s).FindDerives(Identity)
	values := derives.FindValue(Identity, Attribute_Name_Type)
	for _, v := range values {
		if v, ok := v.(annotation.String); ok {
			ret.Type = v.Value
			break
		}
	}
	return ret, remainComments
}
