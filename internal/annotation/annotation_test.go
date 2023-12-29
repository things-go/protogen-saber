package annotation

import (
	"reflect"
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *Derive
		wantErr bool
	}{
		{
			name:  "仅identity",
			input: "#[ident]",
			want: &Derive{
				Identity: "ident",
				Attrs:    nil,
				// Empty:      struct{}{},
			},
			wantErr: false,
		},
		{
			name:  "单个键值对 -> 值-字符串",
			input: `#[ident(k1="v1")]`,
			want: &Derive{
				Identity: "ident",
				Attrs: []*NameValue{
					{
						Name:  "k1",
						Value: String{Value: "v1"},
					},
				},
				// Empty:   struct{}{},
			},
			wantErr: false,
		},
		{
			name:  "单个键值对 -> 值-字符串数组",
			input: `#[ident(k1=["1", "2","3"])]`,
			want: &Derive{
				Identity: "ident",
				Attrs: []*NameValue{
					{
						Name: "k1",
						Value: StringList{
							Value: []string{"1", "2", "3"},
						},
					},
				},
				// Empty:   struct{}{},
			},
			wantErr: false,
		},
		{
			name:  "单个键值对 -> 值-整数",
			input: `#[ident(k1=1)]`,
			want: &Derive{
				Identity: "ident",
				Attrs: []*NameValue{
					{
						Name:  "k1",
						Value: Integer{Value: 1},
					},
				},
				// Empty:   struct{}{},
			},
			wantErr: false,
		},
		{
			name:  "单个键值对 -> 值-整数数组",
			input: `#[ident(k1=[1, 2, 3])]`,
			want: &Derive{
				Identity: "ident",
				Attrs: []*NameValue{
					{
						Name: "k1",
						Value: IntegerList{
							Value: []int64{1, 2, 3},
						},
					},
				},
				// Empty:   struct{}{},
			},
			wantErr: false,
		},
		{
			name:  "单个键值对 -> 值-浮点数",
			input: `#[ident(k1=1.1)]`,
			want: &Derive{
				Identity: "ident",
				Attrs: []*NameValue{
					{
						Name:  "k1",
						Value: Float{Value: 1.1},
					},
				},
				// Empty:   struct{}{},
			},
			wantErr: false,
		},
		{
			name:  "单个键值对 -> 值-浮点数组",
			input: `#[ident(k1=[1.1, 2, 3])]`,
			want: &Derive{
				Identity: "ident",
				Attrs: []*NameValue{
					{
						Name: "k1",
						Value: FloatList{
							Value: []float64{1.1, 2, 3},
						},
					},
				},
				// Empty:   struct{}{},
			},
			wantErr: false,
		},
		{
			name:  "单个键值对 -> 值-布尔",
			input: `#[ident(k1=true)]`,
			want: &Derive{
				Identity: "ident",
				Attrs: []*NameValue{
					{
						Name:  "k1",
						Value: Bool{Value: true},
					},
				},
				// Empty:   struct{}{},
			},
			wantErr: false,
		},
		{
			name:  "单个键值对 -> 值-布尔数组",
			input: `#[ident(k1=[true, false, false])]`,
			want: &Derive{
				Identity: "ident",
				Attrs: []*NameValue{
					{
						Name: "k1",
						Value: BoolList{
							Value: []Boolean{true, false, false},
						},
					},
				},
				// Empty:   struct{}{},
			},
			wantErr: false,
		},
		{
			name:  "多个键值对",
			input: `#[ident(k1="v1",k2=1,k3=1.1,k4=false,k5=["1","2","3"],k6=[1,2,3],k7=[1.1,2,3],k8=[true, false, false])]`,
			want: &Derive{
				Identity: "ident",
				Attrs: []*NameValue{
					{
						Name:  "k1",
						Value: String{Value: "v1"},
					},
					{
						Name:  "k2",
						Value: Integer{Value: 1},
					},
					{
						Name:  "k3",
						Value: Float{Value: 1.1},
					},
					{
						Name:  "k4",
						Value: Bool{Value: false},
					},
					{
						Name: "k5",
						Value: StringList{
							Value: []string{"1", "2", "3"},
						},
					},
					{
						Name: "k6",
						Value: IntegerList{
							Value: []int64{1, 2, 3},
						},
					},
					{
						Name: "k7",
						Value: FloatList{
							Value: []float64{1.1, 2, 3},
						},
					},
					{
						Name: "k8",
						Value: BoolList{
							Value: []Boolean{true, false, false},
						},
					},
				},
				// Empty:   struct{}{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Match(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
