// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/license_control/config_request.proto

package license_control

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
	V1                   *ConfigRequest_V1 `protobuf:"bytes,3,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f48c4def2b47dd4, []int{0}
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
	return fileDescriptor_9f48c4def2b47dd4, []int{0, 0}
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
	Mlsa                 *shared.Mlsa                       `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Tls                  *shared.TLSCredentials             `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Service              *ConfigRequest_V1_System_Service   `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Telemetry            *ConfigRequest_V1_System_Telemetry `protobuf:"bytes,4,opt,name=telemetry,proto3" json:"telemetry,omitempty" toml:"telemetry,omitempty" mapstructure:"telemetry,omitempty"`
	Log                  *ConfigRequest_V1_System_Log       `protobuf:"bytes,5,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                              `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f48c4def2b47dd4, []int{0, 0, 0}
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

func (m *ConfigRequest_V1_System) GetTelemetry() *ConfigRequest_V1_System_Telemetry {
	if m != nil {
		return m.Telemetry
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrappers.Int32Value  `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f48c4def2b47dd4, []int{0, 0, 0, 0}
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

type ConfigRequest_V1_System_Telemetry struct {
	OptOut               *wrappers.BoolValue   `protobuf:"bytes,1,opt,name=opt_out,json=optOut,proto3" json:"opt_out,omitempty" toml:"opt_out,omitempty" mapstructure:"opt_out,omitempty"`
	Url                  *wrappers.StringValue `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty" toml:"url,omitempty" mapstructure:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Telemetry) Reset()         { *m = ConfigRequest_V1_System_Telemetry{} }
func (m *ConfigRequest_V1_System_Telemetry) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Telemetry) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Telemetry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f48c4def2b47dd4, []int{0, 0, 0, 1}
}

func (m *ConfigRequest_V1_System_Telemetry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Telemetry.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Telemetry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Telemetry.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Telemetry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Telemetry.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Telemetry) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Telemetry.Size(m)
}
func (m *ConfigRequest_V1_System_Telemetry) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Telemetry.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Telemetry proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Telemetry) GetOptOut() *wrappers.BoolValue {
	if m != nil {
		return m.OptOut
	}
	return nil
}

func (m *ConfigRequest_V1_System_Telemetry) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
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
	return fileDescriptor_9f48c4def2b47dd4, []int{0, 0, 0, 2}
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

type ConfigRequest_V1_Service struct {
	LicensePath          *wrappers.StringValue `protobuf:"bytes,1,opt,name=license_path,json=licensePath,proto3" json:"license_path,omitempty" toml:"license_path,omitempty" mapstructure:"license_path,omitempty"`
	License              *wrappers.StringValue `protobuf:"bytes,2,opt,name=license,proto3" json:"license,omitempty" toml:"license,omitempty" mapstructure:"license,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f48c4def2b47dd4, []int{0, 0, 1}
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

func (m *ConfigRequest_V1_Service) GetLicensePath() *wrappers.StringValue {
	if m != nil {
		return m.LicensePath
	}
	return nil
}

func (m *ConfigRequest_V1_Service) GetLicense() *wrappers.StringValue {
	if m != nil {
		return m.License
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.domain.license_control.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.domain.license_control.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.domain.license_control.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.domain.license_control.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Telemetry)(nil), "chef.automate.domain.license_control.ConfigRequest.V1.System.Telemetry")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.domain.license_control.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.domain.license_control.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("config/license_control/config_request.proto", fileDescriptor_9f48c4def2b47dd4)
}

var fileDescriptor_9f48c4def2b47dd4 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0xd5, 0xa6, 0x6b, 0xb7, 0x67, 0x20, 0x4d, 0x3e, 0xb0, 0x10, 0x5e, 0x34, 0x21, 0x0e,
	0x08, 0x54, 0x87, 0xb6, 0x63, 0x07, 0x04, 0x4c, 0x6c, 0x82, 0x09, 0x34, 0xb4, 0x29, 0x9d, 0x76,
	0xe0, 0x52, 0xb9, 0x99, 0x9b, 0x46, 0x72, 0xfc, 0x04, 0xdb, 0x29, 0xda, 0x15, 0x38, 0xf2, 0x65,
	0xe0, 0x23, 0xec, 0x2b, 0xed, 0x0b, 0xa0, 0x38, 0xce, 0x60, 0x1b, 0x9a, 0xaa, 0xed, 0xe8, 0xe4,
	0xf9, 0xff, 0xfe, 0xcf, 0x9b, 0x0d, 0xcf, 0x62, 0x94, 0x93, 0x34, 0x09, 0x45, 0x1a, 0x73, 0xa9,
	0xf9, 0x28, 0x46, 0x69, 0x14, 0x8a, 0xb0, 0xfa, 0x3c, 0x52, 0xfc, 0x4b, 0xc1, 0xb5, 0xa1, 0xb9,
	0x42, 0x83, 0xe4, 0x71, 0x3c, 0xe5, 0x13, 0xca, 0x0a, 0x83, 0x19, 0x33, 0x9c, 0x1e, 0x61, 0xc6,
	0x52, 0x49, 0x2f, 0x48, 0x83, 0x87, 0x09, 0x62, 0x22, 0x78, 0x68, 0x35, 0xe3, 0x62, 0x12, 0x7e,
	0x55, 0x2c, 0xcf, 0xb9, 0xd2, 0x15, 0x25, 0x08, 0x9c, 0xa5, 0x9e, 0x32, 0xc5, 0x8f, 0xc2, 0x44,
	0xe0, 0x98, 0x09, 0xf7, 0x6f, 0xf5, 0xfc, 0x3f, 0x23, 0x6a, 0xd1, 0x66, 0xed, 0xda, 0x4d, 0x54,
	0x1e, 0x57, 0xec, 0xb8, 0x9b, 0x70, 0xd9, 0x65, 0xfd, 0xae, 0x13, 0xb1, 0x3c, 0x0d, 0x59, 0xbf,
	0x3c, 0x84, 0x4c, 0x4a, 0x34, 0xcc, 0xa4, 0x28, 0x1d, 0xe0, 0xd1, 0xaf, 0x25, 0xb8, 0xbd, 0x6d,
	0xe3, 0xa2, 0xaa, 0x26, 0xf2, 0x1e, 0x9a, 0xb3, 0x9e, 0xef, 0xad, 0x35, 0x9e, 0x2c, 0xf7, 0x37,
	0xe8, 0x3c, 0xa5, 0xd1, 0x73, 0x00, 0x7a, 0xd8, 0x8b, 0x9a, 0xb3, 0x5e, 0xf0, 0x7b, 0x11, 0x9a,
	0x87, 0x3d, 0xb2, 0x07, 0x9e, 0x3e, 0xd6, 0x7e, 0xc3, 0xf2, 0x5e, 0x5f, 0x8f, 0x47, 0x87, 0xc7,
	0xda, 0xf0, 0x2c, 0x2a, 0x49, 0x64, 0x1f, 0x3c, 0x3d, 0x8b, 0xfd, 0xa6, 0x05, 0xbe, 0xb9, 0x2e,
	0x90, 0xab, 0x59, 0x1a, 0xf3, 0xa8, 0x44, 0x05, 0xdf, 0xdb, 0xd0, 0xae, 0x1c, 0xc8, 0x3a, 0xb4,
	0x32, 0xa1, 0x99, 0x4b, 0x77, 0xed, 0x02, 0x3d, 0x95, 0x13, 0xc5, 0x68, 0xd5, 0x56, 0xfa, 0x49,
	0x68, 0x16, 0xd9, 0x68, 0xf2, 0x0a, 0x3c, 0x23, 0xb4, 0x4b, 0xe9, 0xe9, 0x55, 0xa2, 0x83, 0xdd,
	0xe1, 0xb6, 0xe2, 0x47, 0x5c, 0x9a, 0x94, 0x09, 0x1d, 0x95, 0x32, 0x32, 0x82, 0x8e, 0xae, 0xd2,
	0x71, 0x5d, 0x7f, 0x77, 0xa3, 0x2e, 0x9d, 0xd5, 0x56, 0x53, 0x09, 0x87, 0x25, 0xc3, 0x05, 0xcf,
	0xb8, 0x51, 0xc7, 0x7e, 0xcb, 0x5a, 0xec, 0xdc, 0xcc, 0xe2, 0xa0, 0xc6, 0x45, 0x7f, 0xc9, 0x64,
	0x08, 0x9e, 0xc0, 0xc4, 0x5f, 0xb0, 0x06, 0x6f, 0x6f, 0x66, 0xb0, 0x8b, 0x49, 0x54, 0xd2, 0x82,
	0x1f, 0x0d, 0xe8, 0xb8, 0x82, 0xc8, 0x73, 0x68, 0x4d, 0x51, 0x1b, 0x37, 0x9c, 0xfb, 0xb4, 0xba,
	0x50, 0xb4, 0xbe, 0x50, 0x74, 0x68, 0x54, 0x2a, 0x93, 0x43, 0x26, 0x0a, 0x1e, 0xd9, 0x48, 0xb2,
	0x03, 0xad, 0x1c, 0x95, 0x71, 0x7d, 0xbd, 0x77, 0x49, 0xf1, 0x41, 0x9a, 0x41, 0xdf, 0x0a, 0xb6,
	0xee, 0x9c, 0x9c, 0xfa, 0xe4, 0x6c, 0x12, 0x2b, 0x3f, 0xf7, 0x82, 0x56, 0x79, 0xa5, 0x22, 0x0b,
	0x08, 0x72, 0x58, 0x3a, 0xab, 0x99, 0x0c, 0xa0, 0x83, 0xb9, 0x19, 0x61, 0x51, 0xa7, 0x12, 0x5c,
	0x02, 0x6f, 0x21, 0x8a, 0x2a, 0x91, 0x36, 0xe6, 0x66, 0xaf, 0x30, 0x84, 0x82, 0x57, 0x28, 0xe1,
	0x76, 0xe4, 0xea, 0xdc, 0xcb, 0xc0, 0x00, 0xc1, 0xdb, 0xc5, 0x84, 0xac, 0x43, 0x7b, 0x82, 0x2a,
	0x63, 0xf3, 0x55, 0xed, 0x62, 0x49, 0x1f, 0x16, 0x04, 0x9f, 0xf1, 0xf9, 0xec, 0xaa, 0xd0, 0xe0,
	0xdb, 0x3f, 0x9d, 0xde, 0x84, 0x5b, 0xf5, 0xa4, 0x72, 0x66, 0xa6, 0x73, 0x79, 0x2f, 0x3b, 0xc5,
	0x3e, 0x33, 0x53, 0xb2, 0x01, 0x1d, 0x77, 0x9c, 0x2b, 0x85, 0x3a, 0xf8, 0xe5, 0x83, 0x93, 0x53,
	0xff, 0x2e, 0xac, 0xba, 0x63, 0xd7, 0xad, 0x49, 0xd7, 0x4d, 0xe5, 0x63, 0x6b, 0xb1, 0xb1, 0xe2,
	0x6d, 0xbd, 0xf8, 0x3c, 0x48, 0x52, 0x33, 0x2d, 0xc6, 0x34, 0xc6, 0x2c, 0x2c, 0xf7, 0x2c, 0xac,
	0xf7, 0xcc, 0x3e, 0x76, 0xff, 0x7f, 0xbb, 0xc7, 0x6d, 0x6b, 0x3d, 0xf8, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xd4, 0x95, 0x7d, 0xbe, 0xdc, 0x05, 0x00, 0x00,
}
