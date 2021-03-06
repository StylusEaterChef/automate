// Code generated by protoc-gen-go. DO NOT EDIT.
// source: automate-gateway/api/event_feed/response/event.proto

package response

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Events struct {
	// List of events.
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	// Total count of events.
	TotalEvents          int64    `protobuf:"varint,2,opt,name=total_events,json=totalEvents,proto3" json:"total_events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Events) Reset()         { *m = Events{} }
func (m *Events) String() string { return proto.CompactTextString(m) }
func (*Events) ProtoMessage()    {}
func (*Events) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fdacb4c611862df, []int{0}
}

func (m *Events) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Events.Unmarshal(m, b)
}
func (m *Events) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Events.Marshal(b, m, deterministic)
}
func (m *Events) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Events.Merge(m, src)
}
func (m *Events) XXX_Size() int {
	return xxx_messageInfo_Events.Size(m)
}
func (m *Events) XXX_DiscardUnknown() {
	xxx_messageInfo_Events.DiscardUnknown(m)
}

var xxx_messageInfo_Events proto.InternalMessageInfo

func (m *Events) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Events) GetTotalEvents() int64 {
	if m != nil {
		return m.TotalEvents
	}
	return 0
}

type Event struct {
	// Type of event (cookbook, role, etc).
	EventType string `protobuf:"bytes,1,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	// Type of event task (create, update, delete).
	Task string `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	// Event start time.
	StartTime  *timestamp.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EntityName string               `protobuf:"bytes,4,opt,name=entity_name,json=entityName,proto3" json:"entity_name,omitempty"`
	// Event record requestor type.
	RequestorType string `protobuf:"bytes,5,opt,name=requestor_type,json=requestorType,proto3" json:"requestor_type,omitempty"`
	// Event record requestor name.
	RequestorName string `protobuf:"bytes,6,opt,name=requestor_name,json=requestorName,proto3" json:"requestor_name,omitempty"`
	// Hostname from which the record was gathered.
	ServiceHostname string `protobuf:"bytes,7,opt,name=service_hostname,json=serviceHostname,proto3" json:"service_hostname,omitempty"`
	// Used for grouping events together.
	StartId string `protobuf:"bytes,8,opt,name=start_id,json=startId,proto3" json:"start_id,omitempty"`
	// Used for grouping events together.
	EventCount int32 `protobuf:"varint,9,opt,name=event_count,json=eventCount,proto3" json:"event_count,omitempty"`
	// Used for grouping events together.
	ParentName string `protobuf:"bytes,16,opt,name=parent_name,json=parentName,proto3" json:"parent_name,omitempty"`
	// Used for grouping events together.
	ParentType string `protobuf:"bytes,17,opt,name=parent_type,json=parentType,proto3" json:"parent_type,omitempty"`
	// Used for grouping events together; equal to start_time if not grouped
	EndTime *timestamp.Timestamp `protobuf:"bytes,18,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Used for grouping events together; equal to start_id if not grouped
	EndId                string   `protobuf:"bytes,19,opt,name=end_id,json=endId,proto3" json:"end_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fdacb4c611862df, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *Event) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *Event) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Event) GetEntityName() string {
	if m != nil {
		return m.EntityName
	}
	return ""
}

func (m *Event) GetRequestorType() string {
	if m != nil {
		return m.RequestorType
	}
	return ""
}

func (m *Event) GetRequestorName() string {
	if m != nil {
		return m.RequestorName
	}
	return ""
}

func (m *Event) GetServiceHostname() string {
	if m != nil {
		return m.ServiceHostname
	}
	return ""
}

func (m *Event) GetStartId() string {
	if m != nil {
		return m.StartId
	}
	return ""
}

func (m *Event) GetEventCount() int32 {
	if m != nil {
		return m.EventCount
	}
	return 0
}

func (m *Event) GetParentName() string {
	if m != nil {
		return m.ParentName
	}
	return ""
}

func (m *Event) GetParentType() string {
	if m != nil {
		return m.ParentType
	}
	return ""
}

func (m *Event) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Event) GetEndId() string {
	if m != nil {
		return m.EndId
	}
	return ""
}

type EventCounts struct {
	// Total count of events.
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	// Total count of events per type.
	Counts               []*EventCount `protobuf:"bytes,2,rep,name=counts,proto3" json:"counts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EventCounts) Reset()         { *m = EventCounts{} }
func (m *EventCounts) String() string { return proto.CompactTextString(m) }
func (*EventCounts) ProtoMessage()    {}
func (*EventCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fdacb4c611862df, []int{2}
}

func (m *EventCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventCounts.Unmarshal(m, b)
}
func (m *EventCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventCounts.Marshal(b, m, deterministic)
}
func (m *EventCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCounts.Merge(m, src)
}
func (m *EventCounts) XXX_Size() int {
	return xxx_messageInfo_EventCounts.Size(m)
}
func (m *EventCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCounts.DiscardUnknown(m)
}

var xxx_messageInfo_EventCounts proto.InternalMessageInfo

func (m *EventCounts) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *EventCounts) GetCounts() []*EventCount {
	if m != nil {
		return m.Counts
	}
	return nil
}

type EventCount struct {
	// Event name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Count of events.
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventCount) Reset()         { *m = EventCount{} }
func (m *EventCount) String() string { return proto.CompactTextString(m) }
func (*EventCount) ProtoMessage()    {}
func (*EventCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fdacb4c611862df, []int{3}
}

func (m *EventCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventCount.Unmarshal(m, b)
}
func (m *EventCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventCount.Marshal(b, m, deterministic)
}
func (m *EventCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCount.Merge(m, src)
}
func (m *EventCount) XXX_Size() int {
	return xxx_messageInfo_EventCount.Size(m)
}
func (m *EventCount) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCount.DiscardUnknown(m)
}

var xxx_messageInfo_EventCount proto.InternalMessageInfo

func (m *EventCount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventCount) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Events)(nil), "chef.automate.api.event_feed.response.Events")
	proto.RegisterType((*Event)(nil), "chef.automate.api.event_feed.response.Event")
	proto.RegisterType((*EventCounts)(nil), "chef.automate.api.event_feed.response.EventCounts")
	proto.RegisterType((*EventCount)(nil), "chef.automate.api.event_feed.response.EventCount")
}

func init() {
	proto.RegisterFile("automate-gateway/api/event_feed/response/event.proto", fileDescriptor_5fdacb4c611862df)
}

var fileDescriptor_5fdacb4c611862df = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x25, 0xa6, 0xc9, 0x6e, 0x6e, 0xfc, 0xa8, 0xa3, 0x42, 0x2c, 0x48, 0xe3, 0x42, 0x21, 0x82,
	0x4e, 0xb0, 0x7e, 0x80, 0xaf, 0x6a, 0xc1, 0x05, 0xe9, 0x43, 0xe8, 0x93, 0x2f, 0x61, 0x36, 0xb9,
	0xbb, 0x1b, 0x6c, 0x66, 0xd2, 0xcc, 0xa4, 0xb2, 0xbf, 0xd5, 0x3f, 0x23, 0x73, 0x27, 0xe9, 0xaa,
	0x2f, 0xd6, 0xb7, 0xb9, 0x87, 0x73, 0xcf, 0xbd, 0xe7, 0xe4, 0x06, 0xde, 0x8a, 0xc1, 0xa8, 0x56,
	0x18, 0x7c, 0xb5, 0x11, 0x06, 0x7f, 0x88, 0x5d, 0x2e, 0xba, 0x26, 0xc7, 0x6b, 0x94, 0xa6, 0x5c,
	0x23, 0xd6, 0x79, 0x8f, 0xba, 0x53, 0x52, 0xa3, 0xc3, 0x78, 0xd7, 0x2b, 0xa3, 0xd8, 0x49, 0xb5,
	0xc5, 0x35, 0x9f, 0x5a, 0xb9, 0xe8, 0x1a, 0xbe, 0x6f, 0xe1, 0x53, 0xcb, 0xd1, 0xf1, 0x46, 0xa9,
	0xcd, 0x25, 0xe6, 0xd4, 0xb4, 0x1a, 0xd6, 0xb9, 0x69, 0x5a, 0xd4, 0x46, 0xb4, 0x9d, 0xd3, 0x59,
	0x5c, 0x41, 0x78, 0x66, 0xfb, 0x34, 0xfb, 0x0c, 0x21, 0x29, 0xe8, 0xc4, 0x4b, 0xfd, 0x2c, 0x3e,
	0x7d, 0xc9, 0x6f, 0x35, 0x82, 0x53, 0x7b, 0x31, 0xf6, 0xb2, 0xe7, 0x70, 0xd7, 0x28, 0x23, 0x2e,
	0xcb, 0x51, 0xeb, 0x4e, 0xea, 0x65, 0x7e, 0x11, 0x13, 0xe6, 0x06, 0x2d, 0x7e, 0xfa, 0x10, 0xd0,
	0x93, 0x3d, 0x03, 0x70, 0x8a, 0x66, 0xd7, 0x61, 0xe2, 0xa5, 0x5e, 0x16, 0x15, 0x11, 0x21, 0x17,
	0xbb, 0x0e, 0x19, 0x83, 0x03, 0x23, 0xf4, 0x77, 0xd2, 0x88, 0x0a, 0x7a, 0xb3, 0x0f, 0x00, 0xda,
	0x88, 0xde, 0x94, 0xd6, 0x48, 0xe2, 0xa7, 0x5e, 0x16, 0x9f, 0x1e, 0x71, 0xe7, 0x92, 0x4f, 0x2e,
	0xf9, 0xc5, 0xe4, 0xb2, 0x88, 0x88, 0x6d, 0x6b, 0x76, 0x0c, 0x31, 0x4a, 0xd3, 0x98, 0x5d, 0x29,
	0x45, 0x8b, 0xc9, 0x01, 0xa9, 0x82, 0x83, 0xce, 0x45, 0x8b, 0xec, 0x04, 0xee, 0xf7, 0x78, 0x35,
	0xa0, 0x36, 0xaa, 0x77, 0x2b, 0x05, 0xc4, 0xb9, 0x77, 0x83, 0xd2, 0x5a, 0x7f, 0xd0, 0x48, 0x2a,
	0xfc, 0x8b, 0x46, 0x6a, 0x2f, 0xe0, 0x50, 0x63, 0x7f, 0xdd, 0x54, 0x58, 0x6e, 0x95, 0x36, 0x44,
	0x9c, 0x11, 0xf1, 0xc1, 0x88, 0x7f, 0x19, 0x61, 0xf6, 0x14, 0xe6, 0xce, 0x54, 0x53, 0x27, 0x73,
	0xa2, 0xcc, 0xa8, 0x5e, 0xd6, 0xb4, 0x34, 0x45, 0x54, 0xa9, 0x41, 0x9a, 0x24, 0x4a, 0xbd, 0x2c,
	0x28, 0x5c, 0x6a, 0x9f, 0x2c, 0x62, 0x09, 0x9d, 0xe8, 0x2d, 0x83, 0x26, 0x1c, 0x3a, 0x57, 0x0e,
	0xa2, 0x3d, 0xf6, 0x04, 0xb2, 0xf4, 0xf0, 0x77, 0x02, 0xf9, 0x79, 0x07, 0x73, 0x94, 0xb5, 0x0b,
	0x94, 0xfd, 0x33, 0xd0, 0x19, 0xca, 0x9a, 0xe2, 0x7c, 0x02, 0xa1, 0x6d, 0x6b, 0xea, 0xe4, 0x11,
	0x49, 0x06, 0x28, 0xeb, 0x65, 0xbd, 0x90, 0x10, 0x9f, 0xdd, 0x6c, 0xa7, 0xd9, 0x63, 0x08, 0xe8,
	0xdb, 0xd3, 0xd7, 0xf5, 0x0b, 0x57, 0xb0, 0x25, 0x84, 0xe4, 0xc7, 0xde, 0x87, 0xbd, 0xb5, 0xd7,
	0xff, 0x73, 0x6b, 0xa4, 0x5c, 0x8c, 0x02, 0x8b, 0xf7, 0x00, 0x7b, 0xd4, 0x9e, 0x0c, 0xc5, 0xe0,
	0x6e, 0x89, 0xde, 0x76, 0x05, 0x17, 0x9e, 0xbb, 0x45, 0x57, 0x7c, 0x3c, 0xff, 0xf6, 0x75, 0xd3,
	0x98, 0xed, 0xb0, 0xe2, 0x95, 0x6a, 0x73, 0x3b, 0x3e, 0x9f, 0xc6, 0xe7, 0x95, 0x6a, 0x3b, 0x25,
	0xed, 0xb1, 0xe6, 0xb7, 0xfd, 0x39, 0x57, 0x21, 0x65, 0xf5, 0xe6, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa4, 0xa8, 0xf7, 0xaf, 0xcf, 0x03, 0x00, 0x00,
}
