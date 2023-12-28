package annotation

import (
	"github.com/alecthomas/participle/v2"
)

var parser = participle.MustBuild[Annotation](
	participle.Unquote(),
	participle.Union[Value](
		String{}, Integer{}, Float{}, Bool{},
		StringList{}, IntegerList{}, FloatList{}, BoolList{},
	),
)

// Annotation an specified identifier and it's attribute list.
// `#[ident]`
// `#[ident(k1=v1,k2=v2)]`
type Annotation struct {
	Identifier string       `parser:"'#' '[' @Ident"`
	Attrs      []*NameValue `parser:"('(' (@@ (',' @@)*)? ')')?"`
	Empty      struct{}     `parser:"']'"`
}

// `#[ident]` only, not contain any attributes.
func (a *Annotation) IsHeadless() bool {
	return len(a.Attrs) == 0
}

// NameValue like `#[ident(name=value]`
type NameValue struct {
	// name
	Name string `parser:"@Ident '='"`
	// one of follow
	// String, Integer, Float, Bool,
	// StringList, IntegerList, FloatList, BoolList,
	Value Value `parser:"@@"`
}

// Match 匹配注解
// `#[ident]`
// `#[ident(k1=v1,k2=v2)]`
func Match(s string) (*Annotation, error) {
	return parser.ParseString("", s)
}

type Annotations []*Annotation

// ContainHeadless contain headless
func (a Annotations) ContainHeadless(identifier string) bool {
	for _, v := range a {
		if v.Identifier == identifier && v.IsHeadless() {
			return true
		}
	}
	return false
}

func (a Annotations) Find(identifier string) Annotations {
	ret := make(Annotations, 0, len(a))
	for _, v := range a {
		if v.Identifier == identifier {
			ret = append(ret, v)
		}
	}
	return ret
}

func (a Annotations) FindValue(identifier, name string) []Value {
	ret := make([]Value, 0, len(a))
	for _, v := range a {
		if v.Identifier == identifier {
			for _, vv := range v.Attrs {
				if vv.Name == name {
					ret = append(ret, vv.Value)
				}
			}
		}
	}
	return ret
}
