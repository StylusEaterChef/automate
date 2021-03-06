// Code generated by protoc-gen-go. DO NOT EDIT.
// source: external/infra_proxy/response/servers.proto

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

type CreateServer struct {
	// Chef Infra Server.
	Server               *Server  `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServer) Reset()         { *m = CreateServer{} }
func (m *CreateServer) String() string { return proto.CompactTextString(m) }
func (*CreateServer) ProtoMessage()    {}
func (*CreateServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_3049f47ec4ad0dd8, []int{0}
}

func (m *CreateServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServer.Unmarshal(m, b)
}
func (m *CreateServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServer.Marshal(b, m, deterministic)
}
func (m *CreateServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServer.Merge(m, src)
}
func (m *CreateServer) XXX_Size() int {
	return xxx_messageInfo_CreateServer.Size(m)
}
func (m *CreateServer) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServer.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServer proto.InternalMessageInfo

func (m *CreateServer) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

type DeleteServer struct {
	// Chef Infra Server.
	Server               *Server  `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteServer) Reset()         { *m = DeleteServer{} }
func (m *DeleteServer) String() string { return proto.CompactTextString(m) }
func (*DeleteServer) ProtoMessage()    {}
func (*DeleteServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_3049f47ec4ad0dd8, []int{1}
}

func (m *DeleteServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteServer.Unmarshal(m, b)
}
func (m *DeleteServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteServer.Marshal(b, m, deterministic)
}
func (m *DeleteServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteServer.Merge(m, src)
}
func (m *DeleteServer) XXX_Size() int {
	return xxx_messageInfo_DeleteServer.Size(m)
}
func (m *DeleteServer) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteServer.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteServer proto.InternalMessageInfo

func (m *DeleteServer) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

type UpdateServer struct {
	// Chef Infra Server.
	Server               *Server  `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateServer) Reset()         { *m = UpdateServer{} }
func (m *UpdateServer) String() string { return proto.CompactTextString(m) }
func (*UpdateServer) ProtoMessage()    {}
func (*UpdateServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_3049f47ec4ad0dd8, []int{2}
}

func (m *UpdateServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateServer.Unmarshal(m, b)
}
func (m *UpdateServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateServer.Marshal(b, m, deterministic)
}
func (m *UpdateServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateServer.Merge(m, src)
}
func (m *UpdateServer) XXX_Size() int {
	return xxx_messageInfo_UpdateServer.Size(m)
}
func (m *UpdateServer) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateServer.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateServer proto.InternalMessageInfo

func (m *UpdateServer) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

type GetServers struct {
	// List of Chef Infra Servers.
	Servers              []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetServers) Reset()         { *m = GetServers{} }
func (m *GetServers) String() string { return proto.CompactTextString(m) }
func (*GetServers) ProtoMessage()    {}
func (*GetServers) Descriptor() ([]byte, []int) {
	return fileDescriptor_3049f47ec4ad0dd8, []int{3}
}

func (m *GetServers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServers.Unmarshal(m, b)
}
func (m *GetServers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServers.Marshal(b, m, deterministic)
}
func (m *GetServers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServers.Merge(m, src)
}
func (m *GetServers) XXX_Size() int {
	return xxx_messageInfo_GetServers.Size(m)
}
func (m *GetServers) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServers.DiscardUnknown(m)
}

var xxx_messageInfo_GetServers proto.InternalMessageInfo

func (m *GetServers) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

type GetServer struct {
	// Chef Infra Server.
	Server               *Server  `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServer) Reset()         { *m = GetServer{} }
func (m *GetServer) String() string { return proto.CompactTextString(m) }
func (*GetServer) ProtoMessage()    {}
func (*GetServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_3049f47ec4ad0dd8, []int{4}
}

func (m *GetServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServer.Unmarshal(m, b)
}
func (m *GetServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServer.Marshal(b, m, deterministic)
}
func (m *GetServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServer.Merge(m, src)
}
func (m *GetServer) XXX_Size() int {
	return xxx_messageInfo_GetServer.Size(m)
}
func (m *GetServer) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServer.DiscardUnknown(m)
}

var xxx_messageInfo_GetServer proto.InternalMessageInfo

func (m *GetServer) GetServer() *Server {
	if m != nil {
		return m.Server
	}
	return nil
}

type Server struct {
	// Chef Infra Server ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Chef Infra Server name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Chef Infra Server FQDN.
	Fqdn string `protobuf:"bytes,4,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// Chef Infra Server IP address.
	IpAddress string `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// Chef organizations count associated with Chef Infra Server.
	OrgsCount            int32    `protobuf:"varint,6,opt,name=orgs_count,json=orgsCount,proto3" json:"orgs_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_3049f47ec4ad0dd8, []int{5}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Server) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Server) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *Server) GetOrgsCount() int32 {
	if m != nil {
		return m.OrgsCount
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateServer)(nil), "chef.automate.api.infra_proxy.response.CreateServer")
	proto.RegisterType((*DeleteServer)(nil), "chef.automate.api.infra_proxy.response.DeleteServer")
	proto.RegisterType((*UpdateServer)(nil), "chef.automate.api.infra_proxy.response.UpdateServer")
	proto.RegisterType((*GetServers)(nil), "chef.automate.api.infra_proxy.response.GetServers")
	proto.RegisterType((*GetServer)(nil), "chef.automate.api.infra_proxy.response.GetServer")
	proto.RegisterType((*Server)(nil), "chef.automate.api.infra_proxy.response.Server")
}

func init() {
	proto.RegisterFile("external/infra_proxy/response/servers.proto", fileDescriptor_3049f47ec4ad0dd8)
}

var fileDescriptor_3049f47ec4ad0dd8 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0x87, 0x69, 0xbd, 0xab, 0x34, 0x8a, 0x43, 0xa6, 0x2c, 0x42, 0xb9, 0x41, 0x0e, 0x84, 0x04,
	0x74, 0x13, 0x17, 0x3d, 0x51, 0xe7, 0x1e, 0x3a, 0xb8, 0x94, 0x5c, 0xf3, 0x7a, 0x17, 0xb8, 0x26,
	0x31, 0x49, 0xe5, 0x5c, 0xfc, 0xdb, 0x25, 0x69, 0x2b, 0x4e, 0x22, 0xe2, 0x6d, 0x8f, 0xef, 0xf1,
	0xfb, 0x78, 0xef, 0xf1, 0xd0, 0x39, 0xec, 0x3c, 0x58, 0xc5, 0xb7, 0x4c, 0xaa, 0xc6, 0xf2, 0xca,
	0x58, 0xbd, 0x7b, 0x67, 0x16, 0x9c, 0xd1, 0xca, 0x01, 0x73, 0x60, 0xdf, 0xc0, 0x3a, 0x6a, 0xac,
	0xf6, 0x1a, 0x9f, 0xd5, 0x1b, 0x68, 0x28, 0xef, 0xbc, 0x6e, 0xb9, 0x07, 0xca, 0x8d, 0xa4, 0xdf,
	0x52, 0x74, 0x4c, 0xcd, 0x9e, 0xd1, 0xf1, 0xc2, 0x02, 0xf7, 0xb0, 0x8c, 0x71, 0x7c, 0x8f, 0xb2,
	0x5e, 0x44, 0x92, 0x22, 0x99, 0x1f, 0x5d, 0x50, 0xfa, 0x3b, 0x11, 0xed, 0xf3, 0xe5, 0x90, 0x0e,
	0xde, 0x3b, 0xd8, 0xc2, 0x3e, 0xbc, 0x4f, 0x46, 0xec, 0x63, 0x5e, 0xf4, 0x00, 0xbe, 0x87, 0x0e,
	0x3f, 0xa2, 0xc3, 0xe1, 0x9c, 0x24, 0x29, 0x0e, 0xfe, 0xa0, 0x1d, 0xe3, 0xb3, 0x25, 0xca, 0xbf,
	0xbc, 0xff, 0x36, 0xec, 0x07, 0xca, 0x06, 0xe3, 0x09, 0x4a, 0xa5, 0x88, 0xb6, 0xbc, 0x4c, 0xa5,
	0xc0, 0x18, 0x4d, 0x14, 0x6f, 0x81, 0xa4, 0x91, 0xc4, 0x3a, 0xb0, 0xe6, 0x55, 0x28, 0x32, 0xe9,
	0x59, 0xa8, 0xf1, 0x29, 0x42, 0xd2, 0x54, 0x5c, 0x08, 0x0b, 0xce, 0x91, 0x69, 0xec, 0xe4, 0xd2,
	0xdc, 0xf4, 0x20, 0xb4, 0xb5, 0x5d, 0xbb, 0xaa, 0xd6, 0x9d, 0xf2, 0x24, 0x2b, 0x92, 0xf9, 0xb4,
	0xcc, 0x03, 0x59, 0x04, 0x70, 0x7b, 0xfd, 0x72, 0xb5, 0x96, 0x7e, 0xd3, 0xad, 0x68, 0xad, 0x5b,
	0x16, 0x76, 0x60, 0xe3, 0x0e, 0x8c, 0x1b, 0xc9, 0x7e, 0x7c, 0xd4, 0x55, 0x16, 0x3f, 0xf4, 0xf2,
	0x33, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x31, 0xa2, 0x2e, 0xd0, 0x02, 0x00, 0x00,
}
