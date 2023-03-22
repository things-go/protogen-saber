package main

import (
	"embed"
	"io"
	"text/template"

	"github.com/things-go/protogen-saber/internal/infra"
)

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

//go:embed seaql.tpl
var Static embed.FS

var TemplateFuncs = template.FuncMap{
	"add":            func(a, b int) int { return a + b },
	"snakecase":      func(s string) string { return infra.SnakeCase(s, false) },
	"kebabcase":      func(s string) string { return infra.Kebab(s, false) },
	"camelcase":      func(s string) string { return infra.CamelCase(s) },
	"smallcamelcase": func(s string) string { return infra.SmallCamelCase(s, false) },
}
var enumTemplate = template.Must(template.New("components").
	Funcs(TemplateFuncs).
	ParseFS(Static, "seaql.tpl")).
	Lookup("seaql.tpl")

type File struct {
	Version       string
	ProtocVersion string
	IsDeprecated  bool
	Source        string
	Table         Table
}

func (e *File) execute(t *template.Template, w io.Writer) error {
	return t.Execute(w, e)
}
