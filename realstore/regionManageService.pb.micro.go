// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: regionManageService.proto

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

// Api Endpoints for RegionManageService service

func NewRegionManageServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RegionManageService service

type RegionManageService interface {
	Create(ctx context.Context, in *RegionManage, opts ...client.CallOption) (*RegionManageResponse, error)
	Update(ctx context.Context, in *RegionManage, opts ...client.CallOption) (*RegionManageResponse, error)
	Delete(ctx context.Context, in *RegionManage, opts ...client.CallOption) (*RegionManageResponse, error)
	Detail(ctx context.Context, in *RegionManage, opts ...client.CallOption) (*RegionManageResponse, error)
	List(ctx context.Context, in *RegionManageRequest, opts ...client.CallOption) (*RegionManageResponse, error)
	Search(ctx context.Context, in *RegionManageRequest, opts ...client.CallOption) (*RegionManageResponse, error)
}

type regionManageService struct {
	c    client.Client
	name string
}

func NewRegionManageService(name string, c client.Client) RegionManageService {
	return &regionManageService{
		c:    c,
		name: name,
	}
}

func (c *regionManageService) Create(ctx context.Context, in *RegionManage, opts ...client.CallOption) (*RegionManageResponse, error) {
	req := c.c.NewRequest(c.name, "RegionManageService.Create", in)
	out := new(RegionManageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionManageService) Update(ctx context.Context, in *RegionManage, opts ...client.CallOption) (*RegionManageResponse, error) {
	req := c.c.NewRequest(c.name, "RegionManageService.Update", in)
	out := new(RegionManageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionManageService) Delete(ctx context.Context, in *RegionManage, opts ...client.CallOption) (*RegionManageResponse, error) {
	req := c.c.NewRequest(c.name, "RegionManageService.Delete", in)
	out := new(RegionManageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionManageService) Detail(ctx context.Context, in *RegionManage, opts ...client.CallOption) (*RegionManageResponse, error) {
	req := c.c.NewRequest(c.name, "RegionManageService.Detail", in)
	out := new(RegionManageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionManageService) List(ctx context.Context, in *RegionManageRequest, opts ...client.CallOption) (*RegionManageResponse, error) {
	req := c.c.NewRequest(c.name, "RegionManageService.List", in)
	out := new(RegionManageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionManageService) Search(ctx context.Context, in *RegionManageRequest, opts ...client.CallOption) (*RegionManageResponse, error) {
	req := c.c.NewRequest(c.name, "RegionManageService.Search", in)
	out := new(RegionManageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RegionManageService service

type RegionManageServiceHandler interface {
	Create(context.Context, *RegionManage, *RegionManageResponse) error
	Update(context.Context, *RegionManage, *RegionManageResponse) error
	Delete(context.Context, *RegionManage, *RegionManageResponse) error
	Detail(context.Context, *RegionManage, *RegionManageResponse) error
	List(context.Context, *RegionManageRequest, *RegionManageResponse) error
	Search(context.Context, *RegionManageRequest, *RegionManageResponse) error
}

func RegisterRegionManageServiceHandler(s server.Server, hdlr RegionManageServiceHandler, opts ...server.HandlerOption) error {
	type regionManageService interface {
		Create(ctx context.Context, in *RegionManage, out *RegionManageResponse) error
		Update(ctx context.Context, in *RegionManage, out *RegionManageResponse) error
		Delete(ctx context.Context, in *RegionManage, out *RegionManageResponse) error
		Detail(ctx context.Context, in *RegionManage, out *RegionManageResponse) error
		List(ctx context.Context, in *RegionManageRequest, out *RegionManageResponse) error
		Search(ctx context.Context, in *RegionManageRequest, out *RegionManageResponse) error
	}
	type RegionManageService struct {
		regionManageService
	}
	h := &regionManageServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RegionManageService{h}, opts...))
}

type regionManageServiceHandler struct {
	RegionManageServiceHandler
}

func (h *regionManageServiceHandler) Create(ctx context.Context, in *RegionManage, out *RegionManageResponse) error {
	return h.RegionManageServiceHandler.Create(ctx, in, out)
}

func (h *regionManageServiceHandler) Update(ctx context.Context, in *RegionManage, out *RegionManageResponse) error {
	return h.RegionManageServiceHandler.Update(ctx, in, out)
}

func (h *regionManageServiceHandler) Delete(ctx context.Context, in *RegionManage, out *RegionManageResponse) error {
	return h.RegionManageServiceHandler.Delete(ctx, in, out)
}

func (h *regionManageServiceHandler) Detail(ctx context.Context, in *RegionManage, out *RegionManageResponse) error {
	return h.RegionManageServiceHandler.Detail(ctx, in, out)
}

func (h *regionManageServiceHandler) List(ctx context.Context, in *RegionManageRequest, out *RegionManageResponse) error {
	return h.RegionManageServiceHandler.List(ctx, in, out)
}

func (h *regionManageServiceHandler) Search(ctx context.Context, in *RegionManageRequest, out *RegionManageResponse) error {
	return h.RegionManageServiceHandler.Search(ctx, in, out)
}