// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: orderAllotService.proto

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

// Api Endpoints for OrderAllotService service

func NewOrderAllotServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderAllotService service

type OrderAllotService interface {
	Create(ctx context.Context, in *OrderAllot, opts ...client.CallOption) (*OrderAllotResponse, error)
	Update(ctx context.Context, in *OrderAllot, opts ...client.CallOption) (*OrderAllotResponse, error)
	Delete(ctx context.Context, in *OrderAllot, opts ...client.CallOption) (*OrderAllotResponse, error)
	Detail(ctx context.Context, in *OrderAllot, opts ...client.CallOption) (*OrderAllotResponse, error)
	List(ctx context.Context, in *OrderAllotRequest, opts ...client.CallOption) (*OrderAllotResponse, error)
	Search(ctx context.Context, in *OrderAllotRequest, opts ...client.CallOption) (*OrderAllotResponse, error)
}

type orderAllotService struct {
	c    client.Client
	name string
}

func NewOrderAllotService(name string, c client.Client) OrderAllotService {
	return &orderAllotService{
		c:    c,
		name: name,
	}
}

func (c *orderAllotService) Create(ctx context.Context, in *OrderAllot, opts ...client.CallOption) (*OrderAllotResponse, error) {
	req := c.c.NewRequest(c.name, "OrderAllotService.Create", in)
	out := new(OrderAllotResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAllotService) Update(ctx context.Context, in *OrderAllot, opts ...client.CallOption) (*OrderAllotResponse, error) {
	req := c.c.NewRequest(c.name, "OrderAllotService.Update", in)
	out := new(OrderAllotResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAllotService) Delete(ctx context.Context, in *OrderAllot, opts ...client.CallOption) (*OrderAllotResponse, error) {
	req := c.c.NewRequest(c.name, "OrderAllotService.Delete", in)
	out := new(OrderAllotResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAllotService) Detail(ctx context.Context, in *OrderAllot, opts ...client.CallOption) (*OrderAllotResponse, error) {
	req := c.c.NewRequest(c.name, "OrderAllotService.Detail", in)
	out := new(OrderAllotResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAllotService) List(ctx context.Context, in *OrderAllotRequest, opts ...client.CallOption) (*OrderAllotResponse, error) {
	req := c.c.NewRequest(c.name, "OrderAllotService.List", in)
	out := new(OrderAllotResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAllotService) Search(ctx context.Context, in *OrderAllotRequest, opts ...client.CallOption) (*OrderAllotResponse, error) {
	req := c.c.NewRequest(c.name, "OrderAllotService.Search", in)
	out := new(OrderAllotResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderAllotService service

type OrderAllotServiceHandler interface {
	Create(context.Context, *OrderAllot, *OrderAllotResponse) error
	Update(context.Context, *OrderAllot, *OrderAllotResponse) error
	Delete(context.Context, *OrderAllot, *OrderAllotResponse) error
	Detail(context.Context, *OrderAllot, *OrderAllotResponse) error
	List(context.Context, *OrderAllotRequest, *OrderAllotResponse) error
	Search(context.Context, *OrderAllotRequest, *OrderAllotResponse) error
}

func RegisterOrderAllotServiceHandler(s server.Server, hdlr OrderAllotServiceHandler, opts ...server.HandlerOption) error {
	type orderAllotService interface {
		Create(ctx context.Context, in *OrderAllot, out *OrderAllotResponse) error
		Update(ctx context.Context, in *OrderAllot, out *OrderAllotResponse) error
		Delete(ctx context.Context, in *OrderAllot, out *OrderAllotResponse) error
		Detail(ctx context.Context, in *OrderAllot, out *OrderAllotResponse) error
		List(ctx context.Context, in *OrderAllotRequest, out *OrderAllotResponse) error
		Search(ctx context.Context, in *OrderAllotRequest, out *OrderAllotResponse) error
	}
	type OrderAllotService struct {
		orderAllotService
	}
	h := &orderAllotServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderAllotService{h}, opts...))
}

type orderAllotServiceHandler struct {
	OrderAllotServiceHandler
}

func (h *orderAllotServiceHandler) Create(ctx context.Context, in *OrderAllot, out *OrderAllotResponse) error {
	return h.OrderAllotServiceHandler.Create(ctx, in, out)
}

func (h *orderAllotServiceHandler) Update(ctx context.Context, in *OrderAllot, out *OrderAllotResponse) error {
	return h.OrderAllotServiceHandler.Update(ctx, in, out)
}

func (h *orderAllotServiceHandler) Delete(ctx context.Context, in *OrderAllot, out *OrderAllotResponse) error {
	return h.OrderAllotServiceHandler.Delete(ctx, in, out)
}

func (h *orderAllotServiceHandler) Detail(ctx context.Context, in *OrderAllot, out *OrderAllotResponse) error {
	return h.OrderAllotServiceHandler.Detail(ctx, in, out)
}

func (h *orderAllotServiceHandler) List(ctx context.Context, in *OrderAllotRequest, out *OrderAllotResponse) error {
	return h.OrderAllotServiceHandler.List(ctx, in, out)
}

func (h *orderAllotServiceHandler) Search(ctx context.Context, in *OrderAllotRequest, out *OrderAllotResponse) error {
	return h.OrderAllotServiceHandler.Search(ctx, in, out)
}