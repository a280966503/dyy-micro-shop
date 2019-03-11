// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/areas/content.proto

/*
Package content is a generated protocol buffer package.

It is generated from these files:
	proto/areas/content.proto

It has these top-level messages:
	Ids
	ReqContent
	RespContent
	ContentModel
	Resp
*/
package content

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

// Client API for Content service

type ContentService interface {
	Search(ctx context.Context, in *ReqContent, opts ...client.CallOption) (*RespContent, error)
	Add(ctx context.Context, in *ContentModel, opts ...client.CallOption) (*Resp, error)
	Delete(ctx context.Context, in *Ids, opts ...client.CallOption) (*Resp, error)
	FindOne(ctx context.Context, in *ContentModel, opts ...client.CallOption) (*ContentModel, error)
	Update(ctx context.Context, in *ContentModel, opts ...client.CallOption) (*Resp, error)
}

type contentService struct {
	c    client.Client
	name string
}

func NewContentService(name string, c client.Client) ContentService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "content"
	}
	return &contentService{
		c:    c,
		name: name,
	}
}

func (c *contentService) Search(ctx context.Context, in *ReqContent, opts ...client.CallOption) (*RespContent, error) {
	req := c.c.NewRequest(c.name, "Content.Search", in)
	out := new(RespContent)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentService) Add(ctx context.Context, in *ContentModel, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "Content.Add", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentService) Delete(ctx context.Context, in *Ids, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "Content.Delete", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentService) FindOne(ctx context.Context, in *ContentModel, opts ...client.CallOption) (*ContentModel, error) {
	req := c.c.NewRequest(c.name, "Content.FindOne", in)
	out := new(ContentModel)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentService) Update(ctx context.Context, in *ContentModel, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "Content.Update", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Content service

type ContentHandler interface {
	Search(context.Context, *ReqContent, *RespContent) error
	Add(context.Context, *ContentModel, *Resp) error
	Delete(context.Context, *Ids, *Resp) error
	FindOne(context.Context, *ContentModel, *ContentModel) error
	Update(context.Context, *ContentModel, *Resp) error
}

func RegisterContentHandler(s server.Server, hdlr ContentHandler, opts ...server.HandlerOption) error {
	type content interface {
		Search(ctx context.Context, in *ReqContent, out *RespContent) error
		Add(ctx context.Context, in *ContentModel, out *Resp) error
		Delete(ctx context.Context, in *Ids, out *Resp) error
		FindOne(ctx context.Context, in *ContentModel, out *ContentModel) error
		Update(ctx context.Context, in *ContentModel, out *Resp) error
	}
	type Content struct {
		content
	}
	h := &contentHandler{hdlr}
	return s.Handle(s.NewHandler(&Content{h}, opts...))
}

type contentHandler struct {
	ContentHandler
}

func (h *contentHandler) Search(ctx context.Context, in *ReqContent, out *RespContent) error {
	return h.ContentHandler.Search(ctx, in, out)
}

func (h *contentHandler) Add(ctx context.Context, in *ContentModel, out *Resp) error {
	return h.ContentHandler.Add(ctx, in, out)
}

func (h *contentHandler) Delete(ctx context.Context, in *Ids, out *Resp) error {
	return h.ContentHandler.Delete(ctx, in, out)
}

func (h *contentHandler) FindOne(ctx context.Context, in *ContentModel, out *ContentModel) error {
	return h.ContentHandler.FindOne(ctx, in, out)
}

func (h *contentHandler) Update(ctx context.Context, in *ContentModel, out *Resp) error {
	return h.ContentHandler.Update(ctx, in, out)
}
