// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: traceLogService.proto

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

// Api Endpoints for TraceLogService service

func NewTraceLogServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TraceLogService service

type TraceLogService interface {
	//物流信息详情
	Detail(ctx context.Context, in *TraceLog, opts ...client.CallOption) (*TraceLogResponse, error)
	//查询物流信息
	Search(ctx context.Context, in *TraceLogRequest, opts ...client.CallOption) (*TraceLogResponse, error)
}

type traceLogService struct {
	c    client.Client
	name string
}

func NewTraceLogService(name string, c client.Client) TraceLogService {
	return &traceLogService{
		c:    c,
		name: name,
	}
}

func (c *traceLogService) Detail(ctx context.Context, in *TraceLog, opts ...client.CallOption) (*TraceLogResponse, error) {
	req := c.c.NewRequest(c.name, "TraceLogService.Detail", in)
	out := new(TraceLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceLogService) Search(ctx context.Context, in *TraceLogRequest, opts ...client.CallOption) (*TraceLogResponse, error) {
	req := c.c.NewRequest(c.name, "TraceLogService.Search", in)
	out := new(TraceLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TraceLogService service

type TraceLogServiceHandler interface {
	//物流信息详情
	Detail(context.Context, *TraceLog, *TraceLogResponse) error
	//查询物流信息
	Search(context.Context, *TraceLogRequest, *TraceLogResponse) error
}

func RegisterTraceLogServiceHandler(s server.Server, hdlr TraceLogServiceHandler, opts ...server.HandlerOption) error {
	type traceLogService interface {
		Detail(ctx context.Context, in *TraceLog, out *TraceLogResponse) error
		Search(ctx context.Context, in *TraceLogRequest, out *TraceLogResponse) error
	}
	type TraceLogService struct {
		traceLogService
	}
	h := &traceLogServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TraceLogService{h}, opts...))
}

type traceLogServiceHandler struct {
	TraceLogServiceHandler
}

func (h *traceLogServiceHandler) Detail(ctx context.Context, in *TraceLog, out *TraceLogResponse) error {
	return h.TraceLogServiceHandler.Detail(ctx, in, out)
}

func (h *traceLogServiceHandler) Search(ctx context.Context, in *TraceLogRequest, out *TraceLogResponse) error {
	return h.TraceLogServiceHandler.Search(ctx, in, out)
}