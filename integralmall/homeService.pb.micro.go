// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: homeService.proto

package services

import (
	fmt "fmt"
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

// Api Endpoints for HomeService service

func NewHomeServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for HomeService service

type HomeService interface {
	Index(ctx context.Context, in *HomeRequest, opts ...client.CallOption) (*HomeResponse, error)
}

type homeService struct {
	c    client.Client
	name string
}

func NewHomeService(name string, c client.Client) HomeService {
	return &homeService{
		c:    c,
		name: name,
	}
}

func (c *homeService) Index(ctx context.Context, in *HomeRequest, opts ...client.CallOption) (*HomeResponse, error) {
	req := c.c.NewRequest(c.name, "HomeService.Index", in)
	out := new(HomeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HomeService service

type HomeServiceHandler interface {
	Index(context.Context, *HomeRequest, *HomeResponse) error
}

func RegisterHomeServiceHandler(s server.Server, hdlr HomeServiceHandler, opts ...server.HandlerOption) error {
	type homeService interface {
		Index(ctx context.Context, in *HomeRequest, out *HomeResponse) error
	}
	type HomeService struct {
		homeService
	}
	h := &homeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&HomeService{h}, opts...))
}

type homeServiceHandler struct {
	HomeServiceHandler
}

func (h *homeServiceHandler) Index(ctx context.Context, in *HomeRequest, out *HomeResponse) error {
	return h.HomeServiceHandler.Index(ctx, in, out)
}
