// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: dictybase/identity/identity.proto

package identity

import (
	jsonapi "github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Definition for various fields
type IdentityAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An unique identifier provided by the third party.
	// Generally it's an email id, however it could be something else specifically
	// provided by an provider.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Name of the provider, for example, orcid, google, facebook etc.
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	// The id of the user to which this identity is connected.
	// This id could be used to fetch a complete user response from the user service
	UserId int64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Timestamp for creation and update
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *IdentityAttributes) Reset() {
	*x = IdentityAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_identity_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityAttributes) ProtoMessage() {}

func (x *IdentityAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_identity_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityAttributes.ProtoReflect.Descriptor instead.
func (*IdentityAttributes) Descriptor() ([]byte, []int) {
	return file_dictybase_identity_identity_proto_rawDescGZIP(), []int{0}
}

func (x *IdentityAttributes) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *IdentityAttributes) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *IdentityAttributes) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IdentityAttributes) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *IdentityAttributes) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type IdentityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Unique id
	Id         int64               `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Attributes *IdentityAttributes `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
	Links      *jsonapi.Links      `protobuf:"bytes,4,opt,name=links,proto3" json:"links,omitempty"`
}

func (x *IdentityData) Reset() {
	*x = IdentityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_identity_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityData) ProtoMessage() {}

func (x *IdentityData) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_identity_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityData.ProtoReflect.Descriptor instead.
func (*IdentityData) Descriptor() ([]byte, []int) {
	return file_dictybase_identity_identity_proto_rawDescGZIP(), []int{1}
}

func (x *IdentityData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IdentityData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IdentityData) GetAttributes() *IdentityAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *IdentityData) GetLinks() *jsonapi.Links {
	if x != nil {
		return x.Links
	}
	return nil
}

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *IdentityData  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Links *jsonapi.Links `protobuf:"bytes,4,opt,name=links,proto3" json:"links,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_identity_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_identity_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_dictybase_identity_identity_proto_rawDescGZIP(), []int{2}
}

func (x *Identity) GetData() *IdentityData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Identity) GetLinks() *jsonapi.Links {
	if x != nil {
		return x.Links
	}
	return nil
}

type IdentityProviderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An unique identifier provided by the third party.
	// Generally it's an email id, however it could be something else specifically
	// provided by an provider.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Name of the provider, for example, orcid, google, facebook etc.
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *IdentityProviderReq) Reset() {
	*x = IdentityProviderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_identity_identity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityProviderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProviderReq) ProtoMessage() {}

func (x *IdentityProviderReq) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_identity_identity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProviderReq.ProtoReflect.Descriptor instead.
func (*IdentityProviderReq) Descriptor() ([]byte, []int) {
	return file_dictybase_identity_identity_proto_rawDescGZIP(), []int{3}
}

func (x *IdentityProviderReq) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *IdentityProviderReq) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type CreateIdentityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *CreateIdentityReq_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateIdentityReq) Reset() {
	*x = CreateIdentityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_identity_identity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIdentityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIdentityReq) ProtoMessage() {}

func (x *CreateIdentityReq) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_identity_identity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIdentityReq.ProtoReflect.Descriptor instead.
func (*CreateIdentityReq) Descriptor() ([]byte, []int) {
	return file_dictybase_identity_identity_proto_rawDescGZIP(), []int{4}
}

func (x *CreateIdentityReq) GetData() *CreateIdentityReq_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type NewIdentityAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An unique identifier provided by the third party.
	// Generally it's an email id, however it could be something else specifically
	// provided by an provider.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Name of the provider, for example, orcid, google, facebook etc.
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	// The id of the user to which this identity is connected.
	// This id could be used to fetch a complete user response from the user service
	UserId int64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *NewIdentityAttributes) Reset() {
	*x = NewIdentityAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_identity_identity_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewIdentityAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewIdentityAttributes) ProtoMessage() {}

func (x *NewIdentityAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_identity_identity_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewIdentityAttributes.ProtoReflect.Descriptor instead.
func (*NewIdentityAttributes) Descriptor() ([]byte, []int) {
	return file_dictybase_identity_identity_proto_rawDescGZIP(), []int{5}
}

func (x *NewIdentityAttributes) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *NewIdentityAttributes) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *NewIdentityAttributes) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateIdentityReq_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource name
	Type       string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Attributes *NewIdentityAttributes `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *CreateIdentityReq_Data) Reset() {
	*x = CreateIdentityReq_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_identity_identity_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIdentityReq_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIdentityReq_Data) ProtoMessage() {}

func (x *CreateIdentityReq_Data) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_identity_identity_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIdentityReq_Data.ProtoReflect.Descriptor instead.
func (*CreateIdentityReq_Data) Descriptor() ([]byte, []int) {
	return file_dictybase_identity_identity_proto_rawDescGZIP(), []int{4, 0}
}

func (x *CreateIdentityReq_Data) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateIdentityReq_Data) GetAttributes() *NewIdentityAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

var File_dictybase_identity_identity_proto protoreflect.FileDescriptor

var file_dictybase_identity_identity_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x64, 0x69, 0x63, 0x74,
	0x79, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xdf, 0x01, 0x0a, 0x12, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xae, 0x01, 0x0a, 0x0c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x46, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x69, 0x63,
	0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x32,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x73,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x05, 0x6c, 0x69, 0x6e,
	0x6b, 0x73, 0x22, 0x74, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x34,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64,
	0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x51, 0x0a, 0x13, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x22, 0xba, 0x01, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x12, 0x3e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x65, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x49, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x4e, 0x65, 0x77, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0x6c, 0x0a, 0x15, 0x4e, 0x65, 0x77, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0xa5, 0x04, 0x0a, 0x0f, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1c,
	0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4f,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x20, 0x2e,
	0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x73,
	0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x68, 0x0a, 0x15, 0x45, 0x78, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x27, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x24, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x25, 0x2e, 0x64, 0x69,
	0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x20, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x12, 0x27, 0x2e, 0x64, 0x69,
	0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6a, 0x73, 0x6f, 0x6e,
	0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x7a, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x79,
	0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x0d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x79, 0x42, 0x61, 0x73, 0x65, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x63, 0x74,
	0x79, 0x62, 0x61, 0x73, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x3b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0xf8, 0x01, 0x01, 0xa2, 0x02,
	0x08, 0x44, 0x49, 0x43, 0x54, 0x59, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_dictybase_identity_identity_proto_rawDescOnce sync.Once
	file_dictybase_identity_identity_proto_rawDescData = file_dictybase_identity_identity_proto_rawDesc
)

func file_dictybase_identity_identity_proto_rawDescGZIP() []byte {
	file_dictybase_identity_identity_proto_rawDescOnce.Do(func() {
		file_dictybase_identity_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_dictybase_identity_identity_proto_rawDescData)
	})
	return file_dictybase_identity_identity_proto_rawDescData
}

var file_dictybase_identity_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_dictybase_identity_identity_proto_goTypes = []interface{}{
	(*IdentityAttributes)(nil),       // 0: dictybase.identity.IdentityAttributes
	(*IdentityData)(nil),             // 1: dictybase.identity.IdentityData
	(*Identity)(nil),                 // 2: dictybase.identity.Identity
	(*IdentityProviderReq)(nil),      // 3: dictybase.identity.IdentityProviderReq
	(*CreateIdentityReq)(nil),        // 4: dictybase.identity.CreateIdentityReq
	(*NewIdentityAttributes)(nil),    // 5: dictybase.identity.NewIdentityAttributes
	(*CreateIdentityReq_Data)(nil),   // 6: dictybase.identity.CreateIdentityReq.Data
	(*timestamppb.Timestamp)(nil),    // 7: google.protobuf.Timestamp
	(*jsonapi.Links)(nil),            // 8: dictybase.api.jsonapi.Links
	(*jsonapi.IdRequest)(nil),        // 9: dictybase.api.jsonapi.IdRequest
	(*jsonapi.HealthzIdRequest)(nil), // 10: dictybase.api.jsonapi.HealthzIdRequest
	(*jsonapi.ExistResponse)(nil),    // 11: dictybase.api.jsonapi.ExistResponse
	(*emptypb.Empty)(nil),            // 12: google.protobuf.Empty
}
var file_dictybase_identity_identity_proto_depIdxs = []int32{
	7,  // 0: dictybase.identity.IdentityAttributes.created_at:type_name -> google.protobuf.Timestamp
	7,  // 1: dictybase.identity.IdentityAttributes.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 2: dictybase.identity.IdentityData.attributes:type_name -> dictybase.identity.IdentityAttributes
	8,  // 3: dictybase.identity.IdentityData.links:type_name -> dictybase.api.jsonapi.Links
	1,  // 4: dictybase.identity.Identity.data:type_name -> dictybase.identity.IdentityData
	8,  // 5: dictybase.identity.Identity.links:type_name -> dictybase.api.jsonapi.Links
	6,  // 6: dictybase.identity.CreateIdentityReq.data:type_name -> dictybase.identity.CreateIdentityReq.Data
	5,  // 7: dictybase.identity.CreateIdentityReq.Data.attributes:type_name -> dictybase.identity.NewIdentityAttributes
	3,  // 8: dictybase.identity.IdentityService.GetIdentityFromProvider:input_type -> dictybase.identity.IdentityProviderReq
	9,  // 9: dictybase.identity.IdentityService.GetIdentity:input_type -> dictybase.api.jsonapi.IdRequest
	3,  // 10: dictybase.identity.IdentityService.ExistProviderIdentity:input_type -> dictybase.identity.IdentityProviderReq
	4,  // 11: dictybase.identity.IdentityService.CreateIdentity:input_type -> dictybase.identity.CreateIdentityReq
	9,  // 12: dictybase.identity.IdentityService.DeleteIdentity:input_type -> dictybase.api.jsonapi.IdRequest
	10, // 13: dictybase.identity.IdentityService.Healthz:input_type -> dictybase.api.jsonapi.HealthzIdRequest
	2,  // 14: dictybase.identity.IdentityService.GetIdentityFromProvider:output_type -> dictybase.identity.Identity
	2,  // 15: dictybase.identity.IdentityService.GetIdentity:output_type -> dictybase.identity.Identity
	11, // 16: dictybase.identity.IdentityService.ExistProviderIdentity:output_type -> dictybase.api.jsonapi.ExistResponse
	2,  // 17: dictybase.identity.IdentityService.CreateIdentity:output_type -> dictybase.identity.Identity
	12, // 18: dictybase.identity.IdentityService.DeleteIdentity:output_type -> google.protobuf.Empty
	12, // 19: dictybase.identity.IdentityService.Healthz:output_type -> google.protobuf.Empty
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_dictybase_identity_identity_proto_init() }
func file_dictybase_identity_identity_proto_init() {
	if File_dictybase_identity_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dictybase_identity_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityAttributes); i {
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
		file_dictybase_identity_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityData); i {
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
		file_dictybase_identity_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_dictybase_identity_identity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityProviderReq); i {
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
		file_dictybase_identity_identity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIdentityReq); i {
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
		file_dictybase_identity_identity_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewIdentityAttributes); i {
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
		file_dictybase_identity_identity_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIdentityReq_Data); i {
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
			RawDescriptor: file_dictybase_identity_identity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dictybase_identity_identity_proto_goTypes,
		DependencyIndexes: file_dictybase_identity_identity_proto_depIdxs,
		MessageInfos:      file_dictybase_identity_identity_proto_msgTypes,
	}.Build()
	File_dictybase_identity_identity_proto = out.File
	file_dictybase_identity_identity_proto_rawDesc = nil
	file_dictybase_identity_identity_proto_goTypes = nil
	file_dictybase_identity_identity_proto_depIdxs = nil
}
