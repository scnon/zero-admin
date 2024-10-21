// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: apps/business/rpc/business.proto

package business

import (
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

type BusinessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone   string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	TgId    string `protobuf:"bytes,3,opt,name=tgId,proto3" json:"tgId,omitempty"`
	AdminId int64  `protobuf:"varint,4,opt,name=adminId,proto3" json:"adminId,omitempty"`
}

func (x *BusinessInfo) Reset() {
	*x = BusinessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_business_rpc_business_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessInfo) ProtoMessage() {}

func (x *BusinessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_apps_business_rpc_business_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessInfo.ProtoReflect.Descriptor instead.
func (*BusinessInfo) Descriptor() ([]byte, []int) {
	return file_apps_business_rpc_business_proto_rawDescGZIP(), []int{0}
}

func (x *BusinessInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BusinessInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *BusinessInfo) GetTgId() string {
	if x != nil {
		return x.TgId
	}
	return ""
}

func (x *BusinessInfo) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

type AddBusinessReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone   string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	TgId    string `protobuf:"bytes,2,opt,name=tgId,proto3" json:"tgId,omitempty"`
	AdminId int64  `protobuf:"varint,3,opt,name=adminId,proto3" json:"adminId,omitempty"`
}

func (x *AddBusinessReq) Reset() {
	*x = AddBusinessReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_business_rpc_business_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBusinessReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBusinessReq) ProtoMessage() {}

func (x *AddBusinessReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_business_rpc_business_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBusinessReq.ProtoReflect.Descriptor instead.
func (*AddBusinessReq) Descriptor() ([]byte, []int) {
	return file_apps_business_rpc_business_proto_rawDescGZIP(), []int{1}
}

func (x *AddBusinessReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AddBusinessReq) GetTgId() string {
	if x != nil {
		return x.TgId
	}
	return ""
}

func (x *AddBusinessReq) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

type DeleteBusinessReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteBusinessReq) Reset() {
	*x = DeleteBusinessReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_business_rpc_business_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBusinessReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBusinessReq) ProtoMessage() {}

func (x *DeleteBusinessReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_business_rpc_business_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBusinessReq.ProtoReflect.Descriptor instead.
func (*DeleteBusinessReq) Descriptor() ([]byte, []int) {
	return file_apps_business_rpc_business_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBusinessReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteBusinessResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBusinessResp) Reset() {
	*x = DeleteBusinessResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_business_rpc_business_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBusinessResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBusinessResp) ProtoMessage() {}

func (x *DeleteBusinessResp) ProtoReflect() protoreflect.Message {
	mi := &file_apps_business_rpc_business_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBusinessResp.ProtoReflect.Descriptor instead.
func (*DeleteBusinessResp) Descriptor() ([]byte, []int) {
	return file_apps_business_rpc_business_proto_rawDescGZIP(), []int{3}
}

type GetBusinessReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids      []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	AdminIds []int64 `protobuf:"varint,2,rep,packed,name=adminIds,proto3" json:"adminIds,omitempty"`
}

func (x *GetBusinessReq) Reset() {
	*x = GetBusinessReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_business_rpc_business_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusinessReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessReq) ProtoMessage() {}

func (x *GetBusinessReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_business_rpc_business_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessReq.ProtoReflect.Descriptor instead.
func (*GetBusinessReq) Descriptor() ([]byte, []int) {
	return file_apps_business_rpc_business_proto_rawDescGZIP(), []int{4}
}

func (x *GetBusinessReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetBusinessReq) GetAdminIds() []int64 {
	if x != nil {
		return x.AdminIds
	}
	return nil
}

type GetBusinessResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Business []*BusinessInfo `protobuf:"bytes,1,rep,name=business,proto3" json:"business,omitempty"`
}

func (x *GetBusinessResp) Reset() {
	*x = GetBusinessResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_business_rpc_business_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusinessResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessResp) ProtoMessage() {}

func (x *GetBusinessResp) ProtoReflect() protoreflect.Message {
	mi := &file_apps_business_rpc_business_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessResp.ProtoReflect.Descriptor instead.
func (*GetBusinessResp) Descriptor() ([]byte, []int) {
	return file_apps_business_rpc_business_proto_rawDescGZIP(), []int{5}
}

func (x *GetBusinessResp) GetBusiness() []*BusinessInfo {
	if x != nil {
		return x.Business
	}
	return nil
}

var File_apps_business_rpc_business_proto protoreflect.FileDescriptor

var file_apps_business_rpc_business_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x62, 0x0a, 0x0c,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64,
	0x22, 0x54, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x67, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x14, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x3e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x49, 0x64, 0x73, 0x22, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x32, 0xa6, 0x02, 0x0a, 0x08, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x18, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x16, 0x2e, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x4d,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x12, 0x1b, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x18, 0x2e, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_business_rpc_business_proto_rawDescOnce sync.Once
	file_apps_business_rpc_business_proto_rawDescData = file_apps_business_rpc_business_proto_rawDesc
)

func file_apps_business_rpc_business_proto_rawDescGZIP() []byte {
	file_apps_business_rpc_business_proto_rawDescOnce.Do(func() {
		file_apps_business_rpc_business_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_business_rpc_business_proto_rawDescData)
	})
	return file_apps_business_rpc_business_proto_rawDescData
}

var file_apps_business_rpc_business_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_business_rpc_business_proto_goTypes = []interface{}{
	(*BusinessInfo)(nil),       // 0: business.BusinessInfo
	(*AddBusinessReq)(nil),     // 1: business.AddBusinessReq
	(*DeleteBusinessReq)(nil),  // 2: business.DeleteBusinessReq
	(*DeleteBusinessResp)(nil), // 3: business.DeleteBusinessResp
	(*GetBusinessReq)(nil),     // 4: business.GetBusinessReq
	(*GetBusinessResp)(nil),    // 5: business.GetBusinessResp
}
var file_apps_business_rpc_business_proto_depIdxs = []int32{
	0, // 0: business.GetBusinessResp.business:type_name -> business.BusinessInfo
	1, // 1: business.Business.AddBusiness:input_type -> business.AddBusinessReq
	0, // 2: business.Business.UpdateBusiness:input_type -> business.BusinessInfo
	2, // 3: business.Business.DeleteBusiness:input_type -> business.DeleteBusinessReq
	4, // 4: business.Business.GetBusiness:input_type -> business.GetBusinessReq
	0, // 5: business.Business.AddBusiness:output_type -> business.BusinessInfo
	0, // 6: business.Business.UpdateBusiness:output_type -> business.BusinessInfo
	3, // 7: business.Business.DeleteBusiness:output_type -> business.DeleteBusinessResp
	5, // 8: business.Business.GetBusiness:output_type -> business.GetBusinessResp
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apps_business_rpc_business_proto_init() }
func file_apps_business_rpc_business_proto_init() {
	if File_apps_business_rpc_business_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_business_rpc_business_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessInfo); i {
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
		file_apps_business_rpc_business_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBusinessReq); i {
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
		file_apps_business_rpc_business_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBusinessReq); i {
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
		file_apps_business_rpc_business_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBusinessResp); i {
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
		file_apps_business_rpc_business_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBusinessReq); i {
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
		file_apps_business_rpc_business_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBusinessResp); i {
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
			RawDescriptor: file_apps_business_rpc_business_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_business_rpc_business_proto_goTypes,
		DependencyIndexes: file_apps_business_rpc_business_proto_depIdxs,
		MessageInfos:      file_apps_business_rpc_business_proto_msgTypes,
	}.Build()
	File_apps_business_rpc_business_proto = out.File
	file_apps_business_rpc_business_proto_rawDesc = nil
	file_apps_business_rpc_business_proto_goTypes = nil
	file_apps_business_rpc_business_proto_depIdxs = nil
}
