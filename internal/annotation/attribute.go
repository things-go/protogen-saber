package annotation

// NameValue like `#[ident(name=value]`
type NameValue struct {
	// name
	Name string `parser:"@Ident '='"`
	// one of follow
	// String, Integer, Float, Bool,
	// StringList, IntegerList, FloatList, BoolList,
	Value ValueType `parser:"@@"`
}

type ValueType interface {
	Type() string
}

type String struct {
	Value string `parser:"@String"`
}

func (String) Type() string { return "string" }

type Integer struct {
	Value int64 `parser:"@Int"`
}

func (Integer) Type() string { return "integer" }

type Float struct {
	Value float64 `parser:"@Float"`
}

func (Float) Type() string { return "float" }

type Bool struct {
	Value Boolean `parser:"@('true' | 'false')"`
}

func (Bool) Type() string { return "bool" }

type Boolean bool

func (b *Boolean) Capture(values []string) error {
	*b = values[0] == "true"
	return nil
}

type StringList struct {
	Value []string `parser:"'[' (@String (',' @String)*)? ']'"`
}

func (StringList) Type() string { return "slice<string>" }

type IntegerList struct {
	Value []int64 `parser:"'[' (@Int (',' @Int)*)? ']'"`
}

func (IntegerList) Type() string { return "slice<integer>" }

// NOTE: FloatList float list. must be first is float.
type FloatList struct {
	Value []float64 `parser:"'[' (@Float (',' (@Float | @Int))*)? ']'"`
}

func (FloatList) Type() string { return "slice<float>" }

type BoolList struct {
	Value []Boolean `parser:"'[' (@('true' | 'false') (',' @('true' | 'false'))*)? ']'"`
}

func (BoolList) Type() string { return "slice<bool>" }
