package protoutil

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

// Source proto file source
type Source struct {
	Path         string
	IsDeprecated bool
}

func ProtocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

// ToArrayString convert to array string format [0:aaa,1:bbb,3:ccc]
func ToArrayString(values map[int]string) string {
	if len(values) == 0 {
		return "[]"
	}

	keys := make([]int, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}

	b := strings.Builder{}
	b.WriteString("[")
	sort.Ints(keys)
	for i, k := range keys {
		if i != 0 {
			b.WriteString(",")
		}
		b.WriteString(strconv.Itoa(k))
		b.WriteString(":")
		b.WriteString(values[k])
	}
	b.WriteString("]")
	return b.String()
}
