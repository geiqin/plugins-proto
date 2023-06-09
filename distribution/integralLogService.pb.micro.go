// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: integralLogService.proto

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

// Api Endpoints for IntegralLogService service

func NewIntegralLogServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for IntegralLogService service

type IntegralLogService interface {
	// 获得积分详情
	Detail(ctx context.Context, in *IntegralLog, opts ...client.CallOption) (*IntegralLogResponse, error)
	// 查询积分列表
	Search(ctx context.Context, in *IntegralLogRequest, opts ...client.CallOption) (*IntegralLogResponse, error)
}

type integralLogService struct {
	c    client.Client
	name string
}

func NewIntegralLogService(name string, c client.Client) IntegralLogService {
	return &integralLogService{
		c:    c,
		name: name,
	}
}

func (c *integralLogService) Detail(ctx context.Context, in *IntegralLog, opts ...client.CallOption) (*IntegralLogResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralLogService.Detail", in)
	out := new(IntegralLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integralLogService) Search(ctx context.Context, in *IntegralLogRequest, opts ...client.CallOption) (*IntegralLogResponse, error) {
	req := c.c.NewRequest(c.name, "IntegralLogService.Search", in)
	out := new(IntegralLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IntegralLogService service

type IntegralLogServiceHandler interface {
	// 获得积分详情
	Detail(context.Context, *IntegralLog, *IntegralLogResponse) error
	// 查询积分列表
	Search(context.Context, *IntegralLogRequest, *IntegralLogResponse) error
}

func RegisterIntegralLogServiceHandler(s server.Server, hdlr IntegralLogServiceHandler, opts ...server.HandlerOption) error {
	type integralLogService interface {
		Detail(ctx context.Context, in *IntegralLog, out *IntegralLogResponse) error
		Search(ctx context.Context, in *IntegralLogRequest, out *IntegralLogResponse) error
	}
	type IntegralLogService struct {
		integralLogService
	}
	h := &integralLogServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&IntegralLogService{h}, opts...))
}

type integralLogServiceHandler struct {
	IntegralLogServiceHandler
}

func (h *integralLogServiceHandler) Detail(ctx context.Context, in *IntegralLog, out *IntegralLogResponse) error {
	return h.IntegralLogServiceHandler.Detail(ctx, in, out)
}

func (h *integralLogServiceHandler) Search(ctx context.Context, in *IntegralLogRequest, out *IntegralLogResponse) error {
	return h.IntegralLogServiceHandler.Search(ctx, in, out)
}
