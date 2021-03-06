// Code generated by protoc-gen-go. DO NOT EDIT.
// source: automate-grpc/protoc-gen-a2-config/api/a2conf/annotations.proto

package a2conf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type ServiceConfig struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceConfig) Reset()         { *m = ServiceConfig{} }
func (m *ServiceConfig) String() string { return proto.CompactTextString(m) }
func (*ServiceConfig) ProtoMessage()    {}
func (*ServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c713123588ddfeba, []int{0}
}

func (m *ServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceConfig.Unmarshal(m, b)
}
func (m *ServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceConfig.Marshal(b, m, deterministic)
}
func (m *ServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceConfig.Merge(m, src)
}
func (m *ServiceConfig) XXX_Size() int {
	return xxx_messageInfo_ServiceConfig.Size(m)
}
func (m *ServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceConfig proto.InternalMessageInfo

func (m *ServiceConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

var E_Port = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Port)(nil),
	Field:         51000,
	Name:          "chef.automate.api.port",
	Tag:           "bytes,51000,opt,name=port",
	Filename:      "automate-grpc/protoc-gen-a2-config/api/a2conf/annotations.proto",
}

var E_ServiceConfig = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*ServiceConfig)(nil),
	Field:         51000,
	Name:          "chef.automate.api.service_config",
	Tag:           "bytes,51000,opt,name=service_config",
	Filename:      "automate-grpc/protoc-gen-a2-config/api/a2conf/annotations.proto",
}

func init() {
	proto.RegisterType((*ServiceConfig)(nil), "chef.automate.api.ServiceConfig")
	proto.RegisterExtension(E_Port)
	proto.RegisterExtension(E_ServiceConfig)
}

func init() {
	proto.RegisterFile("automate-grpc/protoc-gen-a2-config/api/a2conf/annotations.proto", fileDescriptor_c713123588ddfeba)
}

var fileDescriptor_c713123588ddfeba = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd0, 0xc1, 0x4a, 0xf4, 0x30,
	0x10, 0x07, 0x70, 0xca, 0xb7, 0x7c, 0x60, 0x65, 0x05, 0x73, 0x71, 0x59, 0x10, 0x8b, 0x5e, 0xf6,
	0xd2, 0x04, 0xea, 0x45, 0xf6, 0x22, 0x28, 0x78, 0x52, 0xd4, 0xee, 0xcd, 0x8b, 0xa4, 0xd9, 0x69,
	0x36, 0xb2, 0xcd, 0x84, 0x64, 0xea, 0x63, 0xf8, 0x3c, 0x3e, 0x9e, 0x34, 0xa1, 0xa0, 0xae, 0x97,
	0xbd, 0x75, 0xe0, 0xff, 0xff, 0x75, 0x26, 0xf9, 0xb5, 0xec, 0x09, 0x3b, 0x49, 0x50, 0x6a, 0xef,
	0x94, 0x70, 0x1e, 0x09, 0x55, 0xa9, 0xc1, 0x96, 0xb2, 0x2a, 0x15, 0xda, 0xd6, 0x68, 0x21, 0x9d,
	0x11, 0xb2, 0x1a, 0x06, 0x21, 0xad, 0x45, 0x92, 0x64, 0xd0, 0x06, 0x1e, 0xb3, 0xec, 0x58, 0x6d,
	0xa0, 0xe5, 0xa3, 0xc2, 0xa5, 0x33, 0xf3, 0xab, 0xfd, 0x4c, 0x87, 0x9e, 0x12, 0x36, 0x2f, 0x34,
	0xa2, 0xde, 0x42, 0xaa, 0x34, 0x7d, 0x2b, 0xd6, 0x10, 0x94, 0x37, 0x8e, 0xd0, 0xa7, 0xc4, 0xf9,
	0x45, 0x3e, 0x5d, 0x81, 0x7f, 0x37, 0x0a, 0x6e, 0x23, 0xc4, 0x58, 0x3e, 0xb1, 0xb2, 0x83, 0x59,
	0x56, 0x64, 0x8b, 0x83, 0x3a, 0x7e, 0x2f, 0xef, 0xf3, 0xc9, 0x80, 0xb2, 0x53, 0x9e, 0x3c, 0x3e,
	0x7a, 0xfc, 0xce, 0xc0, 0x76, 0xfd, 0xe8, 0xe2, 0x01, 0xb3, 0xcf, 0x8f, 0x7f, 0x45, 0xb6, 0x38,
	0xac, 0x4e, 0xf8, 0xce, 0x0d, 0xfc, 0x09, 0x3d, 0xd5, 0x51, 0x59, 0xbe, 0xe5, 0x47, 0x21, 0xfd,
	0xf2, 0x35, 0x2d, 0xcf, 0xce, 0x76, 0xdc, 0x07, 0x08, 0x41, 0x6a, 0xf8, 0x2d, 0x17, 0x7f, 0xc8,
	0x3f, 0xd6, 0xaf, 0xa7, 0xe1, 0xfb, 0x78, 0xb3, 0x7a, 0x79, 0xd6, 0x86, 0x36, 0x7d, 0xc3, 0x15,
	0x76, 0x62, 0x28, 0x8b, 0xb1, 0x2c, 0x14, 0x76, 0x0e, 0x2d, 0x58, 0x0a, 0x62, 0xaf, 0x07, 0x6e,
	0xfe, 0xc7, 0xc0, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0x0e, 0xb1, 0xe3, 0xec, 0x01,
	0x00, 0x00,
}
