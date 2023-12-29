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

type Annotation struct {
	Derive Derive `parser:"'#' '[' @@ ']'"`
}

// Match 匹配注解
// `#[ident]`
// `#[ident(k1=1,k2="2")]`
// `#[ident(k1=[1,2,3],k2=["1","2","3"])]`
func Match(s string) (*Derive, error) {
	p, err := parser.ParseString("", s)
	if err != nil {
		return nil, err
	}
	return &p.Derive, nil
}
