// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/areas/upload.proto

package upload

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ReqUpload struct {
	Img                  []byte   `protobuf:"bytes,1,opt,name=img,proto3" json:"img,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUpload) Reset()         { *m = ReqUpload{} }
func (m *ReqUpload) String() string { return proto.CompactTextString(m) }
func (*ReqUpload) ProtoMessage()    {}
func (*ReqUpload) Descriptor() ([]byte, []int) {
	return fileDescriptor_49559f894ba3519a, []int{0}
}

func (m *ReqUpload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUpload.Unmarshal(m, b)
}
func (m *ReqUpload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUpload.Marshal(b, m, deterministic)
}
func (m *ReqUpload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUpload.Merge(m, src)
}
func (m *ReqUpload) XXX_Size() int {
	return xxx_messageInfo_ReqUpload.Size(m)
}
func (m *ReqUpload) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUpload.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUpload proto.InternalMessageInfo

func (m *ReqUpload) GetImg() []byte {
	if m != nil {
		return m.Img
	}
	return nil
}

type Resp struct {
	Flag                 bool     `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_49559f894ba3519a, []int{1}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetFlag() bool {
	if m != nil {
		return m.Flag
	}
	return false
}

func (m *Resp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqUpload)(nil), "ReqUpload")
	proto.RegisterType((*Resp)(nil), "Resp")
}

func init() { proto.RegisterFile("proto/areas/upload.proto", fileDescriptor_49559f894ba3519a) }

var fileDescriptor_49559f894ba3519a = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2c, 0x4a, 0x4d, 0x2c, 0xd6, 0x2f, 0x2d, 0xc8, 0xc9, 0x4f, 0x4c, 0xd1, 0x03,
	0x0b, 0x29, 0xc9, 0x72, 0x71, 0x06, 0xa5, 0x16, 0x86, 0x82, 0x85, 0x84, 0x04, 0xb8, 0x98, 0x33,
	0x73, 0xd3, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x40, 0x4c, 0x25, 0x13, 0x2e, 0x96, 0xa0,
	0xd4, 0xe2, 0x02, 0x21, 0x21, 0x2e, 0x96, 0xb4, 0x9c, 0x44, 0x88, 0x14, 0x47, 0x10, 0x98, 0x2d,
	0x24, 0xc1, 0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0xe3, 0x1a, 0x69, 0x73, 0xb1, 0x41, 0x4d, 0x54, 0xe4, 0xe2, 0x82, 0xb0, 0xdc, 0x32,
	0x73, 0x52, 0x85, 0xb8, 0xf4, 0xe0, 0x76, 0x49, 0xb1, 0xea, 0x81, 0x0c, 0x56, 0x62, 0x48, 0x62,
	0x03, 0x3b, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x13, 0xcc, 0xa6, 0xa4, 0x00, 0x00,
	0x00,
}
