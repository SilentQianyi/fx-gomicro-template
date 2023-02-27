// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: proto/hello_world/v1/hello_world.proto

package helloWorld

import (
	v1 "chat/api/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloWorldStatus int32

const (
	HelloWorldStatus_HelloWorldStatusNone    HelloWorldStatus = 0
	HelloWorldStatus_HelloWorldStatusDoing   HelloWorldStatus = 1
	HelloWorldStatus_HelloWorldStatusDone    HelloWorldStatus = 2
	HelloWorldStatus_HelloWorldStatusUnknown HelloWorldStatus = 100
)

// Enum value maps for HelloWorldStatus.
var (
	HelloWorldStatus_name = map[int32]string{
		0:   "HelloWorldStatusNone",
		1:   "HelloWorldStatusDoing",
		2:   "HelloWorldStatusDone",
		100: "HelloWorldStatusUnknown",
	}
	HelloWorldStatus_value = map[string]int32{
		"HelloWorldStatusNone":    0,
		"HelloWorldStatusDoing":   1,
		"HelloWorldStatusDone":    2,
		"HelloWorldStatusUnknown": 100,
	}
)

func (x HelloWorldStatus) Enum() *HelloWorldStatus {
	p := new(HelloWorldStatus)
	*p = x
	return p
}

func (x HelloWorldStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HelloWorldStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_hello_world_v1_hello_world_proto_enumTypes[0].Descriptor()
}

func (HelloWorldStatus) Type() protoreflect.EnumType {
	return &file_proto_hello_world_v1_hello_world_proto_enumTypes[0]
}

func (x HelloWorldStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HelloWorldStatus.Descriptor instead.
func (HelloWorldStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_hello_world_v1_hello_world_proto_rawDescGZIP(), []int{0}
}

type HelloWorldModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Common   *v1.HelloWorldCommon `protobuf:"bytes,2,opt,name=common,proto3" json:"common,omitempty"`
	Language string               `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
	Status   HelloWorldStatus     `protobuf:"varint,4,opt,name=status,proto3,enum=helloWorld.v1.HelloWorldStatus" json:"status,omitempty"`
	Deleted  bool                 `protobuf:"varint,5,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *HelloWorldModel) Reset() {
	*x = HelloWorldModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldModel) ProtoMessage() {}

func (x *HelloWorldModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldModel.ProtoReflect.Descriptor instead.
func (*HelloWorldModel) Descriptor() ([]byte, []int) {
	return file_proto_hello_world_v1_hello_world_proto_rawDescGZIP(), []int{0}
}

func (x *HelloWorldModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HelloWorldModel) GetCommon() *v1.HelloWorldCommon {
	if x != nil {
		return x.Common
	}
	return nil
}

func (x *HelloWorldModel) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *HelloWorldModel) GetStatus() HelloWorldStatus {
	if x != nil {
		return x.Status
	}
	return HelloWorldStatus_HelloWorldStatusNone
}

func (x *HelloWorldModel) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type HelloWorldListModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*HelloWorldModel `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *HelloWorldListModel) Reset() {
	*x = HelloWorldListModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWorldListModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWorldListModel) ProtoMessage() {}

func (x *HelloWorldListModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWorldListModel.ProtoReflect.Descriptor instead.
func (*HelloWorldListModel) Descriptor() ([]byte, []int) {
	return file_proto_hello_world_v1_hello_world_proto_rawDescGZIP(), []int{1}
}

func (x *HelloWorldListModel) GetData() []*HelloWorldModel {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetHelloWorldByLanguageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *GetHelloWorldByLanguageRequest) Reset() {
	*x = GetHelloWorldByLanguageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHelloWorldByLanguageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHelloWorldByLanguageRequest) ProtoMessage() {}

func (x *GetHelloWorldByLanguageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHelloWorldByLanguageRequest.ProtoReflect.Descriptor instead.
func (*GetHelloWorldByLanguageRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_world_v1_hello_world_proto_rawDescGZIP(), []int{2}
}

func (x *GetHelloWorldByLanguageRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

var File_proto_hello_world_v1_hello_world_proto protoreflect.FileDescriptor

var file_proto_hello_world_v1_hello_world_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x0f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x49, 0x0a,
	0x13, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x79, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2a, 0x7e, 0x0a, 0x10, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f,
	0x6e, 0x65, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72,
	0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x6f, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12,
	0x18, 0x0a, 0x14, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x64, 0x32, 0x7a, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x12, 0x6c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x79, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x2d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x79, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x42, 0x23, 0x5a, 0x21, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hello_world_v1_hello_world_proto_rawDescOnce sync.Once
	file_proto_hello_world_v1_hello_world_proto_rawDescData = file_proto_hello_world_v1_hello_world_proto_rawDesc
)

func file_proto_hello_world_v1_hello_world_proto_rawDescGZIP() []byte {
	file_proto_hello_world_v1_hello_world_proto_rawDescOnce.Do(func() {
		file_proto_hello_world_v1_hello_world_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hello_world_v1_hello_world_proto_rawDescData)
	})
	return file_proto_hello_world_v1_hello_world_proto_rawDescData
}

var file_proto_hello_world_v1_hello_world_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_hello_world_v1_hello_world_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_hello_world_v1_hello_world_proto_goTypes = []interface{}{
	(HelloWorldStatus)(0),                  // 0: helloWorld.v1.HelloWorldStatus
	(*HelloWorldModel)(nil),                // 1: helloWorld.v1.HelloWorldModel
	(*HelloWorldListModel)(nil),            // 2: helloWorld.v1.HelloWorldListModel
	(*GetHelloWorldByLanguageRequest)(nil), // 3: helloWorld.v1.GetHelloWorldByLanguageRequest
	(*v1.HelloWorldCommon)(nil),            // 4: common.v1.HelloWorldCommon
}
var file_proto_hello_world_v1_hello_world_proto_depIdxs = []int32{
	4, // 0: helloWorld.v1.HelloWorldModel.common:type_name -> common.v1.HelloWorldCommon
	0, // 1: helloWorld.v1.HelloWorldModel.status:type_name -> helloWorld.v1.HelloWorldStatus
	1, // 2: helloWorld.v1.HelloWorldListModel.data:type_name -> helloWorld.v1.HelloWorldModel
	3, // 3: helloWorld.v1.HelloWorld.GetHelloWorldByLanguage:input_type -> helloWorld.v1.GetHelloWorldByLanguageRequest
	2, // 4: helloWorld.v1.HelloWorld.GetHelloWorldByLanguage:output_type -> helloWorld.v1.HelloWorldListModel
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_hello_world_v1_hello_world_proto_init() }
func file_proto_hello_world_v1_hello_world_proto_init() {
	if File_proto_hello_world_v1_hello_world_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hello_world_v1_hello_world_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldModel); i {
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
		file_proto_hello_world_v1_hello_world_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWorldListModel); i {
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
		file_proto_hello_world_v1_hello_world_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHelloWorldByLanguageRequest); i {
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
			RawDescriptor: file_proto_hello_world_v1_hello_world_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_world_v1_hello_world_proto_goTypes,
		DependencyIndexes: file_proto_hello_world_v1_hello_world_proto_depIdxs,
		EnumInfos:         file_proto_hello_world_v1_hello_world_proto_enumTypes,
		MessageInfos:      file_proto_hello_world_v1_hello_world_proto_msgTypes,
	}.Build()
	File_proto_hello_world_v1_hello_world_proto = out.File
	file_proto_hello_world_v1_hello_world_proto_rawDesc = nil
	file_proto_hello_world_v1_hello_world_proto_goTypes = nil
	file_proto_hello_world_v1_hello_world_proto_depIdxs = nil
}
