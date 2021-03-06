// Code generated by protoc-gen-go. DO NOT EDIT.
// source: automate-gateway/api/iam/v2/common/tokens.proto

package common

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

type Token struct {
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name for the token.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Unique, optionally user-specified value.
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// Active state. Defaults to true.
	// If set to false, token will not authenticate.
	Active bool `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	// Created timestamp.
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Updated timestamp.
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// List of projects this token belongs to. May be empty.
	Projects             []string `protobuf:"bytes,7,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e02ea53f6bf2368, []int{0}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Token) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Token) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Token) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *Token) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Token) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *Token) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*Token)(nil), "chef.automate.api.iam.v2.Token")
}

func init() {
	proto.RegisterFile("automate-gateway/api/iam/v2/common/tokens.proto", fileDescriptor_5e02ea53f6bf2368)
}

var fileDescriptor_5e02ea53f6bf2368 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xb6, 0x09, 0x8d, 0x07, 0x06, 0x0b, 0x21, 0x0b, 0x09, 0x29, 0x62, 0xca, 0x82,
	0x2d, 0x95, 0x5f, 0x50, 0x06, 0xc4, 0x1c, 0x31, 0xb1, 0xa0, 0xab, 0x73, 0xb4, 0x06, 0xec, 0xb3,
	0x92, 0x4b, 0x10, 0xbf, 0x8b, 0x3f, 0x88, 0xe2, 0xa6, 0x9d, 0xd9, 0xee, 0x7d, 0xef, 0xde, 0xf0,
	0x9e, 0x30, 0x30, 0x30, 0x79, 0x60, 0xbc, 0xdf, 0x03, 0xe3, 0x37, 0xfc, 0x18, 0x88, 0xce, 0x38,
	0xf0, 0x66, 0xdc, 0x18, 0x4b, 0xde, 0x53, 0x30, 0x4c, 0x9f, 0x18, 0x7a, 0x1d, 0x3b, 0x62, 0x92,
	0xca, 0x1e, 0xf0, 0x5d, 0x9f, 0x52, 0x1a, 0xa2, 0xd3, 0x0e, 0xbc, 0x1e, 0x37, 0x77, 0xbf, 0x99,
	0xc8, 0x5f, 0xa6, 0x57, 0x79, 0x29, 0x16, 0xae, 0x55, 0x59, 0x95, 0xd5, 0x65, 0xb3, 0x70, 0xad,
	0x94, 0x62, 0x15, 0xc0, 0xa3, 0x5a, 0x24, 0x92, 0x6e, 0x79, 0x25, 0xf2, 0x11, 0xbe, 0x06, 0x54,
	0xcb, 0x04, 0x8f, 0x42, 0x5e, 0x8b, 0x02, 0x2c, 0xbb, 0x11, 0xd5, 0xaa, 0xca, 0xea, 0x75, 0x33,
	0x2b, 0x79, 0x2b, 0x84, 0xed, 0x10, 0x18, 0xdb, 0x37, 0x60, 0x95, 0xa7, 0x48, 0x39, 0x93, 0x2d,
	0x4f, 0xf6, 0x10, 0xdb, 0x93, 0x5d, 0x1c, 0xed, 0x99, 0x6c, 0x59, 0xde, 0x88, 0x75, 0xec, 0xe8,
	0x03, 0x2d, 0xf7, 0xea, 0xa2, 0x5a, 0xd6, 0x65, 0x73, 0xd6, 0x8f, 0xcf, 0xaf, 0x4f, 0x7b, 0xc7,
	0x87, 0x61, 0xa7, 0x2d, 0x79, 0x33, 0x95, 0x3b, 0x4f, 0x32, 0xf5, 0x8f, 0x14, 0x30, 0x70, 0xff,
	0x8f, 0x99, 0x76, 0x45, 0x1a, 0xe8, 0xe1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xa1, 0x12, 0x74,
	0x53, 0x01, 0x00, 0x00,
}
