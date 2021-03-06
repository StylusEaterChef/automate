// Code generated by protoc-gen-go. DO NOT EDIT.
// source: external/infra_proxy/response/orgs.proto

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

type CreateOrg struct {
	// Chef organization.
	Org                  *Org     `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOrg) Reset()         { *m = CreateOrg{} }
func (m *CreateOrg) String() string { return proto.CompactTextString(m) }
func (*CreateOrg) ProtoMessage()    {}
func (*CreateOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e8985271c7ebca, []int{0}
}

func (m *CreateOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrg.Unmarshal(m, b)
}
func (m *CreateOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrg.Marshal(b, m, deterministic)
}
func (m *CreateOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrg.Merge(m, src)
}
func (m *CreateOrg) XXX_Size() int {
	return xxx_messageInfo_CreateOrg.Size(m)
}
func (m *CreateOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrg proto.InternalMessageInfo

func (m *CreateOrg) GetOrg() *Org {
	if m != nil {
		return m.Org
	}
	return nil
}

type DeleteOrg struct {
	// Chef organization.
	Org                  *Org     `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteOrg) Reset()         { *m = DeleteOrg{} }
func (m *DeleteOrg) String() string { return proto.CompactTextString(m) }
func (*DeleteOrg) ProtoMessage()    {}
func (*DeleteOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e8985271c7ebca, []int{1}
}

func (m *DeleteOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteOrg.Unmarshal(m, b)
}
func (m *DeleteOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteOrg.Marshal(b, m, deterministic)
}
func (m *DeleteOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteOrg.Merge(m, src)
}
func (m *DeleteOrg) XXX_Size() int {
	return xxx_messageInfo_DeleteOrg.Size(m)
}
func (m *DeleteOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteOrg.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteOrg proto.InternalMessageInfo

func (m *DeleteOrg) GetOrg() *Org {
	if m != nil {
		return m.Org
	}
	return nil
}

type UpdateOrg struct {
	// Chef organization.
	Org                  *Org     `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateOrg) Reset()         { *m = UpdateOrg{} }
func (m *UpdateOrg) String() string { return proto.CompactTextString(m) }
func (*UpdateOrg) ProtoMessage()    {}
func (*UpdateOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e8985271c7ebca, []int{2}
}

func (m *UpdateOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateOrg.Unmarshal(m, b)
}
func (m *UpdateOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateOrg.Marshal(b, m, deterministic)
}
func (m *UpdateOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateOrg.Merge(m, src)
}
func (m *UpdateOrg) XXX_Size() int {
	return xxx_messageInfo_UpdateOrg.Size(m)
}
func (m *UpdateOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateOrg.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateOrg proto.InternalMessageInfo

func (m *UpdateOrg) GetOrg() *Org {
	if m != nil {
		return m.Org
	}
	return nil
}

type GetOrgs struct {
	// Chef organization list.
	Orgs                 []*Org   `protobuf:"bytes,1,rep,name=orgs,proto3" json:"orgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrgs) Reset()         { *m = GetOrgs{} }
func (m *GetOrgs) String() string { return proto.CompactTextString(m) }
func (*GetOrgs) ProtoMessage()    {}
func (*GetOrgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e8985271c7ebca, []int{3}
}

func (m *GetOrgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrgs.Unmarshal(m, b)
}
func (m *GetOrgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrgs.Marshal(b, m, deterministic)
}
func (m *GetOrgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrgs.Merge(m, src)
}
func (m *GetOrgs) XXX_Size() int {
	return xxx_messageInfo_GetOrgs.Size(m)
}
func (m *GetOrgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrgs.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrgs proto.InternalMessageInfo

func (m *GetOrgs) GetOrgs() []*Org {
	if m != nil {
		return m.Orgs
	}
	return nil
}

type GetOrg struct {
	// Chef organization.
	Org                  *Org     `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrg) Reset()         { *m = GetOrg{} }
func (m *GetOrg) String() string { return proto.CompactTextString(m) }
func (*GetOrg) ProtoMessage()    {}
func (*GetOrg) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e8985271c7ebca, []int{4}
}

func (m *GetOrg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrg.Unmarshal(m, b)
}
func (m *GetOrg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrg.Marshal(b, m, deterministic)
}
func (m *GetOrg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrg.Merge(m, src)
}
func (m *GetOrg) XXX_Size() int {
	return xxx_messageInfo_GetOrg.Size(m)
}
func (m *GetOrg) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrg.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrg proto.InternalMessageInfo

func (m *GetOrg) GetOrg() *Org {
	if m != nil {
		return m.Org
	}
	return nil
}

type Org struct {
	// Chef organization ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Chef organization name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Chef organization admin user.
	AdminUser string `protobuf:"bytes,3,opt,name=admin_user,json=adminUser,proto3" json:"admin_user,omitempty"`
	// Chef organization credential ID.
	CredentialId string `protobuf:"bytes,4,opt,name=credential_id,json=credentialId,proto3" json:"credential_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,5,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// List of projects this chef organization belongs to. May be empty.
	Projects             []string `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Org) Reset()         { *m = Org{} }
func (m *Org) String() string { return proto.CompactTextString(m) }
func (*Org) ProtoMessage()    {}
func (*Org) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e8985271c7ebca, []int{5}
}

func (m *Org) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Org.Unmarshal(m, b)
}
func (m *Org) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Org.Marshal(b, m, deterministic)
}
func (m *Org) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Org.Merge(m, src)
}
func (m *Org) XXX_Size() int {
	return xxx_messageInfo_Org.Size(m)
}
func (m *Org) XXX_DiscardUnknown() {
	xxx_messageInfo_Org.DiscardUnknown(m)
}

var xxx_messageInfo_Org proto.InternalMessageInfo

func (m *Org) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Org) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Org) GetAdminUser() string {
	if m != nil {
		return m.AdminUser
	}
	return ""
}

func (m *Org) GetCredentialId() string {
	if m != nil {
		return m.CredentialId
	}
	return ""
}

func (m *Org) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *Org) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

type ResetOrgAdminKey struct {
	// Chef organization.
	Org                  *Org     `protobuf:"bytes,1,opt,name=org,proto3" json:"org,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetOrgAdminKey) Reset()         { *m = ResetOrgAdminKey{} }
func (m *ResetOrgAdminKey) String() string { return proto.CompactTextString(m) }
func (*ResetOrgAdminKey) ProtoMessage()    {}
func (*ResetOrgAdminKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e8985271c7ebca, []int{6}
}

func (m *ResetOrgAdminKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetOrgAdminKey.Unmarshal(m, b)
}
func (m *ResetOrgAdminKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetOrgAdminKey.Marshal(b, m, deterministic)
}
func (m *ResetOrgAdminKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetOrgAdminKey.Merge(m, src)
}
func (m *ResetOrgAdminKey) XXX_Size() int {
	return xxx_messageInfo_ResetOrgAdminKey.Size(m)
}
func (m *ResetOrgAdminKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetOrgAdminKey.DiscardUnknown(m)
}

var xxx_messageInfo_ResetOrgAdminKey proto.InternalMessageInfo

func (m *ResetOrgAdminKey) GetOrg() *Org {
	if m != nil {
		return m.Org
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateOrg)(nil), "chef.automate.api.infra_proxy.response.CreateOrg")
	proto.RegisterType((*DeleteOrg)(nil), "chef.automate.api.infra_proxy.response.DeleteOrg")
	proto.RegisterType((*UpdateOrg)(nil), "chef.automate.api.infra_proxy.response.UpdateOrg")
	proto.RegisterType((*GetOrgs)(nil), "chef.automate.api.infra_proxy.response.GetOrgs")
	proto.RegisterType((*GetOrg)(nil), "chef.automate.api.infra_proxy.response.GetOrg")
	proto.RegisterType((*Org)(nil), "chef.automate.api.infra_proxy.response.Org")
	proto.RegisterType((*ResetOrgAdminKey)(nil), "chef.automate.api.infra_proxy.response.ResetOrgAdminKey")
}

func init() {
	proto.RegisterFile("external/infra_proxy/response/orgs.proto", fileDescriptor_94e8985271c7ebca)
}

var fileDescriptor_94e8985271c7ebca = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xd9, 0x6e, 0xad, 0xdd, 0xf1, 0x0f, 0x92, 0xd3, 0xa2, 0x08, 0xa5, 0x82, 0x14, 0x84,
	0x2c, 0xe8, 0x4d, 0x14, 0xf1, 0x0f, 0x94, 0xea, 0xa1, 0x58, 0xe8, 0xc5, 0x4b, 0x49, 0x37, 0xd3,
	0x6d, 0xa4, 0xbb, 0x09, 0x93, 0x54, 0xda, 0x2f, 0xe4, 0xe7, 0x94, 0xa4, 0x14, 0x3d, 0x89, 0xd2,
	0xde, 0x92, 0x37, 0xef, 0xfd, 0x32, 0x13, 0x06, 0x3a, 0xb8, 0x70, 0x48, 0x95, 0x98, 0x65, 0xaa,
	0x9a, 0x90, 0x18, 0x19, 0xd2, 0x8b, 0x65, 0x46, 0x68, 0x8d, 0xae, 0x2c, 0x66, 0x9a, 0x0a, 0xcb,
	0x0d, 0x69, 0xa7, 0xd9, 0x79, 0x3e, 0xc5, 0x09, 0x17, 0x73, 0xa7, 0x4b, 0xe1, 0x90, 0x0b, 0xa3,
	0xf8, 0x8f, 0x08, 0x5f, 0x47, 0xda, 0xcf, 0x90, 0x3c, 0x12, 0x0a, 0x87, 0x7d, 0x2a, 0xd8, 0x2d,
	0xc4, 0x9a, 0x8a, 0x34, 0x6a, 0x45, 0x9d, 0xbd, 0xcb, 0x0b, 0xfe, 0x37, 0x04, 0xef, 0x53, 0x31,
	0xf0, 0x39, 0xcf, 0x7a, 0xc2, 0x19, 0x6e, 0x8b, 0x35, 0x34, 0x72, 0x5b, 0x7d, 0xed, 0x76, 0xd1,
	0xf5, 0xa9, 0xb0, 0xec, 0x0e, 0xea, 0xfe, 0x93, 0xd2, 0xa8, 0x15, 0xff, 0x17, 0x15, 0x82, 0xed,
	0x2e, 0x34, 0x56, 0xac, 0x4d, 0x9b, 0xfa, 0x8c, 0x20, 0xf6, 0x98, 0x43, 0xa8, 0x29, 0x19, 0x28,
	0xc9, 0xa0, 0xa6, 0x24, 0x63, 0x50, 0xaf, 0x44, 0x89, 0x69, 0x2d, 0x28, 0xe1, 0xcc, 0x4e, 0x01,
	0x84, 0x2c, 0x55, 0x35, 0x9a, 0x5b, 0xa4, 0x34, 0x0e, 0x95, 0x24, 0x28, 0x43, 0x8b, 0xc4, 0xce,
	0xe0, 0x20, 0x27, 0x94, 0x58, 0x39, 0x25, 0x66, 0x23, 0x25, 0xd3, 0x7a, 0x70, 0xec, 0x7f, 0x8b,
	0x3d, 0xc9, 0x4e, 0x20, 0xb1, 0x48, 0x1f, 0x48, 0xde, 0xb0, 0x13, 0x0c, 0xcd, 0x95, 0xd0, 0x93,
	0xec, 0x18, 0x9a, 0x86, 0xf4, 0x3b, 0xe6, 0xce, 0xa6, 0x8d, 0x56, 0xec, 0x6b, 0xeb, 0x7b, 0xfb,
	0x15, 0x8e, 0x06, 0x68, 0xc3, 0xcc, 0xf7, 0xfe, 0xc9, 0x17, 0x5c, 0x6e, 0x38, 0xfb, 0xc3, 0xcd,
	0xdb, 0x75, 0xa1, 0xdc, 0x74, 0x3e, 0xe6, 0xb9, 0x2e, 0x33, 0x9f, 0xce, 0xd6, 0xe9, 0x4c, 0x18,
	0x95, 0xfd, 0xba, 0xe5, 0xe3, 0x46, 0xd8, 0xf0, 0xab, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2,
	0xe3, 0x8c, 0xfb, 0x0d, 0x03, 0x00, 0x00,
}
