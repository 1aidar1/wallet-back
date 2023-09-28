// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/proto/wallet_storage_ds/wallet_storage_ds.proto

package wallet_storage_ds

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type StatusResponse_Status int32

const (
	StatusResponse_STATUS_SUCCESS StatusResponse_Status = 0
	StatusResponse_STATUS_ERROR   StatusResponse_Status = 1
)

// Enum value maps for StatusResponse_Status.
var (
	StatusResponse_Status_name = map[int32]string{
		0: "STATUS_SUCCESS",
		1: "STATUS_ERROR",
	}
	StatusResponse_Status_value = map[string]int32{
		"STATUS_SUCCESS": 0,
		"STATUS_ERROR":   1,
	}
)

func (x StatusResponse_Status) Enum() *StatusResponse_Status {
	p := new(StatusResponse_Status)
	*p = x
	return p
}

func (x StatusResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_enumTypes[0].Descriptor()
}

func (StatusResponse_Status) Type() protoreflect.EnumType {
	return &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_enumTypes[0]
}

func (x StatusResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusResponse_Status.Descriptor instead.
func (StatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescGZIP(), []int{1, 0}
}

type EntityId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EntityId) Reset() {
	*x = EntityId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityId) ProtoMessage() {}

func (x *EntityId) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityId.ProtoReflect.Descriptor instead.
func (*EntityId) Descriptor() ([]byte, []int) {
	return file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescGZIP(), []int{0}
}

func (x *EntityId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    StatusResponse_Status `protobuf:"varint,2,opt,name=status,proto3,enum=wallet_storage_ds.StatusResponse_Status" json:"status,omitempty"`
	ErrorCode string                `protobuf:"bytes,3,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescGZIP(), []int{1}
}

func (x *StatusResponse) GetStatus() StatusResponse_Status {
	if x != nil {
		return x.Status
	}
	return StatusResponse_STATUS_SUCCESS
}

func (x *StatusResponse) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

type ConsumerItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items         []*ConsumerItems_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	RequestStatus *StatusResponse       `protobuf:"bytes,2,opt,name=requestStatus,proto3" json:"requestStatus,omitempty"`
}

func (x *ConsumerItems) Reset() {
	*x = ConsumerItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerItems) ProtoMessage() {}

func (x *ConsumerItems) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerItems.ProtoReflect.Descriptor instead.
func (*ConsumerItems) Descriptor() ([]byte, []int) {
	return file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescGZIP(), []int{2}
}

func (x *ConsumerItems) GetItems() []*ConsumerItems_Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ConsumerItems) GetRequestStatus() *StatusResponse {
	if x != nil {
		return x.RequestStatus
	}
	return nil
}

type ConsumerCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Slug             string   `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Secret           string   `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	WhiteListMethods []string `protobuf:"bytes,5,rep,name=whiteListMethods,proto3" json:"whiteListMethods,omitempty"`
}

func (x *ConsumerCreateRequest) Reset() {
	*x = ConsumerCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerCreateRequest) ProtoMessage() {}

func (x *ConsumerCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerCreateRequest.ProtoReflect.Descriptor instead.
func (*ConsumerCreateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescGZIP(), []int{3}
}

func (x *ConsumerCreateRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ConsumerCreateRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *ConsumerCreateRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *ConsumerCreateRequest) GetWhiteListMethods() []string {
	if x != nil {
		return x.WhiteListMethods
	}
	return nil
}

type ConsumerCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RequestStatus *StatusResponse `protobuf:"bytes,2,opt,name=requestStatus,proto3" json:"requestStatus,omitempty"`
}

func (x *ConsumerCreateResponse) Reset() {
	*x = ConsumerCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerCreateResponse) ProtoMessage() {}

func (x *ConsumerCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerCreateResponse.ProtoReflect.Descriptor instead.
func (*ConsumerCreateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescGZIP(), []int{4}
}

func (x *ConsumerCreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConsumerCreateResponse) GetRequestStatus() *StatusResponse {
	if x != nil {
		return x.RequestStatus
	}
	return nil
}

type ConsumerUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // which consumer update to
	Code             string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Slug             string   `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Secret           string   `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	WhiteListMethods []string `protobuf:"bytes,5,rep,name=whiteListMethods,proto3" json:"whiteListMethods,omitempty"`
}

func (x *ConsumerUpdateRequest) Reset() {
	*x = ConsumerUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerUpdateRequest) ProtoMessage() {}

func (x *ConsumerUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerUpdateRequest.ProtoReflect.Descriptor instead.
func (*ConsumerUpdateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescGZIP(), []int{5}
}

func (x *ConsumerUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConsumerUpdateRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ConsumerUpdateRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *ConsumerUpdateRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *ConsumerUpdateRequest) GetWhiteListMethods() []string {
	if x != nil {
		return x.WhiteListMethods
	}
	return nil
}

type MethodListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Methods []string `protobuf:"bytes,1,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *MethodListResponse) Reset() {
	*x = MethodListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodListResponse) ProtoMessage() {}

func (x *MethodListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodListResponse.ProtoReflect.Descriptor instead.
func (*MethodListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescGZIP(), []int{6}
}

func (x *MethodListResponse) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

type ConsumerItems_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string               `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Slug             string               `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Secret           string               `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	WhiteListMethods []string             `protobuf:"bytes,5,rep,name=whiteListMethods,proto3" json:"whiteListMethods,omitempty"`
	CreatedAt        *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ConsumerItems_Item) Reset() {
	*x = ConsumerItems_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerItems_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerItems_Item) ProtoMessage() {}

func (x *ConsumerItems_Item) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerItems_Item.ProtoReflect.Descriptor instead.
func (*ConsumerItems_Item) Descriptor() ([]byte, []int) {
	return file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ConsumerItems_Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ConsumerItems_Item) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ConsumerItems_Item) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *ConsumerItems_Item) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *ConsumerItems_Item) GetWhiteListMethods() []string {
	if x != nil {
		return x.WhiteListMethods
	}
	return nil
}

func (x *ConsumerItems_Item) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto protoreflect.FileDescriptor

var file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDesc = []byte{
	0x0a, 0x33, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2f, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xa0, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x22, 0xd5, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x3b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x47, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0d,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0xbd, 0x01,
	0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x77, 0x68, 0x69, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x10, 0x77, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x83, 0x01,
	0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x77, 0x68, 0x69, 0x74, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x77, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x22, 0x71, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x47, 0x0a,
	0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x2a, 0x0a, 0x10, 0x77, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x77, 0x68, 0x69, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x22, 0x2e, 0x0a, 0x12,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x32, 0x90, 0x04, 0x0a,
	0x10, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x72,
	0x6d, 0x12, 0x48, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x65, 0x0a, 0x0e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x61, 0x64, 0x12, 0x1b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x1a,
	0x20, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x64, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x12, 0x5d, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1b, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x1a,
	0x21, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x5f, 0x64, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescOnce sync.Once
	file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescData = file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDesc
)

func file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescGZIP() []byte {
	file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescOnce.Do(func() {
		file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescData)
	})
	return file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDescData
}

var file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_goTypes = []interface{}{
	(StatusResponse_Status)(0),     // 0: wallet_storage_ds.StatusResponse.Status
	(*EntityId)(nil),               // 1: wallet_storage_ds.EntityId
	(*StatusResponse)(nil),         // 2: wallet_storage_ds.StatusResponse
	(*ConsumerItems)(nil),          // 3: wallet_storage_ds.ConsumerItems
	(*ConsumerCreateRequest)(nil),  // 4: wallet_storage_ds.ConsumerCreateRequest
	(*ConsumerCreateResponse)(nil), // 5: wallet_storage_ds.ConsumerCreateResponse
	(*ConsumerUpdateRequest)(nil),  // 6: wallet_storage_ds.ConsumerUpdateRequest
	(*MethodListResponse)(nil),     // 7: wallet_storage_ds.MethodListResponse
	(*ConsumerItems_Item)(nil),     // 8: wallet_storage_ds.ConsumerItems.Item
	(*timestamp.Timestamp)(nil),    // 9: google.protobuf.Timestamp
	(*empty.Empty)(nil),            // 10: google.protobuf.Empty
}
var file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_depIdxs = []int32{
	0,  // 0: wallet_storage_ds.StatusResponse.status:type_name -> wallet_storage_ds.StatusResponse.Status
	8,  // 1: wallet_storage_ds.ConsumerItems.items:type_name -> wallet_storage_ds.ConsumerItems.Item
	2,  // 2: wallet_storage_ds.ConsumerItems.requestStatus:type_name -> wallet_storage_ds.StatusResponse
	2,  // 3: wallet_storage_ds.ConsumerCreateResponse.requestStatus:type_name -> wallet_storage_ds.StatusResponse
	9,  // 4: wallet_storage_ds.ConsumerItems.Item.created_at:type_name -> google.protobuf.Timestamp
	10, // 5: wallet_storage_ds.WalletStorageArm.ConsumerList:input_type -> google.protobuf.Empty
	4,  // 6: wallet_storage_ds.WalletStorageArm.ConsumerCreate:input_type -> wallet_storage_ds.ConsumerCreateRequest
	1,  // 7: wallet_storage_ds.WalletStorageArm.ConsumerRead:input_type -> wallet_storage_ds.EntityId
	6,  // 8: wallet_storage_ds.WalletStorageArm.ConsumerUpdate:input_type -> wallet_storage_ds.ConsumerUpdateRequest
	1,  // 9: wallet_storage_ds.WalletStorageArm.ConsumerDelete:input_type -> wallet_storage_ds.EntityId
	10, // 10: wallet_storage_ds.WalletStorageArm.MethodList:input_type -> google.protobuf.Empty
	3,  // 11: wallet_storage_ds.WalletStorageArm.ConsumerList:output_type -> wallet_storage_ds.ConsumerItems
	5,  // 12: wallet_storage_ds.WalletStorageArm.ConsumerCreate:output_type -> wallet_storage_ds.ConsumerCreateResponse
	3,  // 13: wallet_storage_ds.WalletStorageArm.ConsumerRead:output_type -> wallet_storage_ds.ConsumerItems
	2,  // 14: wallet_storage_ds.WalletStorageArm.ConsumerUpdate:output_type -> wallet_storage_ds.StatusResponse
	2,  // 15: wallet_storage_ds.WalletStorageArm.ConsumerDelete:output_type -> wallet_storage_ds.StatusResponse
	7,  // 16: wallet_storage_ds.WalletStorageArm.MethodList:output_type -> wallet_storage_ds.MethodListResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_init() }
func file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_init() {
	if File_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityId); i {
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
		file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerItems); i {
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
		file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerCreateRequest); i {
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
		file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerCreateResponse); i {
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
		file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerUpdateRequest); i {
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
		file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodListResponse); i {
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
		file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerItems_Item); i {
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
			RawDescriptor: file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_goTypes,
		DependencyIndexes: file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_depIdxs,
		EnumInfos:         file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_enumTypes,
		MessageInfos:      file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_msgTypes,
	}.Build()
	File_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto = out.File
	file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_rawDesc = nil
	file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_goTypes = nil
	file_pkg_proto_wallet_storage_ds_wallet_storage_ds_proto_depIdxs = nil
}
