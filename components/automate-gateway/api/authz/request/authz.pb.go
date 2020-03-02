// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/authz/request/authz.proto

package request

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

type IntrospectAllReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectAllReq) Reset()         { *m = IntrospectAllReq{} }
func (m *IntrospectAllReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectAllReq) ProtoMessage()    {}
func (*IntrospectAllReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa62939b9d22bb69, []int{0}
}

func (m *IntrospectAllReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectAllReq.Unmarshal(m, b)
}
func (m *IntrospectAllReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectAllReq.Marshal(b, m, deterministic)
}
func (m *IntrospectAllReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectAllReq.Merge(m, src)
}
func (m *IntrospectAllReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectAllReq.Size(m)
}
func (m *IntrospectAllReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectAllReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectAllReq proto.InternalMessageInfo

type IntrospectSomeReq struct {
	Paths                []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectSomeReq) Reset()         { *m = IntrospectSomeReq{} }
func (m *IntrospectSomeReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectSomeReq) ProtoMessage()    {}
func (*IntrospectSomeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa62939b9d22bb69, []int{1}
}

func (m *IntrospectSomeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectSomeReq.Unmarshal(m, b)
}
func (m *IntrospectSomeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectSomeReq.Marshal(b, m, deterministic)
}
func (m *IntrospectSomeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectSomeReq.Merge(m, src)
}
func (m *IntrospectSomeReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectSomeReq.Size(m)
}
func (m *IntrospectSomeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectSomeReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectSomeReq proto.InternalMessageInfo

func (m *IntrospectSomeReq) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

type IntrospectReq struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Parameters           []string `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntrospectReq) Reset()         { *m = IntrospectReq{} }
func (m *IntrospectReq) String() string { return proto.CompactTextString(m) }
func (*IntrospectReq) ProtoMessage()    {}
func (*IntrospectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa62939b9d22bb69, []int{2}
}

func (m *IntrospectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntrospectReq.Unmarshal(m, b)
}
func (m *IntrospectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntrospectReq.Marshal(b, m, deterministic)
}
func (m *IntrospectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntrospectReq.Merge(m, src)
}
func (m *IntrospectReq) XXX_Size() int {
	return xxx_messageInfo_IntrospectReq.Size(m)
}
func (m *IntrospectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_IntrospectReq.DiscardUnknown(m)
}

var xxx_messageInfo_IntrospectReq proto.InternalMessageInfo

func (m *IntrospectReq) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *IntrospectReq) GetParameters() []string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func init() {
	proto.RegisterType((*IntrospectAllReq)(nil), "chef.automate.api.authz.request.IntrospectAllReq")
	proto.RegisterType((*IntrospectSomeReq)(nil), "chef.automate.api.authz.request.IntrospectSomeReq")
	proto.RegisterType((*IntrospectReq)(nil), "chef.automate.api.authz.request.IntrospectReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/authz/request/authz.proto", fileDescriptor_fa62939b9d22bb69)
}

var fileDescriptor_fa62939b9d22bb69 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xbf, 0x4b, 0xc5, 0x30,
	0x10, 0xc7, 0xa9, 0xbf, 0xe0, 0x1d, 0x08, 0x1a, 0x1c, 0x3a, 0xe9, 0xa3, 0xd3, 0x73, 0x30, 0x19,
	0x9c, 0x1c, 0x55, 0x10, 0x5d, 0xe3, 0xe6, 0x76, 0x2f, 0x9c, 0x2f, 0x85, 0xa6, 0x97, 0x26, 0x57,
	0x44, 0xff, 0x7a, 0x49, 0xab, 0xd6, 0xf5, 0x6d, 0xf7, 0xfd, 0x05, 0xf7, 0x81, 0x3b, 0xc7, 0x21,
	0x72, 0x4f, 0xbd, 0x64, 0x83, 0xa3, 0x70, 0x40, 0xa1, 0x9b, 0x1d, 0x0a, 0x7d, 0xe0, 0xa7, 0xc1,
	0xd8, 0x16, 0xd3, 0x7f, 0x99, 0x44, 0xc3, 0x48, 0x59, 0x66, 0xa5, 0x63, 0x62, 0x61, 0x75, 0xe5,
	0x3c, 0xbd, 0xeb, 0xdf, 0x91, 0xc6, 0xd8, 0xea, 0x39, 0xfe, 0x29, 0x37, 0x0a, 0xce, 0x5e, 0x7a,
	0x49, 0x9c, 0x23, 0x39, 0xb9, 0xef, 0x3a, 0x4b, 0x43, 0x73, 0x0d, 0xe7, 0x8b, 0xf7, 0xca, 0x81,
	0x2c, 0x0d, 0xea, 0x02, 0x8e, 0x23, 0x8a, 0xcf, 0x75, 0xb5, 0x3e, 0xdc, 0xac, 0xec, 0x2c, 0x9a,
	0x47, 0x38, 0x5d, 0xaa, 0xa5, 0xa6, 0xe0, 0xa8, 0x24, 0x75, 0xb5, 0xae, 0x36, 0x2b, 0x3b, 0xdd,
	0xea, 0x12, 0x20, 0x62, 0xc2, 0x40, 0x42, 0x29, 0xd7, 0x07, 0xd3, 0xfe, 0x9f, 0xf3, 0xf0, 0xfc,
	0xf6, 0xb4, 0x6b, 0xc5, 0x8f, 0x5b, 0xed, 0x38, 0x98, 0xf2, 0xf1, 0x1f, 0xa6, 0xd9, 0x0b, 0x7d,
	0x7b, 0x32, 0x51, 0xdf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xf3, 0xd3, 0x1b, 0x32, 0x01,
	0x00, 0x00,
}
