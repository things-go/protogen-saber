package protoenum

import (
	"strings"

	"github.com/things-go/protogen-saber/internal/infra"
	"github.com/things-go/protogen-saber/internal/protoutil"
	"google.golang.org/protobuf/compiler/protogen"
)

// annotation const value
const (
	annotation_Path        = "enum"
	annotation_Key_Mapping = "mapping"
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
		annotates, remainComments := protoutil.NewCommentLines(pe.Comments.Leading).
			FindAnnotation(annotation_Path)
		if len(annotates) == 0 {
			continue
		}

		emName := string(pe.Desc.Name())
		emValueMp := make(map[int]string, len(pe.Values))
		emValues := make([]*EnumValue, 0, len(pe.Values))
		for _, v := range pe.Values {
			mappingValue := ""
			comment := strings.TrimSpace(strings.TrimSuffix(string(v.Comments.Leading), "\n"))
			// 先判断注解, 再判断扩展
			annotateValues, _ := protoutil.NewCommentLines(v.Comments.Leading).
				FindAnnotationValues(annotation_Path, annotation_Key_Mapping)
			if len(annotateValues) > 0 && annotateValues[0] != "" {
				mappingValue = annotateValues[0]
			} else {
				mappingValue = comment
			}

			comment = strings.ReplaceAll(strings.ReplaceAll(comment, "\n", ","), `"`, `\"`)
			mappingValue = strings.ReplaceAll(strings.ReplaceAll(mappingValue, "\n", ","), `"`, `\"`)

			enumValueName := string(v.Desc.Name())
			ev := &EnumValue{
				Number:     int(v.Desc.Number()),
				Value:      enumValueName,
				CamelValue: infra.CamelCase(enumValueName),
				TrimValue:  strings.TrimPrefix(strings.TrimPrefix(enumValueName, emName), "_"),
				Mapping:    mappingValue,
				Comment:    comment,
			}
			//* duplicate
			if _, ev.IsDuplicate = emValueMp[ev.Number]; !ev.IsDuplicate {
				emValueMp[ev.Number] = mappingValue
			}
			emValues = append(emValues, ev)
		}

		comment := remainComments.Append(protoutil.ToArrayString(emValueMp)).String()
		enums = append(enums, &Enum{
			MessageName: nestedMessageName,
			Name:        emName,
			Comment:     comment,
			Values:      emValues,
		})
	}
	return enums
}

// IntoEnumComment generates enum comment if it exists
func IntoEnumComment(pe *protogen.Enum) string {
	if pe == nil || len(pe.Values) == 0 {
		return ""
	}
	annotate, _ := protoutil.NewCommentLines(pe.Comments.Leading).FindAnnotation(annotation_Path)
	if len(annotate) == 0 {
		return ""
	}

	emValueMp := make(map[int]string, len(pe.Values))
	for _, v := range pe.Values {
		mappingValue := ""
		annotateVal, _ := protoutil.NewCommentLines(v.Comments.Leading).FindAnnotationValues(annotation_Path, annotation_Key_Mapping)
		if len(annotateVal) > 0 && annotateVal[0] != "" {
			mappingValue = annotateVal[0]
		} else {
			mappingValue = strings.TrimSpace(strings.TrimSuffix(string(v.Comments.Leading), "\n"))
		}
		mappingValue = strings.ReplaceAll(strings.ReplaceAll(mappingValue, "\n", ","), `"`, `\"`)
		emValueMp[int(v.Desc.Number())] = mappingValue
	}
	return protoutil.ToArrayString(emValueMp)
}
