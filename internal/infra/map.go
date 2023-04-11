package infra

import (
	"sort"
	"strconv"
	"strings"
)

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
