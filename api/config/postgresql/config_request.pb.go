// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/postgresql/config_request.proto

package pg

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
	return fileDescriptor_9f765ec9632e5fcf, []int{0}
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
	return fileDescriptor_9f765ec9632e5fcf, []int{0, 0}
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

type ConfigRequest_V1_Service struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f765ec9632e5fcf, []int{0, 0, 0}
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

type ConfigRequest_V1_System struct {
	Service              *ConfigRequest_V1_System_Service   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Pg                   *ConfigRequest_V1_System_PGConfig  `protobuf:"bytes,2,opt,name=pg,proto3" json:"pg,omitempty" toml:"pg,omitempty" mapstructure:"pg,omitempty"`
	Logger               *ConfigRequest_V1_System_Logger    `protobuf:"bytes,3,opt,name=logger,proto3" json:"logger,omitempty" toml:"logger,omitempty" mapstructure:"logger,omitempty"`
	Superuser            *ConfigRequest_V1_System_Superuser `protobuf:"bytes,4,opt,name=superuser,proto3" json:"superuser,omitempty" toml:"superuser,omitempty" mapstructure:"superuser,omitempty"`
	Tls                  *shared.TLSCredentials             `protobuf:"bytes,5,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Disable              *wrappers.BoolValue                `protobuf:"bytes,6,opt,name=disable,proto3" json:"disable,omitempty" toml:"disable,omitempty" mapstructure:"disable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                              `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f765ec9632e5fcf, []int{0, 0, 1}
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

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetPg() *ConfigRequest_V1_System_PGConfig {
	if m != nil {
		return m.Pg
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLogger() *ConfigRequest_V1_System_Logger {
	if m != nil {
		return m.Logger
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetSuperuser() *ConfigRequest_V1_System_Superuser {
	if m != nil {
		return m.Superuser
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetDisable() *wrappers.BoolValue {
	if m != nil {
		return m.Disable
	}
	return nil
}

type ConfigRequest_V1_System_Logger struct {
	Level                *wrappers.StringValue `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Logger) Reset()         { *m = ConfigRequest_V1_System_Logger{} }
func (m *ConfigRequest_V1_System_Logger) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Logger) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Logger) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f765ec9632e5fcf, []int{0, 0, 1, 0}
}

func (m *ConfigRequest_V1_System_Logger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Logger.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Logger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Logger.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Logger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Logger.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Logger) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Logger.Size(m)
}
func (m *ConfigRequest_V1_System_Logger) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Logger.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Logger proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Logger) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f765ec9632e5fcf, []int{0, 0, 1, 1}
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

type ConfigRequest_V1_System_PGConfig struct {
	MaxWalSize                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=max_wal_size,json=maxWalSize,proto3" json:"max_wal_size,omitempty" toml:"max_wal_size,omitempty" mapstructure:"max_wal_size,omitempty"`
	MinWalSize                 *wrappers.StringValue `protobuf:"bytes,4,opt,name=min_wal_size,json=minWalSize,proto3" json:"min_wal_size,omitempty" toml:"min_wal_size,omitempty" mapstructure:"min_wal_size,omitempty"`
	WalKeepSegments            *wrappers.Int32Value  `protobuf:"bytes,5,opt,name=wal_keep_segments,json=walKeepSegments,proto3" json:"wal_keep_segments,omitempty" toml:"wal_keep_segments,omitempty" mapstructure:"wal_keep_segments,omitempty"`
	CheckpointTimeout          *wrappers.StringValue `protobuf:"bytes,6,opt,name=checkpoint_timeout,json=checkpointTimeout,proto3" json:"checkpoint_timeout,omitempty" toml:"checkpoint_timeout,omitempty" mapstructure:"checkpoint_timeout,omitempty"`
	CheckpointCompletionTarget *wrappers.FloatValue  `protobuf:"bytes,7,opt,name=checkpoint_completion_target,json=checkpointCompletionTarget,proto3" json:"checkpoint_completion_target,omitempty" toml:"checkpoint_completion_target,omitempty" mapstructure:"checkpoint_completion_target,omitempty"`
	CheckpointWarning          *wrappers.StringValue `protobuf:"bytes,8,opt,name=checkpoint_warning,json=checkpointWarning,proto3" json:"checkpoint_warning,omitempty" toml:"checkpoint_warning,omitempty" mapstructure:"checkpoint_warning,omitempty"`
	EffectiveCacheSize         *wrappers.StringValue `protobuf:"bytes,9,opt,name=effective_cache_size,json=effectiveCacheSize,proto3" json:"effective_cache_size,omitempty" toml:"effective_cache_size,omitempty" mapstructure:"effective_cache_size,omitempty"`
	MaxConnections             *wrappers.Int32Value  `protobuf:"bytes,10,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty" toml:"max_connections,omitempty" mapstructure:"max_connections,omitempty"`
	MaxLocksPerTransaction     *wrappers.Int32Value  `protobuf:"bytes,11,opt,name=max_locks_per_transaction,json=maxLocksPerTransaction,proto3" json:"max_locks_per_transaction,omitempty" toml:"max_locks_per_transaction,omitempty" mapstructure:"max_locks_per_transaction,omitempty"`
	Md5AuthCidrAddresses       []string              `protobuf:"bytes,12,rep,name=md5_auth_cidr_addresses,json=md5AuthCidrAddresses,proto3" json:"md5_auth_cidr_addresses,omitempty" toml:"md5_auth_cidr_addresses,omitempty" mapstructure:"md5_auth_cidr_addresses,omitempty"`
	SharedBuffers              *wrappers.StringValue `protobuf:"bytes,13,opt,name=shared_buffers,json=sharedBuffers,proto3" json:"shared_buffers,omitempty" toml:"shared_buffers,omitempty" mapstructure:"shared_buffers,omitempty"`
	WorkMem                    *wrappers.StringValue `protobuf:"bytes,14,opt,name=work_mem,json=workMem,proto3" json:"work_mem,omitempty" toml:"work_mem,omitempty" mapstructure:"work_mem,omitempty"`
	SslCiphers                 *wrappers.StringValue `protobuf:"bytes,15,opt,name=ssl_ciphers,json=sslCiphers,proto3" json:"ssl_ciphers,omitempty" toml:"ssl_ciphers,omitempty" mapstructure:"ssl_ciphers,omitempty"`
	LogDisconnections          *wrappers.StringValue `protobuf:"bytes,16,opt,name=log_disconnections,json=logDisconnections,proto3" json:"log_disconnections,omitempty" toml:"log_disconnections,omitempty" mapstructure:"log_disconnections,omitempty"`
	ClientMinMessages          *wrappers.StringValue `protobuf:"bytes,17,opt,name=client_min_messages,json=clientMinMessages,proto3" json:"client_min_messages,omitempty" toml:"client_min_messages,omitempty" mapstructure:"client_min_messages,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized           []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache              int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_PGConfig) Reset()         { *m = ConfigRequest_V1_System_PGConfig{} }
func (m *ConfigRequest_V1_System_PGConfig) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_PGConfig) ProtoMessage()    {}
func (*ConfigRequest_V1_System_PGConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f765ec9632e5fcf, []int{0, 0, 1, 2}
}

func (m *ConfigRequest_V1_System_PGConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_PGConfig.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_PGConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_PGConfig.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_PGConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_PGConfig.Merge(m, src)
}
func (m *ConfigRequest_V1_System_PGConfig) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_PGConfig.Size(m)
}
func (m *ConfigRequest_V1_System_PGConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_PGConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_PGConfig proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_PGConfig) GetMaxWalSize() *wrappers.StringValue {
	if m != nil {
		return m.MaxWalSize
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetMinWalSize() *wrappers.StringValue {
	if m != nil {
		return m.MinWalSize
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetWalKeepSegments() *wrappers.Int32Value {
	if m != nil {
		return m.WalKeepSegments
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetCheckpointTimeout() *wrappers.StringValue {
	if m != nil {
		return m.CheckpointTimeout
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetCheckpointCompletionTarget() *wrappers.FloatValue {
	if m != nil {
		return m.CheckpointCompletionTarget
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetCheckpointWarning() *wrappers.StringValue {
	if m != nil {
		return m.CheckpointWarning
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetEffectiveCacheSize() *wrappers.StringValue {
	if m != nil {
		return m.EffectiveCacheSize
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetMaxConnections() *wrappers.Int32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetMaxLocksPerTransaction() *wrappers.Int32Value {
	if m != nil {
		return m.MaxLocksPerTransaction
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetMd5AuthCidrAddresses() []string {
	if m != nil {
		return m.Md5AuthCidrAddresses
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetSharedBuffers() *wrappers.StringValue {
	if m != nil {
		return m.SharedBuffers
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetWorkMem() *wrappers.StringValue {
	if m != nil {
		return m.WorkMem
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetSslCiphers() *wrappers.StringValue {
	if m != nil {
		return m.SslCiphers
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetLogDisconnections() *wrappers.StringValue {
	if m != nil {
		return m.LogDisconnections
	}
	return nil
}

func (m *ConfigRequest_V1_System_PGConfig) GetClientMinMessages() *wrappers.StringValue {
	if m != nil {
		return m.ClientMinMessages
	}
	return nil
}

type ConfigRequest_V1_System_Superuser struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Superuser) Reset()         { *m = ConfigRequest_V1_System_Superuser{} }
func (m *ConfigRequest_V1_System_Superuser) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Superuser) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Superuser) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f765ec9632e5fcf, []int{0, 0, 1, 3}
}

func (m *ConfigRequest_V1_System_Superuser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Superuser.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Superuser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Superuser.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Superuser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Superuser.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Superuser) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Superuser.Size(m)
}
func (m *ConfigRequest_V1_System_Superuser) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Superuser.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Superuser proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Superuser) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.infra.postgresql.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.infra.postgresql.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.infra.postgresql.ConfigRequest.V1.Service")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.infra.postgresql.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Logger)(nil), "chef.automate.infra.postgresql.ConfigRequest.V1.System.Logger")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.infra.postgresql.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_PGConfig)(nil), "chef.automate.infra.postgresql.ConfigRequest.V1.System.PGConfig")
	proto.RegisterType((*ConfigRequest_V1_System_Superuser)(nil), "chef.automate.infra.postgresql.ConfigRequest.V1.System.Superuser")
}

func init() {
	proto.RegisterFile("config/postgresql/config_request.proto", fileDescriptor_9f765ec9632e5fcf)
}

var fileDescriptor_9f765ec9632e5fcf = []byte{
	// 935 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xd1, 0x6e, 0xdb, 0x36,
	0x14, 0x40, 0x11, 0xdb, 0x89, 0x6d, 0xa6, 0x49, 0x1c, 0xb6, 0x5d, 0x55, 0xad, 0x28, 0x8a, 0x3d,
	0x0c, 0x45, 0x81, 0xc8, 0xb5, 0xb3, 0xa0, 0xc3, 0xd6, 0xb5, 0x4d, 0x5c, 0xb4, 0x68, 0x9a, 0x0c,
	0x81, 0x1c, 0xb8, 0xd8, 0x80, 0x41, 0xa0, 0xa5, 0x6b, 0x9a, 0x08, 0x45, 0x2a, 0x24, 0x65, 0xa7,
	0x7d, 0xdd, 0x4f, 0xed, 0x65, 0x2f, 0xfd, 0x81, 0x7d, 0xc0, 0xde, 0xf7, 0x03, 0xfd, 0x81, 0x41,
	0xa2, 0x64, 0xb7, 0xcb, 0x9a, 0x3a, 0x79, 0x93, 0xc8, 0x7b, 0x0e, 0xaf, 0x74, 0xaf, 0x28, 0xa2,
	0x6f, 0x43, 0x29, 0x46, 0x8c, 0xb6, 0x13, 0xa9, 0x0d, 0x55, 0xa0, 0x4f, 0x79, 0xdb, 0x8e, 0x04,
	0x0a, 0x4e, 0x53, 0xd0, 0xc6, 0x4b, 0x94, 0x34, 0x12, 0xdf, 0x0d, 0xc7, 0x30, 0xf2, 0x48, 0x6a,
	0x64, 0x4c, 0x0c, 0x78, 0x4c, 0x8c, 0x14, 0xf1, 0xe6, 0x90, 0x7b, 0xab, 0xf0, 0xe8, 0x31, 0x51,
	0x10, 0xb5, 0x0d, 0xd7, 0x16, 0x74, 0x9f, 0x96, 0xcc, 0x16, 0x55, 0x49, 0xd8, 0xce, 0x07, 0xc3,
	0x2d, 0x0a, 0x62, 0x8b, 0x74, 0xb7, 0x0a, 0x88, 0x24, 0xac, 0x4d, 0xba, 0xd9, 0x4d, 0x9b, 0x08,
	0x21, 0x0d, 0x31, 0x4c, 0x8a, 0x52, 0x70, 0x97, 0x4a, 0x49, 0x39, 0x58, 0x72, 0x98, 0x8e, 0xda,
	0x53, 0x45, 0x92, 0x04, 0x54, 0x31, 0xff, 0xcd, 0xdf, 0x2d, 0xb4, 0xd6, 0xcb, 0x3d, 0xbe, 0xcd,
	0x18, 0x3f, 0x43, 0x95, 0x49, 0xc7, 0xa9, 0xde, 0x5b, 0xba, 0xbf, 0xda, 0x7d, 0xe8, 0x5d, 0x9c,
	0xb8, 0xf7, 0x09, 0xea, 0x0d, 0x3a, 0x7e, 0x65, 0xd2, 0x71, 0xff, 0xd9, 0x40, 0x95, 0x41, 0x07,
	0xbf, 0x42, 0x55, 0xfd, 0x56, 0x3b, 0x4b, 0xb9, 0xe9, 0xd1, 0x65, 0x4d, 0x5e, 0xff, 0xad, 0x36,
	0x10, 0xfb, 0x99, 0x03, 0xef, 0xa3, 0xaa, 0x9e, 0x84, 0x4e, 0x25, 0x57, 0x7d, 0x7f, 0x79, 0x15,
	0xa8, 0x09, 0x0b, 0xc1, 0xcf, 0x24, 0x6e, 0x13, 0xd5, 0x8b, 0x7b, 0xf7, 0xcf, 0x75, 0xb4, 0x62,
	0x97, 0xc1, 0xbf, 0xa0, 0xba, 0xb6, 0xa3, 0x45, 0xc2, 0x4f, 0xaf, 0x98, 0xf0, 0x6c, 0xb1, 0xd2,
	0x87, 0x8f, 0x50, 0x25, 0xa1, 0x45, 0xee, 0xcf, 0xae, 0x6a, 0x3d, 0x7a, 0x59, 0xcc, 0x54, 0x12,
	0x8a, 0x07, 0x68, 0x85, 0x4b, 0x4a, 0x41, 0x15, 0x65, 0x7a, 0x72, 0x55, 0xeb, 0x41, 0x6e, 0xf1,
	0x0b, 0x1b, 0x0e, 0x50, 0x53, 0xa7, 0x09, 0xa8, 0x54, 0x83, 0x72, 0x6a, 0xb9, 0x7a, 0xf7, 0xca,
	0xaf, 0xa1, 0x14, 0xf9, 0x73, 0x27, 0x7e, 0x8c, 0xaa, 0x86, 0x6b, 0x67, 0x39, 0x57, 0x3f, 0xf8,
	0x5f, 0xb5, 0x6d, 0x6a, 0xef, 0xf8, 0xa0, 0xdf, 0x53, 0x10, 0x81, 0x30, 0x8c, 0x70, 0xed, 0x67,
	0x18, 0xfe, 0x0e, 0xd5, 0x23, 0xa6, 0xc9, 0x90, 0x83, 0xb3, 0x92, 0x1b, 0x5c, 0xcf, 0x76, 0xb7,
	0x57, 0x76, 0xb7, 0xb7, 0x27, 0x25, 0x1f, 0x10, 0x9e, 0x82, 0x5f, 0x86, 0xba, 0x8f, 0xd1, 0x8a,
	0x7d, 0x4c, 0xdc, 0x45, 0xcb, 0x1c, 0x26, 0xc0, 0x8b, 0x0a, 0xdf, 0x39, 0x47, 0xf7, 0x8d, 0x62,
	0x82, 0x5a, 0xde, 0x86, 0xba, 0xbf, 0x2f, 0xcd, 0xda, 0x05, 0x3f, 0x44, 0xb5, 0xb1, 0xd4, 0x66,
	0x21, 0x3c, 0x8f, 0xc4, 0x2f, 0x50, 0x2d, 0x91, 0xca, 0x14, 0xc5, 0xff, 0xfa, 0x1c, 0xf1, 0x4a,
	0x98, 0xed, 0x6e, 0x0e, 0xec, 0xdd, 0x7c, 0xff, 0xc1, 0xd9, 0x44, 0xcb, 0x09, 0xd5, 0xa7, 0xbc,
	0xf5, 0xc7, 0x03, 0xd7, 0x5e, 0xf9, 0x39, 0xef, 0xfe, 0xd5, 0x40, 0x8d, 0xb2, 0x03, 0xf0, 0x13,
	0x74, 0x2d, 0x26, 0x67, 0xc1, 0x94, 0xf0, 0x40, 0xb3, 0x77, 0x50, 0xf4, 0xc0, 0xc5, 0xe9, 0xa0,
	0x98, 0x9c, 0xbd, 0x21, 0xbc, 0xcf, 0xde, 0x41, 0xce, 0x33, 0x31, 0xe7, 0x6b, 0x0b, 0xf1, 0x4c,
	0x94, 0xfc, 0x4b, 0xb4, 0x99, 0xb1, 0x27, 0x00, 0x49, 0xa0, 0x81, 0xc6, 0x20, 0x4c, 0x59, 0xd2,
	0x8b, 0x9e, 0xd0, 0xdf, 0x98, 0x12, 0xfe, 0x1a, 0x20, 0xe9, 0x17, 0x0c, 0x7e, 0x8d, 0x70, 0x38,
	0x86, 0xf0, 0x24, 0x91, 0x4c, 0x98, 0xc0, 0xb0, 0x18, 0x64, 0x6a, 0x8a, 0xd2, 0x5e, 0x9c, 0xce,
	0xe6, 0x9c, 0x3b, 0xb6, 0x18, 0xfe, 0x0d, 0xdd, 0xf9, 0x48, 0x16, 0xca, 0x38, 0xe1, 0x90, 0x6d,
	0x84, 0x81, 0x21, 0x8a, 0x82, 0x71, 0xea, 0x9f, 0x49, 0xf0, 0x05, 0x97, 0xc4, 0x58, 0xab, 0x3b,
	0x17, 0xf4, 0x66, 0xfc, 0x71, 0x8e, 0xff, 0x27, 0xd7, 0x29, 0x51, 0x82, 0x09, 0xea, 0x34, 0x2e,
	0x97, 0xeb, 0x1b, 0x8b, 0xe1, 0x9f, 0xd1, 0x0d, 0x18, 0x8d, 0x20, 0x34, 0x6c, 0x02, 0x41, 0x48,
	0xc2, 0x31, 0xd8, 0x4a, 0x34, 0x17, 0xd0, 0xe1, 0x19, 0xd9, 0xcb, 0xc0, 0xbc, 0x22, 0xcf, 0xd1,
	0x46, 0xd6, 0x11, 0xa1, 0x14, 0x22, 0x9b, 0x92, 0x42, 0x3b, 0xe8, 0xcb, 0xf5, 0x58, 0x8f, 0xc9,
	0x59, 0x6f, 0x8e, 0xe0, 0x01, 0xba, 0x9d, 0x59, 0xb8, 0x0c, 0x4f, 0x74, 0x90, 0x80, 0x0a, 0x8c,
	0x22, 0x42, 0x93, 0x7c, 0xd6, 0x59, 0xfd, 0xb2, 0xef, 0xab, 0x98, 0x9c, 0x1d, 0x64, 0xf0, 0x11,
	0xa8, 0xe3, 0x39, 0x8a, 0x77, 0xd0, 0xad, 0x38, 0xda, 0x09, 0x48, 0x6a, 0xc6, 0x41, 0xc8, 0x22,
	0x15, 0x90, 0x28, 0x52, 0xa0, 0x35, 0x68, 0xe7, 0xda, 0xbd, 0xea, 0xfd, 0xa6, 0x7f, 0x23, 0x8e,
	0x76, 0x76, 0x53, 0x33, 0xee, 0xb1, 0x48, 0xed, 0x96, 0x73, 0xb8, 0x87, 0xd6, 0xed, 0xef, 0x30,
	0x18, 0xa6, 0xa3, 0x11, 0x28, 0xed, 0xac, 0x2d, 0xf0, 0x7a, 0xd6, 0x2c, 0xb3, 0x67, 0x11, 0xfc,
	0x08, 0x35, 0xa6, 0x52, 0x9d, 0x04, 0x31, 0xc4, 0xce, 0xfa, 0x02, 0x78, 0x3d, 0x8b, 0x3e, 0x84,
	0x18, 0xff, 0x84, 0x56, 0xb5, 0xe6, 0x41, 0xc8, 0x92, 0x71, 0xb6, 0xf4, 0xc6, 0x22, 0xdf, 0x88,
	0xd6, 0xbc, 0x67, 0xe3, 0xb3, 0x76, 0xe1, 0x92, 0x06, 0x11, 0xd3, 0x1f, 0x17, 0xa5, 0xb5, 0x48,
	0xbb, 0x70, 0x49, 0x9f, 0x7f, 0x82, 0xe1, 0x03, 0x74, 0x3d, 0xe4, 0x0c, 0x84, 0x09, 0xb2, 0xef,
	0x36, 0x06, 0xad, 0x09, 0x05, 0xed, 0x6c, 0x2e, 0xd4, 0x7c, 0x39, 0x78, 0xc8, 0xc4, 0x61, 0x81,
	0xed, 0xd7, 0x1a, 0x4b, 0xad, 0xaa, 0xdb, 0x43, 0xcd, 0xd9, 0x0e, 0x9d, 0x6d, 0x6c, 0x82, 0xc4,
	0xb0, 0xd8, 0xc6, 0x96, 0x45, 0xee, 0xd7, 0x1a, 0x95, 0x56, 0xf5, 0x87, 0xdb, 0xef, 0x3f, 0x38,
	0x37, 0xd1, 0xf5, 0xd9, 0x19, 0x65, 0xfe, 0x5b, 0xb0, 0xab, 0xec, 0x6d, 0xff, 0xda, 0xa1, 0xcc,
	0x8c, 0xd3, 0xa1, 0x17, 0xca, 0xb8, 0x9d, 0x6d, 0xf7, 0xed, 0x32, 0x38, 0x3f, 0xb6, 0x9c, 0x3b,
	0x3e, 0xfd, 0x98, 0xd0, 0xe1, 0x4a, 0xbe, 0xee, 0xf6, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4d,
	0x00, 0x39, 0x94, 0x5d, 0x09, 0x00, 0x00,
}
