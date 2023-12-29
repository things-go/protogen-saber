package protoutil

import (
	"strings"

	"github.com/things-go/protogen-saber/internal/annotation"
	"google.golang.org/protobuf/compiler/protogen"
)

// CommentLines comment the line like `// xxxx`
type CommentLines []string

func NewCommentLines(s protogen.Comments) CommentLines {
	return strings.Split(strings.TrimSuffix(s.String(), "\n"), "\n")
}

// Derives all match the derive and remaining comment lines.
func (c CommentLines) Derives() (annotation.Derives, CommentLines) {
	remain := make(CommentLines, 0, len(c))
	ret := make([]*annotation.Derive, 0, len(c))
	for _, s := range c {
		if m, err := annotation.Match(strings.TrimSpace(strings.TrimPrefix(s, "//"))); err != nil {
			remain = append(remain, s)
		} else {
			ret = append(ret, m)
		}
	}
	return ret, remain
}

// Annotations find `identity` match the annotation and remaining comment lines.
func (c CommentLines) FindDerives(identity string) (annotation.Derives, CommentLines) {
	remain := make(CommentLines, 0, len(c))
	ret := make([]*annotation.Derive, 0, len(c))
	for _, s := range c {
		if m, err := annotation.Match(strings.TrimSpace(strings.TrimPrefix(s, "//"))); err != nil {
			remain = append(remain, s)
		} else if m.Identity == identity {
			ret = append(ret, m)
		}
	}
	return ret, remain
}

func (c *CommentLines) Append(s string) CommentLines {
	*c = append(*c, "// "+s)
	return *c
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
