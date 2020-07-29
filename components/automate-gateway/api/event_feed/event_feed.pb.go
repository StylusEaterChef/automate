// Code generated by protoc-gen-go. DO NOT EDIT.
// source: automate-gateway/api/event_feed/event_feed.proto

package event_feed

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/event_feed/request"
	response "github.com/chef/automate/components/automate-gateway/api/event_feed/response"
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
	proto.RegisterFile("automate-gateway/api/event_feed/event_feed.proto", fileDescriptor_6aea95401f8de03d)
}

var fileDescriptor_6aea95401f8de03d = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0xb1, 0xce, 0xd3, 0x30,
	0x10, 0x80, 0x55, 0x86, 0x5f, 0x22, 0xaa, 0x90, 0xea, 0x02, 0x2a, 0xa1, 0x53, 0xf7, 0xd8, 0xa5,
	0x05, 0x21, 0x95, 0xad, 0x55, 0x41, 0x62, 0x85, 0x89, 0xa5, 0x72, 0xdd, 0x6b, 0x6a, 0x35, 0xb1,
	0x4d, 0x7c, 0x29, 0xca, 0xca, 0xd8, 0x95, 0xd7, 0x60, 0xa4, 0xcf, 0xc0, 0xc4, 0xc4, 0x2b, 0xf0,
	0x20, 0x28, 0x71, 0xa2, 0x36, 0x42, 0x22, 0xf2, 0xf8, 0x6f, 0xd1, 0xe5, 0xbe, 0xbb, 0xfb, 0x4e,
	0xbe, 0x60, 0xca, 0x73, 0xd4, 0x29, 0x47, 0x88, 0x62, 0x8e, 0xf0, 0x85, 0x17, 0x8c, 0x1b, 0xc9,
	0xe0, 0x04, 0x0a, 0x37, 0x7b, 0x80, 0xdd, 0xcd, 0x27, 0x35, 0x99, 0x46, 0x4d, 0xc6, 0xe2, 0x00,
	0x7b, 0xda, 0x60, 0x94, 0x1b, 0x49, 0xaf, 0x39, 0xe1, 0xbc, 0xab, 0x5e, 0x06, 0x9f, 0x73, 0xb0,
	0xe8, 0x42, 0xae, 0x64, 0xb8, 0xf0, 0x82, 0x2c, 0x66, 0x52, 0xc5, 0xb6, 0x66, 0x5f, 0x76, 0xb3,
	0xd6, 0x68, 0x65, 0xa1, 0xd5, 0xf1, 0x8d, 0x1f, 0xd5, 0x6e, 0x39, 0x8e, 0xb5, 0x8e, 0x13, 0xa8,
	0x10, 0xae, 0x94, 0x46, 0x8e, 0x52, 0xab, 0xe6, 0xef, 0xab, 0x6b, 0xe9, 0xcc, 0x08, 0x56, 0x05,
	0x45, 0x14, 0x83, 0x8a, 0x8c, 0x4e, 0xa4, 0x28, 0x98, 0xe4, 0xe9, 0xbf, 0xd8, 0xec, 0xfb, 0x5d,
	0xf0, 0x70, 0x5d, 0xf6, 0x7a, 0x0b, 0xb0, 0x23, 0x3f, 0x7a, 0x41, 0xff, 0x1d, 0xe0, 0x35, 0xf0,
	0x82, 0xfe, 0x6f, 0xed, 0xb4, 0x5e, 0x10, 0x75, 0x80, 0x4c, 0x10, 0xb2, 0x30, 0xea, 0x42, 0x9c,
	0xa1, 0x63, 0xec, 0x64, 0x79, 0xbe, 0x8c, 0x1e, 0x05, 0xfd, 0x2a, 0x63, 0xe1, 0xac, 0xcf, 0x97,
	0xd1, 0x90, 0x0c, 0x6e, 0x23, 0x8b, 0x44, 0x5a, 0xfc, 0xfa, 0xfb, 0xcf, 0xb7, 0x07, 0x43, 0x32,
	0xa8, 0x16, 0x70, 0x9a, 0xba, 0x15, 0x95, 0x35, 0xc9, 0xaf, 0x5e, 0x40, 0x9a, 0xb1, 0x3f, 0x16,
	0x06, 0x56, 0x3a, 0x57, 0x68, 0xc9, 0x6b, 0x8f, 0xe1, 0x1d, 0x52, 0x2b, 0xcc, 0x7c, 0x14, 0x1c,
	0x39, 0x79, 0xef, 0xe7, 0xf1, 0x9c, 0x3c, 0x6b, 0x79, 0x6c, 0xb0, 0x30, 0xb0, 0x11, 0x6e, 0xf0,
	0x96, 0x0f, 0xb7, 0xc7, 0x7b, 0xe9, 0xc3, 0xed, 0xb1, 0xf1, 0xf9, 0xd9, 0x0b, 0x9e, 0x34, 0x3e,
	0x1f, 0xaa, 0x37, 0xbd, 0xcc, 0xc5, 0x11, 0xd0, 0x92, 0x99, 0x87, 0x92, 0x23, 0x6d, 0x38, 0xf7,
	0xb1, 0xa9, 0xa1, 0xc9, 0xda, 0x4f, 0xe7, 0x29, 0x79, 0xdc, 0xd2, 0xa9, 0x2f, 0x71, 0xb9, 0xfe,
	0xb4, 0x8a, 0x25, 0x1e, 0xf2, 0x2d, 0x15, 0x3a, 0x65, 0xe5, 0x1c, 0xac, 0x99, 0x83, 0x09, 0x9d,
	0x1a, 0xad, 0xca, 0x64, 0xd6, 0x71, 0xe6, 0xdb, 0xbb, 0xea, 0xf8, 0xe6, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x1a, 0x60, 0x50, 0xdc, 0x07, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventFeedClient is the client API for EventFeed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventFeedClient interface {
	//
	//List Events
	//
	//Returns a list of recorded events in Chef Automate, such as Infra Server events (cookbook creation, policyfile updates, and node creation) and Chef Automate internal events (profile installs and scan job creation).
	//Adding a filter makes a list of all events that meet the filter criteria.
	//
	//Example:
	//```
	//eventfeed?collapse=true&filter=organization:The%2520Watchmen&page_size=100&start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventFeed(ctx context.Context, in *request.EventFilter, opts ...grpc.CallOption) (*response.Events, error)
	//
	//List Count of Event Type Occurrence
	//
	//Returns totals for role, cookbook, and organization event types.
	//
	//Example:
	//```
	//event_type_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTypeCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error)
	//
	//List Counts Per Event Task Occurrence
	//
	//Returns totals for update, create, and delete event tasks, which are the actions taken on the event.
	//
	//Example:
	//```
	//event_task_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTaskCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error)
	//
	//List Summary Stats About Events
	//
	//Returns data that populates the guitar strings visual on the top of the event feed.
	//
	//Example:
	//```
	//eventstrings?timezone=America/Denver&hours_between=1&start=2020-06-19&end=2020-06-25
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventStringBuckets(ctx context.Context, in *request.EventStrings, opts ...grpc.CallOption) (*response.EventStrings, error)
}

type eventFeedClient struct {
	cc grpc.ClientConnInterface
}

func NewEventFeedClient(cc grpc.ClientConnInterface) EventFeedClient {
	return &eventFeedClient{cc}
}

func (c *eventFeedClient) GetEventFeed(ctx context.Context, in *request.EventFilter, opts ...grpc.CallOption) (*response.Events, error) {
	out := new(response.Events)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedClient) GetEventTypeCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error) {
	out := new(response.EventCounts)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventTypeCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedClient) GetEventTaskCounts(ctx context.Context, in *request.EventCountsFilter, opts ...grpc.CallOption) (*response.EventCounts, error) {
	out := new(response.EventCounts)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventTaskCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedClient) GetEventStringBuckets(ctx context.Context, in *request.EventStrings, opts ...grpc.CallOption) (*response.EventStrings, error) {
	out := new(response.EventStrings)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeed/GetEventStringBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventFeedServer is the server API for EventFeed service.
type EventFeedServer interface {
	//
	//List Events
	//
	//Returns a list of recorded events in Chef Automate, such as Infra Server events (cookbook creation, policyfile updates, and node creation) and Chef Automate internal events (profile installs and scan job creation).
	//Adding a filter makes a list of all events that meet the filter criteria.
	//
	//Example:
	//```
	//eventfeed?collapse=true&filter=organization:The%2520Watchmen&page_size=100&start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventFeed(context.Context, *request.EventFilter) (*response.Events, error)
	//
	//List Count of Event Type Occurrence
	//
	//Returns totals for role, cookbook, and organization event types.
	//
	//Example:
	//```
	//event_type_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTypeCounts(context.Context, *request.EventCountsFilter) (*response.EventCounts, error)
	//
	//List Counts Per Event Task Occurrence
	//
	//Returns totals for update, create, and delete event tasks, which are the actions taken on the event.
	//
	//Example:
	//```
	//event_task_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTaskCounts(context.Context, *request.EventCountsFilter) (*response.EventCounts, error)
	//
	//List Summary Stats About Events
	//
	//Returns data that populates the guitar strings visual on the top of the event feed.
	//
	//Example:
	//```
	//eventstrings?timezone=America/Denver&hours_between=1&start=2020-06-19&end=2020-06-25
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventStringBuckets(context.Context, *request.EventStrings) (*response.EventStrings, error)
}

// UnimplementedEventFeedServer can be embedded to have forward compatible implementations.
type UnimplementedEventFeedServer struct {
}

func (*UnimplementedEventFeedServer) GetEventFeed(ctx context.Context, req *request.EventFilter) (*response.Events, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventFeed not implemented")
}
func (*UnimplementedEventFeedServer) GetEventTypeCounts(ctx context.Context, req *request.EventCountsFilter) (*response.EventCounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventTypeCounts not implemented")
}
func (*UnimplementedEventFeedServer) GetEventTaskCounts(ctx context.Context, req *request.EventCountsFilter) (*response.EventCounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventTaskCounts not implemented")
}
func (*UnimplementedEventFeedServer) GetEventStringBuckets(ctx context.Context, req *request.EventStrings) (*response.EventStrings, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventStringBuckets not implemented")
}

func RegisterEventFeedServer(s *grpc.Server, srv EventFeedServer) {
	s.RegisterService(&_EventFeed_serviceDesc, srv)
}

func _EventFeed_GetEventFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventFeed(ctx, req.(*request.EventFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeed_GetEventTypeCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventCountsFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventTypeCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventTypeCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventTypeCounts(ctx, req.(*request.EventCountsFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeed_GetEventTaskCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventCountsFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventTaskCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventTaskCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventTaskCounts(ctx, req.(*request.EventCountsFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeed_GetEventStringBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.EventStrings)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServer).GetEventStringBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeed/GetEventStringBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServer).GetEventStringBuckets(ctx, req.(*request.EventStrings))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventFeed_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.event_feed.EventFeed",
	HandlerType: (*EventFeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEventFeed",
			Handler:    _EventFeed_GetEventFeed_Handler,
		},
		{
			MethodName: "GetEventTypeCounts",
			Handler:    _EventFeed_GetEventTypeCounts_Handler,
		},
		{
			MethodName: "GetEventTaskCounts",
			Handler:    _EventFeed_GetEventTaskCounts_Handler,
		},
		{
			MethodName: "GetEventStringBuckets",
			Handler:    _EventFeed_GetEventStringBuckets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "automate-gateway/api/event_feed/event_feed.proto",
}
