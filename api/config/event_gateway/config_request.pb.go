// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/event_gateway/config_request.proto

package event_gateway

import (
	fmt "fmt"
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ConfigRequest struct {
	V1                   *ConfigRequest_V1 `protobuf:"bytes,1,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95c9fe8e6daff5b, []int{0}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if m != nil {
		return m.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	Sys                  *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc                  *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1) Reset()         { *m = ConfigRequest_V1{} }
func (m *ConfigRequest_V1) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1) ProtoMessage()    {}
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95c9fe8e6daff5b, []int{0, 0}
}

func (m *ConfigRequest_V1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1.Unmarshal(m, b)
}
func (m *ConfigRequest_V1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1.Merge(m, src)
}
func (m *ConfigRequest_V1) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1.Size(m)
}
func (m *ConfigRequest_V1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1 proto.InternalMessageInfo

func (m *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if m != nil {
		return m.Sys
	}
	return nil
}

func (m *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if m != nil {
		return m.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	Mlsa                 *shared.Mlsa                     `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Tls                  *shared.TLSCredentials           `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Service              *ConfigRequest_V1_System_Service `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Log                  *ConfigRequest_V1_System_Log     `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	FrontendTls          []*shared.FrontendTLSCredential  `protobuf:"bytes,5,rep,name=frontend_tls,json=frontendTls,proto3" json:"frontend_tls,omitempty" toml:"frontend_tls,omitempty" mapstructure:"frontend_tls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95c9fe8e6daff5b, []int{0, 0, 0}
}

func (m *ConfigRequest_V1_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System.Merge(m, src)
}
func (m *ConfigRequest_V1_System) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System.Size(m)
}
func (m *ConfigRequest_V1_System) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System proto.InternalMessageInfo

func (m *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if m != nil {
		return m.Mlsa
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetFrontendTls() []*shared.FrontendTLSCredential {
	if m != nil {
		return m.FrontendTls
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	GatewayPort          *wrappers.Int32Value  `protobuf:"bytes,3,opt,name=gateway_port,json=gatewayPort,proto3" json:"gateway_port,omitempty" toml:"gateway_port,omitempty" mapstructure:"gateway_port,omitempty"`
	DisableFrontendTls   *wrappers.BoolValue   `protobuf:"bytes,5,opt,name=disable_frontend_tls,json=disableFrontendTls,proto3" json:"disable_frontend_tls,omitempty" toml:"disable_frontend_tls,omitempty" mapstructure:"disable_frontend_tls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95c9fe8e6daff5b, []int{0, 0, 0, 0}
}

func (m *ConfigRequest_V1_System_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Size(m)
}
func (m *ConfigRequest_V1_System_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Service proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Service) GetHost() *wrappers.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetGatewayPort() *wrappers.Int32Value {
	if m != nil {
		return m.GatewayPort
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetDisableFrontendTls() *wrappers.BoolValue {
	if m != nil {
		return m.DisableFrontendTls
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	Format               *wrappers.StringValue `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty" toml:"format,omitempty" mapstructure:"format,omitempty"`
	Level                *wrappers.StringValue `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95c9fe8e6daff5b, []int{0, 0, 0, 1}
}

func (m *ConfigRequest_V1_System_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Log.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Log) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Size(m)
}
func (m *ConfigRequest_V1_System_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Log.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Log proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Log) GetFormat() *wrappers.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *ConfigRequest_V1_System_Log) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

// V1.Service can currently only be used for deployment-service
type ConfigRequest_V1_Service struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_c95c9fe8e6daff5b, []int{0, 0, 1}
}

func (m *ConfigRequest_V1_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Service.Size(m)
}
func (m *ConfigRequest_V1_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Service proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.domain.event_gateway.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.domain.event_gateway.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.domain.event_gateway.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.domain.event_gateway.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.domain.event_gateway.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.domain.event_gateway.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("config/event_gateway/config_request.proto", fileDescriptor_c95c9fe8e6daff5b)
}

var fileDescriptor_c95c9fe8e6daff5b = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0xb5, 0x26, 0xeb, 0xf6, 0x73, 0x37, 0xfd, 0x26, 0x83, 0x20, 0x0a, 0x08, 0xaa, 0x71,
	0x33, 0x90, 0xea, 0x90, 0x6c, 0x57, 0x30, 0x69, 0xd2, 0x86, 0x86, 0x40, 0x1d, 0x7f, 0xb2, 0x69,
	0x17, 0x48, 0xa8, 0x72, 0x53, 0xc7, 0x8d, 0xe4, 0xd8, 0xc1, 0x76, 0x33, 0xed, 0x7d, 0x78, 0x05,
	0x2e, 0xb9, 0xd9, 0x2b, 0xf0, 0x0e, 0xbc, 0x40, 0xef, 0x11, 0x8a, 0xe3, 0x14, 0x75, 0x43, 0x1d,
	0xda, 0xee, 0x92, 0xe3, 0xf3, 0xfd, 0x9c, 0x73, 0xbe, 0x3e, 0x32, 0x78, 0x9a, 0x08, 0x9e, 0x66,
	0x34, 0x20, 0x25, 0xe1, 0x7a, 0x40, 0xb1, 0x26, 0x67, 0xf8, 0x3c, 0xa8, 0x83, 0x03, 0x49, 0xbe,
	0x4c, 0x88, 0xd2, 0xa8, 0x90, 0x42, 0x0b, 0xb8, 0x99, 0x8c, 0x49, 0x8a, 0xf0, 0x44, 0x8b, 0x1c,
	0x6b, 0x82, 0x46, 0x22, 0xc7, 0x19, 0x47, 0x73, 0x42, 0xdf, 0xb7, 0x38, 0x35, 0xc6, 0x92, 0x8c,
	0x02, 0xca, 0xc4, 0x10, 0xb3, 0x5a, 0xef, 0xdf, 0x9f, 0x3f, 0xd3, 0x4c, 0xd9, 0x83, 0xbd, 0x86,
	0xd9, 0xa3, 0xb2, 0x48, 0x02, 0x13, 0x4c, 0x7a, 0x94, 0xf0, 0x1e, 0x8e, 0x7a, 0x56, 0x84, 0x8b,
	0x2c, 0xc0, 0x51, 0xf5, 0x13, 0x60, 0xce, 0x85, 0xc6, 0x3a, 0x13, 0xbc, 0x01, 0x3c, 0xa2, 0x42,
	0x50, 0x46, 0x6a, 0xe5, 0x70, 0x92, 0x06, 0x67, 0x12, 0x17, 0x05, 0x91, 0xf6, 0x7c, 0xf3, 0xc7,
	0x2a, 0x58, 0x3f, 0x30, 0x9c, 0xb8, 0x9e, 0x08, 0xbe, 0x02, 0xad, 0x32, 0xf4, 0x96, 0xba, 0x4b,
	0x5b, 0x9d, 0x68, 0x07, 0x5d, 0x3f, 0x18, 0x9a, 0x93, 0xa3, 0xd3, 0x30, 0x6e, 0x95, 0xa1, 0xff,
	0x73, 0x05, 0xb4, 0x4e, 0x43, 0x78, 0x04, 0x1c, 0x75, 0xae, 0x2c, 0xed, 0xe5, 0x4d, 0x68, 0xe8,
	0xf8, 0x5c, 0x69, 0x92, 0xc7, 0x15, 0x07, 0xbe, 0x03, 0x8e, 0x2a, 0x13, 0xaf, 0x65, 0x70, 0xbb,
	0x37, 0xc3, 0x11, 0x59, 0x66, 0x09, 0x89, 0x2b, 0x90, 0xff, 0xad, 0x0d, 0xda, 0x35, 0x1f, 0xee,
	0x00, 0x37, 0x67, 0x0a, 0xdb, 0x56, 0xbb, 0x97, 0xd8, 0x19, 0x4f, 0x25, 0x46, 0xb5, 0xe1, 0xe8,
	0x88, 0x29, 0x1c, 0x9b, 0x6c, 0xb8, 0x0b, 0x1c, 0xcd, 0x94, 0x6d, 0xe8, 0xd9, 0x22, 0xd1, 0x49,
	0xff, 0xf8, 0x40, 0x92, 0x11, 0xe1, 0x3a, 0xc3, 0x4c, 0xc5, 0x95, 0x0c, 0x7e, 0x06, 0x2b, 0xaa,
	0x6e, 0xc7, 0x73, 0x0c, 0xe1, 0xe0, 0x16, 0x0e, 0xcd, 0x26, 0x6b, 0x98, 0xf0, 0x23, 0x70, 0x98,
	0xa0, 0x9e, 0x6b, 0xd0, 0x7b, 0xb7, 0x41, 0xf7, 0x05, 0x8d, 0x2b, 0x16, 0x3c, 0x01, 0x6b, 0xa9,
	0x14, 0x5c, 0x13, 0x3e, 0x1a, 0x54, 0x83, 0x2f, 0x77, 0x9d, 0xad, 0x4e, 0x14, 0x2e, 0x1a, 0xfc,
	0xd0, 0xe6, 0xcf, 0x19, 0x10, 0x77, 0x1a, 0xcc, 0x09, 0x53, 0xfe, 0xf7, 0x16, 0x58, 0xb1, 0xdd,
	0xc3, 0xe7, 0xc0, 0x1d, 0x0b, 0xa5, 0xed, 0x3d, 0x3c, 0x44, 0xf5, 0xfe, 0xa2, 0x66, 0x7f, 0xd1,
	0xb1, 0x96, 0x19, 0xa7, 0xa7, 0x98, 0x4d, 0x48, 0x6c, 0x32, 0xe1, 0x6b, 0xe0, 0x16, 0x42, 0x6a,
	0x7b, 0x09, 0x0f, 0xae, 0x28, 0xde, 0x70, 0xbd, 0x1d, 0x19, 0xc1, 0xfe, 0xbd, 0x8b, 0xa9, 0x07,
	0x67, 0xa6, 0x6f, 0xfc, 0xea, 0xfa, 0x2e, 0xc7, 0x5a, 0xc5, 0x06, 0x00, 0x13, 0xb0, 0x66, 0x8d,
	0x18, 0x18, 0xa0, 0x73, 0x3d, 0xf0, 0xc9, 0xc5, 0xd4, 0x7b, 0x0c, 0xd6, 0x2a, 0x48, 0x63, 0xe2,
	0xc6, 0xd7, 0xf7, 0xbe, 0x09, 0xf4, 0x6c, 0x20, 0xee, 0xd8, 0x8f, 0x0f, 0x55, 0x91, 0x3e, 0xb8,
	0x3b, 0xca, 0x14, 0x1e, 0x32, 0x32, 0xb8, 0xe4, 0x64, 0x55, 0xcc, 0xbf, 0x52, 0x6c, 0x5f, 0x08,
	0x56, 0x4f, 0x0b, 0xad, 0xee, 0xf0, 0x8f, 0x73, 0x6f, 0xdd, 0x55, 0x77, 0x63, 0xd9, 0x17, 0xc0,
	0xe9, 0x0b, 0x0a, 0x77, 0x40, 0x3b, 0x15, 0x32, 0xc7, 0xff, 0x66, 0x9e, 0xcd, 0x85, 0x11, 0x58,
	0x66, 0xa4, 0x24, 0xcc, 0xfa, 0xb7, 0x58, 0x54, 0xa7, 0xfa, 0xff, 0xcd, 0xee, 0xeb, 0xc5, 0x9d,
	0x8b, 0xa9, 0xf7, 0x3f, 0x58, 0x37, 0x7b, 0xd4, 0x4c, 0xbc, 0xbf, 0xfd, 0x29, 0xa4, 0x99, 0x1e,
	0x4f, 0x86, 0x28, 0x11, 0x79, 0x50, 0x2d, 0x47, 0xd0, 0x2c, 0x87, 0x79, 0xae, 0xfe, 0xf6, 0xb2,
	0x0e, 0xdb, 0xa6, 0xe2, 0xf6, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x26, 0xb8, 0x1b, 0x78,
	0x05, 0x00, 0x00,
}
