// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: configService.proto

package services

import (
	common "github.com/geiqin/micro-kit/protobuf/common"
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

//钱包配置
type ExpressDeliveryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsEnableDelivery string `protobuf:"bytes,1,opt,name=is_enable_delivery,json=isEnableDelivery,proto3" json:"is_enable_delivery"` //是否开启快递发货(0否、1是)
	ChargingType     string `protobuf:"bytes,2,opt,name=charging_type,json=chargingType,proto3" json:"charging_type"`               //计费方式:1-累加运费,2-组合运费
	IsInitData       string `protobuf:"bytes,3,opt,name=is_init_data,json=isInitData,proto3" json:"is_init_data"`                   //是否已初始化数据(0否、1是)
}

func (x *ExpressDeliveryConfig) Reset() {
	*x = ExpressDeliveryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_configService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressDeliveryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressDeliveryConfig) ProtoMessage() {}

func (x *ExpressDeliveryConfig) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ExpressDeliveryConfig.ProtoReflect.Descriptor instead.
func (*ExpressDeliveryConfig) Descriptor() ([]byte, []int) {
	return file_configService_proto_rawDescGZIP(), []int{0}
}

func (x *ExpressDeliveryConfig) GetIsEnableDelivery() string {
	if x != nil {
		return x.IsEnableDelivery
	}
	return ""
}

func (x *ExpressDeliveryConfig) GetChargingType() string {
	if x != nil {
		return x.ChargingType
	}
	return ""
}

func (x *ExpressDeliveryConfig) GetIsInitData() string {
	if x != nil {
		return x.IsInitData
	}
	return ""
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *ExpressDeliveryConfig `protobuf:"bytes,1,opt,name=data,proto3" json:"data"`
	Error *common.Error          `protobuf:"bytes,2,opt,name=error,proto3" json:"error"`
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

func (x *ConfigResponse) GetData() *ExpressDeliveryConfig {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ConfigResponse) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_configService_proto protoreflect.FileDescriptor

var file_configService_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x15, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2c, 0x0a, 0x12,
	0x69, 0x73, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x73, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x68,
	0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x22, 0x6a, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xad, 0x02,
	0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x44, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x18, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x53, 0x61, 0x76,
	0x65, 0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x12, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a,
	0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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
	(*ExpressDeliveryConfig)(nil), // 0: services.ExpressDeliveryConfig
	(*ConfigResponse)(nil),        // 1: services.ConfigResponse
	(*common.Error)(nil),          // 2: common.Error
}
var file_configService_proto_depIdxs = []int32{
	0, // 0: services.ConfigResponse.data:type_name -> services.ExpressDeliveryConfig
	2, // 1: services.ConfigResponse.error:type_name -> common.Error
	0, // 2: services.ConfigService.Index:input_type -> services.ExpressDeliveryConfig
	0, // 3: services.ConfigService.Get:input_type -> services.ExpressDeliveryConfig
	0, // 4: services.ConfigService.Save:input_type -> services.ExpressDeliveryConfig
	0, // 5: services.ConfigService.DeliverySwitch:input_type -> services.ExpressDeliveryConfig
	1, // 6: services.ConfigService.Index:output_type -> services.ConfigResponse
	1, // 7: services.ConfigService.Get:output_type -> services.ConfigResponse
	1, // 8: services.ConfigService.Save:output_type -> services.ConfigResponse
	1, // 9: services.ConfigService.DeliverySwitch:output_type -> services.ConfigResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_configService_proto_init() }
func file_configService_proto_init() {
	if File_configService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_configService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressDeliveryConfig); i {
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