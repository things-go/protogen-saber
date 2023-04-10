package infra

import (
	"sort"
	"strconv"
	"strings"
)

func ToArray(values map[int]string) string {
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
