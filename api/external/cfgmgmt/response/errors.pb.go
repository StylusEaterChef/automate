// Code generated by protoc-gen-go. DO NOT EDIT.
// source: external/cfgmgmt/response/errors.proto

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

// Errors contains a list of the most common Chef Infra error type/message
// combinations among nodes in the active project as filtered according to the
// request.
type Errors struct {
	Errors               []*ErrorCount `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Errors) Reset()         { *m = Errors{} }
func (m *Errors) String() string { return proto.CompactTextString(m) }
func (*Errors) ProtoMessage()    {}
func (*Errors) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a221c6959dde75, []int{0}
}

func (m *Errors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Errors.Unmarshal(m, b)
}
func (m *Errors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Errors.Marshal(b, m, deterministic)
}
func (m *Errors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Errors.Merge(m, src)
}
func (m *Errors) XXX_Size() int {
	return xxx_messageInfo_Errors.Size(m)
}
func (m *Errors) XXX_DiscardUnknown() {
	xxx_messageInfo_Errors.DiscardUnknown(m)
}

var xxx_messageInfo_Errors proto.InternalMessageInfo

func (m *Errors) GetErrors() []*ErrorCount {
	if m != nil {
		return m.Errors
	}
	return nil
}

// ErrorCount gives the number of occurrences (count) of the error specified by
// the type and message among the nodes included by the request parameters
type ErrorCount struct {
	Count                int32    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorCount) Reset()         { *m = ErrorCount{} }
func (m *ErrorCount) String() string { return proto.CompactTextString(m) }
func (*ErrorCount) ProtoMessage()    {}
func (*ErrorCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a221c6959dde75, []int{1}
}

func (m *ErrorCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorCount.Unmarshal(m, b)
}
func (m *ErrorCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorCount.Marshal(b, m, deterministic)
}
func (m *ErrorCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorCount.Merge(m, src)
}
func (m *ErrorCount) XXX_Size() int {
	return xxx_messageInfo_ErrorCount.Size(m)
}
func (m *ErrorCount) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorCount.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorCount proto.InternalMessageInfo

func (m *ErrorCount) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ErrorCount) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ErrorCount) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Errors)(nil), "chef.automate.api.cfgmgmt.response.Errors")
	proto.RegisterType((*ErrorCount)(nil), "chef.automate.api.cfgmgmt.response.ErrorCount")
}

func init() {
	proto.RegisterFile("external/cfgmgmt/response/errors.proto", fileDescriptor_16a221c6959dde75)
}

var fileDescriptor_16a221c6959dde75 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd0, 0xb1, 0x4b, 0xc5, 0x30,
	0x10, 0x06, 0x70, 0xe2, 0xf3, 0x15, 0x3c, 0x75, 0x09, 0x0e, 0x19, 0x4b, 0x05, 0xe9, 0x74, 0x01,
	0x05, 0x71, 0x56, 0x74, 0x13, 0xa4, 0xa3, 0x0e, 0x92, 0x17, 0xee, 0xe5, 0x15, 0x4c, 0x13, 0x92,
	0x2b, 0xe8, 0x7f, 0x2f, 0x4d, 0x5b, 0x9c, 0xe4, 0x6d, 0x97, 0x2f, 0xfc, 0x8e, 0xe3, 0x83, 0x1b,
	0xfa, 0x66, 0x4a, 0x83, 0xf9, 0xd2, 0x76, 0xef, 0xbc, 0xf3, 0xac, 0x13, 0xe5, 0x18, 0x86, 0x4c,
	0x9a, 0x52, 0x0a, 0x29, 0x63, 0x4c, 0x81, 0x83, 0x6c, 0xec, 0x81, 0xf6, 0x68, 0x46, 0x0e, 0xde,
	0x30, 0xa1, 0x89, 0x3d, 0x2e, 0x00, 0x57, 0xd0, 0xbc, 0x41, 0xf5, 0x5c, 0x8c, 0x7c, 0x81, 0x6a,
	0xd6, 0x4a, 0xd4, 0x9b, 0xf6, 0xfc, 0x16, 0xf1, 0x38, 0xc7, 0x62, 0x9f, 0xc2, 0x38, 0x70, 0xb7,
	0xe8, 0xe6, 0x03, 0xe0, 0x2f, 0x95, 0x57, 0xb0, 0xb5, 0xd3, 0xa0, 0x44, 0x2d, 0xda, 0x6d, 0x37,
	0x3f, 0xa4, 0x84, 0x53, 0xfe, 0x89, 0xa4, 0x4e, 0x6a, 0xd1, 0x9e, 0x75, 0x65, 0x96, 0xd7, 0x70,
	0x59, 0x36, 0x7c, 0x7a, 0xca, 0xd9, 0x38, 0x52, 0x9b, 0xf2, 0x79, 0x51, 0xc2, 0xd7, 0x39, 0x7b,
	0x7c, 0x78, 0xbf, 0x77, 0x3d, 0x1f, 0xc6, 0x1d, 0xda, 0xe0, 0xf5, 0x74, 0xa0, 0x5e, 0x0f, 0xd4,
	0x26, 0xf6, 0xfa, 0xdf, 0x66, 0x76, 0x55, 0xe9, 0xe4, 0xee, 0x37, 0x00, 0x00, 0xff, 0xff, 0x92,
	0xac, 0xef, 0x66, 0x3d, 0x01, 0x00, 0x00,
}
