// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interservice/deployment/certificate_authority.proto

package deployment

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RootCertRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *RootCertRequest) Reset()         { *m = RootCertRequest{} }
func (m *RootCertRequest) String() string { return proto.CompactTextString(m) }
func (*RootCertRequest) ProtoMessage()    {}
func (*RootCertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc6d195afcbac2c5, []int{0}
}

func (m *RootCertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RootCertRequest.Unmarshal(m, b)
}
func (m *RootCertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RootCertRequest.Marshal(b, m, deterministic)
}
func (m *RootCertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RootCertRequest.Merge(m, src)
}
func (m *RootCertRequest) XXX_Size() int {
	return xxx_messageInfo_RootCertRequest.Size(m)
}
func (m *RootCertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RootCertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RootCertRequest proto.InternalMessageInfo

type RootCertResponse struct {
	Cert                 string   `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty" toml:"cert,omitempty" mapstructure:"cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *RootCertResponse) Reset()         { *m = RootCertResponse{} }
func (m *RootCertResponse) String() string { return proto.CompactTextString(m) }
func (*RootCertResponse) ProtoMessage()    {}
func (*RootCertResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc6d195afcbac2c5, []int{1}
}

func (m *RootCertResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RootCertResponse.Unmarshal(m, b)
}
func (m *RootCertResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RootCertResponse.Marshal(b, m, deterministic)
}
func (m *RootCertResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RootCertResponse.Merge(m, src)
}
func (m *RootCertResponse) XXX_Size() int {
	return xxx_messageInfo_RootCertResponse.Size(m)
}
func (m *RootCertResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RootCertResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RootCertResponse proto.InternalMessageInfo

func (m *RootCertResponse) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

type RegenerateRootRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *RegenerateRootRequest) Reset()         { *m = RegenerateRootRequest{} }
func (m *RegenerateRootRequest) String() string { return proto.CompactTextString(m) }
func (*RegenerateRootRequest) ProtoMessage()    {}
func (*RegenerateRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc6d195afcbac2c5, []int{2}
}

func (m *RegenerateRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegenerateRootRequest.Unmarshal(m, b)
}
func (m *RegenerateRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegenerateRootRequest.Marshal(b, m, deterministic)
}
func (m *RegenerateRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegenerateRootRequest.Merge(m, src)
}
func (m *RegenerateRootRequest) XXX_Size() int {
	return xxx_messageInfo_RegenerateRootRequest.Size(m)
}
func (m *RegenerateRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegenerateRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegenerateRootRequest proto.InternalMessageInfo

type RegenerateRootResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *RegenerateRootResponse) Reset()         { *m = RegenerateRootResponse{} }
func (m *RegenerateRootResponse) String() string { return proto.CompactTextString(m) }
func (*RegenerateRootResponse) ProtoMessage()    {}
func (*RegenerateRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc6d195afcbac2c5, []int{3}
}

func (m *RegenerateRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegenerateRootResponse.Unmarshal(m, b)
}
func (m *RegenerateRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegenerateRootResponse.Marshal(b, m, deterministic)
}
func (m *RegenerateRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegenerateRootResponse.Merge(m, src)
}
func (m *RegenerateRootResponse) XXX_Size() int {
	return xxx_messageInfo_RegenerateRootResponse.Size(m)
}
func (m *RegenerateRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegenerateRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegenerateRootResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RootCertRequest)(nil), "chef.automate.domain.deployment.RootCertRequest")
	proto.RegisterType((*RootCertResponse)(nil), "chef.automate.domain.deployment.RootCertResponse")
	proto.RegisterType((*RegenerateRootRequest)(nil), "chef.automate.domain.deployment.RegenerateRootRequest")
	proto.RegisterType((*RegenerateRootResponse)(nil), "chef.automate.domain.deployment.RegenerateRootResponse")
}

func init() {
	proto.RegisterFile("interservice/deployment/certificate_authority.proto", fileDescriptor_bc6d195afcbac2c5)
}

var fileDescriptor_bc6d195afcbac2c5 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x88, 0xe0, 0x08, 0x7e, 0x2c, 0x7e, 0x94, 0x5e, 0x94, 0x1c, 0xc4, 0xd3, 0xae,
	0x5a, 0xa9, 0x67, 0xed, 0xc1, 0x7b, 0x8e, 0x5e, 0x64, 0xbb, 0x7d, 0xdb, 0x2c, 0x98, 0x9d, 0xb8,
	0x99, 0x08, 0x3d, 0xfa, 0x07, 0xfc, 0xcd, 0x92, 0xd6, 0x10, 0x2c, 0x96, 0xd2, 0xdb, 0xb2, 0xc3,
	0xfb, 0xcc, 0x33, 0xbc, 0x34, 0xf0, 0x41, 0x10, 0x2b, 0xc4, 0x4f, 0xef, 0x60, 0x26, 0x28, 0xdf,
	0x79, 0x5e, 0x20, 0x88, 0x71, 0x88, 0xe2, 0xa7, 0xde, 0x59, 0xc1, 0x9b, 0xad, 0x25, 0xe7, 0xe8,
	0x65, 0xae, 0xcb, 0xc8, 0xc2, 0xea, 0xd2, 0xe5, 0x98, 0x6a, 0x5b, 0x0b, 0x17, 0x56, 0xa0, 0x27,
	0x5c, 0x58, 0x1f, 0x74, 0x17, 0x4e, 0x4f, 0xe8, 0x28, 0x63, 0x96, 0x11, 0xa2, 0x64, 0xf8, 0xa8,
	0x51, 0x49, 0x7a, 0x4d, 0xc7, 0xdd, 0x57, 0x55, 0x72, 0xa8, 0xa0, 0x14, 0xed, 0x36, 0x6b, 0x7a,
	0xc9, 0x55, 0x72, 0xb3, 0x9f, 0x2d, 0xde, 0xe9, 0x05, 0x9d, 0x65, 0x98, 0x21, 0x20, 0x5a, 0x41,
	0x93, 0x68, 0x01, 0x3d, 0x3a, 0x5f, 0x1d, 0x2c, 0x31, 0xf7, 0xdf, 0x3b, 0x74, 0x3a, 0xea, 0x74,
	0x9f, 0x5a, 0x5b, 0x15, 0xe9, 0xe0, 0x05, 0xd2, 0xae, 0x55, 0xb7, 0x7a, 0x83, 0xb7, 0x5e, 0x91,
	0xee, 0xdf, 0x6d, 0x91, 0xf8, 0xbd, 0xe9, 0x2b, 0xa1, 0xc3, 0xbf, 0x9e, 0x6a, 0xb8, 0x99, 0xf2,
	0xdf, 0xc5, 0xfd, 0xc7, 0xad, 0x73, 0x4b, 0x87, 0xe7, 0xe1, 0xeb, 0xc3, 0xcc, 0x4b, 0x5e, 0x8f,
	0xb5, 0xe3, 0xc2, 0x34, 0x10, 0xd3, 0x42, 0x8c, 0x2d, 0xbd, 0x59, 0xd3, 0xf9, 0x78, 0x6f, 0x51,
	0xef, 0xe0, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x02, 0x4e, 0xb6, 0x15, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CertificateAuthorityClient is the client API for CertificateAuthority service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateAuthorityClient interface {
	GetRootCert(ctx context.Context, in *RootCertRequest, opts ...grpc.CallOption) (*RootCertResponse, error)
	RegenerateRoot(ctx context.Context, in *RegenerateRootRequest, opts ...grpc.CallOption) (*RegenerateRootResponse, error)
}

type certificateAuthorityClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateAuthorityClient(cc grpc.ClientConnInterface) CertificateAuthorityClient {
	return &certificateAuthorityClient{cc}
}

func (c *certificateAuthorityClient) GetRootCert(ctx context.Context, in *RootCertRequest, opts ...grpc.CallOption) (*RootCertResponse, error) {
	out := new(RootCertResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.deployment.CertificateAuthority/GetRootCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityClient) RegenerateRoot(ctx context.Context, in *RegenerateRootRequest, opts ...grpc.CallOption) (*RegenerateRootResponse, error) {
	out := new(RegenerateRootResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.deployment.CertificateAuthority/RegenerateRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateAuthorityServer is the server API for CertificateAuthority service.
type CertificateAuthorityServer interface {
	GetRootCert(context.Context, *RootCertRequest) (*RootCertResponse, error)
	RegenerateRoot(context.Context, *RegenerateRootRequest) (*RegenerateRootResponse, error)
}

// UnimplementedCertificateAuthorityServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateAuthorityServer struct {
}

func (*UnimplementedCertificateAuthorityServer) GetRootCert(ctx context.Context, req *RootCertRequest) (*RootCertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRootCert not implemented")
}
func (*UnimplementedCertificateAuthorityServer) RegenerateRoot(ctx context.Context, req *RegenerateRootRequest) (*RegenerateRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegenerateRoot not implemented")
}

func RegisterCertificateAuthorityServer(s *grpc.Server, srv CertificateAuthorityServer) {
	s.RegisterService(&_CertificateAuthority_serviceDesc, srv)
}

func _CertificateAuthority_GetRootCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RootCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).GetRootCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.deployment.CertificateAuthority/GetRootCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).GetRootCert(ctx, req.(*RootCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthority_RegenerateRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegenerateRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).RegenerateRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.deployment.CertificateAuthority/RegenerateRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).RegenerateRoot(ctx, req.(*RegenerateRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAuthority_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.deployment.CertificateAuthority",
	HandlerType: (*CertificateAuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRootCert",
			Handler:    _CertificateAuthority_GetRootCert_Handler,
		},
		{
			MethodName: "RegenerateRoot",
			Handler:    _CertificateAuthority_RegenerateRoot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/deployment/certificate_authority.proto",
}
