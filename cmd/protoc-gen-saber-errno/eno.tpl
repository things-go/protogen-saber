// Code generated by protoc-gen-saber-errno. DO NOT EDIT.
// versions:
//   - protoc-gen-saber-errno {{.Version}}
//   - protoc                 {{.ProtocVersion}}
{{- if .IsDeprecated}}
// {{.Source}} is a deprecated file.
{{- else}}
// source: {{.Source}}
{{- end}}

package {{.Package}}

import (
	"fmt"
{{- if .Epk}}
	errors "{{.Epk}}"
{{- end}}
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
{{- if .Epk}}
var _ =	errors.New
{{- end}}

type Option interface {
	apply(*errors.Error)
}
type optFunc func(e *errors.Error)
func (o optFunc) apply(e *errors.Error) { o(e) }
func WithMessage(s string) Option {
	return optFunc(func(e *errors.Error) {
		if s != "" {
			e.Message = s
		}
	})
}
func WithDetail(s string) Option {
	return optFunc(func(e *errors.Error) {
		if s != "" {
			e.Detail = s
		}
	})
}
func WithMetadata(k string, v string) Option {
	return optFunc(func(e *errors.Error) {
		if k != "" && v != "" {
			e.Metadata[k] = v
		}
	})
}
func _apply(e *errors.Error, opts ...Option) {
	for _, opt := range opts {
		opt.apply(e)
	}
}

{{- range $e := .Errors}}
{{$enumName := $e.Name}}
{{ range .Values }}
func Is{{.CamelValue}}(err error) bool {
	e := errors.FromError(err)
	return e.Code == {{.Code}}
}
func Err{{.CamelValue}}({{if or (eq .Code 400) (eq .Code 500)}}detail string{{else}}message ...string{{end}}) *errors.Error {
{{- if or (eq .Code 400) (eq .Code 500)}}
	return errors.New({{.Code}}, "{{.Message}}", detail)
{{- else}}
	if len(message) > 0 {
	   return Err{{.CamelValue}}w(WithMessage(message[0]))
	}
    return Err{{.CamelValue}}w()
{{- end}}
}
func Err{{.CamelValue}}f(format string, args ...any) *errors.Error {
{{- if or (eq .Code 400) (eq .Code 500)}}
	 return errors.New({{.Code}}, "{{.Message}}", fmt.Sprintf(format, args...))
{{- else}}
	 return Err{{.CamelValue}}w(WithMessage(fmt.Sprintf(format, args...)))
{{- end}}
}
func Err{{.CamelValue}}w(opt ...Option) *errors.Error {
	e := errors.New({{.Code}}, "{{.Message}}", {{$enumName}}_{{.Value}}.String())
	_apply(e, opt...)
	return e
}
{{- end }}
{{- end }}