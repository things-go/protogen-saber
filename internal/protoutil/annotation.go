package protoutil

import (
	"regexp"
)

// annotation matches the following pattern
// `// #[Enum]` will find `// #[enum]`,`Enum`,“,“,“
// `// #[Enum(mapping="aaa")]` will find `// #[Enum(mapping="aaa")]`, "Enum", `(mapping="aaa")`, `mapping`, `aaa`
var rxAnnotation = regexp.MustCompile(`^//\s*#\[(\w+)\s*(\(\s*(\w+)\s*=\s*"(.*)"\s*\))?\s*\].*$`)

type Annotation struct {
	Path  string
	Key   string
	Value string
}

// MatchAnnotation 匹配注解
// `// #[xxx]`
// `// #[xxx(xx="xxxx")]`
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

// Len
func (c Annotations) Len() int {
	return len(c)
}

// Find all `path` annotation
func (c Annotations) Find(path string) Annotations {
	ret := make([]*Annotation, 0, len(c))
	for _, v := range c {
		if v.Path == path {
			ret = append(ret, v)
		}
	}
	return ret
}

// Find all `path` and `key` annotation value
func (c Annotations) FindValues(path, key string) []string {
	ret := make([]string, 0, len(c))
	for _, v := range c {
		if v.Path == path &&
			v.Key == key {
			ret = append(ret, v.Value)
		}
	}
	return ret
}
