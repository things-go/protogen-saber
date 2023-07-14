{{- range $tb := .Tables}}
-- {{$tb.Comment}}
CREATE TABLE 
	`{{$tb.Name}}` (
	{{- $colen := len $tb.Columns}}
	{{- $idxlen := len $tb.Indexes}}
	{{- range $idx, $e := $tb.Columns}}
		`{{$e.Name}}` {{$e.Type}} COMMENT '{{$e.Comment}}'{{- if eq (add $colen -1) $idx}}{{- if gt $idxlen 0}},{{- end}}{{- else}},{{- end}}
	{{- end}}
	{{- range $idx, $e := $tb.Indexes}}
		{{$e}}{{- if ne (add $idxlen -1)  $idx}},{{- end}}
	{{- end}}
	) ENGINE = {{$tb.Engine}} DEFAULT CHARSET = {{$tb.Charset}} COLLATE = {{$tb.Collate}} COMMENT = '{{$tb.Comment}}';
{{end -}}