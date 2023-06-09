// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: extractionOrderService.proto

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

// Api Endpoints for ExtractionOrderService service

func NewExtractionOrderServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ExtractionOrderService service

type ExtractionOrderService interface {
	Create(ctx context.Context, in *ExtractionOrder, opts ...client.CallOption) (*ExtractionOrderResponse, error)
	Take(ctx context.Context, in *ExtractionOrder, opts ...client.CallOption) (*ExtractionOrderResponse, error)
	Detail(ctx context.Context, in *ExtractionOrder, opts ...client.CallOption) (*ExtractionOrderResponse, error)
	Search(ctx context.Context, in *ExtractionOrderRequest, opts ...client.CallOption) (*ExtractionOrderResponse, error)
}

type extractionOrderService struct {
	c    client.Client
	name string
}

func NewExtractionOrderService(name string, c client.Client) ExtractionOrderService {
	return &extractionOrderService{
		c:    c,
		name: name,
	}
}

func (c *extractionOrderService) Create(ctx context.Context, in *ExtractionOrder, opts ...client.CallOption) (*ExtractionOrderResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionOrderService.Create", in)
	out := new(ExtractionOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionOrderService) Take(ctx context.Context, in *ExtractionOrder, opts ...client.CallOption) (*ExtractionOrderResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionOrderService.Take", in)
	out := new(ExtractionOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionOrderService) Detail(ctx context.Context, in *ExtractionOrder, opts ...client.CallOption) (*ExtractionOrderResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionOrderService.Detail", in)
	out := new(ExtractionOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extractionOrderService) Search(ctx context.Context, in *ExtractionOrderRequest, opts ...client.CallOption) (*ExtractionOrderResponse, error) {
	req := c.c.NewRequest(c.name, "ExtractionOrderService.Search", in)
	out := new(ExtractionOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExtractionOrderService service

type ExtractionOrderServiceHandler interface {
	Create(context.Context, *ExtractionOrder, *ExtractionOrderResponse) error
	Take(context.Context, *ExtractionOrder, *ExtractionOrderResponse) error
	Detail(context.Context, *ExtractionOrder, *ExtractionOrderResponse) error
	Search(context.Context, *ExtractionOrderRequest, *ExtractionOrderResponse) error
}

func RegisterExtractionOrderServiceHandler(s server.Server, hdlr ExtractionOrderServiceHandler, opts ...server.HandlerOption) error {
	type extractionOrderService interface {
		Create(ctx context.Context, in *ExtractionOrder, out *ExtractionOrderResponse) error
		Take(ctx context.Context, in *ExtractionOrder, out *ExtractionOrderResponse) error
		Detail(ctx context.Context, in *ExtractionOrder, out *ExtractionOrderResponse) error
		Search(ctx context.Context, in *ExtractionOrderRequest, out *ExtractionOrderResponse) error
	}
	type ExtractionOrderService struct {
		extractionOrderService
	}
	h := &extractionOrderServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ExtractionOrderService{h}, opts...))
}

type extractionOrderServiceHandler struct {
	ExtractionOrderServiceHandler
}

func (h *extractionOrderServiceHandler) Create(ctx context.Context, in *ExtractionOrder, out *ExtractionOrderResponse) error {
	return h.ExtractionOrderServiceHandler.Create(ctx, in, out)
}

func (h *extractionOrderServiceHandler) Take(ctx context.Context, in *ExtractionOrder, out *ExtractionOrderResponse) error {
	return h.ExtractionOrderServiceHandler.Take(ctx, in, out)
}

func (h *extractionOrderServiceHandler) Detail(ctx context.Context, in *ExtractionOrder, out *ExtractionOrderResponse) error {
	return h.ExtractionOrderServiceHandler.Detail(ctx, in, out)
}

func (h *extractionOrderServiceHandler) Search(ctx context.Context, in *ExtractionOrderRequest, out *ExtractionOrderResponse) error {
	return h.ExtractionOrderServiceHandler.Search(ctx, in, out)
}
