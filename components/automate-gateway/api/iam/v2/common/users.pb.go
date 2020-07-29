// Code generated by protoc-gen-go. DO NOT EDIT.
// source: automate-gateway/api/iam/v2/common/users.proto

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

type User struct {
	// Display name for local user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique ID. Cannot be changed. Used to log in.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Unique ID used to add local users to local teams. Cannot be changed.
	MembershipId         string   `protobuf:"bytes,3,opt,name=membership_id,json=membershipId,proto3" json:"membership_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfdaba45ccb23df4, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetMembershipId() string {
	if m != nil {
		return m.MembershipId
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "chef.automate.api.iam.v2.User")
}

func init() {
	proto.RegisterFile("automate-gateway/api/iam/v2/common/users.proto", fileDescriptor_bfdaba45ccb23df4)
}

var fileDescriptor_bfdaba45ccb23df4 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x40, 0xd9, 0x75, 0x11, 0x0c, 0x6a, 0x91, 0x2a, 0xa5, 0x68, 0x63, 0x63, 0x06, 0xd6, 0x3f,
	0xb0, 0x10, 0xad, 0x04, 0xc1, 0xc6, 0x46, 0x66, 0x37, 0xe3, 0xee, 0x14, 0x93, 0x84, 0x24, 0xbb,
	0xc7, 0xfd, 0xfd, 0xb1, 0x81, 0xbb, 0x2b, 0xaf, 0x1b, 0xde, 0x30, 0xbc, 0x37, 0xca, 0xe2, 0x52,
	0x82, 0x60, 0xa1, 0x97, 0x09, 0x0b, 0xed, 0x70, 0x0f, 0x18, 0x19, 0x18, 0x05, 0xd6, 0x1e, 0xc6,
	0x20, 0x12, 0x3c, 0x2c, 0x99, 0x52, 0xb6, 0x31, 0x85, 0x12, 0xb4, 0x19, 0x67, 0xfa, 0x3f, 0x1d,
	0x59, 0x8c, 0x6c, 0x19, 0xc5, 0xae, 0xfd, 0xe3, 0x97, 0xea, 0x7e, 0x32, 0x25, 0xad, 0x55, 0xe7,
	0x51, 0xc8, 0x34, 0x0f, 0xcd, 0xf3, 0xcd, 0x77, 0x9d, 0xf5, 0xbd, 0x6a, 0xd9, 0x99, 0xb6, 0x92,
	0x96, 0x9d, 0x7e, 0x52, 0x77, 0x42, 0x32, 0x50, 0xca, 0x33, 0xc7, 0x3f, 0x76, 0xe6, 0xaa, 0xae,
	0x6e, 0xcf, 0xf0, 0xd3, 0xbd, 0x7d, 0xfc, 0xbe, 0x4f, 0x5c, 0xe6, 0x65, 0xb0, 0x63, 0x10, 0xd8,
	0xbc, 0x70, 0xf4, 0x6e, 0x65, 0x31, 0x78, 0xf2, 0x25, 0xc3, 0xe5, 0x07, 0x86, 0xeb, 0xda, 0xfe,
	0x7a, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x7a, 0x23, 0xa1, 0xed, 0x00, 0x00, 0x00,
}
