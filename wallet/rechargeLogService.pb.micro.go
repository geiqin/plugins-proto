// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rechargeLogService.proto

package services

import (
	fmt "fmt"
	_ "github.com/geiqin/micro-kit/protobuf/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for RechargeLogService service

func NewRechargeLogServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RechargeLogService service

type RechargeLogService interface {
	//充值-创建
	Create(ctx context.Context, in *RechargeLog, opts ...client.CallOption) (*RechargeLogResponse, error)
	//充值-提交支付
	Pay(ctx context.Context, in *RechargeLog, opts ...client.CallOption) (*RechargeLogResponse, error)
	//充值-详情
	Detail(ctx context.Context, in *RechargeLog, opts ...client.CallOption) (*RechargeLogResponse, error)
	//充值删除
	Delete(ctx context.Context, in *RechargeLog, opts ...client.CallOption) (*RechargeLogResponse, error)
	//获取充值信息
	Get(ctx context.Context, in *RechargeLog, opts ...client.CallOption) (*RechargeLogResponse, error)
	//查询充值列表
	Search(ctx context.Context, in *RechargeLogRequest, opts ...client.CallOption) (*RechargeLogResponse, error)
}

type rechargeLogService struct {
	c    client.Client
	name string
}

func NewRechargeLogService(name string, c client.Client) RechargeLogService {
	return &rechargeLogService{
		c:    c,
		name: name,
	}
}

func (c *rechargeLogService) Create(ctx context.Context, in *RechargeLog, opts ...client.CallOption) (*RechargeLogResponse, error) {
	req := c.c.NewRequest(c.name, "RechargeLogService.Create", in)
	out := new(RechargeLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rechargeLogService) Pay(ctx context.Context, in *RechargeLog, opts ...client.CallOption) (*RechargeLogResponse, error) {
	req := c.c.NewRequest(c.name, "RechargeLogService.Pay", in)
	out := new(RechargeLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rechargeLogService) Detail(ctx context.Context, in *RechargeLog, opts ...client.CallOption) (*RechargeLogResponse, error) {
	req := c.c.NewRequest(c.name, "RechargeLogService.Detail", in)
	out := new(RechargeLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rechargeLogService) Delete(ctx context.Context, in *RechargeLog, opts ...client.CallOption) (*RechargeLogResponse, error) {
	req := c.c.NewRequest(c.name, "RechargeLogService.Delete", in)
	out := new(RechargeLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rechargeLogService) Get(ctx context.Context, in *RechargeLog, opts ...client.CallOption) (*RechargeLogResponse, error) {
	req := c.c.NewRequest(c.name, "RechargeLogService.Get", in)
	out := new(RechargeLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rechargeLogService) Search(ctx context.Context, in *RechargeLogRequest, opts ...client.CallOption) (*RechargeLogResponse, error) {
	req := c.c.NewRequest(c.name, "RechargeLogService.Search", in)
	out := new(RechargeLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RechargeLogService service

type RechargeLogServiceHandler interface {
	//充值-创建
	Create(context.Context, *RechargeLog, *RechargeLogResponse) error
	//充值-提交支付
	Pay(context.Context, *RechargeLog, *RechargeLogResponse) error
	//充值-详情
	Detail(context.Context, *RechargeLog, *RechargeLogResponse) error
	//充值删除
	Delete(context.Context, *RechargeLog, *RechargeLogResponse) error
	//获取充值信息
	Get(context.Context, *RechargeLog, *RechargeLogResponse) error
	//查询充值列表
	Search(context.Context, *RechargeLogRequest, *RechargeLogResponse) error
}

func RegisterRechargeLogServiceHandler(s server.Server, hdlr RechargeLogServiceHandler, opts ...server.HandlerOption) error {
	type rechargeLogService interface {
		Create(ctx context.Context, in *RechargeLog, out *RechargeLogResponse) error
		Pay(ctx context.Context, in *RechargeLog, out *RechargeLogResponse) error
		Detail(ctx context.Context, in *RechargeLog, out *RechargeLogResponse) error
		Delete(ctx context.Context, in *RechargeLog, out *RechargeLogResponse) error
		Get(ctx context.Context, in *RechargeLog, out *RechargeLogResponse) error
		Search(ctx context.Context, in *RechargeLogRequest, out *RechargeLogResponse) error
	}
	type RechargeLogService struct {
		rechargeLogService
	}
	h := &rechargeLogServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RechargeLogService{h}, opts...))
}

type rechargeLogServiceHandler struct {
	RechargeLogServiceHandler
}

func (h *rechargeLogServiceHandler) Create(ctx context.Context, in *RechargeLog, out *RechargeLogResponse) error {
	return h.RechargeLogServiceHandler.Create(ctx, in, out)
}

func (h *rechargeLogServiceHandler) Pay(ctx context.Context, in *RechargeLog, out *RechargeLogResponse) error {
	return h.RechargeLogServiceHandler.Pay(ctx, in, out)
}

func (h *rechargeLogServiceHandler) Detail(ctx context.Context, in *RechargeLog, out *RechargeLogResponse) error {
	return h.RechargeLogServiceHandler.Detail(ctx, in, out)
}

func (h *rechargeLogServiceHandler) Delete(ctx context.Context, in *RechargeLog, out *RechargeLogResponse) error {
	return h.RechargeLogServiceHandler.Delete(ctx, in, out)
}

func (h *rechargeLogServiceHandler) Get(ctx context.Context, in *RechargeLog, out *RechargeLogResponse) error {
	return h.RechargeLogServiceHandler.Get(ctx, in, out)
}

func (h *rechargeLogServiceHandler) Search(ctx context.Context, in *RechargeLogRequest, out *RechargeLogResponse) error {
	return h.RechargeLogServiceHandler.Search(ctx, in, out)
}
