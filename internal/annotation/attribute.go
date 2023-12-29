package annotation

// NameValue like `#[ident(name=value]`
type NameValue struct {
	// name
	Name string `parser:"@Ident '='"`
	// one of follow
	// String, Integer, Float, Bool,
	// StringList, IntegerList, FloatList, BoolList,
	Value Value `parser:"@@"`
}

type Value interface {
	value()
}

type String struct {
	Value string `parser:"@String"`
}

func (String) value() {}

type Integer struct {
	Value int64 `parser:"@Int"`
}

func (Integer) value() {}

type Float struct {
	Value float64 `parser:"@Float"`
}

func (Float) value() {}

type Bool struct {
	Value Boolean `parser:"@('true' | 'false')"`
}

func (Bool) value() {}

type Boolean bool

func (b *Boolean) Capture(values []string) error {
	*b = values[0] == "true"
	return nil
}

type StringList struct {
	Value []string `parser:"'[' (@String (',' @String)*)? ']'"`
}

func (StringList) value() {}

type IntegerList struct {
	Value []int64 `parser:"'[' (@Int (',' @Int)*)? ']'"`
}

func (IntegerList) value() {}

// NOTE: FloatList float list. must be first is float.
type FloatList struct {
	Value []float64 `parser:"'[' (@Float (',' (@Float | @Int))*)? ']'"`
}

func (FloatList) value() {}

type BoolList struct {
	Value []Boolean `parser:"'[' (@('true' | 'false') (',' @('true' | 'false'))*)? ']'"`
}

func (BoolList) value() {}
