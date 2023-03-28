package main

import (
	"embed"
	"errors"
	"io"
	"text/template"

	"github.com/things-go/protogen-saber/internal/infra"
)

// EnumValue 枚举的枚举项
type EnumValue struct {
	Number     int    // 编号
	Value      string // 值
	CamelValue string // 驼峰值
	Mapping    string // 映射值
	Comment    string // 注释
}

// Enum 枚举
// NOTE:
//
//	如果 MessageName 为空, 表明枚举独立, 枚举类型为 ${{Name}}, 枚举值为 ${{Name}}_${{Value}}
//	如果 MessageName 为不为空, 表明枚举嵌套在message里, 枚举类型为 ${{MessageName}}_{{Name}}, 枚举值为 ${{MessageName}}_${{Value}}
type Enum struct {
	MessageName string       // 嵌套消息名
	Name        string       // 名称
	Comment     string       // 注释
	Values      []*EnumValue // 枚举项
}

//go:embed enum.tpl
var Static embed.FS

var TemplateFuncs = template.FuncMap{
	"add":            func(a, b int) int { return a + b },
	"snakecase":      func(s string) string { return infra.SnakeCase(s) },
	"kebabcase":      func(s string) string { return infra.Kebab(s) },
	"camelcase":      func(s string) string { return infra.CamelCase(s) },
	"smallcamelcase": func(s string) string { return infra.SmallCamelCase(s) },
}
var enumTemplate = template.Must(template.New("components").
	Funcs(TemplateFuncs).
	ParseFS(Static, "enum.tpl")).
	Lookup("enum.tpl")

type EnumFile struct {
	Version       string
	ProtocVersion string
	IsDeprecated  bool
	Source        string
	Package       string
	Enums         []*Enum
}

func (e *EnumFile) execute(t *template.Template, w io.Writer) error {
	return t.Execute(w, e)
}

func ParseTemplateFromFile(filename string) (*template.Template, error) {
	if filename == "" {
		return nil, errors.New("required template filename")
	}
	tt, err := template.New("custom").
		Funcs(TemplateFuncs).
		ParseFiles(filename)
	if err != nil {
		return nil, err
	}
	ts := tt.Templates()
	if len(ts) == 0 {
		return nil, errors.New("not found any template")
	}
	return ts[0], nil
}
