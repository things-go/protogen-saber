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
func (d *Derive) IsHeadless() bool {
	return len(d.Attrs) == 0
}

type Derives []*Derive

// ContainHeadless contain headless
func (ds Derives) ContainHeadless(identity string) bool {
	for _, v := range ds {
		if v.Identity == identity && v.IsHeadless() {
			return true
		}
	}
	return false
}

func (ds Derives) Find(identity string) Derives {
	ret := make(Derives, 0, len(ds))
	for _, v := range ds {
		if v.Identity == identity {
			ret = append(ret, v)
		}
	}
	return ret
}

func (ds Derives) FindValue(identity, name string) []ValueType {
	ret := make([]ValueType, 0, len(ds))
	for _, v := range ds {
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
