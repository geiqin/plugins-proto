// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: extractionService.proto

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

type ExtractionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged       int64   `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize    int64   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting     string  `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	Keywords    string  `protobuf:"bytes,4,opt,name=keywords,proto3" json:"keywords"`
	Id          int64   `protobuf:"varint,5,opt,name=id,proto3" json:"id"`
	Name        string  `protobuf:"bytes,6,opt,name=name,proto3" json:"name"`
	Lng         float32 `protobuf:"fixed32,7,opt,name=lng,proto3" json:"lng"`
	Lat         float32 `protobuf:"fixed32,8,opt,name=lat,proto3" json:"lat"`
	ShopId      int64   `protobuf:"varint,10,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`                //多店铺ID
	RealstoreId int64   `protobuf:"varint,11,opt,name=realstore_id,json=realstoreId,proto3" json:"realstore_id"` //多门店ID
	ProvinceId  int64   `protobuf:"varint,12,opt,name=province_id,json=provinceId,proto3" json:"province_id"`    //所在省ID
	CityId      int64   `protobuf:"varint,13,opt,name=city_id,json=cityId,proto3" json:"city_id"`                //所在市ID
	CountyId    int64   `protobuf:"varint,14,opt,name=county_id,json=countyId,proto3" json:"county_id"`          //所在县/区ID
	Ids         []int64 `protobuf:"varint,15,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *ExtractionRequest) Reset() {
	*x = ExtractionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extractionService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractionRequest) ProtoMessage() {}

func (x *ExtractionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_extractionService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractionRequest.ProtoReflect.Descriptor instead.
func (*ExtractionRequest) Descriptor() ([]byte, []int) {
	return file_extractionService_proto_rawDescGZIP(), []int{0}
}

func (x *ExtractionRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *ExtractionRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ExtractionRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *ExtractionRequest) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

func (x *ExtractionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExtractionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExtractionRequest) GetLng() float32 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *ExtractionRequest) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *ExtractionRequest) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *ExtractionRequest) GetRealstoreId() int64 {
	if x != nil {
		return x.RealstoreId
	}
	return 0
}

func (x *ExtractionRequest) GetProvinceId() int64 {
	if x != nil {
		return x.ProvinceId
	}
	return 0
}

func (x *ExtractionRequest) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *ExtractionRequest) GetCountyId() int64 {
	if x != nil {
		return x.CountyId
	}
	return 0
}

func (x *ExtractionRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Extraction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	ShopId        int64             `protobuf:"varint,2,opt,name=shop_id,json=shopId,proto3" json:"shop_id"`                   //多店铺ID
	RealstoreId   int64             `protobuf:"varint,3,opt,name=realstore_id,json=realstoreId,proto3" json:"realstore_id"`    //多门店ID
	Name          string            `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`                                      //姓名
	Alias         string            `protobuf:"bytes,5,opt,name=alias,proto3" json:"alias"`                                    //别名
	Tel           string            `protobuf:"bytes,6,opt,name=tel,proto3" json:"tel"`                                        //电话
	ProvinceId    int64             `protobuf:"varint,7,opt,name=province_id,json=provinceId,proto3" json:"province_id"`       //所在省ID
	CityId        int64             `protobuf:"varint,8,opt,name=city_id,json=cityId,proto3" json:"city_id"`                   //所在市ID
	CountyId      int64             `protobuf:"varint,9,opt,name=county_id,json=countyId,proto3" json:"county_id"`             //所在县/区ID
	ProvinceName  string            `protobuf:"bytes,10,opt,name=province_name,json=provinceName,proto3" json:"province_name"` //所在省
	CityName      string            `protobuf:"bytes,11,opt,name=city_name,json=cityName,proto3" json:"city_name"`             //所在市
	CountyName    string            `protobuf:"bytes,12,opt,name=county_name,json=countyName,proto3" json:"county_name"`       //所在县/区
	Address       string            `protobuf:"bytes,13,opt,name=address,proto3" json:"address"`                               //详细地址
	Lng           float32           `protobuf:"fixed32,14,opt,name=lng,proto3" json:"lng"`                                     //经度
	Lat           float32           `protobuf:"fixed32,15,opt,name=lat,proto3" json:"lat"`                                     //纬度
	LogoUrl       string            `protobuf:"bytes,16,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url"`
	Memo          string            `protobuf:"bytes,17,opt,name=memo,proto3" json:"memo"`
	Config        *ExtractionConfig `protobuf:"bytes,18,opt,name=config,proto3" json:"config"` //配置信息
	CreatedAt     string            `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt     string            `protobuf:"bytes,22,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	DistanceValue float32           `protobuf:"fixed32,27,opt,name=distance_value,json=distanceValue,proto3" json:"distance_value"` //距离数据（离用户位置）
	DistanceUnit  string            `protobuf:"bytes,28,opt,name=distance_unit,json=distanceUnit,proto3" json:"distance_unit"`      //距离单位（离用户位置）
}

func (x *Extraction) Reset() {
	*x = Extraction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extractionService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extraction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extraction) ProtoMessage() {}

func (x *Extraction) ProtoReflect() protoreflect.Message {
	mi := &file_extractionService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extraction.ProtoReflect.Descriptor instead.
func (*Extraction) Descriptor() ([]byte, []int) {
	return file_extractionService_proto_rawDescGZIP(), []int{1}
}

func (x *Extraction) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Extraction) GetShopId() int64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *Extraction) GetRealstoreId() int64 {
	if x != nil {
		return x.RealstoreId
	}
	return 0
}

func (x *Extraction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Extraction) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *Extraction) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *Extraction) GetProvinceId() int64 {
	if x != nil {
		return x.ProvinceId
	}
	return 0
}

func (x *Extraction) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *Extraction) GetCountyId() int64 {
	if x != nil {
		return x.CountyId
	}
	return 0
}

func (x *Extraction) GetProvinceName() string {
	if x != nil {
		return x.ProvinceName
	}
	return ""
}

func (x *Extraction) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *Extraction) GetCountyName() string {
	if x != nil {
		return x.CountyName
	}
	return ""
}

func (x *Extraction) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Extraction) GetLng() float32 {
	if x != nil {
		return x.Lng
	}
	return 0
}

func (x *Extraction) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Extraction) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *Extraction) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Extraction) GetConfig() *ExtractionConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Extraction) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Extraction) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Extraction) GetDistanceValue() float32 {
	if x != nil {
		return x.DistanceValue
	}
	return 0
}

func (x *Extraction) GetDistanceUnit() string {
	if x != nil {
		return x.DistanceUnit
	}
	return ""
}

type ExtractionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reception              string            `protobuf:"bytes,9,opt,name=reception,proto3" json:"reception"`
	ReceptionRepeatWeeks   int32             `protobuf:"varint,10,opt,name=reception_repeat_weeks,json=receptionRepeatWeeks,proto3" json:"reception_repeat_weeks"`
	ReceptionRepeatWeekArr []int32           `protobuf:"varint,11,rep,packed,name=reception_repeat_week_arr,json=receptionRepeatWeekArr,proto3" json:"reception_repeat_week_arr"`
	IsFetchTime            bool              `protobuf:"varint,12,opt,name=is_fetch_time,json=isFetchTime,proto3" json:"is_fetch_time"`
	Fetch                  string            `protobuf:"bytes,13,opt,name=fetch,proto3" json:"fetch"`
	FetchRepeatWeeks       int32             `protobuf:"varint,14,opt,name=fetch_repeat_weeks,json=fetchRepeatWeeks,proto3" json:"fetch_repeat_weeks"`
	FetchRepeatWeekArr     []int32           `protobuf:"varint,15,rep,packed,name=fetch_repeat_week_arr,json=fetchRepeatWeekArr,proto3" json:"fetch_repeat_week_arr"`
	SubFetchTime           string            `protobuf:"bytes,16,opt,name=sub_fetch_time,json=subFetchTime,proto3" json:"sub_fetch_time"`
	Appointment            string            `protobuf:"bytes,17,opt,name=appointment,proto3" json:"appointment"`
	AppointmentNum         int32             `protobuf:"varint,18,opt,name=appointment_num,json=appointmentNum,proto3" json:"appointment_num"`
	MaxAppointmentNum      int32             `protobuf:"varint,19,opt,name=max_appointment_num,json=maxAppointmentNum,proto3" json:"max_appointment_num"`
	ReceptionTimes         []*ExtractionTime `protobuf:"bytes,23,rep,name=reception_times,json=receptionTimes,proto3" json:"reception_times"`
	FetchTimes             []*ExtractionTime `protobuf:"bytes,24,rep,name=fetch_times,json=fetchTimes,proto3" json:"fetch_times"`
}

func (x *ExtractionConfig) Reset() {
	*x = ExtractionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extractionService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractionConfig) ProtoMessage() {}

func (x *ExtractionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_extractionService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractionConfig.ProtoReflect.Descriptor instead.
func (*ExtractionConfig) Descriptor() ([]byte, []int) {
	return file_extractionService_proto_rawDescGZIP(), []int{2}
}

func (x *ExtractionConfig) GetReception() string {
	if x != nil {
		return x.Reception
	}
	return ""
}

func (x *ExtractionConfig) GetReceptionRepeatWeeks() int32 {
	if x != nil {
		return x.ReceptionRepeatWeeks
	}
	return 0
}

func (x *ExtractionConfig) GetReceptionRepeatWeekArr() []int32 {
	if x != nil {
		return x.ReceptionRepeatWeekArr
	}
	return nil
}

func (x *ExtractionConfig) GetIsFetchTime() bool {
	if x != nil {
		return x.IsFetchTime
	}
	return false
}

func (x *ExtractionConfig) GetFetch() string {
	if x != nil {
		return x.Fetch
	}
	return ""
}

func (x *ExtractionConfig) GetFetchRepeatWeeks() int32 {
	if x != nil {
		return x.FetchRepeatWeeks
	}
	return 0
}

func (x *ExtractionConfig) GetFetchRepeatWeekArr() []int32 {
	if x != nil {
		return x.FetchRepeatWeekArr
	}
	return nil
}

func (x *ExtractionConfig) GetSubFetchTime() string {
	if x != nil {
		return x.SubFetchTime
	}
	return ""
}

func (x *ExtractionConfig) GetAppointment() string {
	if x != nil {
		return x.Appointment
	}
	return ""
}

func (x *ExtractionConfig) GetAppointmentNum() int32 {
	if x != nil {
		return x.AppointmentNum
	}
	return 0
}

func (x *ExtractionConfig) GetMaxAppointmentNum() int32 {
	if x != nil {
		return x.MaxAppointmentNum
	}
	return 0
}

func (x *ExtractionConfig) GetReceptionTimes() []*ExtractionTime {
	if x != nil {
		return x.ReceptionTimes
	}
	return nil
}

func (x *ExtractionConfig) GetFetchTimes() []*ExtractionTime {
	if x != nil {
		return x.FetchTimes
	}
	return nil
}

type ExtractionTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type"`
	StartTime string `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time"`
	EndTime   string `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time"`
}

func (x *ExtractionTime) Reset() {
	*x = ExtractionTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extractionService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractionTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractionTime) ProtoMessage() {}

func (x *ExtractionTime) ProtoReflect() protoreflect.Message {
	mi := &file_extractionService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractionTime.ProtoReflect.Descriptor instead.
func (*ExtractionTime) Descriptor() ([]byte, []int) {
	return file_extractionService_proto_rawDescGZIP(), []int{3}
}

func (x *ExtractionTime) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ExtractionTime) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *ExtractionTime) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type ExtractionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Extraction   `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Extraction `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Info   string        `protobuf:"bytes,4,opt,name=info,proto3" json:"info"`
}

func (x *ExtractionResponse) Reset() {
	*x = ExtractionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extractionService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractionResponse) ProtoMessage() {}

func (x *ExtractionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_extractionService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractionResponse.ProtoReflect.Descriptor instead.
func (*ExtractionResponse) Descriptor() ([]byte, []int) {
	return file_extractionService_proto_rawDescGZIP(), []int{4}
}

func (x *ExtractionResponse) GetEntity() *Extraction {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ExtractionResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *ExtractionResponse) GetItems() []*Extraction {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ExtractionResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

var File_extractionService_proto protoreflect.FileDescriptor

var file_extractionService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x02, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x61, 0x67, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6e, 0x67,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x6c, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65,
	0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x22, 0xf9, 0x04, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x72, 0x65, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6c, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x12, 0x32, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x22, 0xdb,
	0x04, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x14, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x72, 0x65, 0x63, 0x65, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b,
	0x5f, 0x61, 0x72, 0x72, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x05, 0x52, 0x16, 0x72, 0x65, 0x63, 0x65,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x41,
	0x72, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x65, 0x74, 0x63, 0x68, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x65, 0x74, 0x63, 0x68, 0x12, 0x2c, 0x0a, 0x12,
	0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x77, 0x65, 0x65,
	0x6b, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x66, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x66, 0x65,
	0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x5f,
	0x61, 0x72, 0x72, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x05, 0x52, 0x12, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x41, 0x72, 0x72, 0x12, 0x24, 0x0a,
	0x0e, 0x73, 0x75, 0x62, 0x5f, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x46, 0x65, 0x74, 0x63, 0x68, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x2e,
	0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61, 0x78,
	0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x41,
	0x0a, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x12, 0x39, 0x0a, 0x0b, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x18, 0x18, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x0a, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x0e,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa7, 0x01, 0x0a,
	0x12, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x72, 0x52,
	0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x32, 0xef, 0x03, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b,
	0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extractionService_proto_rawDescOnce sync.Once
	file_extractionService_proto_rawDescData = file_extractionService_proto_rawDesc
)

func file_extractionService_proto_rawDescGZIP() []byte {
	file_extractionService_proto_rawDescOnce.Do(func() {
		file_extractionService_proto_rawDescData = protoimpl.X.CompressGZIP(file_extractionService_proto_rawDescData)
	})
	return file_extractionService_proto_rawDescData
}

var file_extractionService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_extractionService_proto_goTypes = []interface{}{
	(*ExtractionRequest)(nil),  // 0: services.ExtractionRequest
	(*Extraction)(nil),         // 1: services.Extraction
	(*ExtractionConfig)(nil),   // 2: services.ExtractionConfig
	(*ExtractionTime)(nil),     // 3: services.ExtractionTime
	(*ExtractionResponse)(nil), // 4: services.ExtractionResponse
	(*common.Pager)(nil),       // 5: common.Pager
}
var file_extractionService_proto_depIdxs = []int32{
	2,  // 0: services.Extraction.config:type_name -> services.ExtractionConfig
	3,  // 1: services.ExtractionConfig.reception_times:type_name -> services.ExtractionTime
	3,  // 2: services.ExtractionConfig.fetch_times:type_name -> services.ExtractionTime
	1,  // 3: services.ExtractionResponse.entity:type_name -> services.Extraction
	5,  // 4: services.ExtractionResponse.pager:type_name -> common.Pager
	1,  // 5: services.ExtractionResponse.items:type_name -> services.Extraction
	1,  // 6: services.ExtractionService.Create:input_type -> services.Extraction
	1,  // 7: services.ExtractionService.Update:input_type -> services.Extraction
	0,  // 8: services.ExtractionService.Delete:input_type -> services.ExtractionRequest
	1,  // 9: services.ExtractionService.Get:input_type -> services.Extraction
	0,  // 10: services.ExtractionService.Search:input_type -> services.ExtractionRequest
	0,  // 11: services.ExtractionService.List:input_type -> services.ExtractionRequest
	0,  // 12: services.ExtractionService.NearestList:input_type -> services.ExtractionRequest
	4,  // 13: services.ExtractionService.Create:output_type -> services.ExtractionResponse
	4,  // 14: services.ExtractionService.Update:output_type -> services.ExtractionResponse
	4,  // 15: services.ExtractionService.Delete:output_type -> services.ExtractionResponse
	4,  // 16: services.ExtractionService.Get:output_type -> services.ExtractionResponse
	4,  // 17: services.ExtractionService.Search:output_type -> services.ExtractionResponse
	4,  // 18: services.ExtractionService.List:output_type -> services.ExtractionResponse
	4,  // 19: services.ExtractionService.NearestList:output_type -> services.ExtractionResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_extractionService_proto_init() }
func file_extractionService_proto_init() {
	if File_extractionService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extractionService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractionRequest); i {
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
		file_extractionService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extraction); i {
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
		file_extractionService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractionConfig); i {
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
		file_extractionService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractionTime); i {
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
		file_extractionService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractionResponse); i {
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
			RawDescriptor: file_extractionService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_extractionService_proto_goTypes,
		DependencyIndexes: file_extractionService_proto_depIdxs,
		MessageInfos:      file_extractionService_proto_msgTypes,
	}.Build()
	File_extractionService_proto = out.File
	file_extractionService_proto_rawDesc = nil
	file_extractionService_proto_goTypes = nil
	file_extractionService_proto_depIdxs = nil
}
