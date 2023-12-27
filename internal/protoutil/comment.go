package protoutil

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

// CommentLines comment the line like `// xxxx`
type CommentLines []string

func NewCommentLines(s protogen.Comments) CommentLines {
	return strings.Split(strings.TrimSuffix(s.String(), "\n"), "\n")
}

func (c *CommentLines) Append(s string) CommentLines {
	*c = append(*c, "// "+s)
	return *c
}

// Annotations all match the annotation.
func (c CommentLines) Annotations() Annotations {
	ret := make([]*Annotation, 0, len(c))
	for _, v := range c {
		if m := MatchAnnotation(v); m != nil {
			ret = append(ret, m)
		}
	}
	return ret
}

// FindAnnotation all `path` annotation and remaining comment lines.
// return the remaining comment line except the annotation
func (c CommentLines) FindAnnotation(path string) (Annotations, CommentLines) {
	remain := make(CommentLines, 0, len(c))
	ret := make([]*Annotation, 0, len(c))
	for _, v := range c {
		m := MatchAnnotation(v)
		if m != nil &&
			m.Path == path {
			ret = append(ret, m)
		} else {
			remain = append(remain, v)
		}
	}
	return ret, remain
}

// FindAnnotation all `path` and `key` annotation value and remaining comment lines.
// return the remaining comment line except the annotation
func (c CommentLines) FindAnnotationValues(path, key string) ([]string, CommentLines) {
	remain := make(CommentLines, 0, len(c))
	ms := make([]string, 0, len(c))
	for _, v := range c {
		m := MatchAnnotation(v)
		if m != nil &&
			m.Path == path && m.Key == key {
			ms = append(ms, m.Value)
		} else {
			remain = append(remain, v)
		}
	}
	return ms, remain
}

func (c CommentLines) String() string {
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

// LineString one line string.
func (c CommentLines) LineString() string {
	if len(c) == 0 {
		return ""
	}
	var b []byte
	for i, line := range c {
		b = append(b, []byte(strings.TrimSpace(strings.TrimPrefix(line, "//")))...)
		if i+1 < len(c) {
			b = append(b, ","...)
		}
	}
	return string(b)
}
