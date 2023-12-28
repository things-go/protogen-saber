package protoenum

import (
	"github.com/things-go/protogen-saber/internal/annotation"
	"github.com/things-go/protogen-saber/internal/protoutil"
	"google.golang.org/protobuf/compiler/protogen"
)

// annotation const value
const (
	Identifier             = "enum"
	Attribute_Name_Mapping = "mapping"
)

type EnumAnnotation struct {
	Enabled bool
}

func ParseAnnotationEnum(s protogen.Comments) (*EnumAnnotation, protoutil.CommentLines) {
	ret := &EnumAnnotation{Enabled: false}
	annotes, remainComments := protoutil.NewCommentLines(s).FindAnnotations(Identifier)
	ret.Enabled = annotes.ContainHeadless(Identifier)
	return ret, remainComments
}

type EnumValueAnnotation struct {
	Mapping string
}

func ParseAnnotationEnumValue(s protogen.Comments) (*EnumValueAnnotation, protoutil.CommentLines) {
	ret := &EnumValueAnnotation{Mapping: ""}
	annotates, remainComments := protoutil.NewCommentLines(s).FindAnnotations(Identifier)
	values := annotates.FindValue(Identifier, Attribute_Name_Mapping)
	for _, v := range values {
		if v, ok := v.(annotation.String); ok {
			ret.Mapping = v.Value
			break
		}
	}
	return ret, remainComments
}
