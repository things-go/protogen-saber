package protoenum

import (
	"strings"

	"github.com/things-go/protogen-saber/internal/infra"
	"github.com/things-go/protogen-saber/protosaber/enumerate"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// EnumValue 枚举的枚举项
type EnumValue struct {
	Number      int    // 编号
	Value       string // 值,例: Status_Enabled
	CamelValue  string // 驼峰值,例: StatusEnabled
	TrimValue   string // 值截断EnumName前缀,例: Enabled(EnumName=Status)
	Mapping     string // 映射值
	Comment     string // 注释
	IsDuplicate bool   // 是否是副本
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
		isEnabled := proto.GetExtension(pe.Desc.Options(), enumerate.E_Enabled)
		ok := isEnabled.(bool)
		if !ok {
			continue
		}
		enumName := string(pe.Desc.Name())

		eValueMp := make(map[int]string, len(pe.Values))
		eValues := make([]*EnumValue, 0, len(pe.Values))
		for _, v := range pe.Values {
			mpv := proto.GetExtension(v.Desc.Options(), enumerate.E_Mapping)
			mappingValue, _ := mpv.(string)
			comment := strings.TrimSpace(strings.TrimSuffix(string(v.Comments.Leading), "\n"))
			if mappingValue == "" {
				mappingValue = comment
			}
			comment = strings.ReplaceAll(strings.ReplaceAll(comment, "\n", ","), `"`, `\"`)
			mappingValue = strings.ReplaceAll(strings.ReplaceAll(mappingValue, "\n", ","), `"`, `\"`)

			enumValueName := string(v.Desc.Name())
			ev := &EnumValue{
				Number:     int(v.Desc.Number()),
				Value:      enumValueName,
				CamelValue: infra.CamelCase(enumValueName),
				TrimValue:  strings.TrimPrefix(strings.TrimPrefix(enumValueName, enumName), "_"),
				Mapping:    mappingValue,
				Comment:    comment,
			}
			//* duplicate
			if _, ev.IsDuplicate = eValueMp[ev.Number]; !ev.IsDuplicate {
				eValueMp[ev.Number] = mappingValue
			}
			eValues = append(eValues, ev)
		}

		comment := strings.TrimSpace(strings.ReplaceAll(string(pe.Comments.Leading), "\n", ""))
		bb := infra.ToArrayString(eValueMp)
		if comment == "" {
			comment = bb
		} else {
			comment += ", " + bb
		}
		enums = append(enums, &Enum{
			MessageName: nestedMessageName,
			Name:        enumName,
			Comment:     comment,
			Values:      eValues,
		})
	}
	return enums
}

// IntoEnumComment generates enum comment if it exists
func IntoEnumComment(pe *protogen.Enum) string {
	if pe == nil || len(pe.Values) == 0 {
		return ""
	}
	isEnabled := proto.GetExtension(pe.Desc.Options(), enumerate.E_Enabled)
	ok := isEnabled.(bool)
	if !ok {
		return ""
	}

	eValueMp := make(map[int]string, len(pe.Values))
	for _, v := range pe.Values {
		mpv := proto.GetExtension(v.Desc.Options(), enumerate.E_Mapping)
		mappingValue, _ := mpv.(string)
		comment := strings.TrimSpace(strings.TrimSuffix(string(v.Comments.Leading), "\n"))
		if mappingValue == "" {
			mappingValue = comment
		}
		mappingValue = strings.ReplaceAll(strings.ReplaceAll(mappingValue, "\n", ","), `"`, `\"`)
		eValueMp[int(v.Desc.Number())] = mappingValue
	}
	return infra.ToArrayString(eValueMp)
}
