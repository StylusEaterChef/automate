// Code generated by protoc-gen-go. DO NOT EDIT.
// source: external/common/response.proto

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

type ExportData struct {
	// Exported reports in JSON or CSV.
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExportData) Reset()         { *m = ExportData{} }
func (m *ExportData) String() string { return proto.CompactTextString(m) }
func (*ExportData) ProtoMessage()    {}
func (*ExportData) Descriptor() ([]byte, []int) {
	return fileDescriptor_69a4762be7c41afe, []int{0}
}

func (m *ExportData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportData.Unmarshal(m, b)
}
func (m *ExportData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportData.Marshal(b, m, deterministic)
}
func (m *ExportData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportData.Merge(m, src)
}
func (m *ExportData) XXX_Size() int {
	return xxx_messageInfo_ExportData.Size(m)
}
func (m *ExportData) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportData.DiscardUnknown(m)
}

var xxx_messageInfo_ExportData proto.InternalMessageInfo

func (m *ExportData) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*ExportData)(nil), "chef.automate.api.common.ExportData")
}

func init() {
	proto.RegisterFile("external/common/response.proto", fileDescriptor_69a4762be7c41afe)
}

var fileDescriptor_69a4762be7c41afe = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xad, 0x28, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2f, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x48, 0xce, 0x48, 0x4d,
	0xd3, 0x4b, 0x2c, 0x2d, 0xc9, 0xcf, 0x4d, 0x2c, 0x49, 0xd5, 0x4b, 0x2c, 0xc8, 0xd4, 0x83, 0x28,
	0x54, 0x52, 0xe3, 0xe2, 0x72, 0xad, 0x28, 0xc8, 0x2f, 0x2a, 0x71, 0x49, 0x2c, 0x49, 0x14, 0x92,
	0xe0, 0x62, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09,
	0x82, 0x71, 0x9d, 0xf4, 0xa2, 0x74, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0x40, 0x1a, 0xf5, 0x41,
	0xc6, 0xe9, 0xc3, 0x8c, 0xd3, 0x4f, 0x2c, 0xc8, 0xd4, 0x47, 0x73, 0x40, 0x12, 0x1b, 0xd8, 0x62,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x5c, 0x4c, 0x12, 0x9a, 0x00, 0x00, 0x00,
}
