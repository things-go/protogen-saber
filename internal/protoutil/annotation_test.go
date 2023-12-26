package protoutil

import (
	"reflect"
	"regexp"
	"testing"
)

func TestEnumRegexp(t *testing.T) {
	testCases := []struct {
		name string
		rx   *regexp.Regexp
		arg  string
		want []string
	}{
		{
			name: "单个 - 标准",
			rx:   rxAnnotation,
			arg:  "// #[Enum]",
			want: []string{"// #[Enum]", "Enum", "", "", ""},
		},
		{
			name: "单个 - //后面不带空格",
			rx:   rxAnnotation,
			arg:  "//#[Enum]",
			want: []string{"//#[Enum]", "Enum", "", "", ""},
		},
		{
			name: "单个 - //后面带多个空格",
			rx:   rxAnnotation,
			arg:  "//    #[Enum]",
			want: []string{"//    #[Enum]", "Enum", "", "", ""},
		},
		{
			name: "单个 - 各种带空格或不带空格的乱七八槽",
			rx:   rxAnnotation,
			arg:  "//    #[Enum]   ",
			want: []string{"//    #[Enum]   ", "Enum", "", "", ""},
		},
		{
			name: "单个 - 尾部带乱七八槽",
			rx:   rxAnnotation,
			arg:  "// #[Enum]dfadfadf",
			want: []string{"// #[Enum]dfadfadf", "Enum", "", "", ""},
		},
		{
			name: "复合 - 1",
			rx:   rxAnnotation,
			arg:  `// #[Enum(mapping="aaa")]`,
			want: []string{`// #[Enum(mapping="aaa")]`, "Enum", `(mapping="aaa")`, "mapping", "aaa"},
		},
		{
			name: "复合 - 2",
			rx:   rxAnnotation,
			arg:  `// #[Enum( mapping="aaa" )]`,
			want: []string{`// #[Enum( mapping="aaa" )]`, "Enum", `( mapping="aaa" )`, "mapping", "aaa"},
		},
		{
			name: "复合 - 3",
			rx:   rxAnnotation,
			arg:  `// #[Enum( mapping  =  "aaa" )]`,
			want: []string{`// #[Enum( mapping  =  "aaa" )]`, "Enum", `( mapping  =  "aaa" )`, "mapping", "aaa"},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			match := tt.rx.FindStringSubmatch(tt.arg)
			if !reflect.DeepEqual(match, tt.want) {
				t.Errorf("expected want: %v, got %v", tt.want, match)
			}
		})
	}
}

func TestMatchAnnotation(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want *Annotation
	}{
		{
			name: "非法",
			arg:  "// sd #[Enum]",
			want: nil,
		},
		{
			name: "单个 - 标准",
			arg:  "// #[Enum]",
			want: &Annotation{"Enum", "", ""},
		},
		{
			name: "单个 - //后面不带空格",
			arg:  "//#[Enum]",
			want: &Annotation{"Enum", "", ""},
		},
		{
			name: "单个 - //后面带多个空格",
			arg:  "//    #[Enum]",
			want: &Annotation{"Enum", "", ""},
		},
		{
			name: "单个 - 各种带空格或不带空格的乱七八槽",
			arg:  "//    #[Enum]   ",
			want: &Annotation{"Enum", "", ""},
		},
		{
			name: "单个 - 尾部带乱七八槽",
			arg:  "// #[Enum]dfadfadf",
			want: &Annotation{"Enum", "", ""},
		},
		{
			name: "复合 - 1",
			arg:  `// #[Enum(mapping="aaa")]`,
			want: &Annotation{"Enum", "mapping", "aaa"},
		},
		{
			name: "复合 - 2",
			arg:  `// #[Enum( mapping="aaa" )]`,
			want: &Annotation{"Enum", "mapping", "aaa"},
		},
		{
			name: "复合 - 3",
			arg:  `// #[Enum( mapping  =  "aaa" )]`,
			want: &Annotation{"Enum", "mapping", "aaa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchAnnotation(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatchAnnotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
