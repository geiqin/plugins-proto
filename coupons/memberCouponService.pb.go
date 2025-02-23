// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: memberCouponService.proto

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

type MemberCoupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                                //
	Code         string      `protobuf:"bytes,2,opt,name=code,proto3" json:"code"`                                             //凭证编号
	CouponId     int64       `protobuf:"varint,3,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id"`                    //优惠券id
	MemberId     int64       `protobuf:"varint,4,opt,name=member_id,json=memberId,proto3" json:"member_id"`                    //用户id
	StartTime    string      `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time"`                  //开始时间
	EndTime      string      `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time"`                        //结束时间
	ReceiveScene string      `protobuf:"bytes,7,opt,name=receive_scene,json=receiveScene,proto3" json:"receive_scene"`         //领取场景
	UsedOrderId  int64       `protobuf:"varint,8,opt,name=used_order_id,json=usedOrderId,proto3" json:"used_order_id"`         //使用订单id
	UsedAmount   int64       `protobuf:"varint,9,opt,name=used_amount,json=usedAmount,proto3" json:"used_amount"`              //使用金额
	UsedTime     string      `protobuf:"bytes,10,opt,name=used_time,json=usedTime,proto3" json:"used_time"`                    //使用时间
	Status       string      `protobuf:"bytes,11,opt,name=status,proto3" json:"status"`                                        //优惠劵状态
	CreatedAt    string      `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at"`                 //创建时间
	UpdatedAt    string      `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`                 //修改时间
	Coupon       *Coupon     `protobuf:"bytes,14,opt,name=coupon,proto3" json:"coupon"`                                        //优惠券
	Member       *MemberInfo `protobuf:"bytes,15,opt,name=member,proto3" json:"member"`                                        //用户
	UseOrder     *OrderInfo  `protobuf:"bytes,16,opt,name=use_order,json=useOrder,proto3" json:"use_order"`                    //使用关联的订单
	BuyGoodsIds  []int64     `protobuf:"varint,17,rep,packed,name=buy_goods_ids,json=buyGoodsIds,proto3" json:"buy_goods_ids"` //附加：购买商品ID
}

func (x *MemberCoupon) Reset() {
	*x = MemberCoupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberCouponService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberCoupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCoupon) ProtoMessage() {}

func (x *MemberCoupon) ProtoReflect() protoreflect.Message {
	mi := &file_memberCouponService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCoupon.ProtoReflect.Descriptor instead.
func (*MemberCoupon) Descriptor() ([]byte, []int) {
	return file_memberCouponService_proto_rawDescGZIP(), []int{0}
}

func (x *MemberCoupon) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberCoupon) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MemberCoupon) GetCouponId() int64 {
	if x != nil {
		return x.CouponId
	}
	return 0
}

func (x *MemberCoupon) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberCoupon) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *MemberCoupon) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *MemberCoupon) GetReceiveScene() string {
	if x != nil {
		return x.ReceiveScene
	}
	return ""
}

func (x *MemberCoupon) GetUsedOrderId() int64 {
	if x != nil {
		return x.UsedOrderId
	}
	return 0
}

func (x *MemberCoupon) GetUsedAmount() int64 {
	if x != nil {
		return x.UsedAmount
	}
	return 0
}

func (x *MemberCoupon) GetUsedTime() string {
	if x != nil {
		return x.UsedTime
	}
	return ""
}

func (x *MemberCoupon) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *MemberCoupon) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *MemberCoupon) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *MemberCoupon) GetCoupon() *Coupon {
	if x != nil {
		return x.Coupon
	}
	return nil
}

func (x *MemberCoupon) GetMember() *MemberInfo {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *MemberCoupon) GetUseOrder() *OrderInfo {
	if x != nil {
		return x.UseOrder
	}
	return nil
}

func (x *MemberCoupon) GetBuyGoodsIds() []int64 {
	if x != nil {
		return x.BuyGoodsIds
	}
	return nil
}

type MemberCouponRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top       int32    `protobuf:"varint,1,opt,name=top,proto3" json:"top"`
	Paged     int64    `protobuf:"varint,2,opt,name=paged,proto3" json:"paged"`
	PageSize  int64    `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Keywords  string   `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Sorts     []string `protobuf:"bytes,5,rep,name=sorts,proto3" json:"sorts"`
	DateRange []string `protobuf:"bytes,6,rep,name=date_range,json=dateRange,proto3" json:"date_range"`
	Ids       []int64  `protobuf:"varint,7,rep,packed,name=ids,proto3" json:"ids"`
	Id        int64    `protobuf:"varint,8,opt,name=id,proto3" json:"id"`
	//以下为自定义参数
	Code         string `protobuf:"bytes,11,opt,name=code,proto3" json:"code"`                                     //凭证编号
	CouponId     int64  `protobuf:"varint,12,opt,name=coupon_id,json=couponId,proto3" json:"coupon_id"`            //优惠券id
	MemberId     int64  `protobuf:"varint,13,opt,name=member_id,json=memberId,proto3" json:"member_id"`            //用户id
	ReceiveScene string `protobuf:"bytes,14,opt,name=receive_scene,json=receiveScene,proto3" json:"receive_scene"` //领取场景
	UsedOrderId  int64  `protobuf:"varint,15,opt,name=used_order_id,json=usedOrderId,proto3" json:"used_order_id"` //使用订单id
	Status       string `protobuf:"bytes,16,opt,name=status,proto3" json:"status"`                                 //优惠劵状态
}

func (x *MemberCouponRequest) Reset() {
	*x = MemberCouponRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberCouponService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberCouponRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCouponRequest) ProtoMessage() {}

func (x *MemberCouponRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memberCouponService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCouponRequest.ProtoReflect.Descriptor instead.
func (*MemberCouponRequest) Descriptor() ([]byte, []int) {
	return file_memberCouponService_proto_rawDescGZIP(), []int{1}
}

func (x *MemberCouponRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *MemberCouponRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *MemberCouponRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MemberCouponRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *MemberCouponRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *MemberCouponRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *MemberCouponRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *MemberCouponRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberCouponRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *MemberCouponRequest) GetCouponId() int64 {
	if x != nil {
		return x.CouponId
	}
	return 0
}

func (x *MemberCouponRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberCouponRequest) GetReceiveScene() string {
	if x != nil {
		return x.ReceiveScene
	}
	return ""
}

func (x *MemberCouponRequest) GetUsedOrderId() int64 {
	if x != nil {
		return x.UsedOrderId
	}
	return 0
}

func (x *MemberCouponRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type MemberCouponIndexResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pager         *common.Pager   `protobuf:"bytes,1,opt,name=pager,proto3" json:"pager"`
	Items         []*MemberCoupon `protobuf:"bytes,2,rep,name=items,proto3" json:"items"`
	NotUse        []*MemberCoupon `protobuf:"bytes,3,rep,name=not_use,json=notUse,proto3" json:"not_use"`
	AlreadyUse    []*MemberCoupon `protobuf:"bytes,4,rep,name=already_use,json=alreadyUse,proto3" json:"already_use"`
	AlreadyExpire []*MemberCoupon `protobuf:"bytes,5,rep,name=already_expire,json=alreadyExpire,proto3" json:"already_expire"`
	Msg           string          `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg"`
}

func (x *MemberCouponIndexResponse) Reset() {
	*x = MemberCouponIndexResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberCouponService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberCouponIndexResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCouponIndexResponse) ProtoMessage() {}

func (x *MemberCouponIndexResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memberCouponService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCouponIndexResponse.ProtoReflect.Descriptor instead.
func (*MemberCouponIndexResponse) Descriptor() ([]byte, []int) {
	return file_memberCouponService_proto_rawDescGZIP(), []int{2}
}

func (x *MemberCouponIndexResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MemberCouponIndexResponse) GetItems() []*MemberCoupon {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MemberCouponIndexResponse) GetNotUse() []*MemberCoupon {
	if x != nil {
		return x.NotUse
	}
	return nil
}

func (x *MemberCouponIndexResponse) GetAlreadyUse() []*MemberCoupon {
	if x != nil {
		return x.AlreadyUse
	}
	return nil
}

func (x *MemberCouponIndexResponse) GetAlreadyExpire() []*MemberCoupon {
	if x != nil {
		return x.AlreadyExpire
	}
	return nil
}

func (x *MemberCouponIndexResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type MemberCouponResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *MemberCoupon   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*MemberCoupon `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string          `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
}

func (x *MemberCouponResponse) Reset() {
	*x = MemberCouponResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memberCouponService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberCouponResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberCouponResponse) ProtoMessage() {}

func (x *MemberCouponResponse) ProtoReflect() protoreflect.Message {
	mi := &file_memberCouponService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberCouponResponse.ProtoReflect.Descriptor instead.
func (*MemberCouponResponse) Descriptor() ([]byte, []int) {
	return file_memberCouponService_proto_rawDescGZIP(), []int{3}
}

func (x *MemberCouponResponse) GetEntity() *MemberCoupon {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MemberCouponResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MemberCouponResponse) GetItems() []*MemberCoupon {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MemberCouponResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_memberCouponService_proto protoreflect.FileDescriptor

var file_memberCouponService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x04, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x73, 0x65,
	0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x75, 0x79, 0x5f, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x75, 0x79,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x73, 0x22, 0xfc, 0x02, 0x0a, 0x13, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74,
	0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x5f, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x75,
	0x73, 0x65, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa9, 0x02, 0x0a, 0x19, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x5f,
	0x75, 0x73, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x61, 0x6c, 0x72,
	0x65, 0x61, 0x64, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x0a, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x55,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x52, 0x0d, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0xab, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x32, 0x86, 0x03, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x1a, 0x1e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x1a,
	0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x55,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_memberCouponService_proto_rawDescOnce sync.Once
	file_memberCouponService_proto_rawDescData = file_memberCouponService_proto_rawDesc
)

func file_memberCouponService_proto_rawDescGZIP() []byte {
	file_memberCouponService_proto_rawDescOnce.Do(func() {
		file_memberCouponService_proto_rawDescData = protoimpl.X.CompressGZIP(file_memberCouponService_proto_rawDescData)
	})
	return file_memberCouponService_proto_rawDescData
}

var file_memberCouponService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_memberCouponService_proto_goTypes = []interface{}{
	(*MemberCoupon)(nil),              // 0: services.MemberCoupon
	(*MemberCouponRequest)(nil),       // 1: services.MemberCouponRequest
	(*MemberCouponIndexResponse)(nil), // 2: services.MemberCouponIndexResponse
	(*MemberCouponResponse)(nil),      // 3: services.MemberCouponResponse
	(*Coupon)(nil),                    // 4: services.Coupon
	(*MemberInfo)(nil),                // 5: services.MemberInfo
	(*OrderInfo)(nil),                 // 6: services.OrderInfo
	(*common.Pager)(nil),              // 7: common.Pager
}
var file_memberCouponService_proto_depIdxs = []int32{
	4,  // 0: services.MemberCoupon.coupon:type_name -> services.Coupon
	5,  // 1: services.MemberCoupon.member:type_name -> services.MemberInfo
	6,  // 2: services.MemberCoupon.use_order:type_name -> services.OrderInfo
	7,  // 3: services.MemberCouponIndexResponse.pager:type_name -> common.Pager
	0,  // 4: services.MemberCouponIndexResponse.items:type_name -> services.MemberCoupon
	0,  // 5: services.MemberCouponIndexResponse.not_use:type_name -> services.MemberCoupon
	0,  // 6: services.MemberCouponIndexResponse.already_use:type_name -> services.MemberCoupon
	0,  // 7: services.MemberCouponIndexResponse.already_expire:type_name -> services.MemberCoupon
	0,  // 8: services.MemberCouponResponse.entity:type_name -> services.MemberCoupon
	7,  // 9: services.MemberCouponResponse.pager:type_name -> common.Pager
	0,  // 10: services.MemberCouponResponse.items:type_name -> services.MemberCoupon
	0,  // 11: services.MemberCouponService.Delete:input_type -> services.MemberCoupon
	0,  // 12: services.MemberCouponService.Detail:input_type -> services.MemberCoupon
	1,  // 13: services.MemberCouponService.Index:input_type -> services.MemberCouponRequest
	1,  // 14: services.MemberCouponService.Search:input_type -> services.MemberCouponRequest
	1,  // 15: services.MemberCouponService.UsableList:input_type -> services.MemberCouponRequest
	3,  // 16: services.MemberCouponService.Delete:output_type -> services.MemberCouponResponse
	3,  // 17: services.MemberCouponService.Detail:output_type -> services.MemberCouponResponse
	2,  // 18: services.MemberCouponService.Index:output_type -> services.MemberCouponIndexResponse
	3,  // 19: services.MemberCouponService.Search:output_type -> services.MemberCouponResponse
	3,  // 20: services.MemberCouponService.UsableList:output_type -> services.MemberCouponResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_memberCouponService_proto_init() }
func file_memberCouponService_proto_init() {
	if File_memberCouponService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	file_couponService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_memberCouponService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberCoupon); i {
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
		file_memberCouponService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberCouponRequest); i {
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
		file_memberCouponService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberCouponIndexResponse); i {
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
		file_memberCouponService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberCouponResponse); i {
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
			RawDescriptor: file_memberCouponService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memberCouponService_proto_goTypes,
		DependencyIndexes: file_memberCouponService_proto_depIdxs,
		MessageInfos:      file_memberCouponService_proto_msgTypes,
	}.Build()
	File_memberCouponService_proto = out.File
	file_memberCouponService_proto_rawDesc = nil
	file_memberCouponService_proto_goTypes = nil
	file_memberCouponService_proto_depIdxs = nil
}
