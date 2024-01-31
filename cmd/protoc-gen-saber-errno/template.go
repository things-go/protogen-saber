package main

import (
	"embed"
	"errors"
	"html/template"
	"io"

	"github.com/things-go/protogen-saber/internal/protoerrno"
)

//go:embed eno.tpl est.tpl
var Static embed.FS
var (
	tpl = template.Must(template.New("components").
		ParseFS(Static, "eno.tpl", "est.tpl"))
	enoTpl = tpl.Lookup("eno.tpl")
	estTpl = tpl.Lookup("est.tpl")
)

type ErrnoFile struct {
	Version       string
	ProtocVersion string
	IsDeprecated  bool
	Source        string
	Package       string
	Epk           string
	Errors        []*protoerrno.Enum
}

func GetUsedTemplate() (*template.Template, error) {
	switch args.CustomTemplate {
	case "builtin-eno":
		return enoTpl, nil
	case "builtin-est":
		return estTpl, nil
	default:
		t, err := ParseTemplateFromFile(args.CustomTemplate)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
}

func (e *ErrnoFile) execute(t *template.Template, w io.Writer) error {
	return t.Execute(w, e)
}

func ParseTemplateFromFile(filename string) (*template.Template, error) {
	if filename == "" {
		return nil, errors.New("required template filename")
	}
	tt, err := template.New("custom").ParseFiles(filename)
	if err != nil {
		return nil, err
	}
	ts := tt.Templates()
	if len(ts) == 0 {
		return nil, errors.New("not found any template")
	}
	return ts[0], nil
}
