// Code generated by protoc-gen-go. DO NOT EDIT.
// source: external/infra_proxy/request/policyfiles.proto

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

type Policyfiles struct {
	// Chef Organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId             string   `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policyfiles) Reset()         { *m = Policyfiles{} }
func (m *Policyfiles) String() string { return proto.CompactTextString(m) }
func (*Policyfiles) ProtoMessage()    {}
func (*Policyfiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_859b464031f99f11, []int{0}
}

func (m *Policyfiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policyfiles.Unmarshal(m, b)
}
func (m *Policyfiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policyfiles.Marshal(b, m, deterministic)
}
func (m *Policyfiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policyfiles.Merge(m, src)
}
func (m *Policyfiles) XXX_Size() int {
	return xxx_messageInfo_Policyfiles.Size(m)
}
func (m *Policyfiles) XXX_DiscardUnknown() {
	xxx_messageInfo_Policyfiles.DiscardUnknown(m)
}

var xxx_messageInfo_Policyfiles proto.InternalMessageInfo

func (m *Policyfiles) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Policyfiles) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type Policyfile struct {
	// Chef Organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Policyfile name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Policyfile revision ID.
	RevisionId           string   `protobuf:"bytes,4,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policyfile) Reset()         { *m = Policyfile{} }
func (m *Policyfile) String() string { return proto.CompactTextString(m) }
func (*Policyfile) ProtoMessage()    {}
func (*Policyfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_859b464031f99f11, []int{1}
}

func (m *Policyfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policyfile.Unmarshal(m, b)
}
func (m *Policyfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policyfile.Marshal(b, m, deterministic)
}
func (m *Policyfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policyfile.Merge(m, src)
}
func (m *Policyfile) XXX_Size() int {
	return xxx_messageInfo_Policyfile.Size(m)
}
func (m *Policyfile) XXX_DiscardUnknown() {
	xxx_messageInfo_Policyfile.DiscardUnknown(m)
}

var xxx_messageInfo_Policyfile proto.InternalMessageInfo

func (m *Policyfile) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Policyfile) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *Policyfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policyfile) GetRevisionId() string {
	if m != nil {
		return m.RevisionId
	}
	return ""
}

func init() {
	proto.RegisterType((*Policyfiles)(nil), "chef.automate.api.infra_proxy.request.Policyfiles")
	proto.RegisterType((*Policyfile)(nil), "chef.automate.api.infra_proxy.request.Policyfile")
}

func init() {
	proto.RegisterFile("external/infra_proxy/request/policyfiles.proto", fileDescriptor_859b464031f99f11)
}

var fileDescriptor_859b464031f99f11 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x31, 0x4b, 0x03, 0x41,
	0x10, 0x85, 0x89, 0xc6, 0x60, 0x26, 0xdd, 0x82, 0x10, 0xb0, 0x50, 0x02, 0x82, 0xd5, 0x6e, 0x61,
	0x25, 0x56, 0xda, 0x5d, 0x27, 0x96, 0x36, 0x61, 0x73, 0x3b, 0x77, 0x19, 0xb8, 0xdb, 0x59, 0x67,
	0x77, 0x43, 0xf2, 0xef, 0xe5, 0x56, 0x8f, 0xb3, 0xb2, 0x48, 0x37, 0xbc, 0x8f, 0x6f, 0xe0, 0x3d,
	0xd0, 0x78, 0x4c, 0x28, 0xde, 0x76, 0x86, 0x7c, 0x23, 0x76, 0x1b, 0x84, 0x8f, 0x27, 0x23, 0xf8,
	0x95, 0x31, 0x26, 0x13, 0xb8, 0xa3, 0xfa, 0xd4, 0x50, 0x87, 0x51, 0x07, 0xe1, 0xc4, 0xea, 0xa1,
	0xde, 0x63, 0xa3, 0x6d, 0x4e, 0xdc, 0xdb, 0x84, 0xda, 0x06, 0xd2, 0x7f, 0x44, 0xfd, 0x2b, 0x6e,
	0x5e, 0x61, 0xf5, 0x3e, 0xb9, 0xea, 0x06, 0x16, 0x2c, 0xed, 0x96, 0xdc, 0x7a, 0x76, 0x3f, 0x7b,
	0x5c, 0x7e, 0x5c, 0xb1, 0xb4, 0x95, 0x53, 0xb7, 0xb0, 0x8c, 0x28, 0x07, 0x94, 0x81, 0x5c, 0x14,
	0x72, 0xfd, 0x13, 0x54, 0x6e, 0x93, 0x01, 0xa6, 0x17, 0xe7, 0x7c, 0x50, 0x0a, 0xe6, 0xde, 0xf6,
	0xb8, 0xbe, 0x2c, 0x79, 0xb9, 0xd5, 0x1d, 0xac, 0x04, 0x0f, 0x14, 0x89, 0xfd, 0xa0, 0xcc, 0x0b,
	0x82, 0x31, 0xaa, 0xdc, 0xdb, 0xcb, 0xe7, 0x73, 0x4b, 0x69, 0x9f, 0x77, 0xba, 0xe6, 0xde, 0x0c,
	0x6d, 0xcd, 0xd8, 0xd6, 0xd8, 0x40, 0xe6, 0xbf, 0xbd, 0x76, 0x8b, 0x32, 0xd2, 0xd3, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x55, 0xdc, 0x5c, 0x4a, 0x56, 0x01, 0x00, 0x00,
}
