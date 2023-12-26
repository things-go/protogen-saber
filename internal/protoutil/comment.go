package protoutil

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

type Comments []string

func NewComments(s protogen.Comments) Comments {
	return strings.Split(strings.TrimSuffix(s.String(), "\n"), "\n")
}

func (c Comments) Append(s string) Comments {
	return append(c, "// "+s)
}

func (c Comments) Annotations() Annotations {
	ms := make([]*Annotation, 0, len(c))
	for _, v := range c {
		if m := MatchAnnotation(v); m != nil {
			ms = append(ms, m)
		}
	}
	return ms
}

func (c Comments) FindAnnotation(path string) Annotations {
	ms := make([]*Annotation, 0, len(c))
	for _, v := range c {
		m := MatchAnnotation(v)
		if m != nil && strings.EqualFold(m.Path, path) {
			ms = append(ms, m)
		}
	}
	return ms
}

func (c Comments) FindAnnotationValues(path, key string) []string {
	ms := make([]string, 0, len(c))
	for _, v := range c {
		m := MatchAnnotation(v)
		if m != nil && strings.EqualFold(m.Path, path) && strings.EqualFold(m.Key, key) {
			ms = append(ms, m.Value)
		}
	}
	return ms
}

func (c Comments) String() string {
	if len(c) == 0 {
		return ""
	}
	var b []byte
	for i, line := range c {
		b = append(b, line...)
		if i+1 < len(c) {
			b = append(b, "\n"...)
		}
	}
	return string(b)
}
