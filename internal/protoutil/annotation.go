package protoutil

import (
	"regexp"
	"strings"
)

// annotation matches the following pattern
// `// #[Enum]“ will get `// #[enum]`,`Enum`,“,“,“
// `// #[Enum(mapping="aaa")]` will get `// #[Enum(mapping="aaa")]`, "Enum", `(mapping="aaa")`, `mapping`, `aaa`
var rxAnnotation = regexp.MustCompile(`^//\s*#\[(\w+)\s*(\(\s*(\w+)\s*=\s*"(.*)"\s*\))?\s*\].*$`)

type Annotation struct {
	Path  string
	Key   string
	Value string
}

func MatchAnnotation(s string) *Annotation {
	matches := rxAnnotation.FindStringSubmatch(s)
	if len(matches) != 5 {
		return nil
	}
	return &Annotation{
		Path:  matches[1],
		Key:   matches[3],
		Value: matches[4],
	}
}

type Annotations []*Annotation

func (c Annotations) Len() int {
	return len(c)
}

func (c Annotations) FindValues(path, key string) []string {
	vs := make([]string, 0, len(c))
	for _, v := range c {
		if strings.EqualFold(v.Path, path) && strings.EqualFold(v.Key, key) {
			vs = append(vs, v.Value)
		}
	}
	return vs
}

func (c Annotations) Find(path string) Annotations {
	vs := make([]*Annotation, 0, len(c))
	for _, v := range c {
		if strings.EqualFold(v.Path, path) {
			vs = append(vs, v)
		}
	}
	return vs
}
