// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interservice/cfgmgmt/response/suggestions.proto

package response

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

type Suggestion struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty" toml:"text,omitempty" mapstructure:"text,omitempty"`
	Score                float32  `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty" toml:"score,omitempty" mapstructure:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Suggestion) Reset()         { *m = Suggestion{} }
func (m *Suggestion) String() string { return proto.CompactTextString(m) }
func (*Suggestion) ProtoMessage()    {}
func (*Suggestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce12e7373590c04b, []int{0}
}

func (m *Suggestion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Suggestion.Unmarshal(m, b)
}
func (m *Suggestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Suggestion.Marshal(b, m, deterministic)
}
func (m *Suggestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Suggestion.Merge(m, src)
}
func (m *Suggestion) XXX_Size() int {
	return xxx_messageInfo_Suggestion.Size(m)
}
func (m *Suggestion) XXX_DiscardUnknown() {
	xxx_messageInfo_Suggestion.DiscardUnknown(m)
}

var xxx_messageInfo_Suggestion proto.InternalMessageInfo

func (m *Suggestion) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Suggestion) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func init() {
	proto.RegisterType((*Suggestion)(nil), "chef.automate.domain.cfgmgmt.response.suggestion")
}

func init() {
	proto.RegisterFile("interservice/cfgmgmt/response/suggestions.proto", fileDescriptor_ce12e7373590c04b)
}

var fileDescriptor_ce12e7373590c04b = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8e, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x46, 0x69, 0x51, 0xc1, 0x8c, 0xc1, 0xa1, 0x63, 0x11, 0x84, 0x4e, 0xb9, 0x41, 0x70, 0x10,
	0x27, 0x7f, 0x42, 0x47, 0xb7, 0x34, 0x5e, 0xd3, 0x0c, 0xc9, 0x95, 0xdc, 0x55, 0xfc, 0xf9, 0x62,
	0xb5, 0xb8, 0xb9, 0x7d, 0xdf, 0xf0, 0x78, 0x4f, 0x41, 0x48, 0x82, 0x99, 0x31, 0x3f, 0x82, 0x43,
	0x70, 0xbd, 0x8f, 0x3e, 0x0a, 0x64, 0xe4, 0x91, 0x12, 0x23, 0xf0, 0xe4, 0x3d, 0xb2, 0x04, 0x4a,
	0x6c, 0xc6, 0x4c, 0x42, 0xfa, 0xe0, 0x06, 0xec, 0x8d, 0x9d, 0x84, 0xa2, 0x15, 0x34, 0x77, 0x8a,
	0x36, 0x24, 0xf3, 0x05, 0xcd, 0x02, 0xee, 0x4f, 0x4a, 0xfd, 0x58, 0xad, 0xd5, 0x4a, 0xf0, 0x29,
	0x55, 0x51, 0x17, 0xcd, 0xb6, 0x9d, 0xb7, 0xde, 0xa9, 0x35, 0x3b, 0xca, 0x58, 0x95, 0x75, 0xd1,
	0x94, 0xed, 0xe7, 0x5c, 0x2f, 0xb7, 0xb3, 0x0f, 0x32, 0x4c, 0x9d, 0x71, 0x14, 0xe1, 0xed, 0x82,
	0xc5, 0x05, 0x76, 0x0c, 0xff, 0x73, 0xbb, 0xcd, 0xdc, 0x78, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x1c, 0x81, 0x0f, 0xd9, 0xd6, 0x00, 0x00, 0x00,
}
