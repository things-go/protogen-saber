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

type EnumDerive struct {
	Enabled bool
}

func ParseDeriveEnum(s protogen.Comments) (*EnumDerive, protoutil.CommentLines) {
	ret := &EnumDerive{Enabled: false}
	derives, remainComments := protoutil.NewCommentLines(s).FindDerives(Identity)
	ret.Enabled = derives.ContainHeadless(Identity)
	return ret, remainComments
}

type EnumValueDerive struct {
	Mapping string
}

func ParseDeriveEnumValue(s protogen.Comments) (*EnumValueDerive, protoutil.CommentLines) {
	ret := &EnumValueDerive{Mapping: ""}
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
