// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/bifrost/config_request.proto

package bifrost

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
	return fileDescriptor_22429cd173537913, []int{0}
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
	return fileDescriptor_22429cd173537913, []int{0, 0}
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
	Log                  *ConfigRequest_V1_System_Log     `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	Network              *ConfigRequest_V1_System_Network `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty" toml:"network,omitempty" mapstructure:"network,omitempty"`
	Sql                  *ConfigRequest_V1_System_Sql     `protobuf:"bytes,4,opt,name=sql,proto3" json:"sql,omitempty" toml:"sql,omitempty" mapstructure:"sql,omitempty"`
	Tls                  *shared.TLSCredentials           `protobuf:"bytes,5,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_22429cd173537913, []int{0, 0, 0}
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

func (m *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetNetwork() *ConfigRequest_V1_System_Network {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetSql() *ConfigRequest_V1_System_Sql {
	if m != nil {
		return m.Sql
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

type ConfigRequest_V1_System_Network struct {
	Port                 *wrappers.Int32Value  `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	ListenIp             *wrappers.StringValue `protobuf:"bytes,2,opt,name=listen_ip,json=listenIp,proto3" json:"listen_ip,omitempty" toml:"listen_ip,omitempty" mapstructure:"listen_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Network) Reset()         { *m = ConfigRequest_V1_System_Network{} }
func (m *ConfigRequest_V1_System_Network) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Network) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Network) Descriptor() ([]byte, []int) {
	return fileDescriptor_22429cd173537913, []int{0, 0, 0, 0}
}

func (m *ConfigRequest_V1_System_Network) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Network.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Network) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Network.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Network) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Network.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Network) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Network.Size(m)
}
func (m *ConfigRequest_V1_System_Network) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Network.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Network proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Network) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *ConfigRequest_V1_System_Network) GetListenIp() *wrappers.StringValue {
	if m != nil {
		return m.ListenIp
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	Level                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	RotationMaxBytes      *wrappers.Int64Value  `protobuf:"bytes,2,opt,name=rotation_max_bytes,json=rotationMaxBytes,proto3" json:"rotation_max_bytes,omitempty" toml:"rotation_max_bytes,omitempty" mapstructure:"rotation_max_bytes,omitempty"`
	RotationMaxFiles      *wrappers.Int32Value  `protobuf:"bytes,3,opt,name=rotation_max_files,json=rotationMaxFiles,proto3" json:"rotation_max_files,omitempty" toml:"rotation_max_files,omitempty" mapstructure:"rotation_max_files,omitempty"`
	MaxErrorLogsPerSecond *wrappers.Int32Value  `protobuf:"bytes,4,opt,name=max_error_logs_per_second,json=maxErrorLogsPerSecond,proto3" json:"max_error_logs_per_second,omitempty" toml:"max_error_logs_per_second,omitempty" mapstructure:"max_error_logs_per_second,omitempty"`
	// TODO(ssd) 2018-07-24: Different log
	// rotation configurables require
	// different units.
	RotationMaxMegabytes *wrappers.Int32Value `protobuf:"bytes,5,opt,name=rotation_max_megabytes,json=rotationMaxMegabytes,proto3" json:"rotation_max_megabytes,omitempty" toml:"rotation_max_megabytes,omitempty" mapstructure:"rotation_max_megabytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_22429cd173537913, []int{0, 0, 0, 1}
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

func (m *ConfigRequest_V1_System_Log) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

func (m *ConfigRequest_V1_System_Log) GetRotationMaxBytes() *wrappers.Int64Value {
	if m != nil {
		return m.RotationMaxBytes
	}
	return nil
}

func (m *ConfigRequest_V1_System_Log) GetRotationMaxFiles() *wrappers.Int32Value {
	if m != nil {
		return m.RotationMaxFiles
	}
	return nil
}

func (m *ConfigRequest_V1_System_Log) GetMaxErrorLogsPerSecond() *wrappers.Int32Value {
	if m != nil {
		return m.MaxErrorLogsPerSecond
	}
	return nil
}

func (m *ConfigRequest_V1_System_Log) GetRotationMaxMegabytes() *wrappers.Int32Value {
	if m != nil {
		return m.RotationMaxMegabytes
	}
	return nil
}

type ConfigRequest_V1_System_Sql struct {
	Timeout              *wrappers.Int32Value `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty" toml:"timeout,omitempty" mapstructure:"timeout,omitempty"`
	PoolInitSize         *wrappers.Int32Value `protobuf:"bytes,2,opt,name=pool_init_size,json=poolInitSize,proto3" json:"pool_init_size,omitempty" toml:"pool_init_size,omitempty" mapstructure:"pool_init_size,omitempty"`
	PoolMaxSize          *wrappers.Int32Value `protobuf:"bytes,3,opt,name=pool_max_size,json=poolMaxSize,proto3" json:"pool_max_size,omitempty" toml:"pool_max_size,omitempty" mapstructure:"pool_max_size,omitempty"`
	PoolQueueMax         *wrappers.Int32Value `protobuf:"bytes,4,opt,name=pool_queue_max,json=poolQueueMax,proto3" json:"pool_queue_max,omitempty" toml:"pool_queue_max,omitempty" mapstructure:"pool_queue_max,omitempty"`
	PoolQueueTimeout     *wrappers.Int32Value `protobuf:"bytes,5,opt,name=pool_queue_timeout,json=poolQueueTimeout,proto3" json:"pool_queue_timeout,omitempty" toml:"pool_queue_timeout,omitempty" mapstructure:"pool_queue_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Sql) Reset()         { *m = ConfigRequest_V1_System_Sql{} }
func (m *ConfigRequest_V1_System_Sql) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Sql) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Sql) Descriptor() ([]byte, []int) {
	return fileDescriptor_22429cd173537913, []int{0, 0, 0, 2}
}

func (m *ConfigRequest_V1_System_Sql) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Sql.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Sql) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Sql.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Sql) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Sql.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Sql) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Sql.Size(m)
}
func (m *ConfigRequest_V1_System_Sql) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Sql.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Sql proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Sql) GetTimeout() *wrappers.Int32Value {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *ConfigRequest_V1_System_Sql) GetPoolInitSize() *wrappers.Int32Value {
	if m != nil {
		return m.PoolInitSize
	}
	return nil
}

func (m *ConfigRequest_V1_System_Sql) GetPoolMaxSize() *wrappers.Int32Value {
	if m != nil {
		return m.PoolMaxSize
	}
	return nil
}

func (m *ConfigRequest_V1_System_Sql) GetPoolQueueMax() *wrappers.Int32Value {
	if m != nil {
		return m.PoolQueueMax
	}
	return nil
}

func (m *ConfigRequest_V1_System_Sql) GetPoolQueueTimeout() *wrappers.Int32Value {
	if m != nil {
		return m.PoolQueueTimeout
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_22429cd173537913, []int{0, 0, 1}
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
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.domain.bifrost.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.domain.bifrost.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.domain.bifrost.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Network)(nil), "chef.automate.domain.bifrost.ConfigRequest.V1.System.Network")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.domain.bifrost.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_System_Sql)(nil), "chef.automate.domain.bifrost.ConfigRequest.V1.System.Sql")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.domain.bifrost.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("config/bifrost/config_request.proto", fileDescriptor_22429cd173537913)
}

var fileDescriptor_22429cd173537913 = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0x87, 0xb5, 0x76, 0x5d, 0x99, 0xc7, 0xd0, 0x64, 0x60, 0x0b, 0x61, 0x42, 0x13, 0x5c, 0x10,
	0xa8, 0x8e, 0xda, 0xfd, 0x91, 0x86, 0x80, 0x89, 0x4d, 0xc0, 0x2a, 0x5a, 0x60, 0xe9, 0x18, 0x12,
	0x97, 0xc8, 0x4d, 0xdd, 0xd4, 0xc2, 0x89, 0x53, 0xdb, 0xe9, 0xba, 0x7d, 0x08, 0x3e, 0xd4, 0xee,
	0xf0, 0x01, 0x38, 0xf2, 0x11, 0xd8, 0x89, 0x1b, 0x72, 0xe2, 0x4c, 0x2b, 0x43, 0xa5, 0x1b, 0xc7,
	0x34, 0x7e, 0x9e, 0xfc, 0xde, 0xd7, 0xaf, 0x6b, 0xf0, 0xc0, 0xe7, 0x51, 0x97, 0x06, 0x4e, 0x9b,
	0x76, 0x05, 0x97, 0xca, 0xc9, 0x1e, 0x3d, 0x41, 0xfa, 0x09, 0x91, 0x0a, 0xc5, 0x82, 0x2b, 0x0e,
	0x97, 0xfd, 0x1e, 0xe9, 0x22, 0x9c, 0x28, 0x1e, 0x62, 0x45, 0x50, 0x87, 0x87, 0x98, 0x46, 0xc8,
	0x20, 0xb6, 0x6d, 0x14, 0xb2, 0x87, 0x05, 0xe9, 0x38, 0x01, 0xe3, 0x6d, 0xcc, 0x32, 0xd2, 0x5e,
	0x1a, 0x7d, 0xa7, 0x98, 0x34, 0x2f, 0xb6, 0x72, 0x5b, 0x25, 0x10, 0xb1, 0xef, 0xa4, 0x3f, 0xfa,
	0x95, 0x80, 0x44, 0x15, 0x5c, 0xab, 0x18, 0x08, 0xc7, 0xd4, 0xc1, 0x35, 0xfd, 0xe0, 0xe0, 0x28,
	0xe2, 0x0a, 0x2b, 0xca, 0xa3, 0x5c, 0x70, 0x2f, 0xe0, 0x3c, 0x60, 0x24, 0x23, 0xdb, 0x49, 0xd7,
	0x39, 0x14, 0x38, 0x8e, 0x89, 0x30, 0xef, 0xef, 0xff, 0x98, 0x03, 0xf3, 0x3b, 0xa9, 0xc7, 0xcd,
	0x6a, 0x81, 0xcf, 0x41, 0x61, 0x50, 0xb5, 0xa6, 0x56, 0xa6, 0x1e, 0xce, 0xd5, 0x10, 0x1a, 0x57,
	0x12, 0x1a, 0x01, 0xd1, 0x41, 0xd5, 0x2d, 0x0c, 0xaa, 0xf6, 0x4f, 0x00, 0x0a, 0x07, 0x55, 0xf8,
	0x1a, 0x14, 0xe5, 0x91, 0x34, 0x9e, 0xf5, 0xcb, 0x79, 0x50, 0xeb, 0x48, 0x2a, 0x12, 0xba, 0xda,
	0x00, 0x77, 0x41, 0x51, 0x0e, 0x7c, 0xab, 0x90, 0x8a, 0x36, 0x2e, 0x2b, 0x22, 0x62, 0x40, 0x7d,
	0xe2, 0x6a, 0x85, 0xfd, 0x75, 0x16, 0xcc, 0x64, 0x66, 0xb8, 0x06, 0xa6, 0x43, 0x26, 0xb1, 0x89,
	0xb7, 0xf2, 0x87, 0x95, 0x46, 0x5d, 0x81, 0x51, 0xd6, 0x5e, 0xd4, 0x64, 0x12, 0xbb, 0xe9, 0x6a,
	0xf8, 0x06, 0x14, 0x19, 0x0f, 0x4c, 0x94, 0xcd, 0x2b, 0xd5, 0x84, 0x1a, 0x3c, 0x70, 0xb5, 0x05,
	0x7e, 0x04, 0xe5, 0x88, 0xa8, 0x43, 0x2e, 0x3e, 0x5b, 0xc5, 0x54, 0xf8, 0xec, 0x6a, 0xc2, 0xb7,
	0x99, 0xc4, 0xcd, 0x6d, 0x3a, 0xa5, 0xec, 0x33, 0x6b, 0xfa, 0x7f, 0x52, 0xb6, 0xfa, 0xcc, 0xd5,
	0x16, 0xf8, 0x14, 0x14, 0x15, 0x93, 0x56, 0x29, 0x95, 0x3d, 0x1a, 0xd7, 0xa7, 0xfd, 0x46, 0x6b,
	0x47, 0x90, 0x0e, 0x89, 0x14, 0xc5, 0x4c, 0xba, 0x1a, 0xb3, 0xbf, 0x4c, 0x81, 0xb2, 0xc9, 0x07,
	0x77, 0xc1, 0x74, 0xcc, 0x85, 0x32, 0x2d, 0xbf, 0x8b, 0xb2, 0xc1, 0x44, 0xf9, 0x60, 0xa2, 0x7a,
	0xa4, 0x56, 0x6b, 0x07, 0x98, 0x25, 0x64, 0x7b, 0xe9, 0xe4, 0xd4, 0xba, 0x09, 0xca, 0x32, 0xdb,
	0xbc, 0x85, 0xef, 0xef, 0xec, 0x52, 0x4f, 0xa9, 0x58, 0xba, 0xa9, 0x01, 0x6e, 0x82, 0x59, 0x46,
	0xa5, 0x22, 0x91, 0x47, 0x63, 0xb3, 0x19, 0xcb, 0x17, 0x74, 0x2d, 0x25, 0x68, 0x14, 0xa4, 0x3e,
	0xf7, 0x5a, 0xb6, 0xbc, 0x1e, 0xdb, 0xbf, 0x0a, 0xa0, 0xd8, 0xe0, 0x01, 0xac, 0x81, 0x12, 0x23,
	0x03, 0xc2, 0x4c, 0x9a, 0xf1, 0x78, 0xb6, 0x14, 0xd6, 0x01, 0x14, 0xe6, 0x74, 0x79, 0x21, 0x1e,
	0x7a, 0xed, 0x23, 0x45, 0xa4, 0xf9, 0xfe, 0x5f, 0xcb, 0xd9, 0x58, 0xcb, 0xf8, 0x85, 0x1c, 0x6b,
	0xe2, 0xe1, 0xb6, 0x86, 0x2e, 0xa8, 0xba, 0x94, 0x11, 0x69, 0xc6, 0x60, 0x5c, 0x67, 0x46, 0x54,
	0xaf, 0x34, 0x04, 0x3f, 0x80, 0x3b, 0xda, 0x40, 0x84, 0xe0, 0xc2, 0x63, 0x3c, 0x90, 0x5e, 0x4c,
	0x84, 0x27, 0x89, 0xcf, 0xa3, 0x8e, 0x99, 0x81, 0xb1, 0xc6, 0xdb, 0x21, 0x1e, 0xbe, 0xd4, 0x70,
	0x83, 0x07, 0xf2, 0x3d, 0x11, 0xad, 0x94, 0x84, 0x7b, 0x60, 0x71, 0x24, 0x61, 0x48, 0x02, 0x9c,
	0x15, 0x5c, 0xfa, 0xb7, 0xf3, 0xd6, 0xb9, 0x94, 0xcd, 0x1c, 0xb4, 0xbf, 0x15, 0x40, 0xb1, 0xd5,
	0x67, 0x70, 0x1d, 0x94, 0x15, 0x0d, 0x09, 0x4f, 0x26, 0x99, 0x05, 0x37, 0x5f, 0x0b, 0x5f, 0x80,
	0x1b, 0x31, 0xe7, 0xcc, 0xa3, 0x11, 0x55, 0x9e, 0xa4, 0xc7, 0x64, 0x5c, 0xeb, 0x73, 0xfa, 0xba,
	0x46, 0xea, 0x11, 0x55, 0x2d, 0x7a, 0x4c, 0xe0, 0x16, 0x98, 0x4f, 0x15, 0xba, 0xa0, 0xd4, 0x30,
	0x41, 0xc7, 0xe7, 0x34, 0xd1, 0xc4, 0xc3, 0x54, 0x90, 0x67, 0xe8, 0x27, 0x24, 0x21, 0x5a, 0x33,
	0x49, 0x87, 0xd3, 0x0c, 0x7b, 0x9a, 0x68, 0xe2, 0xa1, 0xde, 0xfa, 0x73, 0x8a, 0xbc, 0x11, 0x13,
	0x34, 0x75, 0xe1, 0x4c, 0xb3, 0x9f, 0x41, 0xf6, 0x2c, 0x28, 0x9b, 0xff, 0xb7, 0x27, 0xcb, 0x27,
	0xa7, 0x96, 0x05, 0x16, 0xcf, 0x6e, 0x0b, 0x5f, 0x56, 0xb8, 0x5f, 0x31, 0xa7, 0x7c, 0xbb, 0xf2,
	0xe9, 0x71, 0x40, 0x55, 0x2f, 0x69, 0x23, 0x9f, 0x87, 0x8e, 0x3e, 0xc3, 0x4e, 0xbe, 0x32, 0xbd,
	0x3d, 0x46, 0x2f, 0xb7, 0xf6, 0x4c, 0xfa, 0xf9, 0xd5, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0xf0, 0xbf, 0xd1, 0xf5, 0x06, 0x00, 0x00,
}
