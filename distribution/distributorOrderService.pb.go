// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: distributorOrderService.proto

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

type DistributorOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                  //ID
	MemberId   int64       `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id"`      //用户ID
	OrderId    int64       `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id"`         //订单ID
	Status     string      `protobuf:"bytes,4,opt,name=status,proto3" json:"status"`                           //状态（0待生效, 1待结算, 2已结算, 3已失效）
	ValidTime  string      `protobuf:"bytes,5,opt,name=valid_time,json=validTime,proto3" json:"valid_time"`    //生效时间
	SettleTime string      `protobuf:"bytes,6,opt,name=settle_time,json=settleTime,proto3" json:"settle_time"` //结算时间
	CreatedAt  string      `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  string      `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Member     *MemberInfo `protobuf:"bytes,10,opt,name=member,proto3" json:"member"`
	Order      *OrderInfo  `protobuf:"bytes,11,opt,name=order,proto3" json:"order"`
	StatusName string      `protobuf:"bytes,12,opt,name=status_name,json=statusName,proto3" json:"status_name"`
}

func (x *DistributorOrder) Reset() {
	*x = DistributorOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distributorOrderService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributorOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributorOrder) ProtoMessage() {}

func (x *DistributorOrder) ProtoReflect() protoreflect.Message {
	mi := &file_distributorOrderService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributorOrder.ProtoReflect.Descriptor instead.
func (*DistributorOrder) Descriptor() ([]byte, []int) {
	return file_distributorOrderService_proto_rawDescGZIP(), []int{0}
}

func (x *DistributorOrder) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DistributorOrder) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *DistributorOrder) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *DistributorOrder) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DistributorOrder) GetValidTime() string {
	if x != nil {
		return x.ValidTime
	}
	return ""
}

func (x *DistributorOrder) GetSettleTime() string {
	if x != nil {
		return x.SettleTime
	}
	return ""
}

func (x *DistributorOrder) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *DistributorOrder) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *DistributorOrder) GetMember() *MemberInfo {
	if x != nil {
		return x.Member
	}
	return nil
}

func (x *DistributorOrder) GetOrder() *OrderInfo {
	if x != nil {
		return x.Order
	}
	return nil
}

func (x *DistributorOrder) GetStatusName() string {
	if x != nil {
		return x.StatusName
	}
	return ""
}

type DistributorOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged      int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize   int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Id         int64   `protobuf:"varint,3,opt,name=id,proto3" json:"id"`
	Mobile     string  `protobuf:"bytes,4,opt,name=mobile,proto3" json:"mobile"`
	OrderId    int64   `protobuf:"varint,5,opt,name=order_id,json=orderId,proto3" json:"order_id"`
	OrderSn    string  `protobuf:"bytes,6,opt,name=order_sn,json=orderSn,proto3" json:"order_sn"`
	Keywords   string  `protobuf:"bytes,7,opt,name=keywords,proto3" json:"keywords"`
	Status     string  `protobuf:"bytes,8,opt,name=status,proto3" json:"status"`
	StartDate  string  `protobuf:"bytes,9,opt,name=start_date,json=startDate,proto3" json:"start_date"`
	EndDate    string  `protobuf:"bytes,10,opt,name=end_date,json=endDate,proto3" json:"end_date"`
	IsSettled  string  `protobuf:"bytes,11,opt,name=is_settled,json=isSettled,proto3" json:"is_settled"`
	MemberId   int64   `protobuf:"varint,12,opt,name=member_id,json=memberId,proto3" json:"member_id"`
	ReferrerId int64   `protobuf:"varint,13,opt,name=referrer_id,json=referrerId,proto3" json:"referrer_id"`
	LevelId    int32   `protobuf:"varint,14,opt,name=level_id,json=levelId,proto3" json:"level_id"`
	Ids        []int64 `protobuf:"varint,15,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *DistributorOrderRequest) Reset() {
	*x = DistributorOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distributorOrderService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributorOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributorOrderRequest) ProtoMessage() {}

func (x *DistributorOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_distributorOrderService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributorOrderRequest.ProtoReflect.Descriptor instead.
func (*DistributorOrderRequest) Descriptor() ([]byte, []int) {
	return file_distributorOrderService_proto_rawDescGZIP(), []int{1}
}

func (x *DistributorOrderRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *DistributorOrderRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *DistributorOrderRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DistributorOrderRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *DistributorOrderRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *DistributorOrderRequest) GetOrderSn() string {
	if x != nil {
		return x.OrderSn
	}
	return ""
}

func (x *DistributorOrderRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *DistributorOrderRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DistributorOrderRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *DistributorOrderRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *DistributorOrderRequest) GetIsSettled() string {
	if x != nil {
		return x.IsSettled
	}
	return ""
}

func (x *DistributorOrderRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *DistributorOrderRequest) GetReferrerId() int64 {
	if x != nil {
		return x.ReferrerId
	}
	return 0
}

func (x *DistributorOrderRequest) GetLevelId() int32 {
	if x != nil {
		return x.LevelId
	}
	return 0
}

func (x *DistributorOrderRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DistributorOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *DistributorOrder   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager       `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*DistributorOrder `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string              `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *DistributorOrderResponse) Reset() {
	*x = DistributorOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_distributorOrderService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributorOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributorOrderResponse) ProtoMessage() {}

func (x *DistributorOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_distributorOrderService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributorOrderResponse.ProtoReflect.Descriptor instead.
func (*DistributorOrderResponse) Descriptor() ([]byte, []int) {
	return file_distributorOrderService_proto_rawDescGZIP(), []int{2}
}

func (x *DistributorOrderResponse) GetEntity() *DistributorOrder {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *DistributorOrderResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *DistributorOrderResponse) GetItems() []*DistributorOrder {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DistributorOrderResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_distributorOrderService_proto protoreflect.FileDescriptor

var file_distributorOrderService_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x62, 0x61,
	0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xea, 0x02, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xa2, 0x03, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x73, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xb9, 0x01, 0x0a, 0x18, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x32, 0x8e, 0x02, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a,
	0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x06, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x09,
	0x50, 0x75, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_distributorOrderService_proto_rawDescOnce sync.Once
	file_distributorOrderService_proto_rawDescData = file_distributorOrderService_proto_rawDesc
)

func file_distributorOrderService_proto_rawDescGZIP() []byte {
	file_distributorOrderService_proto_rawDescOnce.Do(func() {
		file_distributorOrderService_proto_rawDescData = protoimpl.X.CompressGZIP(file_distributorOrderService_proto_rawDescData)
	})
	return file_distributorOrderService_proto_rawDescData
}

var file_distributorOrderService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_distributorOrderService_proto_goTypes = []interface{}{
	(*DistributorOrder)(nil),         // 0: services.DistributorOrder
	(*DistributorOrderRequest)(nil),  // 1: services.DistributorOrderRequest
	(*DistributorOrderResponse)(nil), // 2: services.DistributorOrderResponse
	(*MemberInfo)(nil),               // 3: services.MemberInfo
	(*OrderInfo)(nil),                // 4: services.OrderInfo
	(*common.Pager)(nil),             // 5: common.Pager
}
var file_distributorOrderService_proto_depIdxs = []int32{
	3, // 0: services.DistributorOrder.member:type_name -> services.MemberInfo
	4, // 1: services.DistributorOrder.order:type_name -> services.OrderInfo
	0, // 2: services.DistributorOrderResponse.entity:type_name -> services.DistributorOrder
	5, // 3: services.DistributorOrderResponse.pager:type_name -> common.Pager
	0, // 4: services.DistributorOrderResponse.items:type_name -> services.DistributorOrder
	0, // 5: services.DistributorOrderService.Detail:input_type -> services.DistributorOrder
	1, // 6: services.DistributorOrderService.Search:input_type -> services.DistributorOrderRequest
	1, // 7: services.DistributorOrderService.PullOrder:input_type -> services.DistributorOrderRequest
	2, // 8: services.DistributorOrderService.Detail:output_type -> services.DistributorOrderResponse
	2, // 9: services.DistributorOrderService.Search:output_type -> services.DistributorOrderResponse
	2, // 10: services.DistributorOrderService.PullOrder:output_type -> services.DistributorOrderResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_distributorOrderService_proto_init() }
func file_distributorOrderService_proto_init() {
	if File_distributorOrderService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_distributorOrderService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributorOrder); i {
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
		file_distributorOrderService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributorOrderRequest); i {
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
		file_distributorOrderService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributorOrderResponse); i {
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
			RawDescriptor: file_distributorOrderService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_distributorOrderService_proto_goTypes,
		DependencyIndexes: file_distributorOrderService_proto_depIdxs,
		MessageInfos:      file_distributorOrderService_proto_msgTypes,
	}.Build()
	File_distributorOrderService_proto = out.File
	file_distributorOrderService_proto_rawDesc = nil
	file_distributorOrderService_proto_goTypes = nil
	file_distributorOrderService_proto_depIdxs = nil
}
