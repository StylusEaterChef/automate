// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: external/infra_proxy/response/databags.proto

package response

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DataBags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bags item list.
	DataBags []*DataBagListItem `protobuf:"bytes,2,rep,name=data_bags,json=dataBags,proto3" json:"data_bags,omitempty"`
}

func (x *DataBags) Reset() {
	*x = DataBags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_databags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBags) ProtoMessage() {}

func (x *DataBags) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_databags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBags.ProtoReflect.Descriptor instead.
func (*DataBags) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_databags_proto_rawDescGZIP(), []int{0}
}

func (x *DataBags) GetDataBags() []*DataBagListItem {
	if x != nil {
		return x.DataBags
	}
	return nil
}

type DataBagListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag item name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DataBagListItem) Reset() {
	*x = DataBagListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_databags_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBagListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBagListItem) ProtoMessage() {}

func (x *DataBagListItem) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_databags_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBagListItem.ProtoReflect.Descriptor instead.
func (*DataBagListItem) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_databags_proto_rawDescGZIP(), []int{1}
}

func (x *DataBagListItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DataBag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Data bag item ID.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Stringified json of data bag item.
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DataBag) Reset() {
	*x = DataBag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_databags_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBag) ProtoMessage() {}

func (x *DataBag) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_databags_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBag.ProtoReflect.Descriptor instead.
func (*DataBag) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_databags_proto_rawDescGZIP(), []int{2}
}

func (x *DataBag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataBag) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataBag) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type CreateDataBag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateDataBag) Reset() {
	*x = CreateDataBag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_databags_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDataBag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDataBag) ProtoMessage() {}

func (x *CreateDataBag) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_databags_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDataBag.ProtoReflect.Descriptor instead.
func (*CreateDataBag) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_databags_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDataBag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateDataBagItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Data bag item ID.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateDataBagItem) Reset() {
	*x = CreateDataBagItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_databags_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDataBagItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDataBagItem) ProtoMessage() {}

func (x *CreateDataBagItem) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_databags_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDataBagItem.ProtoReflect.Descriptor instead.
func (*CreateDataBagItem) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_databags_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDataBagItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDataBagItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateDataBagItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Data bag item ID.
	ItemId string `protobuf:"bytes,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *UpdateDataBagItem) Reset() {
	*x = UpdateDataBagItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_infra_proxy_response_databags_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDataBagItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDataBagItem) ProtoMessage() {}

func (x *UpdateDataBagItem) ProtoReflect() protoreflect.Message {
	mi := &file_external_infra_proxy_response_databags_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDataBagItem.ProtoReflect.Descriptor instead.
func (*UpdateDataBagItem) Descriptor() ([]byte, []int) {
	return file_external_infra_proxy_response_databags_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDataBagItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDataBagItem) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

var File_external_infra_proxy_response_databags_proto protoreflect.FileDescriptor

var file_external_infra_proxy_response_databags_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61,
	0x67, 0x73, 0x12, 0x54, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x61, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x61, 0x67, 0x73, 0x22, 0x25, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x41, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x23, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x42, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x40, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61,
	0x67, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_infra_proxy_response_databags_proto_rawDescOnce sync.Once
	file_external_infra_proxy_response_databags_proto_rawDescData = file_external_infra_proxy_response_databags_proto_rawDesc
)

func file_external_infra_proxy_response_databags_proto_rawDescGZIP() []byte {
	file_external_infra_proxy_response_databags_proto_rawDescOnce.Do(func() {
		file_external_infra_proxy_response_databags_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_infra_proxy_response_databags_proto_rawDescData)
	})
	return file_external_infra_proxy_response_databags_proto_rawDescData
}

var file_external_infra_proxy_response_databags_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_external_infra_proxy_response_databags_proto_goTypes = []interface{}{
	(*DataBags)(nil),          // 0: chef.automate.api.infra_proxy.response.DataBags
	(*DataBagListItem)(nil),   // 1: chef.automate.api.infra_proxy.response.DataBagListItem
	(*DataBag)(nil),           // 2: chef.automate.api.infra_proxy.response.DataBag
	(*CreateDataBag)(nil),     // 3: chef.automate.api.infra_proxy.response.CreateDataBag
	(*CreateDataBagItem)(nil), // 4: chef.automate.api.infra_proxy.response.CreateDataBagItem
	(*UpdateDataBagItem)(nil), // 5: chef.automate.api.infra_proxy.response.UpdateDataBagItem
}
var file_external_infra_proxy_response_databags_proto_depIdxs = []int32{
	1, // 0: chef.automate.api.infra_proxy.response.DataBags.data_bags:type_name -> chef.automate.api.infra_proxy.response.DataBagListItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_external_infra_proxy_response_databags_proto_init() }
func file_external_infra_proxy_response_databags_proto_init() {
	if File_external_infra_proxy_response_databags_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_infra_proxy_response_databags_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBags); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_infra_proxy_response_databags_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBagListItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_infra_proxy_response_databags_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_infra_proxy_response_databags_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDataBag); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_infra_proxy_response_databags_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDataBagItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_external_infra_proxy_response_databags_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDataBagItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_infra_proxy_response_databags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_infra_proxy_response_databags_proto_goTypes,
		DependencyIndexes: file_external_infra_proxy_response_databags_proto_depIdxs,
		MessageInfos:      file_external_infra_proxy_response_databags_proto_msgTypes,
	}.Build()
	File_external_infra_proxy_response_databags_proto = out.File
	file_external_infra_proxy_response_databags_proto_rawDesc = nil
	file_external_infra_proxy_response_databags_proto_goTypes = nil
	file_external_infra_proxy_response_databags_proto_depIdxs = nil
}
