// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderGiveIntegralLogService.proto

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

// Api Endpoints for OrderGiveIntegralLogService service

func NewOrderGiveIntegralLogServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderGiveIntegralLogService service

type OrderGiveIntegralLogService interface {
	Detail(ctx context.Context, in *OrderGiveIntegralLog, opts ...client.CallOption) (*OrderGiveIntegralLogResponse, error)
	Search(ctx context.Context, in *OrderGiveIntegralLogRequest, opts ...client.CallOption) (*OrderGiveIntegralLogResponse, error)
}

type orderGiveIntegralLogService struct {
	c    client.Client
	name string
}

func NewOrderGiveIntegralLogService(name string, c client.Client) OrderGiveIntegralLogService {
	return &orderGiveIntegralLogService{
		c:    c,
		name: name,
	}
}

func (c *orderGiveIntegralLogService) Detail(ctx context.Context, in *OrderGiveIntegralLog, opts ...client.CallOption) (*OrderGiveIntegralLogResponse, error) {
	req := c.c.NewRequest(c.name, "OrderGiveIntegralLogService.Detail", in)
	out := new(OrderGiveIntegralLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderGiveIntegralLogService) Search(ctx context.Context, in *OrderGiveIntegralLogRequest, opts ...client.CallOption) (*OrderGiveIntegralLogResponse, error) {
	req := c.c.NewRequest(c.name, "OrderGiveIntegralLogService.Search", in)
	out := new(OrderGiveIntegralLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderGiveIntegralLogService service

type OrderGiveIntegralLogServiceHandler interface {
	Detail(context.Context, *OrderGiveIntegralLog, *OrderGiveIntegralLogResponse) error
	Search(context.Context, *OrderGiveIntegralLogRequest, *OrderGiveIntegralLogResponse) error
}

func RegisterOrderGiveIntegralLogServiceHandler(s server.Server, hdlr OrderGiveIntegralLogServiceHandler, opts ...server.HandlerOption) error {
	type orderGiveIntegralLogService interface {
		Detail(ctx context.Context, in *OrderGiveIntegralLog, out *OrderGiveIntegralLogResponse) error
		Search(ctx context.Context, in *OrderGiveIntegralLogRequest, out *OrderGiveIntegralLogResponse) error
	}
	type OrderGiveIntegralLogService struct {
		orderGiveIntegralLogService
	}
	h := &orderGiveIntegralLogServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderGiveIntegralLogService{h}, opts...))
}

type orderGiveIntegralLogServiceHandler struct {
	OrderGiveIntegralLogServiceHandler
}

func (h *orderGiveIntegralLogServiceHandler) Detail(ctx context.Context, in *OrderGiveIntegralLog, out *OrderGiveIntegralLogResponse) error {
	return h.OrderGiveIntegralLogServiceHandler.Detail(ctx, in, out)
}

func (h *orderGiveIntegralLogServiceHandler) Search(ctx context.Context, in *OrderGiveIntegralLogRequest, out *OrderGiveIntegralLogResponse) error {
	return h.OrderGiveIntegralLogServiceHandler.Search(ctx, in, out)
}