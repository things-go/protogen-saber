// Code generated by protoc-gen-saber-enum. DO NOT EDIT.
// versions:
//   - protoc-gen-saber-enum v0.0.1
//   - protoc             v3.21.2
// source: nested.proto

package enums

// __StatusMapping Status mapping
var __StatusMapping = map[Nested_Status]string{
	0: "未定义",
	1: "nested1",
	2: "nested2",
	3: "nested3",
	4: "nested4",
}

// GetStatusDesc get mapping description
//
//	Status 状态值, [0:未定义,1:nested1,2:nested2,3:nested3,4:nested4]
func GetStatusDesc(t Nested_Status) string {
	return __StatusMapping[t]
}

// __TypeMapping Type mapping
var __TypeMapping = map[Nested_Nested1_Type]string{
	0: "禁用",
	1: "启用",
}

// GetTypeDesc get mapping description
// , [0:禁用,1:启用]
func GetTypeDesc(t Nested_Nested1_Type) string {
	return __TypeMapping[t]
}
