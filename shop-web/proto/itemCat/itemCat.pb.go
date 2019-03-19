// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/areas/item.proto

package itemCat

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

type ItemCatArr struct {
	ItemCats             []*ItemCatModel `protobuf:"bytes,1,rep,name=itemCats,proto3" json:"itemCats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ItemCatArr) Reset()         { *m = ItemCatArr{} }
func (m *ItemCatArr) String() string { return proto.CompactTextString(m) }
func (*ItemCatArr) ProtoMessage()    {}
func (*ItemCatArr) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad96e9f562b409d, []int{0}
}

func (m *ItemCatArr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemCatArr.Unmarshal(m, b)
}
func (m *ItemCatArr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemCatArr.Marshal(b, m, deterministic)
}
func (m *ItemCatArr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemCatArr.Merge(m, src)
}
func (m *ItemCatArr) XXX_Size() int {
	return xxx_messageInfo_ItemCatArr.Size(m)
}
func (m *ItemCatArr) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemCatArr.DiscardUnknown(m)
}

var xxx_messageInfo_ItemCatArr proto.InternalMessageInfo

func (m *ItemCatArr) GetItemCats() []*ItemCatModel {
	if m != nil {
		return m.ItemCats
	}
	return nil
}

type ItemCatModel struct {
	Id                   int64    `protobuf:"zigzag64,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId             int64    `protobuf:"zigzag64,2,opt,name=parentId,proto3" json:"parentId,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TypeId               int64    `protobuf:"zigzag64,4,opt,name=typeId,proto3" json:"typeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemCatModel) Reset()         { *m = ItemCatModel{} }
func (m *ItemCatModel) String() string { return proto.CompactTextString(m) }
func (*ItemCatModel) ProtoMessage()    {}
func (*ItemCatModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ad96e9f562b409d, []int{1}
}

func (m *ItemCatModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemCatModel.Unmarshal(m, b)
}
func (m *ItemCatModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemCatModel.Marshal(b, m, deterministic)
}
func (m *ItemCatModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemCatModel.Merge(m, src)
}
func (m *ItemCatModel) XXX_Size() int {
	return xxx_messageInfo_ItemCatModel.Size(m)
}
func (m *ItemCatModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemCatModel.DiscardUnknown(m)
}

var xxx_messageInfo_ItemCatModel proto.InternalMessageInfo

func (m *ItemCatModel) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ItemCatModel) GetParentId() int64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *ItemCatModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ItemCatModel) GetTypeId() int64 {
	if m != nil {
		return m.TypeId
	}
	return 0
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
	return fileDescriptor_0ad96e9f562b409d, []int{2}
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
	proto.RegisterType((*ItemCatArr)(nil), "ItemCatArr")
	proto.RegisterType((*ItemCatModel)(nil), "ItemCatModel")
	proto.RegisterType((*Resp)(nil), "Resp")
}

func init() { proto.RegisterFile("proto/areas/item.proto", fileDescriptor_0ad96e9f562b409d) }

var fileDescriptor_0ad96e9f562b409d = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x85, 0x6b, 0xc7, 0xc4, 0xce, 0xa5, 0xcd, 0x70, 0x43, 0x51, 0x33, 0x19, 0x2f, 0x75, 0x17,
	0x07, 0xd2, 0x42, 0xe7, 0xb4, 0x50, 0xf0, 0x50, 0x28, 0xfa, 0x07, 0x0a, 0xba, 0x04, 0x83, 0x6c,
	0x0b, 0x49, 0x4b, 0xfe, 0x7d, 0xc9, 0x55, 0x71, 0xdb, 0x25, 0xdb, 0xbd, 0xf7, 0x3d, 0xf1, 0x74,
	0x07, 0x0f, 0xd6, 0x8d, 0x61, 0xdc, 0x28, 0x47, 0xca, 0x6f, 0xba, 0x40, 0xfd, 0xbb, 0x0a, 0x0d,
	0x7b, 0xd5, 0x2b, 0x40, 0xfb, 0x63, 0xec, 0x9c, 0xc3, 0x27, 0x28, 0x22, 0xf6, 0x22, 0x29, 0x67,
	0xf5, 0x72, 0x7b, 0xd7, 0x44, 0xfc, 0x39, 0x6a, 0x32, 0x72, 0xc2, 0xd5, 0x01, 0x6e, 0xff, 0x12,
	0x5c, 0x41, 0xda, 0x69, 0x91, 0x94, 0x49, 0x8d, 0x32, 0xed, 0x34, 0xae, 0xa1, 0xb0, 0xca, 0xd1,
	0x10, 0x5a, 0x2d, 0x52, 0x76, 0x27, 0x8d, 0x08, 0xd9, 0xa0, 0x7a, 0x12, 0xb3, 0x32, 0xa9, 0x17,
	0x92, 0x67, 0xbc, 0x87, 0x79, 0x38, 0x59, 0x6a, 0xb5, 0xc8, 0x38, 0x1d, 0x55, 0xf5, 0x02, 0x99,
	0x24, 0x6f, 0xcf, 0x6f, 0x0e, 0x46, 0x1d, 0xb9, 0xa1, 0x90, 0x3c, 0xa3, 0x80, 0xbc, 0x27, 0xef,
	0xd5, 0x91, 0xb8, 0x62, 0x21, 0x2f, 0x72, 0xbb, 0x87, 0x3c, 0xfe, 0x0e, 0x1b, 0x58, 0x7d, 0x74,
	0x83, 0x7e, 0x3b, 0x7d, 0x5d, 0xea, 0xff, 0xef, 0xb4, 0x5e, 0x36, 0xbf, 0x17, 0xa8, 0x6e, 0xf0,
	0x11, 0xf2, 0x73, 0x7e, 0x67, 0xcc, 0xf5, 0xe0, 0x7e, 0xce, 0x17, 0x7c, 0xfe, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x4f, 0xbc, 0xf9, 0x8a, 0x5e, 0x01, 0x00, 0x00,
}
