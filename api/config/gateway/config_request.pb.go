// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/gateway/config_request.proto

package gateway

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
	return fileDescriptor_b882b1992a7225b6, []int{0}
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
	return fileDescriptor_b882b1992a7225b6, []int{0, 0}
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
	Mlsa                 *shared.Mlsa                           `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Tls                  *shared.TLSCredentials                 `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Service              *ConfigRequest_V1_System_Service       `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Log                  *ConfigRequest_V1_System_Log           `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	DataCollector        *ConfigRequest_V1_System_DataCollector `protobuf:"bytes,5,opt,name=data_collector,json=dataCollector,proto3" json:"data_collector,omitempty" toml:"data_collector,omitempty" mapstructure:"data_collector,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_b882b1992a7225b6, []int{0, 0, 0}
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

func (m *ConfigRequest_V1_System) GetDataCollector() *ConfigRequest_V1_System_DataCollector {
	if m != nil {
		return m.DataCollector
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host                 *wrappers.StringValue        `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrappers.Int32Value         `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	ExternalFqdn         *wrappers.StringValue        `protobuf:"bytes,3,opt,name=external_fqdn,json=externalFqdn,proto3" json:"external_fqdn,omitempty" toml:"external_fqdn,omitempty" mapstructure:"external_fqdn,omitempty"`
	GrpcPort             *wrappers.Int32Value         `protobuf:"bytes,4,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty" toml:"grpc_port,omitempty" mapstructure:"grpc_port,omitempty"`
	TrialLicenseUrl      *wrappers.StringValue        `protobuf:"bytes,5,opt,name=trial_license_url,json=trialLicenseUrl,proto3" json:"trial_license_url,omitempty" toml:"trial_license_url,omitempty" mapstructure:"trial_license_url,omitempty"`
	Log                  *ConfigRequest_V1_System_Log `protobuf:"bytes,6,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	AuthMiddleware       *wrappers.StringValue        `protobuf:"bytes,7,opt,name=auth_middleware,json=authMiddleware,proto3" json:"auth_middleware,omitempty" toml:"auth_middleware,omitempty" mapstructure:"auth_middleware,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                       `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                        `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_b882b1992a7225b6, []int{0, 0, 0, 0}
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

func (m *ConfigRequest_V1_System_Service) GetExternalFqdn() *wrappers.StringValue {
	if m != nil {
		return m.ExternalFqdn
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetGrpcPort() *wrappers.Int32Value {
	if m != nil {
		return m.GrpcPort
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetTrialLicenseUrl() *wrappers.StringValue {
	if m != nil {
		return m.TrialLicenseUrl
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

// Deprecated: Do not use.
func (m *ConfigRequest_V1_System_Service) GetAuthMiddleware() *wrappers.StringValue {
	if m != nil {
		return m.AuthMiddleware
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	Level                *wrappers.StringValue `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	Format               *wrappers.StringValue `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty" toml:"format,omitempty" mapstructure:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_b882b1992a7225b6, []int{0, 0, 0, 1}
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

func (m *ConfigRequest_V1_System_Log) GetFormat() *wrappers.StringValue {
	if m != nil {
		return m.Format
	}
	return nil
}

type ConfigRequest_V1_System_DataCollector struct {
	Limiter              *ConfigRequest_V1_System_DataCollector_Limiter `protobuf:"bytes,1,opt,name=limiter,proto3" json:"limiter,omitempty" toml:"limiter,omitempty" mapstructure:"limiter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                                          `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_DataCollector) Reset()         { *m = ConfigRequest_V1_System_DataCollector{} }
func (m *ConfigRequest_V1_System_DataCollector) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_DataCollector) ProtoMessage()    {}
func (*ConfigRequest_V1_System_DataCollector) Descriptor() ([]byte, []int) {
	return fileDescriptor_b882b1992a7225b6, []int{0, 0, 0, 2}
}

func (m *ConfigRequest_V1_System_DataCollector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_DataCollector.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_DataCollector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_DataCollector.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_DataCollector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_DataCollector.Merge(m, src)
}
func (m *ConfigRequest_V1_System_DataCollector) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_DataCollector.Size(m)
}
func (m *ConfigRequest_V1_System_DataCollector) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_DataCollector.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_DataCollector proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_DataCollector) GetLimiter() *ConfigRequest_V1_System_DataCollector_Limiter {
	if m != nil {
		return m.Limiter
	}
	return nil
}

type ConfigRequest_V1_System_DataCollector_Limiter struct {
	Disable              *wrappers.BoolValue  `protobuf:"bytes,1,opt,name=disable,proto3" json:"disable,omitempty" toml:"disable,omitempty" mapstructure:"disable,omitempty"`
	MaxInflightRequests  *wrappers.Int32Value `protobuf:"bytes,2,opt,name=max_inflight_requests,json=maxInflightRequests,proto3" json:"max_inflight_requests,omitempty" toml:"max_inflight_requests,omitempty" mapstructure:"max_inflight_requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_DataCollector_Limiter) Reset() {
	*m = ConfigRequest_V1_System_DataCollector_Limiter{}
}
func (m *ConfigRequest_V1_System_DataCollector_Limiter) String() string {
	return proto.CompactTextString(m)
}
func (*ConfigRequest_V1_System_DataCollector_Limiter) ProtoMessage() {}
func (*ConfigRequest_V1_System_DataCollector_Limiter) Descriptor() ([]byte, []int) {
	return fileDescriptor_b882b1992a7225b6, []int{0, 0, 0, 2, 0}
}

func (m *ConfigRequest_V1_System_DataCollector_Limiter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_DataCollector_Limiter.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_DataCollector_Limiter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_DataCollector_Limiter.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_DataCollector_Limiter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_DataCollector_Limiter.Merge(m, src)
}
func (m *ConfigRequest_V1_System_DataCollector_Limiter) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_DataCollector_Limiter.Size(m)
}
func (m *ConfigRequest_V1_System_DataCollector_Limiter) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_DataCollector_Limiter.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_DataCollector_Limiter proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_DataCollector_Limiter) GetDisable() *wrappers.BoolValue {
	if m != nil {
		return m.Disable
	}
	return nil
}

func (m *ConfigRequest_V1_System_DataCollector_Limiter) GetMaxInflightRequests() *wrappers.Int32Value {
	if m != nil {
		return m.MaxInflightRequests
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
	return fileDescriptor_b882b1992a7225b6, []int{0, 0, 1}
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
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.api.config.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.api.config.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.api.config.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.api.config.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.api.config.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_System_DataCollector)(nil), "chef.automate.api.config.ConfigRequest.V1.System.DataCollector")
	proto.RegisterType((*ConfigRequest_V1_System_DataCollector_Limiter)(nil), "chef.automate.api.config.ConfigRequest.V1.System.DataCollector.Limiter")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.api.config.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("config/gateway/config_request.proto", fileDescriptor_b882b1992a7225b6)
}

var fileDescriptor_b882b1992a7225b6 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x41, 0x6f, 0xd3, 0x3c,
	0x18, 0xc7, 0xd5, 0x26, 0x6b, 0x57, 0x6f, 0x5d, 0xfb, 0xfa, 0x7d, 0x5f, 0x88, 0x02, 0x42, 0x13,
	0x5c, 0xd0, 0x50, 0x13, 0xda, 0x8d, 0x03, 0x13, 0xd2, 0x44, 0x3b, 0x18, 0x9b, 0x3a, 0x31, 0xa5,
	0xb0, 0x03, 0x97, 0xc8, 0x4d, 0x9c, 0xd4, 0x92, 0x13, 0x67, 0xb6, 0xdb, 0x6d, 0x57, 0x24, 0xee,
	0x9c, 0xf8, 0x40, 0xfb, 0x04, 0xf0, 0x59, 0x76, 0x42, 0x5c, 0x50, 0x12, 0xa7, 0x50, 0x26, 0xba,
	0x55, 0xbb, 0x54, 0x75, 0xfc, 0xfc, 0xfe, 0xf9, 0xdb, 0xcf, 0x3f, 0x7a, 0xc0, 0x23, 0x8f, 0xc5,
	0x01, 0x09, 0xed, 0x10, 0x49, 0x7c, 0x8a, 0xce, 0xed, 0x7c, 0xe9, 0x72, 0x7c, 0x32, 0xc6, 0x42,
	0x5a, 0x09, 0x67, 0x92, 0x41, 0xc3, 0x1b, 0xe1, 0xc0, 0x42, 0x63, 0xc9, 0x22, 0x24, 0xb1, 0x85,
	0x12, 0x62, 0xe5, 0x75, 0xa6, 0xa9, 0x70, 0x31, 0x42, 0x1c, 0xfb, 0x76, 0x48, 0xd9, 0x10, 0xd1,
	0x9c, 0x32, 0xef, 0xce, 0xee, 0x49, 0x2a, 0xd4, 0xc6, 0x4e, 0xa1, 0xd4, 0x0a, 0x79, 0xe2, 0xd9,
	0xd9, 0x43, 0xaf, 0x15, 0xe2, 0xb8, 0x85, 0x3a, 0x2d, 0x05, 0xa1, 0x84, 0xd8, 0xa8, 0x93, 0x2e,
	0x6c, 0x14, 0xc7, 0x4c, 0x22, 0x49, 0x58, 0x5c, 0x08, 0x3c, 0x08, 0x19, 0x0b, 0x29, 0xce, 0xc9,
	0xe1, 0x38, 0xb0, 0x4f, 0x39, 0x4a, 0x12, 0xcc, 0xd5, 0xfe, 0xc3, 0x8f, 0xab, 0xa0, 0xde, 0xcb,
	0x74, 0x9c, 0xfc, 0x1c, 0x70, 0x1b, 0x94, 0x27, 0x6d, 0x43, 0x5b, 0x2f, 0x3d, 0x5e, 0xe9, 0x6c,
	0x58, 0x7f, 0x3b, 0x8e, 0x35, 0x03, 0x59, 0xc7, 0x6d, 0xa7, 0x3c, 0x69, 0x9b, 0x5f, 0x56, 0x40,
	0xf9, 0xb8, 0x0d, 0x7b, 0x40, 0x13, 0xe7, 0xc2, 0x28, 0x65, 0x1a, 0xed, 0x9b, 0x6b, 0x58, 0x83,
	0x73, 0x21, 0x71, 0xe4, 0xa4, 0x34, 0xdc, 0x05, 0x9a, 0x98, 0x78, 0x46, 0x39, 0x13, 0xe9, 0x2c,
	0x22, 0x82, 0xf9, 0x84, 0x78, 0xd8, 0x49, 0x71, 0xf3, 0x7b, 0x0d, 0x54, 0x72, 0x55, 0xb8, 0x05,
	0xf4, 0x88, 0x0a, 0xa4, 0x6c, 0xad, 0xff, 0xa1, 0x48, 0xe2, 0x80, 0xa3, 0x42, 0xf3, 0x90, 0x0a,
	0xe4, 0x64, 0xd5, 0xf0, 0x05, 0xd0, 0x24, 0x15, 0xca, 0xc6, 0xc6, 0x3c, 0xe8, 0x5d, 0x7f, 0xd0,
	0xe3, 0xd8, 0xc7, 0xb1, 0x24, 0x88, 0x0a, 0x27, 0xc5, 0xe0, 0x00, 0x54, 0x45, 0x6e, 0x47, 0xdd,
	0xe8, 0xf3, 0x85, 0x6f, 0x63, 0x7a, 0x9e, 0x42, 0x09, 0xee, 0x01, 0x8d, 0xb2, 0xd0, 0xd0, 0x33,
	0xc1, 0x67, 0x8b, 0x0b, 0xf6, 0x59, 0xe8, 0xa4, 0x0a, 0x30, 0x00, 0x6b, 0x3e, 0x92, 0xc8, 0xf5,
	0x18, 0xa5, 0xd8, 0x93, 0x8c, 0x1b, 0x4b, 0x99, 0xe6, 0xce, 0xe2, 0x9a, 0xbb, 0x48, 0xa2, 0x5e,
	0x21, 0xe3, 0xd4, 0xfd, 0xdf, 0x97, 0xe6, 0x0f, 0x0d, 0x54, 0xd5, 0x29, 0xe0, 0x53, 0xa0, 0x8f,
	0x98, 0x90, 0xaa, 0x0b, 0xf7, 0xad, 0x3c, 0x9f, 0x56, 0x91, 0x4f, 0x6b, 0x20, 0x39, 0x89, 0xc3,
	0x63, 0x44, 0xc7, 0xd8, 0xc9, 0x2a, 0xe1, 0x2b, 0xa0, 0x27, 0x8c, 0x4b, 0xd5, 0x82, 0x7b, 0x57,
	0x88, 0xfd, 0x58, 0x6e, 0x76, 0x32, 0xa0, 0xfb, 0xdf, 0xc5, 0xa5, 0xd1, 0x04, 0xfa, 0x48, 0xca,
	0xa4, 0xf9, 0xb5, 0x61, 0x2e, 0xa5, 0x7f, 0x84, 0x93, 0xe1, 0xf0, 0x25, 0xa8, 0xe3, 0x33, 0x89,
	0x79, 0x8c, 0xa8, 0x1b, 0x9c, 0xf8, 0xb1, 0x6a, 0xc8, 0x7c, 0x07, 0xab, 0x05, 0xf2, 0xfa, 0xc4,
	0x8f, 0xe1, 0x11, 0xa8, 0xa5, 0x9f, 0xa1, 0x9b, 0xd9, 0xd1, 0xaf, 0xb7, 0x73, 0xe7, 0xe2, 0xd2,
	0x80, 0xd3, 0x04, 0x34, 0xbf, 0x35, 0x4c, 0x3d, 0xe5, 0x9d, 0xe5, 0xf4, 0xf7, 0x28, 0x35, 0xf5,
	0x06, 0xfc, 0x23, 0x39, 0x41, 0xd4, 0xa5, 0xc4, 0xc3, 0xb1, 0xc0, 0xee, 0x98, 0x53, 0xd5, 0x84,
	0xf9, 0xc6, 0x1a, 0x19, 0xd6, 0xcf, 0xa9, 0xf7, 0x9c, 0x16, 0xa1, 0xa8, 0xdc, 0x3a, 0x14, 0xfb,
	0xa0, 0x81, 0xc6, 0x72, 0xe4, 0x46, 0xc4, 0xf7, 0x29, 0x3e, 0x45, 0x1c, 0x1b, 0xd5, 0xeb, 0x0d,
	0x75, 0xcb, 0x46, 0xc9, 0x59, 0x4b, 0xc1, 0xc3, 0x29, 0x77, 0xa0, 0x2f, 0xd7, 0x9a, 0xc0, 0x64,
	0x40, 0xeb, 0xb3, 0x10, 0x76, 0xc0, 0x12, 0xc5, 0x13, 0x4c, 0x6f, 0xd4, 0xf9, 0xbc, 0x14, 0x6e,
	0x81, 0x4a, 0xc0, 0x78, 0x84, 0x8a, 0xe6, 0xcf, 0x87, 0x54, 0xad, 0xf9, 0xa9, 0x0c, 0xea, 0x33,
	0x79, 0x84, 0x08, 0x54, 0x29, 0x89, 0x88, 0xc4, 0x5c, 0xbd, 0x7d, 0xef, 0x96, 0x09, 0xb7, 0xfa,
	0xb9, 0x9c, 0x53, 0xe8, 0x9a, 0x9f, 0x4b, 0xa0, 0xaa, 0x1e, 0xc2, 0x2d, 0x50, 0xf5, 0x89, 0x40,
	0x43, 0x8a, 0xd5, 0xeb, 0xcc, 0x2b, 0xbe, 0xbb, 0x8c, 0xd1, 0xdc, 0x75, 0x51, 0x0a, 0xdf, 0x82,
	0xff, 0x23, 0x74, 0xe6, 0x92, 0x38, 0xa0, 0x24, 0x1c, 0xc9, 0x62, 0xb0, 0x88, 0x1b, 0x04, 0xdf,
	0xf9, 0x37, 0x42, 0x67, 0xfb, 0x0a, 0x54, 0xd6, 0x85, 0x59, 0x9b, 0x7e, 0x75, 0xdb, 0x79, 0x12,
	0x9b, 0xbf, 0xa6, 0x49, 0x3e, 0xc3, 0x0e, 0xf4, 0xe5, 0x52, 0x53, 0xeb, 0xb6, 0x3e, 0x3c, 0x09,
	0x89, 0x1c, 0x8d, 0x87, 0x96, 0xc7, 0x22, 0x3b, 0xbd, 0x19, 0xbb, 0xa8, 0xcc, 0xa6, 0xcb, 0xec,
	0xe0, 0x1b, 0x56, 0x32, 0x07, 0x9b, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xff, 0xa7, 0x14, 0x1b,
	0x11, 0x07, 0x00, 0x00,
}
