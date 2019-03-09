// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/areas/seller.proto

package seller

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

//*
//--response
//response.rows;//结果
//respons
//--requeste.total;//更新总记录数
//
//seller/search.do?page=1&rows=10
//{status: "0"}
type ReqSeller struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Rows                 int32    `protobuf:"varint,2,opt,name=rows,proto3" json:"rows,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSeller) Reset()         { *m = ReqSeller{} }
func (m *ReqSeller) String() string { return proto.CompactTextString(m) }
func (*ReqSeller) ProtoMessage()    {}
func (*ReqSeller) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d725fb596284bf3, []int{0}
}

func (m *ReqSeller) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSeller.Unmarshal(m, b)
}
func (m *ReqSeller) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSeller.Marshal(b, m, deterministic)
}
func (m *ReqSeller) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSeller.Merge(m, src)
}
func (m *ReqSeller) XXX_Size() int {
	return xxx_messageInfo_ReqSeller.Size(m)
}
func (m *ReqSeller) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSeller.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSeller proto.InternalMessageInfo

func (m *ReqSeller) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReqSeller) GetRows() int32 {
	if m != nil {
		return m.Rows
	}
	return 0
}

func (m *ReqSeller) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RespSeller struct {
	Total                int64    `protobuf:"zigzag64,1,opt,name=total,proto3" json:"total,omitempty"`
	Rows                 []*Rows  `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSeller) Reset()         { *m = RespSeller{} }
func (m *RespSeller) String() string { return proto.CompactTextString(m) }
func (*RespSeller) ProtoMessage()    {}
func (*RespSeller) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d725fb596284bf3, []int{1}
}

func (m *RespSeller) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSeller.Unmarshal(m, b)
}
func (m *RespSeller) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSeller.Marshal(b, m, deterministic)
}
func (m *RespSeller) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSeller.Merge(m, src)
}
func (m *RespSeller) XXX_Size() int {
	return xxx_messageInfo_RespSeller.Size(m)
}
func (m *RespSeller) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSeller.DiscardUnknown(m)
}

var xxx_messageInfo_RespSeller proto.InternalMessageInfo

func (m *RespSeller) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RespSeller) GetRows() []*Rows {
	if m != nil {
		return m.Rows
	}
	return nil
}

type ReqId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqId) Reset()         { *m = ReqId{} }
func (m *ReqId) String() string { return proto.CompactTextString(m) }
func (*ReqId) ProtoMessage()    {}
func (*ReqId) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d725fb596284bf3, []int{2}
}

func (m *ReqId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqId.Unmarshal(m, b)
}
func (m *ReqId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqId.Marshal(b, m, deterministic)
}
func (m *ReqId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqId.Merge(m, src)
}
func (m *ReqId) XXX_Size() int {
	return xxx_messageInfo_ReqId.Size(m)
}
func (m *ReqId) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqId.DiscardUnknown(m)
}

var xxx_messageInfo_ReqId proto.InternalMessageInfo

func (m *ReqId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Rows struct {
	SellerId             string   `protobuf:"bytes,1,opt,name=sellerId,proto3" json:"sellerId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NickName             string   `protobuf:"bytes,3,opt,name=nickName,proto3" json:"nickName,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Mobile               string   `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Telephone            string   `protobuf:"bytes,7,opt,name=telephone,proto3" json:"telephone,omitempty"`
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	AddressDetail        string   `protobuf:"bytes,9,opt,name=addressDetail,proto3" json:"addressDetail,omitempty"`
	LinkmanName          string   `protobuf:"bytes,10,opt,name=linkmanName,proto3" json:"linkmanName,omitempty"`
	LinkmanQq            string   `protobuf:"bytes,11,opt,name=linkmanQq,proto3" json:"linkmanQq,omitempty"`
	LinkmanMobile        string   `protobuf:"bytes,12,opt,name=linkmanMobile,proto3" json:"linkmanMobile,omitempty"`
	LinkmanEmail         string   `protobuf:"bytes,13,opt,name=linkmanEmail,proto3" json:"linkmanEmail,omitempty"`
	LicenseNumber        string   `protobuf:"bytes,14,opt,name=licenseNumber,proto3" json:"licenseNumber,omitempty"`
	TaxNumber            string   `protobuf:"bytes,15,opt,name=taxNumber,proto3" json:"taxNumber,omitempty"`
	OrgNumber            string   `protobuf:"bytes,16,opt,name=orgNumber,proto3" json:"orgNumber,omitempty"`
	Address              int64    `protobuf:"zigzag64,17,opt,name=address,proto3" json:"address,omitempty"`
	LogoPic              string   `protobuf:"bytes,18,opt,name=logoPic,proto3" json:"logoPic,omitempty"`
	Brief                string   `protobuf:"bytes,19,opt,name=brief,proto3" json:"brief,omitempty"`
	CreateTime           string   `protobuf:"bytes,20,opt,name=createTime,proto3" json:"createTime,omitempty"`
	LegalPerson          string   `protobuf:"bytes,21,opt,name=legalPerson,proto3" json:"legalPerson,omitempty"`
	LegalPersonCardId    string   `protobuf:"bytes,22,opt,name=legalPersonCardId,proto3" json:"legalPersonCardId,omitempty"`
	BankUser             string   `protobuf:"bytes,23,opt,name=bankUser,proto3" json:"bankUser,omitempty"`
	BankName             string   `protobuf:"bytes,24,opt,name=bankName,proto3" json:"bankName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rows) Reset()         { *m = Rows{} }
func (m *Rows) String() string { return proto.CompactTextString(m) }
func (*Rows) ProtoMessage()    {}
func (*Rows) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d725fb596284bf3, []int{3}
}

func (m *Rows) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rows.Unmarshal(m, b)
}
func (m *Rows) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rows.Marshal(b, m, deterministic)
}
func (m *Rows) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rows.Merge(m, src)
}
func (m *Rows) XXX_Size() int {
	return xxx_messageInfo_Rows.Size(m)
}
func (m *Rows) XXX_DiscardUnknown() {
	xxx_messageInfo_Rows.DiscardUnknown(m)
}

var xxx_messageInfo_Rows proto.InternalMessageInfo

func (m *Rows) GetSellerId() string {
	if m != nil {
		return m.SellerId
	}
	return ""
}

func (m *Rows) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rows) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *Rows) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Rows) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Rows) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Rows) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *Rows) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Rows) GetAddressDetail() string {
	if m != nil {
		return m.AddressDetail
	}
	return ""
}

func (m *Rows) GetLinkmanName() string {
	if m != nil {
		return m.LinkmanName
	}
	return ""
}

func (m *Rows) GetLinkmanQq() string {
	if m != nil {
		return m.LinkmanQq
	}
	return ""
}

func (m *Rows) GetLinkmanMobile() string {
	if m != nil {
		return m.LinkmanMobile
	}
	return ""
}

func (m *Rows) GetLinkmanEmail() string {
	if m != nil {
		return m.LinkmanEmail
	}
	return ""
}

func (m *Rows) GetLicenseNumber() string {
	if m != nil {
		return m.LicenseNumber
	}
	return ""
}

func (m *Rows) GetTaxNumber() string {
	if m != nil {
		return m.TaxNumber
	}
	return ""
}

func (m *Rows) GetOrgNumber() string {
	if m != nil {
		return m.OrgNumber
	}
	return ""
}

func (m *Rows) GetAddress() int64 {
	if m != nil {
		return m.Address
	}
	return 0
}

func (m *Rows) GetLogoPic() string {
	if m != nil {
		return m.LogoPic
	}
	return ""
}

func (m *Rows) GetBrief() string {
	if m != nil {
		return m.Brief
	}
	return ""
}

func (m *Rows) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *Rows) GetLegalPerson() string {
	if m != nil {
		return m.LegalPerson
	}
	return ""
}

func (m *Rows) GetLegalPersonCardId() string {
	if m != nil {
		return m.LegalPersonCardId
	}
	return ""
}

func (m *Rows) GetBankUser() string {
	if m != nil {
		return m.BankUser
	}
	return ""
}

func (m *Rows) GetBankName() string {
	if m != nil {
		return m.BankName
	}
	return ""
}

type ReqIdAndStatus struct {
	SellerId             string   `protobuf:"bytes,1,opt,name=sellerId,proto3" json:"sellerId,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqIdAndStatus) Reset()         { *m = ReqIdAndStatus{} }
func (m *ReqIdAndStatus) String() string { return proto.CompactTextString(m) }
func (*ReqIdAndStatus) ProtoMessage()    {}
func (*ReqIdAndStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d725fb596284bf3, []int{4}
}

func (m *ReqIdAndStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqIdAndStatus.Unmarshal(m, b)
}
func (m *ReqIdAndStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqIdAndStatus.Marshal(b, m, deterministic)
}
func (m *ReqIdAndStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqIdAndStatus.Merge(m, src)
}
func (m *ReqIdAndStatus) XXX_Size() int {
	return xxx_messageInfo_ReqIdAndStatus.Size(m)
}
func (m *ReqIdAndStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqIdAndStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReqIdAndStatus proto.InternalMessageInfo

func (m *ReqIdAndStatus) GetSellerId() string {
	if m != nil {
		return m.SellerId
	}
	return ""
}

func (m *ReqIdAndStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RespReturn struct {
	Flag                 bool     `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespReturn) Reset()         { *m = RespReturn{} }
func (m *RespReturn) String() string { return proto.CompactTextString(m) }
func (*RespReturn) ProtoMessage()    {}
func (*RespReturn) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d725fb596284bf3, []int{5}
}

func (m *RespReturn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespReturn.Unmarshal(m, b)
}
func (m *RespReturn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespReturn.Marshal(b, m, deterministic)
}
func (m *RespReturn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespReturn.Merge(m, src)
}
func (m *RespReturn) XXX_Size() int {
	return xxx_messageInfo_RespReturn.Size(m)
}
func (m *RespReturn) XXX_DiscardUnknown() {
	xxx_messageInfo_RespReturn.DiscardUnknown(m)
}

var xxx_messageInfo_RespReturn proto.InternalMessageInfo

func (m *RespReturn) GetFlag() bool {
	if m != nil {
		return m.Flag
	}
	return false
}

func (m *RespReturn) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqSeller)(nil), "ReqSeller")
	proto.RegisterType((*RespSeller)(nil), "RespSeller")
	proto.RegisterType((*ReqId)(nil), "ReqId")
	proto.RegisterType((*Rows)(nil), "Rows")
	proto.RegisterType((*ReqIdAndStatus)(nil), "ReqIdAndStatus")
	proto.RegisterType((*RespReturn)(nil), "RespReturn")
}

func init() { proto.RegisterFile("proto/areas/seller.proto", fileDescriptor_1d725fb596284bf3) }

var fileDescriptor_1d725fb596284bf3 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x6d, 0xd3, 0x3a, 0x89, 0x27, 0x69, 0xfa, 0x75, 0xbf, 0xd2, 0x2e, 0x11, 0x42, 0x91, 0xe1,
	0x22, 0x17, 0xc8, 0x95, 0xca, 0x1d, 0x12, 0x17, 0x88, 0x82, 0x14, 0x21, 0x4a, 0x71, 0xe9, 0x03,
	0x6c, 0xe2, 0x69, 0x6a, 0xc5, 0xde, 0x75, 0x76, 0x1d, 0x05, 0x5e, 0x94, 0xe7, 0x41, 0x3b, 0xbb,
	0x76, 0x1c, 0x21, 0x71, 0x37, 0xe7, 0x9c, 0xd9, 0xf9, 0xf3, 0x8c, 0x81, 0x97, 0x5a, 0x55, 0xea,
	0x4a, 0x68, 0x14, 0xe6, 0xca, 0x60, 0x9e, 0xa3, 0x8e, 0x89, 0x8a, 0xbe, 0x40, 0x98, 0xe0, 0xfa,
	0x9e, 0x28, 0xc6, 0xe0, 0xb8, 0x14, 0x4b, 0xe4, 0x87, 0x93, 0xc3, 0x69, 0x90, 0x90, 0x6d, 0x39,
	0xad, 0xb6, 0x86, 0x77, 0x1c, 0x67, 0x6d, 0x76, 0x01, 0x5d, 0x53, 0x89, 0x6a, 0x63, 0xf8, 0xd1,
	0xe4, 0x70, 0x1a, 0x26, 0x1e, 0x45, 0xef, 0x01, 0x12, 0x34, 0xa5, 0x8f, 0x76, 0x0e, 0x41, 0xa5,
	0x2a, 0x91, 0x53, 0x38, 0x96, 0x38, 0xc0, 0x9e, 0x37, 0xf1, 0x8e, 0xa6, 0x83, 0xeb, 0x20, 0x4e,
	0xd4, 0xd6, 0xb8, 0xb0, 0xd1, 0x25, 0x04, 0x09, 0xae, 0x67, 0x29, 0x1b, 0x41, 0x27, 0x4b, 0xe9,
	0x59, 0x98, 0x74, 0xb2, 0x34, 0xfa, 0x1d, 0xc0, 0xb1, 0xf5, 0x63, 0x63, 0xe8, 0xbb, 0xea, 0x67,
	0xb5, 0xdc, 0x60, 0x5b, 0xa8, 0x14, 0x05, 0x52, 0xa1, 0x61, 0x42, 0xb6, 0xf5, 0x97, 0xd9, 0x62,
	0x75, 0x6b, 0x79, 0x57, 0x6a, 0x83, 0xad, 0x56, 0x0a, 0x63, 0xb6, 0x4a, 0xa7, 0xfc, 0xd8, 0x69,
	0x35, 0xb6, 0xa5, 0x63, 0x21, 0xb2, 0x9c, 0x07, 0x24, 0x38, 0x60, 0xdb, 0x2e, 0xd4, 0x3c, 0xcb,
	0x91, 0x77, 0x5d, 0xdb, 0x0e, 0xb1, 0x17, 0x10, 0x56, 0x98, 0x63, 0xf9, 0xa4, 0x24, 0xf2, 0x1e,
	0x49, 0x3b, 0xa2, 0x35, 0xac, 0x7e, 0x7b, 0x58, 0xec, 0x35, 0x9c, 0x88, 0x34, 0xd5, 0x68, 0xcc,
	0x0d, 0x56, 0x36, 0x57, 0x48, 0xf2, 0x3e, 0xc9, 0x26, 0x30, 0xc8, 0x33, 0xb9, 0x2a, 0x84, 0xa4,
	0x26, 0x80, 0x7c, 0xda, 0x94, 0xcd, 0xee, 0xe1, 0xf7, 0x35, 0x1f, 0xb8, 0xec, 0x0d, 0x61, 0xb3,
	0x78, 0xf0, 0xd5, 0x95, 0x3e, 0x74, 0x59, 0xf6, 0x48, 0x16, 0xc1, 0xd0, 0x13, 0x9f, 0xa8, 0xed,
	0x13, 0x72, 0xda, 0xe3, 0x5c, 0xa4, 0x05, 0x4a, 0x83, 0xb7, 0x9b, 0x62, 0x8e, 0x9a, 0x8f, 0xea,
	0x48, 0x2d, 0x92, 0x66, 0x21, 0x7e, 0x7a, 0x8f, 0x53, 0x3f, 0x8b, 0x9a, 0xb0, 0xaa, 0xd2, 0x4b,
	0xaf, 0xfe, 0xe7, 0xd4, 0x86, 0x60, 0x1c, 0x7a, 0xbe, 0x79, 0x7e, 0x46, 0x2b, 0x53, 0x43, 0xab,
	0xe4, 0x6a, 0xa9, 0xee, 0xb2, 0x05, 0x67, 0xf4, 0xaa, 0x86, 0xf6, 0x4b, 0xcd, 0x75, 0x86, 0x8f,
	0xfc, 0x7f, 0xf7, 0xa5, 0x08, 0xb0, 0x97, 0x00, 0x0b, 0x8d, 0xa2, 0xc2, 0x1f, 0x59, 0x81, 0xfc,
	0x9c, 0xa4, 0x16, 0x43, 0x53, 0xc5, 0xa5, 0xc8, 0xef, 0x50, 0x1b, 0x25, 0xf9, 0x33, 0x3f, 0xd5,
	0x1d, 0xc5, 0xde, 0xc0, 0x59, 0x0b, 0x7e, 0x14, 0x3a, 0x9d, 0xa5, 0xfc, 0x82, 0xfc, 0xfe, 0x16,
	0xec, 0x2e, 0xcd, 0x85, 0x5c, 0x3d, 0x18, 0xd4, 0xfc, 0xd2, 0xed, 0x52, 0x8d, 0x6b, 0x8d, 0x3e,
	0x1f, 0xdf, 0x69, 0x16, 0x47, 0x37, 0x30, 0xa2, 0x8d, 0xff, 0x20, 0xd3, 0x7b, 0xb7, 0x15, 0xff,
	0xda, 0xf0, 0xdd, 0x26, 0x75, 0xf6, 0xce, 0xee, 0x9d, 0x3b, 0xbb, 0x04, 0xab, 0x8d, 0x96, 0xf6,
	0x0e, 0x1e, 0x73, 0xb1, 0xa4, 0xd7, 0xfd, 0x84, 0x6c, 0x3b, 0xbf, 0x02, 0x8d, 0xb1, 0xb7, 0xed,
	0x9e, 0xd6, 0xf0, 0xfa, 0x17, 0x74, 0xfd, 0xb9, 0xbe, 0xb2, 0x96, 0xd0, 0x8b, 0x27, 0x06, 0x71,
	0xf3, 0x4b, 0x18, 0x0f, 0xe2, 0xdd, 0x45, 0x47, 0x07, 0x6c, 0x0c, 0xbd, 0xcf, 0x99, 0x4c, 0xbf,
	0x49, 0x64, 0xdd, 0x98, 0x4a, 0x1f, 0xbb, 0x13, 0x8e, 0x0e, 0x58, 0x0c, 0xc3, 0x87, 0x32, 0x15,
	0x15, 0xfa, 0x56, 0x4e, 0xe3, 0xfd, 0xde, 0x7c, 0x2c, 0x57, 0x66, 0x74, 0x30, 0xef, 0xd2, 0x1f,
	0xe8, 0xed, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0xc3, 0xb5, 0xf7, 0x9d, 0x04, 0x00, 0x00,
}