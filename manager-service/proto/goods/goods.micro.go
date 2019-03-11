// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/areas/goods.proto

/*
Package goods is a generated protocol buffer package.

It is generated from these files:
	proto/areas/goods.proto

It has these top-level messages:
	Ids
	ReqSearch
	RespSearch
	GoodsModel
	Resp
*/
package goods

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Goods service

type GoodsService interface {
	Search(ctx context.Context, in *ReqSearch, opts ...client.CallOption) (*RespSearch, error)
	Delete(ctx context.Context, in *Ids, opts ...client.CallOption) (*Resp, error)
	UpdateStatus(ctx context.Context, in *Ids, opts ...client.CallOption) (*Resp, error)
}

type goodsService struct {
	c    client.Client
	name string
}

func NewGoodsService(name string, c client.Client) GoodsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "goods"
	}
	return &goodsService{
		c:    c,
		name: name,
	}
}

func (c *goodsService) Search(ctx context.Context, in *ReqSearch, opts ...client.CallOption) (*RespSearch, error) {
	req := c.c.NewRequest(c.name, "Goods.Search", in)
	out := new(RespSearch)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsService) Delete(ctx context.Context, in *Ids, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "Goods.Delete", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsService) UpdateStatus(ctx context.Context, in *Ids, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "Goods.UpdateStatus", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Goods service

type GoodsHandler interface {
	Search(context.Context, *ReqSearch, *RespSearch) error
	Delete(context.Context, *Ids, *Resp) error
	UpdateStatus(context.Context, *Ids, *Resp) error
}

func RegisterGoodsHandler(s server.Server, hdlr GoodsHandler, opts ...server.HandlerOption) error {
	type goods interface {
		Search(ctx context.Context, in *ReqSearch, out *RespSearch) error
		Delete(ctx context.Context, in *Ids, out *Resp) error
		UpdateStatus(ctx context.Context, in *Ids, out *Resp) error
	}
	type Goods struct {
		goods
	}
	h := &goodsHandler{hdlr}
	return s.Handle(s.NewHandler(&Goods{h}, opts...))
}

type goodsHandler struct {
	GoodsHandler
}

func (h *goodsHandler) Search(ctx context.Context, in *ReqSearch, out *RespSearch) error {
	return h.GoodsHandler.Search(ctx, in, out)
}

func (h *goodsHandler) Delete(ctx context.Context, in *Ids, out *Resp) error {
	return h.GoodsHandler.Delete(ctx, in, out)
}

func (h *goodsHandler) UpdateStatus(ctx context.Context, in *Ids, out *Resp) error {
	return h.GoodsHandler.UpdateStatus(ctx, in, out)
}