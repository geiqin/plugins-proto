// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cosumeService.proto

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

// Api Endpoints for ConsumeService service

func NewConsumeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ConsumeService service

type ConsumeService interface {
	Create(ctx context.Context, in *Consume, opts ...client.CallOption) (*ConsumeResponse, error)
	Update(ctx context.Context, in *Consume, opts ...client.CallOption) (*ConsumeResponse, error)
	Delete(ctx context.Context, in *Consume, opts ...client.CallOption) (*ConsumeResponse, error)
	Switch(ctx context.Context, in *Consume, opts ...client.CallOption) (*ConsumeResponse, error)
	Detail(ctx context.Context, in *Consume, opts ...client.CallOption) (*ConsumeResponse, error)
	Search(ctx context.Context, in *ConsumeRequest, opts ...client.CallOption) (*ConsumeResponse, error)
	List(ctx context.Context, in *ConsumeRequest, opts ...client.CallOption) (*ConsumeResponse, error)
}

type consumeService struct {
	c    client.Client
	name string
}

func NewConsumeService(name string, c client.Client) ConsumeService {
	return &consumeService{
		c:    c,
		name: name,
	}
}

func (c *consumeService) Create(ctx context.Context, in *Consume, opts ...client.CallOption) (*ConsumeResponse, error) {
	req := c.c.NewRequest(c.name, "ConsumeService.Create", in)
	out := new(ConsumeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumeService) Update(ctx context.Context, in *Consume, opts ...client.CallOption) (*ConsumeResponse, error) {
	req := c.c.NewRequest(c.name, "ConsumeService.Update", in)
	out := new(ConsumeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumeService) Delete(ctx context.Context, in *Consume, opts ...client.CallOption) (*ConsumeResponse, error) {
	req := c.c.NewRequest(c.name, "ConsumeService.Delete", in)
	out := new(ConsumeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumeService) Switch(ctx context.Context, in *Consume, opts ...client.CallOption) (*ConsumeResponse, error) {
	req := c.c.NewRequest(c.name, "ConsumeService.Switch", in)
	out := new(ConsumeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumeService) Detail(ctx context.Context, in *Consume, opts ...client.CallOption) (*ConsumeResponse, error) {
	req := c.c.NewRequest(c.name, "ConsumeService.Detail", in)
	out := new(ConsumeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumeService) Search(ctx context.Context, in *ConsumeRequest, opts ...client.CallOption) (*ConsumeResponse, error) {
	req := c.c.NewRequest(c.name, "ConsumeService.Search", in)
	out := new(ConsumeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumeService) List(ctx context.Context, in *ConsumeRequest, opts ...client.CallOption) (*ConsumeResponse, error) {
	req := c.c.NewRequest(c.name, "ConsumeService.List", in)
	out := new(ConsumeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConsumeService service

type ConsumeServiceHandler interface {
	Create(context.Context, *Consume, *ConsumeResponse) error
	Update(context.Context, *Consume, *ConsumeResponse) error
	Delete(context.Context, *Consume, *ConsumeResponse) error
	Switch(context.Context, *Consume, *ConsumeResponse) error
	Detail(context.Context, *Consume, *ConsumeResponse) error
	Search(context.Context, *ConsumeRequest, *ConsumeResponse) error
	List(context.Context, *ConsumeRequest, *ConsumeResponse) error
}

func RegisterConsumeServiceHandler(s server.Server, hdlr ConsumeServiceHandler, opts ...server.HandlerOption) error {
	type consumeService interface {
		Create(ctx context.Context, in *Consume, out *ConsumeResponse) error
		Update(ctx context.Context, in *Consume, out *ConsumeResponse) error
		Delete(ctx context.Context, in *Consume, out *ConsumeResponse) error
		Switch(ctx context.Context, in *Consume, out *ConsumeResponse) error
		Detail(ctx context.Context, in *Consume, out *ConsumeResponse) error
		Search(ctx context.Context, in *ConsumeRequest, out *ConsumeResponse) error
		List(ctx context.Context, in *ConsumeRequest, out *ConsumeResponse) error
	}
	type ConsumeService struct {
		consumeService
	}
	h := &consumeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ConsumeService{h}, opts...))
}

type consumeServiceHandler struct {
	ConsumeServiceHandler
}

func (h *consumeServiceHandler) Create(ctx context.Context, in *Consume, out *ConsumeResponse) error {
	return h.ConsumeServiceHandler.Create(ctx, in, out)
}

func (h *consumeServiceHandler) Update(ctx context.Context, in *Consume, out *ConsumeResponse) error {
	return h.ConsumeServiceHandler.Update(ctx, in, out)
}

func (h *consumeServiceHandler) Delete(ctx context.Context, in *Consume, out *ConsumeResponse) error {
	return h.ConsumeServiceHandler.Delete(ctx, in, out)
}

func (h *consumeServiceHandler) Switch(ctx context.Context, in *Consume, out *ConsumeResponse) error {
	return h.ConsumeServiceHandler.Switch(ctx, in, out)
}

func (h *consumeServiceHandler) Detail(ctx context.Context, in *Consume, out *ConsumeResponse) error {
	return h.ConsumeServiceHandler.Detail(ctx, in, out)
}

func (h *consumeServiceHandler) Search(ctx context.Context, in *ConsumeRequest, out *ConsumeResponse) error {
	return h.ConsumeServiceHandler.Search(ctx, in, out)
}

func (h *consumeServiceHandler) List(ctx context.Context, in *ConsumeRequest, out *ConsumeResponse) error {
	return h.ConsumeServiceHandler.List(ctx, in, out)
}