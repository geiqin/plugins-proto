// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: combinationService.proto

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

type Combination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                            //ID
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`                                         //套餐名称
	Type         string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`                                         //套餐类型
	BeginDate    string   `protobuf:"bytes,4,opt,name=begin_date,json=beginDate,proto3" json:"begin_date"`              //开始日期
	EndDate      string   `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date"`                    //结束日期
	Stock        int32    `protobuf:"varint,6,opt,name=stock,proto3" json:"stock"`                                      //可售库存
	LimitBuyType string   `protobuf:"bytes,7,opt,name=limit_buy_type,json=limitBuyType,proto3" json:"limit_buy_type"`   //限购类型
	LimitBuyNum  int32    `protobuf:"varint,8,opt,name=limit_buy_num,json=limitBuyNum,proto3" json:"limit_buy_num"`     //限购数量
	IsSyncAddTag string   `protobuf:"bytes,9,opt,name=is_sync_add_tag,json=isSyncAddTag,proto3" json:"is_sync_add_tag"` //是否同步标签
	SyncAddTags  []string `protobuf:"bytes,10,rep,name=sync_add_tags,json=syncAddTags,proto3" json:"sync_add_tags"`     //同步的标签
	Description  string   `protobuf:"bytes,11,opt,name=description,proto3" json:"description"`                          //活动描述
	Status       string   `protobuf:"bytes,12,opt,name=status,proto3" json:"status"`                                    //状态
	CreatedAt    string   `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at"`             //创建时间
	UpdatedAt    string   `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`             //修改时间
}

func (x *Combination) Reset() {
	*x = Combination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combinationService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Combination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Combination) ProtoMessage() {}

func (x *Combination) ProtoReflect() protoreflect.Message {
	mi := &file_combinationService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Combination.ProtoReflect.Descriptor instead.
func (*Combination) Descriptor() ([]byte, []int) {
	return file_combinationService_proto_rawDescGZIP(), []int{0}
}

func (x *Combination) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Combination) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Combination) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Combination) GetBeginDate() string {
	if x != nil {
		return x.BeginDate
	}
	return ""
}

func (x *Combination) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Combination) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Combination) GetLimitBuyType() string {
	if x != nil {
		return x.LimitBuyType
	}
	return ""
}

func (x *Combination) GetLimitBuyNum() int32 {
	if x != nil {
		return x.LimitBuyNum
	}
	return 0
}

func (x *Combination) GetIsSyncAddTag() string {
	if x != nil {
		return x.IsSyncAddTag
	}
	return ""
}

func (x *Combination) GetSyncAddTags() []string {
	if x != nil {
		return x.SyncAddTags
	}
	return nil
}

func (x *Combination) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Combination) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Combination) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Combination) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CombinationRequest struct {
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
	Name         string `protobuf:"bytes,11,opt,name=name,proto3" json:"name"`                                         //套餐名称
	Type         string `protobuf:"bytes,12,opt,name=type,proto3" json:"type"`                                         //套餐类型
	LimitBuyType string `protobuf:"bytes,13,opt,name=limit_buy_type,json=limitBuyType,proto3" json:"limit_buy_type"`   //限购类型
	IsSyncAddTag string `protobuf:"bytes,14,opt,name=is_sync_add_tag,json=isSyncAddTag,proto3" json:"is_sync_add_tag"` //是否同步标签
	Status       string `protobuf:"bytes,15,opt,name=status,proto3" json:"status"`                                     //状态
}

func (x *CombinationRequest) Reset() {
	*x = CombinationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combinationService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombinationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombinationRequest) ProtoMessage() {}

func (x *CombinationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_combinationService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombinationRequest.ProtoReflect.Descriptor instead.
func (*CombinationRequest) Descriptor() ([]byte, []int) {
	return file_combinationService_proto_rawDescGZIP(), []int{1}
}

func (x *CombinationRequest) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *CombinationRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *CombinationRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CombinationRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *CombinationRequest) GetSorts() []string {
	if x != nil {
		return x.Sorts
	}
	return nil
}

func (x *CombinationRequest) GetDateRange() []string {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *CombinationRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *CombinationRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CombinationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CombinationRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CombinationRequest) GetLimitBuyType() string {
	if x != nil {
		return x.LimitBuyType
	}
	return ""
}

func (x *CombinationRequest) GetIsSyncAddTag() string {
	if x != nil {
		return x.IsSyncAddTag
	}
	return ""
}

func (x *CombinationRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CombinationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Combination   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager  `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Combination `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Msg    string         `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg"`
}

func (x *CombinationResponse) Reset() {
	*x = CombinationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combinationService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombinationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombinationResponse) ProtoMessage() {}

func (x *CombinationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_combinationService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombinationResponse.ProtoReflect.Descriptor instead.
func (*CombinationResponse) Descriptor() ([]byte, []int) {
	return file_combinationService_proto_rawDescGZIP(), []int{2}
}

func (x *CombinationResponse) GetEntity() *Combination {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *CombinationResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *CombinationResponse) GetItems() []*Combination {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CombinationResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_combinationService_proto protoreflect.FileDescriptor

var file_combinationService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x03, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x62,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x24, 0x0a, 0x0e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x62, 0x75, 0x79, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x75,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x62,
	0x75, 0x79, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x42, 0x75, 0x79, 0x4e, 0x75, 0x6d, 0x12, 0x25, 0x0a, 0x0f, 0x69, 0x73, 0x5f,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x64, 0x64, 0x54, 0x61, 0x67,
	0x12, 0x22, 0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x41, 0x64, 0x64,
	0x54, 0x61, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd9, 0x02, 0x0a,
	0x12, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x74, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x62, 0x75, 0x79,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x42, 0x75, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0f, 0x69, 0x73, 0x5f,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x64, 0x64, 0x54, 0x61, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x62,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52, 0x05, 0x70,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x32, 0xec, 0x03, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1d, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x62, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_combinationService_proto_rawDescOnce sync.Once
	file_combinationService_proto_rawDescData = file_combinationService_proto_rawDesc
)

func file_combinationService_proto_rawDescGZIP() []byte {
	file_combinationService_proto_rawDescOnce.Do(func() {
		file_combinationService_proto_rawDescData = protoimpl.X.CompressGZIP(file_combinationService_proto_rawDescData)
	})
	return file_combinationService_proto_rawDescData
}

var file_combinationService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_combinationService_proto_goTypes = []interface{}{
	(*Combination)(nil),         // 0: services.Combination
	(*CombinationRequest)(nil),  // 1: services.CombinationRequest
	(*CombinationResponse)(nil), // 2: services.CombinationResponse
	(*common.Pager)(nil),        // 3: common.Pager
}
var file_combinationService_proto_depIdxs = []int32{
	0,  // 0: services.CombinationResponse.entity:type_name -> services.Combination
	3,  // 1: services.CombinationResponse.pager:type_name -> common.Pager
	0,  // 2: services.CombinationResponse.items:type_name -> services.Combination
	0,  // 3: services.CombinationService.Create:input_type -> services.Combination
	0,  // 4: services.CombinationService.Update:input_type -> services.Combination
	0,  // 5: services.CombinationService.Delete:input_type -> services.Combination
	0,  // 6: services.CombinationService.Stop:input_type -> services.Combination
	0,  // 7: services.CombinationService.Detail:input_type -> services.Combination
	1,  // 8: services.CombinationService.Search:input_type -> services.CombinationRequest
	1,  // 9: services.CombinationService.List:input_type -> services.CombinationRequest
	2,  // 10: services.CombinationService.Create:output_type -> services.CombinationResponse
	2,  // 11: services.CombinationService.Update:output_type -> services.CombinationResponse
	2,  // 12: services.CombinationService.Delete:output_type -> services.CombinationResponse
	2,  // 13: services.CombinationService.Stop:output_type -> services.CombinationResponse
	2,  // 14: services.CombinationService.Detail:output_type -> services.CombinationResponse
	2,  // 15: services.CombinationService.Search:output_type -> services.CombinationResponse
	2,  // 16: services.CombinationService.List:output_type -> services.CombinationResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_combinationService_proto_init() }
func file_combinationService_proto_init() {
	if File_combinationService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_combinationService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Combination); i {
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
		file_combinationService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombinationRequest); i {
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
		file_combinationService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombinationResponse); i {
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
			RawDescriptor: file_combinationService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_combinationService_proto_goTypes,
		DependencyIndexes: file_combinationService_proto_depIdxs,
		MessageInfos:      file_combinationService_proto_msgTypes,
	}.Build()
	File_combinationService_proto = out.File
	file_combinationService_proto_rawDesc = nil
	file_combinationService_proto_goTypes = nil
	file_combinationService_proto_depIdxs = nil
}
