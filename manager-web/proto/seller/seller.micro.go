// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/areas/seller.proto

/*
Package seller is a generated protocol buffer package.

It is generated from these files:
	proto/areas/seller.proto

It has these top-level messages:
	ReqSeller
	RespSeller
	ReqId
	Rows
	ReqIdAndStatus
	RespReturn
*/
package seller

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

// Client API for Seller service

type SellerService interface {
	//    rpc MicroSeller(ReqSeller)returns(RespSeller){}
	Search(ctx context.Context, in *ReqSeller, opts ...client.CallOption) (*RespSeller, error)
	FindOne(ctx context.Context, in *ReqId, opts ...client.CallOption) (*Rows, error)
	UpdateStatus(ctx context.Context, in *ReqIdAndStatus, opts ...client.CallOption) (*RespReturn, error)
}

type sellerService struct {
	c    client.Client
	name string
}

func NewSellerService(name string, c client.Client) SellerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "seller"
	}
	return &sellerService{
		c:    c,
		name: name,
	}
}

func (c *sellerService) Search(ctx context.Context, in *ReqSeller, opts ...client.CallOption) (*RespSeller, error) {
	req := c.c.NewRequest(c.name, "Seller.Search", in)
	out := new(RespSeller)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerService) FindOne(ctx context.Context, in *ReqId, opts ...client.CallOption) (*Rows, error) {
	req := c.c.NewRequest(c.name, "Seller.FindOne", in)
	out := new(Rows)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerService) UpdateStatus(ctx context.Context, in *ReqIdAndStatus, opts ...client.CallOption) (*RespReturn, error) {
	req := c.c.NewRequest(c.name, "Seller.UpdateStatus", in)
	out := new(RespReturn)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Seller service

type SellerHandler interface {
	//    rpc MicroSeller(ReqSeller)returns(RespSeller){}
	Search(context.Context, *ReqSeller, *RespSeller) error
	FindOne(context.Context, *ReqId, *Rows) error
	UpdateStatus(context.Context, *ReqIdAndStatus, *RespReturn) error
}

func RegisterSellerHandler(s server.Server, hdlr SellerHandler, opts ...server.HandlerOption) error {
	type seller interface {
		Search(ctx context.Context, in *ReqSeller, out *RespSeller) error
		FindOne(ctx context.Context, in *ReqId, out *Rows) error
		UpdateStatus(ctx context.Context, in *ReqIdAndStatus, out *RespReturn) error
	}
	type Seller struct {
		seller
	}
	h := &sellerHandler{hdlr}
	return s.Handle(s.NewHandler(&Seller{h}, opts...))
}

type sellerHandler struct {
	SellerHandler
}

func (h *sellerHandler) Search(ctx context.Context, in *ReqSeller, out *RespSeller) error {
	return h.SellerHandler.Search(ctx, in, out)
}

func (h *sellerHandler) FindOne(ctx context.Context, in *ReqId, out *Rows) error {
	return h.SellerHandler.FindOne(ctx, in, out)
}

func (h *sellerHandler) UpdateStatus(ctx context.Context, in *ReqIdAndStatus, out *RespReturn) error {
	return h.SellerHandler.UpdateStatus(ctx, in, out)
}
