// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.2
// source: protosaber/seaql/seaql.proto

package seaql

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 表名
	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// 引擎
	Engine string `protobuf:"bytes,2,opt,name=engine,proto3" json:"engine,omitempty"`
	// 字符集
	Charset string `protobuf:"bytes,3,opt,name=charset,proto3" json:"charset,omitempty"`
	// 排序规则
	Collate string `protobuf:"bytes,4,opt,name=collate,proto3" json:"collate,omitempty"`
	// 索引
	Index []string `protobuf:"bytes,10,rep,name=index,proto3" json:"index,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protosaber_seaql_seaql_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_protosaber_seaql_seaql_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_protosaber_seaql_seaql_proto_rawDescGZIP(), []int{0}
}

func (x *Options) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *Options) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *Options) GetCharset() string {
	if x != nil {
		return x.Charset
	}
	return ""
}

func (x *Options) GetCollate() string {
	if x != nil {
		return x.Collate
	}
	return ""
}

func (x *Options) GetIndex() []string {
	if x != nil {
		return x.Index
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protosaber_seaql_seaql_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_protosaber_seaql_seaql_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_protosaber_seaql_seaql_proto_rawDescGZIP(), []int{1}
}

func (x *Field) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var file_protosaber_seaql_seaql_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*Options)(nil),
		Field:         95272001,
		Name:          "things_go.seaql.options",
		Tag:           "bytes,95272001,opt,name=options",
		Filename:      "protosaber/seaql/seaql.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*Field)(nil),
		Field:         95272101,
		Name:          "things_go.seaql.field",
		Tag:           "bytes,95272101,opt,name=field",
		Filename:      "protosaber/seaql/seaql.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional things_go.seaql.Options options = 95272001;
	E_Options = &file_protosaber_seaql_seaql_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional things_go.seaql.Field field = 95272101;
	E_Field = &file_protosaber_seaql_seaql_proto_extTypes[1]
)

var File_protosaber_seaql_seaql_proto protoreflect.FileDescriptor

var file_protosaber_seaql_seaql_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x61, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x61,
	0x71, 0x6c, 0x2f, 0x73, 0x65, 0x61, 0x71, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x67, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x71, 0x6c, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x1b,
	0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x56, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc1, 0xf8, 0xb6, 0x2d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x67, 0x6f, 0x2e, 0x73, 0x65, 0x61,
	0x71, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x3a, 0x4e, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa5, 0xf9, 0xb6, 0x2d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x5f, 0x67, 0x6f,
	0x2e, 0x73, 0x65, 0x61, 0x71, 0x6c, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x42, 0x52, 0x0a, 0x12, 0x63, 0x6e, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x2d, 0x67, 0x6f, 0x2e, 0x73, 0x65, 0x61, 0x71, 0x6c, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2d, 0x67,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x61, 0x62, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x61, 0x62, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x61, 0x71,
	0x6c, 0x3b, 0x73, 0x65, 0x61, 0x71, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protosaber_seaql_seaql_proto_rawDescOnce sync.Once
	file_protosaber_seaql_seaql_proto_rawDescData = file_protosaber_seaql_seaql_proto_rawDesc
)

func file_protosaber_seaql_seaql_proto_rawDescGZIP() []byte {
	file_protosaber_seaql_seaql_proto_rawDescOnce.Do(func() {
		file_protosaber_seaql_seaql_proto_rawDescData = protoimpl.X.CompressGZIP(file_protosaber_seaql_seaql_proto_rawDescData)
	})
	return file_protosaber_seaql_seaql_proto_rawDescData
}

var file_protosaber_seaql_seaql_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protosaber_seaql_seaql_proto_goTypes = []interface{}{
	(*Options)(nil),                     // 0: things_go.seaql.Options
	(*Field)(nil),                       // 1: things_go.seaql.Field
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 3: google.protobuf.FieldOptions
}
var file_protosaber_seaql_seaql_proto_depIdxs = []int32{
	2, // 0: things_go.seaql.options:extendee -> google.protobuf.MessageOptions
	3, // 1: things_go.seaql.field:extendee -> google.protobuf.FieldOptions
	0, // 2: things_go.seaql.options:type_name -> things_go.seaql.Options
	1, // 3: things_go.seaql.field:type_name -> things_go.seaql.Field
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protosaber_seaql_seaql_proto_init() }
func file_protosaber_seaql_seaql_proto_init() {
	if File_protosaber_seaql_seaql_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protosaber_seaql_seaql_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protosaber_seaql_seaql_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protosaber_seaql_seaql_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_protosaber_seaql_seaql_proto_goTypes,
		DependencyIndexes: file_protosaber_seaql_seaql_proto_depIdxs,
		MessageInfos:      file_protosaber_seaql_seaql_proto_msgTypes,
		ExtensionInfos:    file_protosaber_seaql_seaql_proto_extTypes,
	}.Build()
	File_protosaber_seaql_seaql_proto = out.File
	file_protosaber_seaql_seaql_proto_rawDesc = nil
	file_protosaber_seaql_seaql_proto_goTypes = nil
	file_protosaber_seaql_seaql_proto_depIdxs = nil
}