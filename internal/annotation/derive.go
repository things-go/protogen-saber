package annotation

// Derive an specified identity and it's attribute list.
// `#[ident]`
// `#[ident(k1=1,k2="2")]`
// `#[ident(k1=[1,2,3],k2=["1","2","3"])]`
type Derive struct {
	Identity string       `parser:"@Ident"`
	Attrs    []*NameValue `parser:"('(' (@@ (',' @@)*)? ')')?"`
}

// `#[ident]` only, not contain any attributes.
func (a *Derive) IsHeadless() bool {
	return len(a.Attrs) == 0
}

type Derives []*Derive

// ContainHeadless contain headless
func (a Derives) ContainHeadless(identity string) bool {
	for _, v := range a {
		if v.Identity == identity && v.IsHeadless() {
			return true
		}
	}
	return false
}

func (a Derives) Find(identity string) Derives {
	ret := make(Derives, 0, len(a))
	for _, v := range a {
		if v.Identity == identity {
			ret = append(ret, v)
		}
	}
	return ret
}

func (a Derives) FindValue(identity, name string) []Value {
	ret := make([]Value, 0, len(a))
	for _, v := range a {
		if v.Identity == identity {
			for _, vv := range v.Attrs {
				if vv.Name == name {
					ret = append(ret, vv.Value)
				}
			}
		}
	}
	return ret
}
