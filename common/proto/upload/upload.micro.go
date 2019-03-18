// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/areas/upload.proto

/*
Package upload is a generated protocol buffer package.

It is generated from these files:
	proto/areas/upload.proto

It has these top-level messages:
	ReqUpload
	Resp
*/
package upload

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

// Client API for Upload service

type UploadService interface {
	UploadFile(ctx context.Context, in *ReqUpload, opts ...client.CallOption) (*Resp, error)
}

type uploadService struct {
	c    client.Client
	name string
}

func NewUploadService(name string, c client.Client) UploadService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "upload"
	}
	return &uploadService{
		c:    c,
		name: name,
	}
}

func (c *uploadService) UploadFile(ctx context.Context, in *ReqUpload, opts ...client.CallOption) (*Resp, error) {
	req := c.c.NewRequest(c.name, "Upload.UploadFile", in)
	out := new(Resp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Upload service

type UploadHandler interface {
	UploadFile(context.Context, *ReqUpload, *Resp) error
}

func RegisterUploadHandler(s server.Server, hdlr UploadHandler, opts ...server.HandlerOption) error {
	type upload interface {
		UploadFile(ctx context.Context, in *ReqUpload, out *Resp) error
	}
	type Upload struct {
		upload
	}
	h := &uploadHandler{hdlr}
	return s.Handle(s.NewHandler(&Upload{h}, opts...))
}

type uploadHandler struct {
	UploadHandler
}

func (h *uploadHandler) UploadFile(ctx context.Context, in *ReqUpload, out *Resp) error {
	return h.UploadHandler.UploadFile(ctx, in, out)
}
