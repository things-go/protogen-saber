// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: protosaber/enumerate/enumerate.proto

package enumerate

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_protosaber_enumerate_enumerate_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         95271001,
		Name:          "things_go.enumerate.enabled",
		Tag:           "varint,95271001,opt,name=enabled",
		Filename:      "protosaber/enumerate/enumerate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         95271101,
		Name:          "things_go.enumerate.mapping",
		Tag:           "bytes,95271101,opt,name=mapping",
		Filename:      "protosaber/enumerate/enumerate.proto",
	},
}

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional bool enabled = 95271001;
	E_Enabled = &file_protosaber_enumerate_enumerate_proto_extTypes[0]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string mapping = 95271101;
	E_Mapping = &file_protosaber_enumerate_enumerate_proto_extTypes[1]
)

var File_protosaber_enumerate_enumerate_proto protoreflect.FileDescriptor

var file_protosaber_enumerate_enumerate_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x61, 0x62, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x67,
	0x6f, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x39, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9, 0xf0, 0xb6, 0x2d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x3a, 0x3e, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xbd, 0xf1, 0xb6, 0x2d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x5e, 0x0a, 0x16, 0x63, 0x6e, 0x2e, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x2d, 0x67, 0x6f, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x65, 0x6e, 0x2d, 0x73, 0x61, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x61, 0x62, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x3b, 0x65,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protosaber_enumerate_enumerate_proto_goTypes = []interface{}{
	(*descriptorpb.EnumOptions)(nil),      // 0: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 1: google.protobuf.EnumValueOptions
}
var file_protosaber_enumerate_enumerate_proto_depIdxs = []int32{
	0, // 0: things_go.enumerate.enabled:extendee -> google.protobuf.EnumOptions
	1, // 1: things_go.enumerate.mapping:extendee -> google.protobuf.EnumValueOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protosaber_enumerate_enumerate_proto_init() }
func file_protosaber_enumerate_enumerate_proto_init() {
	if File_protosaber_enumerate_enumerate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protosaber_enumerate_enumerate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_protosaber_enumerate_enumerate_proto_goTypes,
		DependencyIndexes: file_protosaber_enumerate_enumerate_proto_depIdxs,
		ExtensionInfos:    file_protosaber_enumerate_enumerate_proto_extTypes,
	}.Build()
	File_protosaber_enumerate_enumerate_proto = out.File
	file_protosaber_enumerate_enumerate_proto_rawDesc = nil
	file_protosaber_enumerate_enumerate_proto_goTypes = nil
	file_protosaber_enumerate_enumerate_proto_depIdxs = nil
}
