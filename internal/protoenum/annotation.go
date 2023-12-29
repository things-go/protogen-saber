package protoenum

import (
	"github.com/things-go/protogen-saber/internal/annotation"
	"github.com/things-go/protogen-saber/internal/protoutil"
	"google.golang.org/protobuf/compiler/protogen"
)

// annotation const value
const (
	Identity               = "enum"
	Attribute_Name_Mapping = "mapping"
)

type EnumAnnotation struct {
	Enabled bool
}

func ParseDeriveEnum(s protogen.Comments) (*EnumAnnotation, protoutil.CommentLines) {
	ret := &EnumAnnotation{Enabled: false}
	derives, remainComments := protoutil.NewCommentLines(s).FindDerives(Identity)
	ret.Enabled = derives.ContainHeadless(Identity)
	return ret, remainComments
}

type EnumValueAnnotation struct {
	Mapping string
}

func ParseDeriveEnumValue(s protogen.Comments) (*EnumValueAnnotation, protoutil.CommentLines) {
	ret := &EnumValueAnnotation{Mapping: ""}
	derives, remainComments := protoutil.NewCommentLines(s).FindDerives(Identity)
	values := derives.FindValue(Identity, Attribute_Name_Mapping)
	for _, v := range values {
		if v, ok := v.(annotation.String); ok {
			ret.Mapping = v.Value
			break
		}
	}
	return ret, remainComments
}
