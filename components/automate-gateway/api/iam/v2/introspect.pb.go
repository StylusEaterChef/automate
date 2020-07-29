// Code generated by protoc-gen-go. DO NOT EDIT.
// source: automate-gateway/api/iam/v2/introspect.proto

package v2

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() {
	proto.RegisterFile("automate-gateway/api/iam/v2/introspect.proto", fileDescriptor_8963f93b47f59f42)
}

var fileDescriptor_8963f93b47f59f42 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0x03, 0x31,
	0x18, 0x86, 0xa9, 0x8a, 0xc3, 0x41, 0x8b, 0x04, 0xb5, 0xe7, 0x51, 0x1c, 0xba, 0x28, 0xd5, 0x26,
	0x50, 0x75, 0xe9, 0x22, 0x75, 0x73, 0x71, 0xa8, 0x4e, 0x82, 0x48, 0x1a, 0x3e, 0xaf, 0x81, 0x4b,
	0xbe, 0xf4, 0x92, 0xab, 0xd4, 0xd1, 0xb1, 0xab, 0xbf, 0xa5, 0x7f, 0x41, 0xdc, 0xfd, 0x0b, 0xfe,
	0x10, 0xb9, 0x2b, 0x6d, 0x29, 0xc7, 0x1d, 0xed, 0x9a, 0x7c, 0xef, 0xf7, 0x3e, 0xef, 0x9b, 0x78,
	0x97, 0x3c, 0x71, 0xa8, 0xb8, 0x83, 0x76, 0xc8, 0x1d, 0xbc, 0xf3, 0x09, 0xe3, 0x46, 0x32, 0xc9,
	0x15, 0x1b, 0x77, 0x98, 0xd4, 0x2e, 0x46, 0x6b, 0x40, 0x38, 0x6a, 0x62, 0x74, 0x48, 0x7c, 0x31,
	0x84, 0x37, 0xba, 0x90, 0x50, 0x6e, 0x24, 0x95, 0x5c, 0xd1, 0x71, 0x27, 0x68, 0x84, 0x88, 0x61,
	0x04, 0x99, 0x9a, 0x6b, 0x8d, 0x8e, 0x3b, 0x89, 0xda, 0xce, 0x75, 0xc1, 0x75, 0x99, 0x4b, 0x0c,
	0xa3, 0x04, 0xac, 0xcb, 0xb9, 0x05, 0x37, 0xe5, 0x2a, 0x6b, 0x50, 0x5b, 0x28, 0x95, 0xc5, 0x46,
	0xb0, 0xec, 0x50, 0xb4, 0x43, 0xd0, 0x6d, 0x83, 0x91, 0x14, 0x93, 0x6c, 0x43, 0x8e, 0xb1, 0xf3,
	0xb3, 0xe7, 0x55, 0x7b, 0x89, 0x1b, 0x62, 0x2c, 0x3f, 0xb2, 0x0b, 0x32, 0xab, 0x78, 0xd5, 0xfb,
	0xe5, 0xf6, 0x5e, 0x14, 0x91, 0x16, 0x2d, 0x2a, 0x80, 0xae, 0x0d, 0xf6, 0x61, 0x14, 0x9c, 0x6f,
	0x32, 0xdb, 0x07, 0x6b, 0x9a, 0x0f, 0xd3, 0x99, 0x7f, 0xe0, 0xd5, 0x24, 0x57, 0xdd, 0x55, 0x98,
	0xe9, 0xcc, 0xaf, 0x93, 0xa3, 0xf5, 0xb3, 0x6e, 0x08, 0xe9, 0xf6, 0xcf, 0xdf, 0xbf, 0xaf, 0x9d,
	0x13, 0x52, 0x4f, 0xeb, 0xb0, 0xf9, 0xb7, 0x22, 0xdf, 0x15, 0xaf, 0xb6, 0xb2, 0x78, 0x44, 0x05,
	0xe4, 0x62, 0x13, 0x98, 0x74, 0x72, 0x3b, 0xf2, 0x97, 0x02, 0x72, 0x9f, 0x1c, 0xe7, 0xc9, 0xd3,
	0xf5, 0x19, 0xfa, 0x69, 0xb3, 0x51, 0x80, 0xfe, 0x6a, 0x17, 0x33, 0xbb, 0xdd, 0x4a, 0x2b, 0x7d,
	0x00, 0x6f, 0xe5, 0x48, 0xce, 0x36, 0xe3, 0xda, 0x26, 0xc0, 0x53, 0x41, 0x80, 0x43, 0x42, 0xf2,
	0x01, 0xe6, 0xbd, 0x37, 0x8b, 0x7a, 0x5f, 0x72, 0xdf, 0xf5, 0x9e, 0x6f, 0x43, 0xe9, 0x86, 0xc9,
	0x80, 0x0a, 0x54, 0x2c, 0x65, 0x61, 0x0b, 0x16, 0x26, 0x50, 0x19, 0xd4, 0xa0, 0x9d, 0x65, 0x25,
	0xdf, 0x7b, 0xb0, 0x9f, 0x7d, 0xca, 0xab, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0xff, 0x6e,
	0x43, 0xa0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	IntrospectAll(ctx context.Context, in *request.IntrospectAllReq, opts ...grpc.CallOption) (*response.IntrospectResp, error)
	IntrospectSome(ctx context.Context, in *request.IntrospectSomeReq, opts ...grpc.CallOption) (*response.IntrospectResp, error)
	Introspect(ctx context.Context, in *request.IntrospectReq, opts ...grpc.CallOption) (*response.IntrospectResp, error)
}

type authorizationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationClient(cc grpc.ClientConnInterface) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) IntrospectAll(ctx context.Context, in *request.IntrospectAllReq, opts ...grpc.CallOption) (*response.IntrospectResp, error) {
	out := new(response.IntrospectResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Authorization/IntrospectAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) IntrospectSome(ctx context.Context, in *request.IntrospectSomeReq, opts ...grpc.CallOption) (*response.IntrospectResp, error) {
	out := new(response.IntrospectResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Authorization/IntrospectSome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) Introspect(ctx context.Context, in *request.IntrospectReq, opts ...grpc.CallOption) (*response.IntrospectResp, error) {
	out := new(response.IntrospectResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Authorization/Introspect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	IntrospectAll(context.Context, *request.IntrospectAllReq) (*response.IntrospectResp, error)
	IntrospectSome(context.Context, *request.IntrospectSomeReq) (*response.IntrospectResp, error)
	Introspect(context.Context, *request.IntrospectReq) (*response.IntrospectResp, error)
}

// UnimplementedAuthorizationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (*UnimplementedAuthorizationServer) IntrospectAll(ctx context.Context, req *request.IntrospectAllReq) (*response.IntrospectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntrospectAll not implemented")
}
func (*UnimplementedAuthorizationServer) IntrospectSome(ctx context.Context, req *request.IntrospectSomeReq) (*response.IntrospectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IntrospectSome not implemented")
}
func (*UnimplementedAuthorizationServer) Introspect(ctx context.Context, req *request.IntrospectReq) (*response.IntrospectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Introspect not implemented")
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_IntrospectAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.IntrospectAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).IntrospectAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Authorization/IntrospectAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).IntrospectAll(ctx, req.(*request.IntrospectAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_IntrospectSome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.IntrospectSomeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).IntrospectSome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Authorization/IntrospectSome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).IntrospectSome(ctx, req.(*request.IntrospectSomeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_Introspect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.IntrospectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Introspect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Authorization/Introspect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Introspect(ctx, req.(*request.IntrospectReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IntrospectAll",
			Handler:    _Authorization_IntrospectAll_Handler,
		},
		{
			MethodName: "IntrospectSome",
			Handler:    _Authorization_IntrospectSome_Handler,
		},
		{
			MethodName: "Introspect",
			Handler:    _Authorization_Introspect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "automate-gateway/api/iam/v2/introspect.proto",
}
