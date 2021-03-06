// Code generated by protoc-gen-go. DO NOT EDIT.
// source: external/infra_proxy/response/roles.proto

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

type Roles struct {
	// List of the roles item.
	Roles                []*RoleListItem `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Roles) Reset()         { *m = Roles{} }
func (m *Roles) String() string { return proto.CompactTextString(m) }
func (*Roles) ProtoMessage()    {}
func (*Roles) Descriptor() ([]byte, []int) {
	return fileDescriptor_6155b399355910ab, []int{0}
}

func (m *Roles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Roles.Unmarshal(m, b)
}
func (m *Roles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Roles.Marshal(b, m, deterministic)
}
func (m *Roles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Roles.Merge(m, src)
}
func (m *Roles) XXX_Size() int {
	return xxx_messageInfo_Roles.Size(m)
}
func (m *Roles) XXX_DiscardUnknown() {
	xxx_messageInfo_Roles.DiscardUnknown(m)
}

var xxx_messageInfo_Roles proto.InternalMessageInfo

func (m *Roles) GetRoles() []*RoleListItem {
	if m != nil {
		return m.Roles
	}
	return nil
}

type RoleListItem struct {
	// Name of the role.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Desscription of the role.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Environment for the role.
	Environments         []string `protobuf:"bytes,3,rep,name=environments,proto3" json:"environments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleListItem) Reset()         { *m = RoleListItem{} }
func (m *RoleListItem) String() string { return proto.CompactTextString(m) }
func (*RoleListItem) ProtoMessage()    {}
func (*RoleListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_6155b399355910ab, []int{1}
}

func (m *RoleListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleListItem.Unmarshal(m, b)
}
func (m *RoleListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleListItem.Marshal(b, m, deterministic)
}
func (m *RoleListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleListItem.Merge(m, src)
}
func (m *RoleListItem) XXX_Size() int {
	return xxx_messageInfo_RoleListItem.Size(m)
}
func (m *RoleListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleListItem.DiscardUnknown(m)
}

var xxx_messageInfo_RoleListItem proto.InternalMessageInfo

func (m *RoleListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoleListItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RoleListItem) GetEnvironments() []string {
	if m != nil {
		return m.Environments
	}
	return nil
}

type Role struct {
	// Name of the role.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the chef object.
	ChefType string `protobuf:"bytes,2,opt,name=chef_type,json=chefType,proto3" json:"chef_type,omitempty"`
	// Descrption of the role.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Role default attributes JSON.
	DefaultAttributes string `protobuf:"bytes,4,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Role override attributes JSON.
	OverrideAttributes string `protobuf:"bytes,5,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
	// Json class name.
	JsonClass string `protobuf:"bytes,6,opt,name=json_class,json=jsonClass,proto3" json:"json_class,omitempty"`
	// Run list for the role.
	RunList []string `protobuf:"bytes,7,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	// List of expanded run list for the role.
	ExpandedRunList      []*ExpandedRunList `protobuf:"bytes,8,rep,name=expanded_run_list,json=expandedRunList,proto3" json:"expanded_run_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_6155b399355910ab, []int{2}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetChefType() string {
	if m != nil {
		return m.ChefType
	}
	return ""
}

func (m *Role) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Role) GetDefaultAttributes() string {
	if m != nil {
		return m.DefaultAttributes
	}
	return ""
}

func (m *Role) GetOverrideAttributes() string {
	if m != nil {
		return m.OverrideAttributes
	}
	return ""
}

func (m *Role) GetJsonClass() string {
	if m != nil {
		return m.JsonClass
	}
	return ""
}

func (m *Role) GetRunList() []string {
	if m != nil {
		return m.RunList
	}
	return nil
}

func (m *Role) GetExpandedRunList() []*ExpandedRunList {
	if m != nil {
		return m.ExpandedRunList
	}
	return nil
}

type ExpandedRunList struct {
	// ID of the run list collection.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// List of the run list.
	RunList              []*RunList `protobuf:"bytes,2,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ExpandedRunList) Reset()         { *m = ExpandedRunList{} }
func (m *ExpandedRunList) String() string { return proto.CompactTextString(m) }
func (*ExpandedRunList) ProtoMessage()    {}
func (*ExpandedRunList) Descriptor() ([]byte, []int) {
	return fileDescriptor_6155b399355910ab, []int{3}
}

func (m *ExpandedRunList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpandedRunList.Unmarshal(m, b)
}
func (m *ExpandedRunList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpandedRunList.Marshal(b, m, deterministic)
}
func (m *ExpandedRunList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpandedRunList.Merge(m, src)
}
func (m *ExpandedRunList) XXX_Size() int {
	return xxx_messageInfo_ExpandedRunList.Size(m)
}
func (m *ExpandedRunList) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpandedRunList.DiscardUnknown(m)
}

var xxx_messageInfo_ExpandedRunList proto.InternalMessageInfo

func (m *ExpandedRunList) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExpandedRunList) GetRunList() []*RunList {
	if m != nil {
		return m.RunList
	}
	return nil
}

type RunList struct {
	// Type of run list item (e.g. 'recipe').
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Name of run list item.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Version of run list item.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// Boolean denoting whether or not the run list item was skipped.
	Skipped bool `protobuf:"varint,4,opt,name=skipped,proto3" json:"skipped,omitempty"`
	// List of the run list.
	Children             []*RunList `protobuf:"bytes,5,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RunList) Reset()         { *m = RunList{} }
func (m *RunList) String() string { return proto.CompactTextString(m) }
func (*RunList) ProtoMessage()    {}
func (*RunList) Descriptor() ([]byte, []int) {
	return fileDescriptor_6155b399355910ab, []int{4}
}

func (m *RunList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunList.Unmarshal(m, b)
}
func (m *RunList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunList.Marshal(b, m, deterministic)
}
func (m *RunList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunList.Merge(m, src)
}
func (m *RunList) XXX_Size() int {
	return xxx_messageInfo_RunList.Size(m)
}
func (m *RunList) XXX_DiscardUnknown() {
	xxx_messageInfo_RunList.DiscardUnknown(m)
}

var xxx_messageInfo_RunList proto.InternalMessageInfo

func (m *RunList) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RunList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RunList) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RunList) GetSkipped() bool {
	if m != nil {
		return m.Skipped
	}
	return false
}

func (m *RunList) GetChildren() []*RunList {
	if m != nil {
		return m.Children
	}
	return nil
}

func init() {
	proto.RegisterType((*Roles)(nil), "chef.automate.api.infra_proxy.response.Roles")
	proto.RegisterType((*RoleListItem)(nil), "chef.automate.api.infra_proxy.response.RoleListItem")
	proto.RegisterType((*Role)(nil), "chef.automate.api.infra_proxy.response.Role")
	proto.RegisterType((*ExpandedRunList)(nil), "chef.automate.api.infra_proxy.response.ExpandedRunList")
	proto.RegisterType((*RunList)(nil), "chef.automate.api.infra_proxy.response.RunList")
}

func init() {
	proto.RegisterFile("external/infra_proxy/response/roles.proto", fileDescriptor_6155b399355910ab)
}

var fileDescriptor_6155b399355910ab = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x55, 0xd3, 0x76, 0x9a, 0xde, 0x19, 0x31, 0x1a, 0xb3, 0x31, 0x42, 0x48, 0x55, 0x16, 0xa8,
	0x2c, 0x88, 0x25, 0x40, 0x42, 0x42, 0x6c, 0x00, 0xb1, 0x60, 0x60, 0x15, 0x58, 0xb1, 0x89, 0xdc,
	0xf8, 0x96, 0x1a, 0x12, 0xdb, 0xb2, 0x9d, 0xaa, 0xfd, 0x2d, 0xbe, 0x86, 0xcf, 0x41, 0xf6, 0x24,
	0x25, 0x43, 0x11, 0xaa, 0x66, 0x17, 0x9f, 0x7b, 0x1e, 0xb6, 0x4f, 0x0c, 0x4f, 0x70, 0xe7, 0xd1,
	0x2a, 0x5e, 0x33, 0xa9, 0xd6, 0x96, 0x97, 0xc6, 0xea, 0xdd, 0x9e, 0x59, 0x74, 0x46, 0x2b, 0x87,
	0xcc, 0xea, 0x1a, 0x5d, 0x6e, 0xac, 0xf6, 0x9a, 0x3c, 0xae, 0x36, 0xb8, 0xce, 0x79, 0xeb, 0x75,
	0xc3, 0x3d, 0xe6, 0xdc, 0xc8, 0x7c, 0xa0, 0xc9, 0x7b, 0x4d, 0xf6, 0x19, 0xa6, 0x45, 0x90, 0x91,
	0x6b, 0x98, 0x46, 0x3d, 0x1d, 0x2d, 0xc6, 0xcb, 0xf3, 0x67, 0x2f, 0xf2, 0xd3, 0x0c, 0xf2, 0xa0,
	0xfe, 0x24, 0x9d, 0xff, 0xe0, 0xb1, 0x29, 0x6e, 0x2c, 0xb2, 0x0d, 0x5c, 0x0c, 0x61, 0x42, 0x60,
	0xa2, 0x78, 0x83, 0x74, 0xb4, 0x18, 0x2d, 0xe7, 0x45, 0xfc, 0x26, 0x0b, 0x38, 0x17, 0xe8, 0x2a,
	0x2b, 0x8d, 0x97, 0x5a, 0xd1, 0x24, 0x8e, 0x86, 0x10, 0xc9, 0xe0, 0x02, 0xd5, 0x56, 0x5a, 0xad,
	0x1a, 0x54, 0xde, 0xd1, 0xf1, 0x62, 0xbc, 0x9c, 0x17, 0xb7, 0xb0, 0xec, 0x57, 0x02, 0x93, 0x10,
	0xf5, 0xcf, 0x88, 0x87, 0x30, 0x0f, 0x87, 0x28, 0xfd, 0xde, 0x60, 0x17, 0x90, 0x06, 0xe0, 0xcb,
	0xde, 0x1c, 0xe5, 0x8f, 0x8f, 0xf3, 0x9f, 0x02, 0x11, 0xb8, 0xe6, 0x6d, 0xed, 0x4b, 0xee, 0xbd,
	0x95, 0xab, 0xd6, 0xa3, 0xa3, 0x93, 0x48, 0xbc, 0xea, 0x26, 0x6f, 0x0e, 0x03, 0xc2, 0xe0, 0xbe,
	0xde, 0xa2, 0xb5, 0x52, 0xe0, 0x90, 0x3f, 0x8d, 0x7c, 0xd2, 0x8f, 0x06, 0x82, 0x47, 0x00, 0xdf,
	0x9d, 0x56, 0x65, 0x55, 0x73, 0xe7, 0xe8, 0x59, 0xe4, 0xcd, 0x03, 0xf2, 0x2e, 0x00, 0xe4, 0x01,
	0xa4, 0xb6, 0x55, 0x65, 0x2d, 0x9d, 0xa7, 0xb3, 0x78, 0xf4, 0x99, 0x6d, 0x55, 0xb8, 0x53, 0x52,
	0xc1, 0x15, 0xee, 0x0c, 0x57, 0x02, 0x45, 0x79, 0xe0, 0xa4, 0xb1, 0xb7, 0x97, 0xa7, 0xf6, 0xf6,
	0xbe, 0x33, 0x28, 0x6e, 0x3c, 0x8b, 0x4b, 0xbc, 0x0d, 0x64, 0x0d, 0x5c, 0xfe, 0xc5, 0x21, 0xf7,
	0x20, 0x91, 0xa2, 0xbb, 0xe2, 0x44, 0x0a, 0x72, 0x3d, 0xd8, 0x62, 0x12, 0xe3, 0xd9, 0xc9, 0xbf,
	0x4d, 0x17, 0xdb, 0x9f, 0x29, 0xfb, 0x39, 0x82, 0x59, 0x9f, 0x43, 0x60, 0x12, 0x3b, 0xeb, 0xca,
	0x0c, 0xdf, 0x87, 0x82, 0x93, 0x41, 0xc1, 0x14, 0x66, 0x5b, 0xb4, 0xee, 0x4f, 0x7f, 0xfd, 0x32,
	0x4c, 0xdc, 0x0f, 0x69, 0x0c, 0x8a, 0x58, 0x58, 0x5a, 0xf4, 0x4b, 0xf2, 0x11, 0xd2, 0x6a, 0x23,
	0x6b, 0x61, 0x51, 0xd1, 0xe9, 0xdd, 0xf6, 0x7c, 0x30, 0x78, 0xfb, 0xfa, 0xeb, 0xab, 0x6f, 0xd2,
	0x6f, 0xda, 0x55, 0x5e, 0xe9, 0x86, 0x05, 0x1b, 0xd6, 0xdb, 0x30, 0x6e, 0x24, 0xfb, 0xef, 0x7b,
	0x5d, 0x9d, 0xc5, 0xa7, 0xfa, 0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xda, 0x0b, 0x53,
	0xd7, 0x03, 0x00, 0x00,
}
