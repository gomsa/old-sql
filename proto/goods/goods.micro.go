// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/goods/goods.proto

package goods

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Goodss service

type GoodssService interface {
	// 获取用户列表
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type goodssService struct {
	c    client.Client
	name string
}

func NewGoodssService(name string, c client.Client) GoodssService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "goods"
	}
	return &goodssService{
		c:    c,
		name: name,
	}
}

func (c *goodssService) List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Goodss.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Goodss service

type GoodssHandler interface {
	// 获取用户列表
	List(context.Context, *Request, *Response) error
}

func RegisterGoodssHandler(s server.Server, hdlr GoodssHandler, opts ...server.HandlerOption) error {
	type goodss interface {
		List(ctx context.Context, in *Request, out *Response) error
	}
	type Goodss struct {
		goodss
	}
	h := &goodssHandler{hdlr}
	return s.Handle(s.NewHandler(&Goodss{h}, opts...))
}

type goodssHandler struct {
	GoodssHandler
}

func (h *goodssHandler) List(ctx context.Context, in *Request, out *Response) error {
	return h.GoodssHandler.List(ctx, in, out)
}