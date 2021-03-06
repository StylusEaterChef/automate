// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interservice/ingest/chef.proto

package ingest

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/api/external/ingest/request"
	response "github.com/chef/automate/api/external/ingest/response"
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

// Version message
//
// The ingest-service version constructed with:
// * Service name
// * Built time
// * Semantic version
// * Git SHA
type Version struct {
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty" toml:"version,omitempty" mapstructure:"version,omitempty"`
	Built                string   `protobuf:"bytes,1,opt,name=built,proto3" json:"built,omitempty" toml:"built,omitempty" mapstructure:"built,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Sha                  string   `protobuf:"bytes,4,opt,name=sha,proto3" json:"sha,omitempty" toml:"sha,omitempty" mapstructure:"sha,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45313ceff872c3e, []int{0}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Version) GetBuilt() string {
	if m != nil {
		return m.Built
	}
	return ""
}

func (m *Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e45313ceff872c3e, []int{1}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Version)(nil), "chef.automate.domain.ingest.Version")
	proto.RegisterType((*VersionRequest)(nil), "chef.automate.domain.ingest.VersionRequest")
}

func init() {
	proto.RegisterFile("interservice/ingest/chef.proto", fileDescriptor_e45313ceff872c3e)
}

var fileDescriptor_e45313ceff872c3e = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0xb6, 0xb0, 0xc2, 0x42, 0x55, 0x31, 0x08, 0x22, 0xb3, 0x5a, 0xc0, 0xed, 0xd2,
	0x52, 0x24, 0xa7, 0x2a, 0x27, 0x0a, 0x17, 0xa0, 0x12, 0x42, 0x02, 0x54, 0xe5, 0xc0, 0x81, 0x0b,
	0xf2, 0x66, 0xa7, 0x59, 0x4b, 0x59, 0x3b, 0xc4, 0x4e, 0xc4, 0x0d, 0x89, 0x57, 0x40, 0xe2, 0x8a,
	0xc4, 0x81, 0xf7, 0xe0, 0x19, 0x78, 0x05, 0x1e, 0x04, 0xc5, 0x76, 0xda, 0x46, 0x9b, 0xed, 0xee,
	0xde, 0x26, 0x93, 0x7f, 0x66, 0x3e, 0xcf, 0xd8, 0x83, 0x06, 0x42, 0x1a, 0x28, 0x34, 0x14, 0x95,
	0x48, 0x20, 0x12, 0x32, 0x05, 0x6d, 0xa2, 0x64, 0x02, 0xa7, 0x2c, 0x2f, 0x94, 0x51, 0xf8, 0xae,
	0xb5, 0x79, 0x69, 0xd4, 0x94, 0x1b, 0x60, 0x63, 0x35, 0xe5, 0x42, 0x32, 0xa7, 0x23, 0xfd, 0x54,
	0xa9, 0x34, 0x83, 0x88, 0xe7, 0x22, 0xe2, 0x52, 0x2a, 0xc3, 0x8d, 0x50, 0x52, 0xbb, 0x50, 0x42,
	0xe1, 0x8b, 0x81, 0x42, 0xf2, 0xac, 0x49, 0x5b, 0xc0, 0xe7, 0xb2, 0x9d, 0x9e, 0xec, 0xcc, 0xd3,
	0xf0, 0xa4, 0x4e, 0xe5, 0x55, 0x0f, 0xe7, 0xa9, 0x32, 0x51, 0x81, 0x04, 0xdd, 0x54, 0xdc, 0x9e,
	0xd5, 0xe9, 0x5c, 0x49, 0x0d, 0x17, 0x4b, 0x0e, 0xe7, 0x8a, 0x5a, 0x35, 0x77, 0xe7, 0xca, 0xda,
	0x45, 0xe9, 0x27, 0xd4, 0xfb, 0x00, 0x85, 0x16, 0x4a, 0xe2, 0x10, 0xf5, 0x2a, 0x67, 0x86, 0x6b,
	0xf7, 0x83, 0xbd, 0x6b, 0x71, 0xf3, 0x89, 0x6f, 0xa1, 0x2b, 0xa3, 0x52, 0x64, 0x26, 0x0c, 0xac,
	0xdf, 0x7d, 0x60, 0x8c, 0x36, 0x24, 0x9f, 0x42, 0xb8, 0x6e, 0x9d, 0xd6, 0xc6, 0x5b, 0x68, 0x5d,
	0x4f, 0x78, 0xb8, 0x61, 0x5d, 0xb5, 0x49, 0xb7, 0xd0, 0xa6, 0x2f, 0x10, 0xbb, 0x63, 0x1f, 0xfe,
	0xe9, 0xa1, 0xeb, 0xaf, 0x26, 0x70, 0xfa, 0xc6, 0x92, 0x41, 0x81, 0x7f, 0x04, 0x68, 0xf3, 0xa4,
	0x50, 0x09, 0x68, 0x5d, 0xfb, 0xe3, 0x52, 0xe2, 0x21, 0x6b, 0x4f, 0x8e, 0xe7, 0xc2, 0x8f, 0x8d,
	0xf9, 0xee, 0xb1, 0xb8, 0x94, 0xe4, 0xe9, 0x65, 0x32, 0x77, 0x60, 0xd6, 0xce, 0x1c, 0x7b, 0x37,
	0xa5, 0xdf, 0xfe, 0xfe, 0xfb, 0xbe, 0xd6, 0xa7, 0x77, 0xec, 0xfc, 0xab, 0x83, 0x08, 0x2a, 0x90,
	0x46, 0xdb, 0x5e, 0x47, 0x45, 0x29, 0x8f, 0x82, 0x7d, 0xfc, 0x2b, 0x40, 0x37, 0x2e, 0x84, 0xbf,
	0xb0, 0x1d, 0xc6, 0x7b, 0x8b, 0xd9, 0x9c, 0x92, 0x3c, 0x5f, 0x0d, 0xcf, 0x45, 0x9d, 0x11, 0x0e,
	0x2d, 0xe1, 0x3d, 0x4a, 0xba, 0x08, 0xdd, 0xb4, 0x6b, 0xc8, 0xdf, 0x01, 0xba, 0xe9, 0x93, 0xbc,
	0xf5, 0xb3, 0x3d, 0x11, 0x32, 0xc5, 0xfb, 0x8b, 0x31, 0x1b, 0x3d, 0x39, 0x5a, 0x1e, 0xb4, 0x89,
	0x39, 0xc3, 0xdc, 0xb5, 0x98, 0x0f, 0x68, 0xbf, 0x0b, 0xb3, 0xb9, 0x6d, 0x35, 0xe8, 0xcf, 0x00,
	0x11, 0x9f, 0xe4, 0x5d, 0x99, 0x19, 0x91, 0x67, 0xf0, 0x5e, 0x8d, 0xe1, 0x18, 0x32, 0x30, 0xa0,
	0xf1, 0xb3, 0xc5, 0xbc, 0xb3, 0x61, 0xfe, 0x4e, 0x91, 0xe3, 0xe5, 0x0f, 0xd0, 0x95, 0xc4, 0x29,
	0xea, 0x56, 0x36, 0xf3, 0x3e, 0xff, 0xbb, 0xcc, 0xbc, 0x9d, 0x72, 0x95, 0x79, 0xcf, 0x56, 0xa7,
	0x8f, 0x6c, 0x23, 0xb7, 0xe9, 0xa0, 0xab, 0x91, 0x52, 0x8d, 0x61, 0x6c, 0xf5, 0x75, 0x2b, 0xbf,
	0x22, 0xf4, 0x1a, 0x4c, 0xf3, 0x70, 0x1f, 0xb3, 0x4b, 0xd6, 0x1c, 0x6b, 0xbf, 0x3e, 0xb2, 0xb3,
	0x8c, 0x98, 0x0e, 0x2c, 0x4b, 0x88, 0x6f, 0x37, 0x2c, 0x7e, 0x8d, 0xf8, 0x8d, 0xf0, 0xf2, 0xf0,
	0xe3, 0x41, 0x2a, 0xcc, 0xa4, 0x1c, 0xb1, 0x44, 0x4d, 0xfd, 0x8d, 0xf4, 0x19, 0x6d, 0x44, 0xc7,
	0x5e, 0x1e, 0x5d, 0xb5, 0x1b, 0xe7, 0xc9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x9f, 0xca,
	0x21, 0xb5, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChefIngesterClient is the client API for ChefIngester service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChefIngesterClient interface {
	ProcessChefRun(ctx context.Context, in *request.Run, opts ...grpc.CallOption) (*response.ProcessChefRunResponse, error)
	ProcessChefAction(ctx context.Context, in *request.Action, opts ...grpc.CallOption) (*response.ProcessChefActionResponse, error)
	ProcessLivenessPing(ctx context.Context, in *request.Liveness, opts ...grpc.CallOption) (*response.ProcessLivenessResponse, error)
	ProcessMultipleNodeDeletes(ctx context.Context, in *request.MultipleNodeDeleteRequest, opts ...grpc.CallOption) (*response.ProcessMultipleNodeDeleteResponse, error)
	ProcessNodeDelete(ctx context.Context, in *request.Delete, opts ...grpc.CallOption) (*response.ProcessNodeDeleteResponse, error)
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*Version, error)
}

type chefIngesterClient struct {
	cc grpc.ClientConnInterface
}

func NewChefIngesterClient(cc grpc.ClientConnInterface) ChefIngesterClient {
	return &chefIngesterClient{cc}
}

func (c *chefIngesterClient) ProcessChefRun(ctx context.Context, in *request.Run, opts ...grpc.CallOption) (*response.ProcessChefRunResponse, error) {
	out := new(response.ProcessChefRunResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/ProcessChefRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessChefAction(ctx context.Context, in *request.Action, opts ...grpc.CallOption) (*response.ProcessChefActionResponse, error) {
	out := new(response.ProcessChefActionResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/ProcessChefAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessLivenessPing(ctx context.Context, in *request.Liveness, opts ...grpc.CallOption) (*response.ProcessLivenessResponse, error) {
	out := new(response.ProcessLivenessResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/ProcessLivenessPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessMultipleNodeDeletes(ctx context.Context, in *request.MultipleNodeDeleteRequest, opts ...grpc.CallOption) (*response.ProcessMultipleNodeDeleteResponse, error) {
	out := new(response.ProcessMultipleNodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/ProcessMultipleNodeDeletes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) ProcessNodeDelete(ctx context.Context, in *request.Delete, opts ...grpc.CallOption) (*response.ProcessNodeDeleteResponse, error) {
	out := new(response.ProcessNodeDeleteResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/ProcessNodeDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chefIngesterClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.ingest.ChefIngester/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChefIngesterServer is the server API for ChefIngester service.
type ChefIngesterServer interface {
	ProcessChefRun(context.Context, *request.Run) (*response.ProcessChefRunResponse, error)
	ProcessChefAction(context.Context, *request.Action) (*response.ProcessChefActionResponse, error)
	ProcessLivenessPing(context.Context, *request.Liveness) (*response.ProcessLivenessResponse, error)
	ProcessMultipleNodeDeletes(context.Context, *request.MultipleNodeDeleteRequest) (*response.ProcessMultipleNodeDeleteResponse, error)
	ProcessNodeDelete(context.Context, *request.Delete) (*response.ProcessNodeDeleteResponse, error)
	GetVersion(context.Context, *VersionRequest) (*Version, error)
}

// UnimplementedChefIngesterServer can be embedded to have forward compatible implementations.
type UnimplementedChefIngesterServer struct {
}

func (*UnimplementedChefIngesterServer) ProcessChefRun(ctx context.Context, req *request.Run) (*response.ProcessChefRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessChefRun not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessChefAction(ctx context.Context, req *request.Action) (*response.ProcessChefActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessChefAction not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessLivenessPing(ctx context.Context, req *request.Liveness) (*response.ProcessLivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessLivenessPing not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessMultipleNodeDeletes(ctx context.Context, req *request.MultipleNodeDeleteRequest) (*response.ProcessMultipleNodeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessMultipleNodeDeletes not implemented")
}
func (*UnimplementedChefIngesterServer) ProcessNodeDelete(ctx context.Context, req *request.Delete) (*response.ProcessNodeDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessNodeDelete not implemented")
}
func (*UnimplementedChefIngesterServer) GetVersion(ctx context.Context, req *VersionRequest) (*Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

func RegisterChefIngesterServer(s *grpc.Server, srv ChefIngesterServer) {
	s.RegisterService(&_ChefIngester_serviceDesc, srv)
}

func _ChefIngester_ProcessChefRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Run)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessChefRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/ProcessChefRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessChefRun(ctx, req.(*request.Run))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessChefAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Action)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessChefAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/ProcessChefAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessChefAction(ctx, req.(*request.Action))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessLivenessPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Liveness)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessLivenessPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/ProcessLivenessPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessLivenessPing(ctx, req.(*request.Liveness))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessMultipleNodeDeletes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.MultipleNodeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessMultipleNodeDeletes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/ProcessMultipleNodeDeletes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessMultipleNodeDeletes(ctx, req.(*request.MultipleNodeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_ProcessNodeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.Delete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).ProcessNodeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/ProcessNodeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).ProcessNodeDelete(ctx, req.(*request.Delete))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChefIngester_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChefIngesterServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.ingest.ChefIngester/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChefIngesterServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChefIngester_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.ingest.ChefIngester",
	HandlerType: (*ChefIngesterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessChefRun",
			Handler:    _ChefIngester_ProcessChefRun_Handler,
		},
		{
			MethodName: "ProcessChefAction",
			Handler:    _ChefIngester_ProcessChefAction_Handler,
		},
		{
			MethodName: "ProcessLivenessPing",
			Handler:    _ChefIngester_ProcessLivenessPing_Handler,
		},
		{
			MethodName: "ProcessMultipleNodeDeletes",
			Handler:    _ChefIngester_ProcessMultipleNodeDeletes_Handler,
		},
		{
			MethodName: "ProcessNodeDelete",
			Handler:    _ChefIngester_ProcessNodeDelete_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _ChefIngester_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/ingest/chef.proto",
}
