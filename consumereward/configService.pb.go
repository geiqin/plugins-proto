// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: configService.proto

package services

import (
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

// 消费有礼配置
type ConsumeRewardConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsEnable string `protobuf:"bytes,1,opt,name=is_enable,json=isEnable,proto3" json:"is_enable"` //是否开启消费有礼(0否、1是)
}

func (x *ConsumeRewardConfig) Reset() {
	*x = ConsumeRewardConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeRewardConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeRewardConfig) ProtoMessage() {}

func (x *ConsumeRewardConfig) ProtoReflect() protoreflect.Message {
	mi := &file_configService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeRewardConfig.ProtoReflect.Descriptor instead.
func (*ConsumeRewardConfig) Descriptor() ([]byte, []int) {
	return file_configService_proto_rawDescGZIP(), []int{0}
}

func (x *ConsumeRewardConfig) GetIsEnable() string {
	if x != nil {
		return x.IsEnable
	}
	return ""
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *ConsumeRewardConfig `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Info string               `protobuf:"bytes,2,opt,name=info,proto3" json:"info"`
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_configService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_configService_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigResponse) GetData() *ConsumeRewardConfig {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ConfigResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_configService_proto protoreflect.FileDescriptor

var file_configService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22,
	0x32, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x22, 0x57, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xd9, 0x01, 0x0a,
	0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x18, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_configService_proto_rawDescOnce sync.Once
	file_configService_proto_rawDescData = file_configService_proto_rawDesc
)

func file_configService_proto_rawDescGZIP() []byte {
	file_configService_proto_rawDescOnce.Do(func() {
		file_configService_proto_rawDescData = protoimpl.X.CompressGZIP(file_configService_proto_rawDescData)
	})
	return file_configService_proto_rawDescData
}

var file_configService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_configService_proto_goTypes = []interface{}{
	(*ConsumeRewardConfig)(nil), // 0: services.ConsumeRewardConfig
	(*ConfigResponse)(nil),      // 1: services.ConfigResponse
}
var file_configService_proto_depIdxs = []int32{
	0, // 0: services.ConfigResponse.data:type_name -> services.ConsumeRewardConfig
	0, // 1: services.ConfigService.Get:input_type -> services.ConsumeRewardConfig
	0, // 2: services.ConfigService.Save:input_type -> services.ConsumeRewardConfig
	0, // 3: services.ConfigService.Switch:input_type -> services.ConsumeRewardConfig
	1, // 4: services.ConfigService.Get:output_type -> services.ConfigResponse
	1, // 5: services.ConfigService.Save:output_type -> services.ConfigResponse
	1, // 6: services.ConfigService.Switch:output_type -> services.ConfigResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_configService_proto_init() }
func file_configService_proto_init() {
	if File_configService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_configService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeRewardConfig); i {
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
		file_configService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
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
			RawDescriptor: file_configService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_configService_proto_goTypes,
		DependencyIndexes: file_configService_proto_depIdxs,
		MessageInfos:      file_configService_proto_msgTypes,
	}.Build()
	File_configService_proto = out.File
	file_configService_proto_rawDesc = nil
	file_configService_proto_goTypes = nil
	file_configService_proto_depIdxs = nil
}