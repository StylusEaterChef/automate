// Code generated by protoc-gen-go. DO NOT EDIT.
// source: automate-gateway/api/iam/v2/common/teams.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type Team struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Projects             []string `protobuf:"bytes,3,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_3410368cbe936077, []int{0}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

func init() {
	proto.RegisterType((*Team)(nil), "chef.automate.api.iam.v2.Team")
}

func init() {
	proto.RegisterFile("automate-gateway/api/iam/v2/common/teams.proto", fileDescriptor_3410368cbe936077)
}

var fileDescriptor_3410368cbe936077 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xcf, 0xb1, 0x6e, 0x02, 0x31,
	0x0c, 0x06, 0x60, 0x71, 0xa0, 0xaa, 0x64, 0xe8, 0x90, 0xe9, 0xc4, 0x84, 0x3a, 0x31, 0xf4, 0x62,
	0x89, 0xbe, 0x41, 0x07, 0xd4, 0x19, 0x75, 0xea, 0x66, 0x82, 0x1b, 0x52, 0xc9, 0x76, 0x74, 0x31,
	0xa0, 0xbe, 0x7d, 0x45, 0xaa, 0xde, 0xca, 0xf6, 0xdb, 0xb2, 0xad, 0xcf, 0x2e, 0xe0, 0xd9, 0x94,
	0xd1, 0x68, 0x48, 0x68, 0x74, 0xc5, 0x1f, 0xc0, 0x92, 0x21, 0x23, 0xc3, 0x65, 0x0b, 0x51, 0x99,
	0x55, 0xc0, 0x08, 0xb9, 0x86, 0x32, 0xaa, 0xa9, 0xef, 0xe3, 0x89, 0xbe, 0xa6, 0xa5, 0x80, 0x25,
	0x87, 0x8c, 0x1c, 0x2e, 0xdb, 0xd5, 0x4b, 0x1b, 0x88, 0x43, 0x22, 0x19, 0xea, 0x15, 0x53, 0xa2,
	0x11, 0xb4, 0x58, 0x56, 0xa9, 0x80, 0x22, 0x6a, 0xd8, 0xf2, 0xdf, 0x9d, 0xe7, 0x9d, 0x5b, 0x7c,
	0x10, 0xb2, 0x7f, 0x72, 0x5d, 0x3e, 0xf6, 0xb3, 0xf5, 0x6c, 0xb3, 0xdc, 0x77, 0xf9, 0xe8, 0xbd,
	0x5b, 0x08, 0x32, 0xf5, 0x5d, 0xeb, 0xb4, 0xec, 0x57, 0xee, 0xb1, 0x8c, 0xfa, 0x4d, 0xd1, 0x6a,
	0x3f, 0x5f, 0xcf, 0x37, 0xcb, 0xfd, 0x54, 0xbf, 0xbd, 0x7f, 0xee, 0x52, 0xb6, 0xd3, 0xf9, 0x10,
	0xa2, 0x32, 0xdc, 0x70, 0xf0, 0x8f, 0xbb, 0xf1, 0x8b, 0x0a, 0x89, 0x55, 0xb8, 0xff, 0xe5, 0xe1,
	0xa1, 0xc1, 0x5e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x46, 0xc4, 0x41, 0x54, 0x12, 0x01, 0x00,
	0x00,
}
