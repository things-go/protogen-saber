package protoerrno

import (
	"strings"

	"github.com/things-go/protogen-saber/internal/annotation"
	"github.com/things-go/protogen-saber/internal/protoutil"
	"google.golang.org/protobuf/compiler/protogen"
)

// annotation const value
const (
	Identity               = "errno"
	Attribute_Name_Status  = "status"
	Attribute_Name_Code    = "code"
	Attribute_Name_Message = "message"
)

type ErrnoDerive struct {
	Enabled bool
	Status  int
}

func ParseDeriveErrno(s protogen.Comments) (*ErrnoDerive, protoutil.CommentLines) {
	derives, remainComments := protoutil.NewCommentLines(s).FindDerives(Identity)
	ret := &ErrnoDerive{
		Enabled: derives.ContainHeadless(Identity),
		Status:  500,
	}
	values := derives.FindValue(Identity, Attribute_Name_Status)
	for _, value := range values {
		if v, ok := value.(annotation.Integer); ok && v.Value > 0 && v.Value < 1000 {
			ret.Status = int(v.Value)
		}
	}
	return ret, remainComments
}

type ErrnoValueDerive struct {
	Status  int
	Code    int
	Message string
}

func ParseDeriveErrnoValue(status, code int, s protogen.Comments) (*ErrnoValueDerive, protoutil.CommentLines) {
	derives, remainComments := protoutil.NewCommentLines(s).FindDerives(Identity)
	ret := &ErrnoValueDerive{
		Status:  status,
		Code:    code,
		Message: strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(remainComments.LineString()), "\n", ","), `"`, `\"`),
	}
	for _, d := range derives {
		for _, v := range d.Attrs {
			switch v.Name {
			case Attribute_Name_Status:
				if v, ok := v.Value.(annotation.Integer); ok {
					ret.Status = int(v.Value)
				}
			case Attribute_Name_Code:
				if v, ok := v.Value.(annotation.Integer); ok {
					ret.Code = int(v.Value)
				}
			case Attribute_Name_Message:
				if v, ok := v.Value.(annotation.String); ok {
					ret.Message = v.Value
				}

			}
		}
	}
	return ret, remainComments
}
