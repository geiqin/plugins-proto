// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: baseInfoService.proto

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

type SpuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                    //ID
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type"`                                 //商品类型
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`                                 //商品名称
	Code       string `protobuf:"bytes,4,opt,name=code,proto3" json:"code"`                                 //SPU编码
	ImageUrl   string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url"`         //商品主图
	IsVirtual  string `protobuf:"bytes,6,opt,name=is_virtual,json=isVirtual,proto3" json:"is_virtual"`      //虚拟商品
	IsManySpec string `protobuf:"bytes,7,opt,name=is_many_spec,json=isManySpec,proto3" json:"is_many_spec"` //多规格商品
	UnitId     int32  `protobuf:"varint,8,opt,name=unit_id,json=unitId,proto3" json:"unit_id"`              //计量单位ID
	UnitName   string `protobuf:"bytes,9,opt,name=unit_name,json=unitName,proto3" json:"unit_name"`         //计量单位名称
	UnitType   string `protobuf:"bytes,10,opt,name=unit_type,json=unitType,proto3" json:"unit_type"`        //计量单位类型
	Inventory  int64  `protobuf:"varint,11,opt,name=inventory,proto3" json:"inventory"`                     //库存
	IsShelve   string `protobuf:"bytes,12,opt,name=is_shelve,json=isShelve,proto3" json:"is_shelve"`        //是否上架
	MinPrice   int64  `protobuf:"varint,13,opt,name=min_price,json=minPrice,proto3" json:"min_price"`       //最低价格
	MaxPrice   int64  `protobuf:"varint,14,opt,name=max_price,json=maxPrice,proto3" json:"max_price"`       //最高价格
	Status     string `protobuf:"bytes,15,opt,name=status,proto3" json:"status"`                            //商品状态
}

func (x *SpuInfo) Reset() {
	*x = SpuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpuInfo) ProtoMessage() {}

func (x *SpuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpuInfo.ProtoReflect.Descriptor instead.
func (*SpuInfo) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{0}
}

func (x *SpuInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SpuInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SpuInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpuInfo) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SpuInfo) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *SpuInfo) GetIsVirtual() string {
	if x != nil {
		return x.IsVirtual
	}
	return ""
}

func (x *SpuInfo) GetIsManySpec() string {
	if x != nil {
		return x.IsManySpec
	}
	return ""
}

func (x *SpuInfo) GetUnitId() int32 {
	if x != nil {
		return x.UnitId
	}
	return 0
}

func (x *SpuInfo) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *SpuInfo) GetUnitType() string {
	if x != nil {
		return x.UnitType
	}
	return ""
}

func (x *SpuInfo) GetInventory() int64 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *SpuInfo) GetIsShelve() string {
	if x != nil {
		return x.IsShelve
	}
	return ""
}

func (x *SpuInfo) GetMinPrice() int64 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *SpuInfo) GetMaxPrice() int64 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *SpuInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// 优惠劵
type CouponInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                           //ID
	ShopId            int64  `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`                                     //店铺id
	CouponType        string `protobuf:"bytes,3,opt,name=coupon_type,json=couponType,proto3" json:"coupon_type"`                          //优惠券类型
	CouponName        string `protobuf:"bytes,4,opt,name=coupon_name,json=couponName,proto3" json:"coupon_name"`                          //名称
	CouponDesc        string `protobuf:"bytes,5,opt,name=coupon_desc,json=couponDesc,proto3" json:"coupon_desc"`                          //名称备注
	UserType          string `protobuf:"bytes,6,opt,name=user_type,json=userType,proto3" json:"user_type"`                                //用户类型
	FaceValue         int64  `protobuf:"varint,7,opt,name=face_value,json=faceValue,proto3" json:"face_value"`                            //优惠内容
	MaxValue          int64  `protobuf:"varint,8,opt,name=max_value,json=maxValue,proto3" json:"max_value"`                               //最多优惠
	MinValue          int64  `protobuf:"varint,9,opt,name=min_value,json=minValue,proto3" json:"min_value"`                               //最少优惠
	ExpireType        string `protobuf:"bytes,13,opt,name=expire_type,json=expireType,proto3" json:"expire_type"`                         //到期类型
	ExpireDays        int32  `protobuf:"varint,14,opt,name=expire_days,json=expireDays,proto3" json:"expire_days"`                        //到期天数
	ExpireStartTime   string `protobuf:"bytes,15,opt,name=expire_start_time,json=expireStartTime,proto3" json:"expire_start_time"`        //生效时间
	ExpireEndTime     string `protobuf:"bytes,16,opt,name=expire_end_time,json=expireEndTime,proto3" json:"expire_end_time"`              //到期时间
	UsageThreshold    string `protobuf:"bytes,17,opt,name=usage_threshold,json=usageThreshold,proto3" json:"usage_threshold"`             //使用门槛
	ThresholdAmount   int64  `protobuf:"varint,18,opt,name=threshold_amount,json=thresholdAmount,proto3" json:"threshold_amount"`         //门槛金额
	LimitReceiveType  string `protobuf:"bytes,19,opt,name=limit_receive_type,json=limitReceiveType,proto3" json:"limit_receive_type"`     //每人限领类型
	LimitReceiveCount int32  `protobuf:"varint,20,opt,name=limit_receive_count,json=limitReceiveCount,proto3" json:"limit_receive_count"` //每人限领次数
	IssueCount        int32  `protobuf:"varint,21,opt,name=issue_count,json=issueCount,proto3" json:"issue_count"`                        //发放总量
	ReceivedCount     int32  `protobuf:"varint,22,opt,name=received_count,json=receivedCount,proto3" json:"received_count"`               //领取数量
	SpuScopeType      string `protobuf:"bytes,23,opt,name=spu_scope_type,json=spuScopeType,proto3" json:"spu_scope_type"`                 //适用商品类型
	IsPublicReceive   string `protobuf:"bytes,24,opt,name=is_public_receive,json=isPublicReceive,proto3" json:"is_public_receive"`        //是否公开领取
	Status            string `protobuf:"bytes,28,opt,name=status,proto3" json:"status"`                                                   //启用状态
}

func (x *CouponInfo) Reset() {
	*x = CouponInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_baseInfoService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponInfo) ProtoMessage() {}

func (x *CouponInfo) ProtoReflect() protoreflect.Message {
	mi := &file_baseInfoService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponInfo.ProtoReflect.Descriptor instead.
func (*CouponInfo) Descriptor() ([]byte, []int) {
	return file_baseInfoService_proto_rawDescGZIP(), []int{1}
}

func (x *CouponInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CouponInfo) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *CouponInfo) GetCouponType() string {
	if x != nil {
		return x.CouponType
	}
	return ""
}

func (x *CouponInfo) GetCouponName() string {
	if x != nil {
		return x.CouponName
	}
	return ""
}

func (x *CouponInfo) GetCouponDesc() string {
	if x != nil {
		return x.CouponDesc
	}
	return ""
}

func (x *CouponInfo) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

func (x *CouponInfo) GetFaceValue() int64 {
	if x != nil {
		return x.FaceValue
	}
	return 0
}

func (x *CouponInfo) GetMaxValue() int64 {
	if x != nil {
		return x.MaxValue
	}
	return 0
}

func (x *CouponInfo) GetMinValue() int64 {
	if x != nil {
		return x.MinValue
	}
	return 0
}

func (x *CouponInfo) GetExpireType() string {
	if x != nil {
		return x.ExpireType
	}
	return ""
}

func (x *CouponInfo) GetExpireDays() int32 {
	if x != nil {
		return x.ExpireDays
	}
	return 0
}

func (x *CouponInfo) GetExpireStartTime() string {
	if x != nil {
		return x.ExpireStartTime
	}
	return ""
}

func (x *CouponInfo) GetExpireEndTime() string {
	if x != nil {
		return x.ExpireEndTime
	}
	return ""
}

func (x *CouponInfo) GetUsageThreshold() string {
	if x != nil {
		return x.UsageThreshold
	}
	return ""
}

func (x *CouponInfo) GetThresholdAmount() int64 {
	if x != nil {
		return x.ThresholdAmount
	}
	return 0
}

func (x *CouponInfo) GetLimitReceiveType() string {
	if x != nil {
		return x.LimitReceiveType
	}
	return ""
}

func (x *CouponInfo) GetLimitReceiveCount() int32 {
	if x != nil {
		return x.LimitReceiveCount
	}
	return 0
}

func (x *CouponInfo) GetIssueCount() int32 {
	if x != nil {
		return x.IssueCount
	}
	return 0
}

func (x *CouponInfo) GetReceivedCount() int32 {
	if x != nil {
		return x.ReceivedCount
	}
	return 0
}

func (x *CouponInfo) GetSpuScopeType() string {
	if x != nil {
		return x.SpuScopeType
	}
	return ""
}

func (x *CouponInfo) GetIsPublicReceive() string {
	if x != nil {
		return x.IsPublicReceive
	}
	return ""
}

func (x *CouponInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_baseInfoService_proto protoreflect.FileDescriptor

var file_baseInfoService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x93, 0x03, 0x0a, 0x07, 0x53, 0x70, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x56, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x79,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x73, 0x4d,
	0x61, 0x6e, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73,
	0x68, 0x65, 0x6c, 0x76, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x53,
	0x68, 0x65, 0x6c, 0x76, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x88, 0x06, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x44, 0x65,
	0x73, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x61, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x69, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6d, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x79, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x70, 0x75, 0x5f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x70, 0x75, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x69, 0x73, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_baseInfoService_proto_rawDescOnce sync.Once
	file_baseInfoService_proto_rawDescData = file_baseInfoService_proto_rawDesc
)

func file_baseInfoService_proto_rawDescGZIP() []byte {
	file_baseInfoService_proto_rawDescOnce.Do(func() {
		file_baseInfoService_proto_rawDescData = protoimpl.X.CompressGZIP(file_baseInfoService_proto_rawDescData)
	})
	return file_baseInfoService_proto_rawDescData
}

var file_baseInfoService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_baseInfoService_proto_goTypes = []interface{}{
	(*SpuInfo)(nil),    // 0: services.SpuInfo
	(*CouponInfo)(nil), // 1: services.CouponInfo
}
var file_baseInfoService_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_baseInfoService_proto_init() }
func file_baseInfoService_proto_init() {
	if File_baseInfoService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_baseInfoService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpuInfo); i {
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
		file_baseInfoService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponInfo); i {
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
			RawDescriptor: file_baseInfoService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_baseInfoService_proto_goTypes,
		DependencyIndexes: file_baseInfoService_proto_depIdxs,
		MessageInfos:      file_baseInfoService_proto_msgTypes,
	}.Build()
	File_baseInfoService_proto = out.File
	file_baseInfoService_proto_rawDesc = nil
	file_baseInfoService_proto_goTypes = nil
	file_baseInfoService_proto_depIdxs = nil
}
