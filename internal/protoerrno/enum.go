package protoerrno

import (
	"strings"

	"github.com/things-go/protogen-saber/internal/infra"
	"google.golang.org/protobuf/compiler/protogen"
)

// EnumValue 枚举的枚举项
type EnumValue struct {
	Number      int    // 编号
	Value       string // 值,例: Status_Enabled
	CamelValue  string // 驼峰值,例: StatusEnabled
	Comment     string // 注释
	IsDuplicate bool   // 是否是副本
	Status      int    // 状态码
	Code        int    // 错误码
	Message     string // 错误信息
}

// Enum 枚举
// NOTE:
//
//	如果 MessageName 为空, 表明枚举独立, 枚举类型为 ${{Name}}, 枚举值为 ${{Name}}_${{Value}}
//	如果 MessageName 为不为空, 表明枚举嵌套在message里, 枚举类型为 ${{MessageName}}_{{Name}}, 枚举值为 ${{MessageName}}_${{Value}}
type Enum struct {
	MessageName string       // 嵌套消息名
	Name        string       // 名称
	Comment     string       // 注释
	Values      []*EnumValue // 枚举项
}

// IntoEnumsFromMessage generates the errors definitions, excluding the package statement.
func IntoEnumsFromMessage(nestedMessageName string, protoMessages []*protogen.Message) []*Enum {
	enums := make([]*Enum, 0, 128)
	for _, pm := range protoMessages {
		tmpNestedMessageName := string(pm.Desc.Name())
		if nestedMessageName != "" {
			tmpNestedMessageName = nestedMessageName + "_" + tmpNestedMessageName
		}
		enums = append(enums, IntoEnums(tmpNestedMessageName, pm.Enums)...)
		enums = append(enums, IntoEnumsFromMessage(tmpNestedMessageName, pm.Messages)...)
	}
	return enums
}

// IntoEnums generates the errors definitions, excluding the package statement.
func IntoEnums(nestedMessageName string, protoEnums []*protogen.Enum) []*Enum {
	enums := make([]*Enum, 0, len(protoEnums))
	for _, pe := range protoEnums {
		if len(pe.Values) == 0 {
			continue
		}
		enumAnnotate, remainComments := ParseDeriveErrno(pe.Comments.Leading)
		if !enumAnnotate.Enabled {
			continue
		}

		emName := string(pe.Desc.Name())
		emValueMp := make(map[int]string, len(pe.Values))
		emValues := make([]*EnumValue, 0, len(pe.Values))
		for _, v := range pe.Values {
			annotateEnumValue, remainComments := ParseDeriveErrnoValue(enumAnnotate.Status, int(v.Desc.Number()), v.Comments.Leading)
			enumValueName := string(v.Desc.Name())
			ev := &EnumValue{
				Number:      int(v.Desc.Number()),
				Value:       enumValueName,
				CamelValue:  infra.CamelCase(enumValueName),
				Comment:     strings.ReplaceAll(strings.ReplaceAll(remainComments.LineString(), "\n", ","), `"`, `\"`),
				IsDuplicate: false,
				Status:      annotateEnumValue.Status,
				Code:        annotateEnumValue.Code,
				Message:     annotateEnumValue.Message,
			}
			//* duplicate
			_, ev.IsDuplicate = emValueMp[ev.Number]

			emValues = append(emValues, ev)
		}
		enums = append(enums, &Enum{
			MessageName: nestedMessageName,
			Name:        emName,
			Comment:     remainComments.String(),
			Values:      emValues,
		})
	}
	return enums
}
