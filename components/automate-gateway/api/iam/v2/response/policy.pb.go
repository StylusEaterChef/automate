// Code generated by protoc-gen-go. DO NOT EDIT.
// source: automate-gateway/api/iam/v2/response/policy.proto

package response

import (
	fmt "fmt"
	common "github.com/chef/automate/components/automate-gateway/api/iam/v2/common"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type CreatePolicyResp struct {
	Policy               *common.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreatePolicyResp) Reset()         { *m = CreatePolicyResp{} }
func (m *CreatePolicyResp) String() string { return proto.CompactTextString(m) }
func (*CreatePolicyResp) ProtoMessage()    {}
func (*CreatePolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{0}
}

func (m *CreatePolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePolicyResp.Unmarshal(m, b)
}
func (m *CreatePolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePolicyResp.Marshal(b, m, deterministic)
}
func (m *CreatePolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePolicyResp.Merge(m, src)
}
func (m *CreatePolicyResp) XXX_Size() int {
	return xxx_messageInfo_CreatePolicyResp.Size(m)
}
func (m *CreatePolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePolicyResp proto.InternalMessageInfo

func (m *CreatePolicyResp) GetPolicy() *common.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type GetPolicyResp struct {
	Policy               *common.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetPolicyResp) Reset()         { *m = GetPolicyResp{} }
func (m *GetPolicyResp) String() string { return proto.CompactTextString(m) }
func (*GetPolicyResp) ProtoMessage()    {}
func (*GetPolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{1}
}

func (m *GetPolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPolicyResp.Unmarshal(m, b)
}
func (m *GetPolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPolicyResp.Marshal(b, m, deterministic)
}
func (m *GetPolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPolicyResp.Merge(m, src)
}
func (m *GetPolicyResp) XXX_Size() int {
	return xxx_messageInfo_GetPolicyResp.Size(m)
}
func (m *GetPolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPolicyResp proto.InternalMessageInfo

func (m *GetPolicyResp) GetPolicy() *common.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ListPoliciesResp struct {
	Policies             []*common.Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListPoliciesResp) Reset()         { *m = ListPoliciesResp{} }
func (m *ListPoliciesResp) String() string { return proto.CompactTextString(m) }
func (*ListPoliciesResp) ProtoMessage()    {}
func (*ListPoliciesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{2}
}

func (m *ListPoliciesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPoliciesResp.Unmarshal(m, b)
}
func (m *ListPoliciesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPoliciesResp.Marshal(b, m, deterministic)
}
func (m *ListPoliciesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPoliciesResp.Merge(m, src)
}
func (m *ListPoliciesResp) XXX_Size() int {
	return xxx_messageInfo_ListPoliciesResp.Size(m)
}
func (m *ListPoliciesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPoliciesResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListPoliciesResp proto.InternalMessageInfo

func (m *ListPoliciesResp) GetPolicies() []*common.Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type DeletePolicyResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePolicyResp) Reset()         { *m = DeletePolicyResp{} }
func (m *DeletePolicyResp) String() string { return proto.CompactTextString(m) }
func (*DeletePolicyResp) ProtoMessage()    {}
func (*DeletePolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{3}
}

func (m *DeletePolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePolicyResp.Unmarshal(m, b)
}
func (m *DeletePolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePolicyResp.Marshal(b, m, deterministic)
}
func (m *DeletePolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePolicyResp.Merge(m, src)
}
func (m *DeletePolicyResp) XXX_Size() int {
	return xxx_messageInfo_DeletePolicyResp.Size(m)
}
func (m *DeletePolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePolicyResp proto.InternalMessageInfo

type UpdatePolicyResp struct {
	Policy               *common.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdatePolicyResp) Reset()         { *m = UpdatePolicyResp{} }
func (m *UpdatePolicyResp) String() string { return proto.CompactTextString(m) }
func (*UpdatePolicyResp) ProtoMessage()    {}
func (*UpdatePolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{4}
}

func (m *UpdatePolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePolicyResp.Unmarshal(m, b)
}
func (m *UpdatePolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePolicyResp.Marshal(b, m, deterministic)
}
func (m *UpdatePolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePolicyResp.Merge(m, src)
}
func (m *UpdatePolicyResp) XXX_Size() int {
	return xxx_messageInfo_UpdatePolicyResp.Size(m)
}
func (m *UpdatePolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePolicyResp proto.InternalMessageInfo

func (m *UpdatePolicyResp) GetPolicy() *common.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type GetPolicyVersionResp struct {
	Version              *common.Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetPolicyVersionResp) Reset()         { *m = GetPolicyVersionResp{} }
func (m *GetPolicyVersionResp) String() string { return proto.CompactTextString(m) }
func (*GetPolicyVersionResp) ProtoMessage()    {}
func (*GetPolicyVersionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{5}
}

func (m *GetPolicyVersionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPolicyVersionResp.Unmarshal(m, b)
}
func (m *GetPolicyVersionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPolicyVersionResp.Marshal(b, m, deterministic)
}
func (m *GetPolicyVersionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPolicyVersionResp.Merge(m, src)
}
func (m *GetPolicyVersionResp) XXX_Size() int {
	return xxx_messageInfo_GetPolicyVersionResp.Size(m)
}
func (m *GetPolicyVersionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPolicyVersionResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPolicyVersionResp proto.InternalMessageInfo

func (m *GetPolicyVersionResp) GetVersion() *common.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type ListPolicyMembersResp struct {
	// List of policy members.
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPolicyMembersResp) Reset()         { *m = ListPolicyMembersResp{} }
func (m *ListPolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*ListPolicyMembersResp) ProtoMessage()    {}
func (*ListPolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{6}
}

func (m *ListPolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPolicyMembersResp.Unmarshal(m, b)
}
func (m *ListPolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *ListPolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPolicyMembersResp.Merge(m, src)
}
func (m *ListPolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_ListPolicyMembersResp.Size(m)
}
func (m *ListPolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListPolicyMembersResp proto.InternalMessageInfo

func (m *ListPolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type ReplacePolicyMembersResp struct {
	// Resulting list of policy members.
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplacePolicyMembersResp) Reset()         { *m = ReplacePolicyMembersResp{} }
func (m *ReplacePolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*ReplacePolicyMembersResp) ProtoMessage()    {}
func (*ReplacePolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{7}
}

func (m *ReplacePolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplacePolicyMembersResp.Unmarshal(m, b)
}
func (m *ReplacePolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplacePolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *ReplacePolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplacePolicyMembersResp.Merge(m, src)
}
func (m *ReplacePolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_ReplacePolicyMembersResp.Size(m)
}
func (m *ReplacePolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplacePolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_ReplacePolicyMembersResp proto.InternalMessageInfo

func (m *ReplacePolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type RemovePolicyMembersResp struct {
	// Resulting list of policy members.
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePolicyMembersResp) Reset()         { *m = RemovePolicyMembersResp{} }
func (m *RemovePolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*RemovePolicyMembersResp) ProtoMessage()    {}
func (*RemovePolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{8}
}

func (m *RemovePolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePolicyMembersResp.Unmarshal(m, b)
}
func (m *RemovePolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *RemovePolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePolicyMembersResp.Merge(m, src)
}
func (m *RemovePolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_RemovePolicyMembersResp.Size(m)
}
func (m *RemovePolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePolicyMembersResp proto.InternalMessageInfo

func (m *RemovePolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type AddPolicyMembersResp struct {
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPolicyMembersResp) Reset()         { *m = AddPolicyMembersResp{} }
func (m *AddPolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*AddPolicyMembersResp) ProtoMessage()    {}
func (*AddPolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{9}
}

func (m *AddPolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPolicyMembersResp.Unmarshal(m, b)
}
func (m *AddPolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *AddPolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPolicyMembersResp.Merge(m, src)
}
func (m *AddPolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_AddPolicyMembersResp.Size(m)
}
func (m *AddPolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddPolicyMembersResp proto.InternalMessageInfo

func (m *AddPolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type CreateRoleResp struct {
	Role                 *common.Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateRoleResp) Reset()         { *m = CreateRoleResp{} }
func (m *CreateRoleResp) String() string { return proto.CompactTextString(m) }
func (*CreateRoleResp) ProtoMessage()    {}
func (*CreateRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{10}
}

func (m *CreateRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleResp.Unmarshal(m, b)
}
func (m *CreateRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleResp.Marshal(b, m, deterministic)
}
func (m *CreateRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleResp.Merge(m, src)
}
func (m *CreateRoleResp) XXX_Size() int {
	return xxx_messageInfo_CreateRoleResp.Size(m)
}
func (m *CreateRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleResp proto.InternalMessageInfo

func (m *CreateRoleResp) GetRole() *common.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type ListRolesResp struct {
	Roles                []*common.Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListRolesResp) Reset()         { *m = ListRolesResp{} }
func (m *ListRolesResp) String() string { return proto.CompactTextString(m) }
func (*ListRolesResp) ProtoMessage()    {}
func (*ListRolesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{11}
}

func (m *ListRolesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesResp.Unmarshal(m, b)
}
func (m *ListRolesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesResp.Marshal(b, m, deterministic)
}
func (m *ListRolesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesResp.Merge(m, src)
}
func (m *ListRolesResp) XXX_Size() int {
	return xxx_messageInfo_ListRolesResp.Size(m)
}
func (m *ListRolesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesResp proto.InternalMessageInfo

func (m *ListRolesResp) GetRoles() []*common.Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type GetRoleResp struct {
	Role                 *common.Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetRoleResp) Reset()         { *m = GetRoleResp{} }
func (m *GetRoleResp) String() string { return proto.CompactTextString(m) }
func (*GetRoleResp) ProtoMessage()    {}
func (*GetRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{12}
}

func (m *GetRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoleResp.Unmarshal(m, b)
}
func (m *GetRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoleResp.Marshal(b, m, deterministic)
}
func (m *GetRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoleResp.Merge(m, src)
}
func (m *GetRoleResp) XXX_Size() int {
	return xxx_messageInfo_GetRoleResp.Size(m)
}
func (m *GetRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoleResp proto.InternalMessageInfo

func (m *GetRoleResp) GetRole() *common.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type DeleteRoleResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoleResp) Reset()         { *m = DeleteRoleResp{} }
func (m *DeleteRoleResp) String() string { return proto.CompactTextString(m) }
func (*DeleteRoleResp) ProtoMessage()    {}
func (*DeleteRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{13}
}

func (m *DeleteRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoleResp.Unmarshal(m, b)
}
func (m *DeleteRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoleResp.Marshal(b, m, deterministic)
}
func (m *DeleteRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoleResp.Merge(m, src)
}
func (m *DeleteRoleResp) XXX_Size() int {
	return xxx_messageInfo_DeleteRoleResp.Size(m)
}
func (m *DeleteRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoleResp proto.InternalMessageInfo

type UpdateRoleResp struct {
	Role                 *common.Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateRoleResp) Reset()         { *m = UpdateRoleResp{} }
func (m *UpdateRoleResp) String() string { return proto.CompactTextString(m) }
func (*UpdateRoleResp) ProtoMessage()    {}
func (*UpdateRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{14}
}

func (m *UpdateRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRoleResp.Unmarshal(m, b)
}
func (m *UpdateRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRoleResp.Marshal(b, m, deterministic)
}
func (m *UpdateRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRoleResp.Merge(m, src)
}
func (m *UpdateRoleResp) XXX_Size() int {
	return xxx_messageInfo_UpdateRoleResp.Size(m)
}
func (m *UpdateRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRoleResp proto.InternalMessageInfo

func (m *UpdateRoleResp) GetRole() *common.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type CreateProjectResp struct {
	Project              *common.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateProjectResp) Reset()         { *m = CreateProjectResp{} }
func (m *CreateProjectResp) String() string { return proto.CompactTextString(m) }
func (*CreateProjectResp) ProtoMessage()    {}
func (*CreateProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{15}
}

func (m *CreateProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProjectResp.Unmarshal(m, b)
}
func (m *CreateProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProjectResp.Marshal(b, m, deterministic)
}
func (m *CreateProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProjectResp.Merge(m, src)
}
func (m *CreateProjectResp) XXX_Size() int {
	return xxx_messageInfo_CreateProjectResp.Size(m)
}
func (m *CreateProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProjectResp proto.InternalMessageInfo

func (m *CreateProjectResp) GetProject() *common.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type UpdateProjectResp struct {
	Project              *common.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateProjectResp) Reset()         { *m = UpdateProjectResp{} }
func (m *UpdateProjectResp) String() string { return proto.CompactTextString(m) }
func (*UpdateProjectResp) ProtoMessage()    {}
func (*UpdateProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{16}
}

func (m *UpdateProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProjectResp.Unmarshal(m, b)
}
func (m *UpdateProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProjectResp.Marshal(b, m, deterministic)
}
func (m *UpdateProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProjectResp.Merge(m, src)
}
func (m *UpdateProjectResp) XXX_Size() int {
	return xxx_messageInfo_UpdateProjectResp.Size(m)
}
func (m *UpdateProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProjectResp proto.InternalMessageInfo

func (m *UpdateProjectResp) GetProject() *common.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type GetProjectResp struct {
	Project              *common.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetProjectResp) Reset()         { *m = GetProjectResp{} }
func (m *GetProjectResp) String() string { return proto.CompactTextString(m) }
func (*GetProjectResp) ProtoMessage()    {}
func (*GetProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{17}
}

func (m *GetProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProjectResp.Unmarshal(m, b)
}
func (m *GetProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProjectResp.Marshal(b, m, deterministic)
}
func (m *GetProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProjectResp.Merge(m, src)
}
func (m *GetProjectResp) XXX_Size() int {
	return xxx_messageInfo_GetProjectResp.Size(m)
}
func (m *GetProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetProjectResp proto.InternalMessageInfo

func (m *GetProjectResp) GetProject() *common.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type ListProjectsResp struct {
	Projects             []*common.Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListProjectsResp) Reset()         { *m = ListProjectsResp{} }
func (m *ListProjectsResp) String() string { return proto.CompactTextString(m) }
func (*ListProjectsResp) ProtoMessage()    {}
func (*ListProjectsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{18}
}

func (m *ListProjectsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsResp.Unmarshal(m, b)
}
func (m *ListProjectsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsResp.Marshal(b, m, deterministic)
}
func (m *ListProjectsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsResp.Merge(m, src)
}
func (m *ListProjectsResp) XXX_Size() int {
	return xxx_messageInfo_ListProjectsResp.Size(m)
}
func (m *ListProjectsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsResp proto.InternalMessageInfo

func (m *ListProjectsResp) GetProjects() []*common.Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type DeleteProjectResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProjectResp) Reset()         { *m = DeleteProjectResp{} }
func (m *DeleteProjectResp) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectResp) ProtoMessage()    {}
func (*DeleteProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{19}
}

func (m *DeleteProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectResp.Unmarshal(m, b)
}
func (m *DeleteProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectResp.Marshal(b, m, deterministic)
}
func (m *DeleteProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectResp.Merge(m, src)
}
func (m *DeleteProjectResp) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectResp.Size(m)
}
func (m *DeleteProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectResp proto.InternalMessageInfo

type UpgradeToV2Resp struct {
	Reports              []string `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpgradeToV2Resp) Reset()         { *m = UpgradeToV2Resp{} }
func (m *UpgradeToV2Resp) String() string { return proto.CompactTextString(m) }
func (*UpgradeToV2Resp) ProtoMessage()    {}
func (*UpgradeToV2Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{20}
}

func (m *UpgradeToV2Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpgradeToV2Resp.Unmarshal(m, b)
}
func (m *UpgradeToV2Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpgradeToV2Resp.Marshal(b, m, deterministic)
}
func (m *UpgradeToV2Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpgradeToV2Resp.Merge(m, src)
}
func (m *UpgradeToV2Resp) XXX_Size() int {
	return xxx_messageInfo_UpgradeToV2Resp.Size(m)
}
func (m *UpgradeToV2Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpgradeToV2Resp.DiscardUnknown(m)
}

var xxx_messageInfo_UpgradeToV2Resp proto.InternalMessageInfo

func (m *UpgradeToV2Resp) GetReports() []string {
	if m != nil {
		return m.Reports
	}
	return nil
}

type ResetToV1Resp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetToV1Resp) Reset()         { *m = ResetToV1Resp{} }
func (m *ResetToV1Resp) String() string { return proto.CompactTextString(m) }
func (*ResetToV1Resp) ProtoMessage()    {}
func (*ResetToV1Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e019340cfee872, []int{21}
}

func (m *ResetToV1Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetToV1Resp.Unmarshal(m, b)
}
func (m *ResetToV1Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetToV1Resp.Marshal(b, m, deterministic)
}
func (m *ResetToV1Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetToV1Resp.Merge(m, src)
}
func (m *ResetToV1Resp) XXX_Size() int {
	return xxx_messageInfo_ResetToV1Resp.Size(m)
}
func (m *ResetToV1Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetToV1Resp.DiscardUnknown(m)
}

var xxx_messageInfo_ResetToV1Resp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreatePolicyResp)(nil), "chef.automate.api.iam.v2.CreatePolicyResp")
	proto.RegisterType((*GetPolicyResp)(nil), "chef.automate.api.iam.v2.GetPolicyResp")
	proto.RegisterType((*ListPoliciesResp)(nil), "chef.automate.api.iam.v2.ListPoliciesResp")
	proto.RegisterType((*DeletePolicyResp)(nil), "chef.automate.api.iam.v2.DeletePolicyResp")
	proto.RegisterType((*UpdatePolicyResp)(nil), "chef.automate.api.iam.v2.UpdatePolicyResp")
	proto.RegisterType((*GetPolicyVersionResp)(nil), "chef.automate.api.iam.v2.GetPolicyVersionResp")
	proto.RegisterType((*ListPolicyMembersResp)(nil), "chef.automate.api.iam.v2.ListPolicyMembersResp")
	proto.RegisterType((*ReplacePolicyMembersResp)(nil), "chef.automate.api.iam.v2.ReplacePolicyMembersResp")
	proto.RegisterType((*RemovePolicyMembersResp)(nil), "chef.automate.api.iam.v2.RemovePolicyMembersResp")
	proto.RegisterType((*AddPolicyMembersResp)(nil), "chef.automate.api.iam.v2.AddPolicyMembersResp")
	proto.RegisterType((*CreateRoleResp)(nil), "chef.automate.api.iam.v2.CreateRoleResp")
	proto.RegisterType((*ListRolesResp)(nil), "chef.automate.api.iam.v2.ListRolesResp")
	proto.RegisterType((*GetRoleResp)(nil), "chef.automate.api.iam.v2.GetRoleResp")
	proto.RegisterType((*DeleteRoleResp)(nil), "chef.automate.api.iam.v2.DeleteRoleResp")
	proto.RegisterType((*UpdateRoleResp)(nil), "chef.automate.api.iam.v2.UpdateRoleResp")
	proto.RegisterType((*CreateProjectResp)(nil), "chef.automate.api.iam.v2.CreateProjectResp")
	proto.RegisterType((*UpdateProjectResp)(nil), "chef.automate.api.iam.v2.UpdateProjectResp")
	proto.RegisterType((*GetProjectResp)(nil), "chef.automate.api.iam.v2.GetProjectResp")
	proto.RegisterType((*ListProjectsResp)(nil), "chef.automate.api.iam.v2.ListProjectsResp")
	proto.RegisterType((*DeleteProjectResp)(nil), "chef.automate.api.iam.v2.DeleteProjectResp")
	proto.RegisterType((*UpgradeToV2Resp)(nil), "chef.automate.api.iam.v2.UpgradeToV2Resp")
	proto.RegisterType((*ResetToV1Resp)(nil), "chef.automate.api.iam.v2.ResetToV1Resp")
}

func init() {
	proto.RegisterFile("automate-gateway/api/iam/v2/response/policy.proto", fileDescriptor_46e019340cfee872)
}

var fileDescriptor_46e019340cfee872 = []byte{
	// 1040 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x98, 0x6f, 0x8b, 0x1b, 0x45,
	0x18, 0xc0, 0xd9, 0x4d, 0x6d, 0xeb, 0x94, 0xbb, 0x5e, 0x97, 0x8a, 0xa1, 0x2f, 0x64, 0x1d, 0x84,
	0x6a, 0x9b, 0x64, 0xc9, 0xf8, 0xaf, 0x8e, 0xed, 0x8b, 0xb5, 0x42, 0x45, 0xee, 0xb0, 0xc6, 0x36,
	0x82, 0x7a, 0x2f, 0xe6, 0x36, 0xcf, 0xa5, 0x5b, 0xb2, 0x3b, 0xdb, 0x9d, 0x49, 0xc2, 0x11, 0x02,
	0x42, 0xa5, 0x62, 0x95, 0x0a, 0xeb, 0x17, 0x10, 0x7c, 0xed, 0x6b, 0xc1, 0x0f, 0x21, 0x8a, 0x22,
	0x52, 0xd0, 0x17, 0x7e, 0x12, 0x99, 0x3f, 0x9b, 0xbb, 0xc4, 0xe4, 0x7a, 0xa7, 0xf6, 0x6c, 0x4b,
	0x5f, 0x25, 0xb3, 0xf3, 0xfc, 0x9b, 0xdf, 0xcc, 0x3c, 0xcf, 0xb3, 0x8b, 0x9a, 0xac, 0x2f, 0x79,
	0xc2, 0x24, 0xd4, 0xbb, 0x4c, 0xc2, 0x90, 0x6d, 0x05, 0x2c, 0x8b, 0x83, 0x98, 0x25, 0xc1, 0x80,
	0x04, 0x39, 0x88, 0x8c, 0xa7, 0x02, 0x82, 0x8c, 0xf7, 0xe2, 0x68, 0xab, 0x91, 0xe5, 0x5c, 0x72,
	0xaf, 0x1a, 0x5d, 0x83, 0xcd, 0x46, 0xa9, 0xd7, 0x60, 0x59, 0xdc, 0x88, 0x59, 0xd2, 0x18, 0x90,
	0x53, 0xc1, 0x6e, 0xc6, 0x22, 0x9e, 0x24, 0x3c, 0x9d, 0x32, 0x75, 0xaa, 0xa6, 0x7f, 0xa2, 0x7a,
	0x17, 0xd2, 0xba, 0x18, 0xb2, 0x6e, 0x17, 0xf2, 0x80, 0x67, 0x32, 0xe6, 0xa9, 0x08, 0x58, 0x9a,
	0x72, 0xc9, 0xf4, 0x7f, 0x23, 0x8d, 0x7f, 0x77, 0xd1, 0xca, 0xc5, 0x1c, 0x98, 0x84, 0xcb, 0xda,
	0x48, 0x0b, 0x44, 0xe6, 0x9d, 0x43, 0x87, 0x8d, 0xc9, 0xaa, 0xe3, 0x3b, 0xcf, 0x1f, 0x23, 0x7e,
	0x63, 0x51, 0x78, 0x0d, 0xab, 0x65, 0xe5, 0xe9, 0x1d, 0xb7, 0x08, 0x6f, 0xbb, 0xe4, 0x53, 0xd7,
	0xfb, 0xc4, 0x1d, 0xe1, 0x94, 0x25, 0x80, 0xa9, 0x8f, 0xd7, 0xb6, 0xfc, 0x76, 0x0c, 0x43, 0xc8,
	0x7d, 0x23, 0x8e, 0x6b, 0x38, 0xee, 0xa8, 0x89, 0xa8, 0x2f, 0x24, 0x4f, 0xea, 0x03, 0x3d, 0x59,
	0xcf, 0xca, 0xc9, 0x04, 0x92, 0x0d, 0xc8, 0x05, 0xa6, 0xfe, 0x87, 0x58, 0x02, 0x4b, 0x68, 0xaf,
	0xc3, 0x32, 0x7a, 0x06, 0xaf, 0xd7, 0x7c, 0x2c, 0x24, 0x93, 0x90, 0x40, 0x2a, 0xf5, 0xf4, 0x08,
	0xe7, 0xbc, 0xa7, 0x7d, 0x18, 0x1b, 0xb8, 0x86, 0xb3, 0x9c, 0x5f, 0x87, 0xc8, 0x4c, 0x97, 0x83,
	0x26, 0xae, 0xf9, 0xe5, 0x7f, 0xa2, 0xcd, 0xc0, 0xe6, 0x26, 0x44, 0x52, 0x29, 0x86, 0xab, 0xab,
	0xef, 0xbc, 0x8f, 0xc7, 0x35, 0x7f, 0xdb, 0xd6, 0x0d, 0x36, 0x6b, 0x87, 0x45, 0x11, 0x64, 0x92,
	0xa5, 0x11, 0x5c, 0x36, 0x8f, 0xe7, 0x1b, 0x59, 0x9f, 0x56, 0x5b, 0x1f, 0xe3, 0xbb, 0x2e, 0x5a,
	0xba, 0x04, 0xf2, 0x31, 0xdc, 0xfb, 0x01, 0xf7, 0xdb, 0x43, 0x68, 0x65, 0x35, 0x16, 0x86, 0x6e,
	0x0c, 0x42, 0xf3, 0x3d, 0x8f, 0x8e, 0x66, 0x76, 0x5c, 0x75, 0xfc, 0xca, 0x9e, 0x08, 0x4f, 0x34,
	0xe8, 0xdd, 0x4a, 0x11, 0xfe, 0x5a, 0x21, 0x3f, 0x57, 0xbc, 0x1f, 0x2b, 0x23, 0x5c, 0x3e, 0x36,
	0x34, 0x16, 0x11, 0xf7, 0x9b, 0xbb, 0x32, 0xaf, 0x37, 0x1f, 0x29, 0xea, 0xca, 0xea, 0x0e, 0x12,
	0x17, 0xf5, 0x8a, 0x4b, 0x12, 0x64, 0x96, 0x84, 0x45, 0x40, 0xe6, 0x22, 0xe0, 0x11, 0xeb, 0x51,
	0x09, 0x42, 0xee, 0x8a, 0x81, 0xf5, 0x3b, 0xb1, 0xe4, 0xf9, 0x8e, 0xf5, 0xce, 0x80, 0xd8, 0x5b,
	0xdc, 0xeb, 0x63, 0xfc, 0x8d, 0x8b, 0x56, 0xde, 0x84, 0x1e, 0xec, 0x4c, 0x76, 0x8f, 0x6f, 0xd5,
	0xec, 0xad, 0xfa, 0xd3, 0x45, 0x2b, 0x57, 0xb3, 0xce, 0x7f, 0x55, 0x12, 0xbe, 0x72, 0x8b, 0xf0,
	0x4b, 0x97, 0x7c, 0xe1, 0x7a, 0x9f, 0x4d, 0xf3, 0x35, 0x3e, 0x3a, 0xb3, 0x9c, 0xfd, 0x29, 0x96,
	0x7d, 0x01, 0xb9, 0x61, 0x99, 0xc2, 0x50, 0x0d, 0x94, 0xc0, 0x36, 0xe0, 0x14, 0x86, 0x6a, 0xb0,
	0x5f, 0xcc, 0xfb, 0xa2, 0x7c, 0x5f, 0x20, 0xdf, 0x74, 0xd0, 0xc9, 0x49, 0x5d, 0x68, 0x43, 0x2e,
	0x62, 0x9e, 0x6a, 0xd0, 0xaf, 0xa3, 0x23, 0x03, 0x33, 0xb4, 0xa4, 0x9f, 0x5d, 0x4c, 0xba, 0xd4,
	0x2b, 0x35, 0x68, 0xa3, 0x08, 0xcf, 0x92, 0x17, 0xbc, 0xd3, 0x23, 0x6c, 0x9f, 0x60, 0x3a, 0xc2,
	0x09, 0xbb, 0xce, 0x73, 0x4c, 0x71, 0x5b, 0x5f, 0xc2, 0x38, 0x35, 0x83, 0x26, 0x1e, 0x8f, 0x71,
	0x86, 0x9e, 0x9a, 0xe4, 0xcf, 0xad, 0x35, 0x83, 0x5d, 0x47, 0x51, 0x45, 0x47, 0xec, 0x2e, 0xe8,
	0x1c, 0xfa, 0x64, 0xab, 0x1c, 0xd2, 0x0b, 0x45, 0x48, 0xc9, 0x39, 0xef, 0x95, 0xd1, 0xa2, 0xfb,
	0x6c, 0x70, 0x0b, 0x05, 0xd6, 0x6c, 0xde, 0x8e, 0x5b, 0x3e, 0xc6, 0x02, 0x55, 0x5b, 0x90, 0xf5,
	0x58, 0x04, 0x07, 0xe8, 0xf4, 0x23, 0xf4, 0x74, 0x0b, 0x12, 0x3e, 0xd8, 0x97, 0xcf, 0xd3, 0x45,
	0xf8, 0x1c, 0xc1, 0x9e, 0x3f, 0x9a, 0x73, 0x1e, 0xa7, 0xac, 0x73, 0x74, 0x32, 0xec, 0x74, 0x0e,
	0x70, 0x39, 0xdf, 0x3b, 0x68, 0xd9, 0xf4, 0x6c, 0x2d, 0xde, 0x03, 0xed, 0x8b, 0xa0, 0x43, 0xea,
	0x90, 0xda, 0x23, 0xf3, 0xcc, 0xe2, 0x23, 0xa3, 0x35, 0xb4, 0x2c, 0x95, 0x45, 0x78, 0x83, 0x70,
	0x2f, 0x19, 0x4d, 0x67, 0x36, 0x7d, 0xd8, 0x6b, 0xfe, 0x9c, 0x1c, 0xff, 0x1e, 0x44, 0x39, 0x48,
	0x7f, 0x8d, 0xa5, 0xac, 0x0b, 0xb9, 0xdf, 0xb2, 0x82, 0x2c, 0xd2, 0x0d, 0xa5, 0x5e, 0x83, 0xd0,
	0x22, 0x82, 0x9e, 0x51, 0x13, 0x31, 0x4b, 0x68, 0x79, 0xec, 0x69, 0x2f, 0x36, 0xc1, 0x7f, 0xed,
	0xa2, 0x25, 0x75, 0xe6, 0x94, 0xb6, 0xe1, 0xf4, 0x12, 0x7a, 0x42, 0xf9, 0x2c, 0xab, 0xf5, 0xbd,
	0x82, 0x37, 0xc2, 0xf4, 0x37, 0xa7, 0x08, 0x7f, 0x71, 0xc8, 0x4f, 0x8e, 0xf7, 0x83, 0x63, 0x6e,
	0xa8, 0xbd, 0xf6, 0x07, 0xb8, 0x18, 0x95, 0x80, 0xff, 0xe6, 0x4e, 0xd5, 0xba, 0x5d, 0x1c, 0x2a,
	0x47, 0x3e, 0x99, 0x75, 0x15, 0xa7, 0x9b, 0x39, 0x53, 0xe5, 0x42, 0x95, 0xa9, 0xef, 0x1c, 0x74,
	0xec, 0x12, 0xc8, 0x87, 0x70, 0x73, 0x6f, 0x39, 0x68, 0xd9, 0x14, 0xd8, 0x32, 0xf8, 0xff, 0x29,
	0x10, 0x75, 0x45, 0x4c, 0x7d, 0x79, 0x08, 0x29, 0xde, 0x71, 0xd0, 0x09, 0xfb, 0x4e, 0x66, 0xa6,
	0xca, 0xc2, 0x60, 0x25, 0xef, 0x5d, 0x18, 0x4a, 0xbd, 0x52, 0x83, 0x9e, 0x2f, 0xc2, 0xd7, 0xc8,
	0xab, 0xde, 0xcb, 0x73, 0x5b, 0x37, 0x5b, 0xb9, 0x54, 0x38, 0x53, 0xad, 0x9b, 0x7d, 0x6e, 0x02,
	0xb2, 0x1d, 0xc1, 0x83, 0x11, 0xd0, 0xe7, 0x0e, 0x5a, 0x56, 0xd5, 0xf3, 0xc1, 0x88, 0xe6, 0x0f,
	0xc7, 0xbe, 0x86, 0xd8, 0x8d, 0xd4, 0xf1, 0x5c, 0x40, 0x47, 0xcb, 0x8d, 0xb5, 0x89, 0x6d, 0x0f,
	0x01, 0x4d, 0x54, 0xe8, 0x6d, 0xa7, 0x08, 0x6f, 0x39, 0xe4, 0xa6, 0xe3, 0x7d, 0xec, 0x8c, 0xa6,
	0x9a, 0x87, 0x7f, 0x14, 0xe3, 0x82, 0xa6, 0xdd, 0x4c, 0x9b, 0x74, 0x34, 0x4f, 0xb1, 0x4e, 0xb0,
	0x4a, 0x48, 0xef, 0xa2, 0x13, 0xb6, 0x6d, 0xde, 0x06, 0xfe, 0x2f, 0x99, 0x9d, 0x45, 0xc7, 0xaf,
	0x66, 0xdd, 0x9c, 0x75, 0xe0, 0x0a, 0x6f, 0x93, 0xb2, 0x5e, 0xe6, 0x90, 0xf1, 0x5c, 0x4e, 0xea,
	0xa5, 0x1d, 0xe2, 0xe3, 0x68, 0xa9, 0x05, 0x02, 0xe4, 0x15, 0xde, 0x6e, 0x2a, 0xd1, 0x37, 0xde,
	0xfe, 0xe0, 0xad, 0x6e, 0x2c, 0xaf, 0xf5, 0x37, 0x1a, 0x11, 0x4f, 0x02, 0x85, 0x75, 0xf2, 0x99,
	0x24, 0x88, 0x78, 0x92, 0xf1, 0x54, 0xb5, 0x81, 0xc1, 0x5e, 0xbe, 0xc3, 0x6c, 0x1c, 0xd6, 0x1f,
	0x42, 0x5e, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x31, 0x8f, 0x5b, 0x9a, 0xb6, 0x11, 0x00, 0x00,
}
