// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content.proto

package content // import "github.com/dictyBase/go-genproto/dictybaseapis/content"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import jsonapi "github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
import _ "github.com/golang/protobuf/ptypes/any"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/mwitkow/go-proto-validators"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import field_mask "google.golang.org/genproto/protobuf/field_mask"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Content struct {
	Data                 *ContentData   `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Links                *jsonapi.Links `protobuf:"bytes,2,opt,name=links" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_cb789752e0a7ca97, []int{0}
}
func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (dst *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(dst, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetData() *ContentData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Content) GetLinks() *jsonapi.Links {
	if m != nil {
		return m.Links
	}
	return nil
}

type ContentData struct {
	// The resource name
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Unique id
	Id                   int64              `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Attributes           *ContentAttributes `protobuf:"bytes,3,opt,name=attributes" json:"attributes,omitempty"`
	Links                *jsonapi.Links     `protobuf:"bytes,4,opt,name=links" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ContentData) Reset()         { *m = ContentData{} }
func (m *ContentData) String() string { return proto.CompactTextString(m) }
func (*ContentData) ProtoMessage()    {}
func (*ContentData) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_cb789752e0a7ca97, []int{1}
}
func (m *ContentData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentData.Unmarshal(m, b)
}
func (m *ContentData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentData.Marshal(b, m, deterministic)
}
func (dst *ContentData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentData.Merge(dst, src)
}
func (m *ContentData) XXX_Size() int {
	return xxx_messageInfo_ContentData.Size(m)
}
func (m *ContentData) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentData.DiscardUnknown(m)
}

var xxx_messageInfo_ContentData proto.InternalMessageInfo

func (m *ContentData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ContentData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ContentData) GetAttributes() *ContentAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *ContentData) GetLinks() *jsonapi.Links {
	if m != nil {
		return m.Links
	}
	return nil
}

// Definition for various content fields
type ContentAttributes struct {
	// page name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// page slug. Look here https://en.wikipedia.org/wiki/Semantic_URL#Slug to know about slug
	Slug string `protobuf:"bytes,2,opt,name=slug" json:"slug,omitempty"`
	// id of the user who created the content
	CreatedBy int64 `protobuf:"varint,3,opt,name=created_by,json=createdBy" json:"created_by,omitempty"`
	// id of the user who updated the content
	UpdatedBy int64 `protobuf:"varint,4,opt,name=updated_by,json=updatedBy" json:"updated_by,omitempty"`
	// Timestamp for creation and update
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	// serialized page content(for example serialized draft js object)
	Content string `protobuf:"bytes,7,opt,name=content" json:"content,omitempty"`
	// namespace for the page
	Namespace            string   `protobuf:"bytes,8,opt,name=namespace" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentAttributes) Reset()         { *m = ContentAttributes{} }
func (m *ContentAttributes) String() string { return proto.CompactTextString(m) }
func (*ContentAttributes) ProtoMessage()    {}
func (*ContentAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_cb789752e0a7ca97, []int{2}
}
func (m *ContentAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentAttributes.Unmarshal(m, b)
}
func (m *ContentAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentAttributes.Marshal(b, m, deterministic)
}
func (dst *ContentAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentAttributes.Merge(dst, src)
}
func (m *ContentAttributes) XXX_Size() int {
	return xxx_messageInfo_ContentAttributes.Size(m)
}
func (m *ContentAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ContentAttributes proto.InternalMessageInfo

func (m *ContentAttributes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContentAttributes) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *ContentAttributes) GetCreatedBy() int64 {
	if m != nil {
		return m.CreatedBy
	}
	return 0
}

func (m *ContentAttributes) GetUpdatedBy() int64 {
	if m != nil {
		return m.UpdatedBy
	}
	return 0
}

func (m *ContentAttributes) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ContentAttributes) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *ContentAttributes) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ContentAttributes) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ContentRequest struct {
	// Url slug
	// Look here https://en.wikipedia.org/wiki/Semantic_URL#Slug to know about slug
	// The slug name should be unique
	Slug                 string   `protobuf:"bytes,1,opt,name=slug" json:"slug,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentRequest) Reset()         { *m = ContentRequest{} }
func (m *ContentRequest) String() string { return proto.CompactTextString(m) }
func (*ContentRequest) ProtoMessage()    {}
func (*ContentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_cb789752e0a7ca97, []int{3}
}
func (m *ContentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentRequest.Unmarshal(m, b)
}
func (m *ContentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentRequest.Marshal(b, m, deterministic)
}
func (dst *ContentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentRequest.Merge(dst, src)
}
func (m *ContentRequest) XXX_Size() int {
	return xxx_messageInfo_ContentRequest.Size(m)
}
func (m *ContentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContentRequest proto.InternalMessageInfo

func (m *ContentRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

type ContentIdRequest struct {
	// Unique id to identify content
	Id                   int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentIdRequest) Reset()         { *m = ContentIdRequest{} }
func (m *ContentIdRequest) String() string { return proto.CompactTextString(m) }
func (*ContentIdRequest) ProtoMessage()    {}
func (*ContentIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_cb789752e0a7ca97, []int{4}
}
func (m *ContentIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentIdRequest.Unmarshal(m, b)
}
func (m *ContentIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentIdRequest.Marshal(b, m, deterministic)
}
func (dst *ContentIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentIdRequest.Merge(dst, src)
}
func (m *ContentIdRequest) XXX_Size() int {
	return xxx_messageInfo_ContentIdRequest.Size(m)
}
func (m *ContentIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContentIdRequest proto.InternalMessageInfo

func (m *ContentIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Definition for fields that are needed for storing the content
type NewContentAttributes struct {
	// page name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// user id who is creating this content
	// The existence of user will be verified(not implemented yet)
	// using the `user` microservice backend.
	CreatedBy int64 `protobuf:"varint,2,opt,name=created_by,json=createdBy" json:"created_by,omitempty"`
	// page content, expected to be serialized `JSON` string.
	Content string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	// namespace for the page, it is prepended to the
	// name to generate an unique slug.
	Namespace            string   `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewContentAttributes) Reset()         { *m = NewContentAttributes{} }
func (m *NewContentAttributes) String() string { return proto.CompactTextString(m) }
func (*NewContentAttributes) ProtoMessage()    {}
func (*NewContentAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_cb789752e0a7ca97, []int{5}
}
func (m *NewContentAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewContentAttributes.Unmarshal(m, b)
}
func (m *NewContentAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewContentAttributes.Marshal(b, m, deterministic)
}
func (dst *NewContentAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewContentAttributes.Merge(dst, src)
}
func (m *NewContentAttributes) XXX_Size() int {
	return xxx_messageInfo_NewContentAttributes.Size(m)
}
func (m *NewContentAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_NewContentAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_NewContentAttributes proto.InternalMessageInfo

func (m *NewContentAttributes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NewContentAttributes) GetCreatedBy() int64 {
	if m != nil {
		return m.CreatedBy
	}
	return 0
}

func (m *NewContentAttributes) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *NewContentAttributes) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Definition for storing new content
type StoreContentRequest struct {
	Data                 *StoreContentRequest_Data `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *StoreContentRequest) Reset()         { *m = StoreContentRequest{} }
func (m *StoreContentRequest) String() string { return proto.CompactTextString(m) }
func (*StoreContentRequest) ProtoMessage()    {}
func (*StoreContentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_cb789752e0a7ca97, []int{6}
}
func (m *StoreContentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreContentRequest.Unmarshal(m, b)
}
func (m *StoreContentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreContentRequest.Marshal(b, m, deterministic)
}
func (dst *StoreContentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreContentRequest.Merge(dst, src)
}
func (m *StoreContentRequest) XXX_Size() int {
	return xxx_messageInfo_StoreContentRequest.Size(m)
}
func (m *StoreContentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreContentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StoreContentRequest proto.InternalMessageInfo

func (m *StoreContentRequest) GetData() *StoreContentRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type StoreContentRequest_Data struct {
	// resource name
	Type                 string                `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Attributes           *NewContentAttributes `protobuf:"bytes,2,opt,name=attributes" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StoreContentRequest_Data) Reset()         { *m = StoreContentRequest_Data{} }
func (m *StoreContentRequest_Data) String() string { return proto.CompactTextString(m) }
func (*StoreContentRequest_Data) ProtoMessage()    {}
func (*StoreContentRequest_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_cb789752e0a7ca97, []int{6, 0}
}
func (m *StoreContentRequest_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StoreContentRequest_Data.Unmarshal(m, b)
}
func (m *StoreContentRequest_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StoreContentRequest_Data.Marshal(b, m, deterministic)
}
func (dst *StoreContentRequest_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoreContentRequest_Data.Merge(dst, src)
}
func (m *StoreContentRequest_Data) XXX_Size() int {
	return xxx_messageInfo_StoreContentRequest_Data.Size(m)
}
func (m *StoreContentRequest_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_StoreContentRequest_Data.DiscardUnknown(m)
}

var xxx_messageInfo_StoreContentRequest_Data proto.InternalMessageInfo

func (m *StoreContentRequest_Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StoreContentRequest_Data) GetAttributes() *NewContentAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Fields that can be updated
// Changing either or both of name and namespace
// attributes alter the slug for the page
type ExistingContentAttributes struct {
	// user id who is updating this content
	// The existence of user will be verified(not implemented yet)
	// using the `user` microservice backend.
	UpdatedBy int64 `protobuf:"varint,1,opt,name=updated_by,json=updatedBy" json:"updated_by,omitempty"`
	// serialized page content(for example serialized draft js object)
	Content              string   `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistingContentAttributes) Reset()         { *m = ExistingContentAttributes{} }
func (m *ExistingContentAttributes) String() string { return proto.CompactTextString(m) }
func (*ExistingContentAttributes) ProtoMessage()    {}
func (*ExistingContentAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_cb789752e0a7ca97, []int{7}
}
func (m *ExistingContentAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistingContentAttributes.Unmarshal(m, b)
}
func (m *ExistingContentAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistingContentAttributes.Marshal(b, m, deterministic)
}
func (dst *ExistingContentAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistingContentAttributes.Merge(dst, src)
}
func (m *ExistingContentAttributes) XXX_Size() int {
	return xxx_messageInfo_ExistingContentAttributes.Size(m)
}
func (m *ExistingContentAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistingContentAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ExistingContentAttributes proto.InternalMessageInfo

func (m *ExistingContentAttributes) GetUpdatedBy() int64 {
	if m != nil {
		return m.UpdatedBy
	}
	return 0
}

func (m *ExistingContentAttributes) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type UpdateContentRequest struct {
	Data *UpdateContentRequest_Data `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Id   int64                      `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	// An optional mask specifying which fields to update.
	// Presence of this field allow partial updates.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateContentRequest) Reset()         { *m = UpdateContentRequest{} }
func (m *UpdateContentRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateContentRequest) ProtoMessage()    {}
func (*UpdateContentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_cb789752e0a7ca97, []int{8}
}
func (m *UpdateContentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateContentRequest.Unmarshal(m, b)
}
func (m *UpdateContentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateContentRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateContentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContentRequest.Merge(dst, src)
}
func (m *UpdateContentRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateContentRequest.Size(m)
}
func (m *UpdateContentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContentRequest proto.InternalMessageInfo

func (m *UpdateContentRequest) GetData() *UpdateContentRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UpdateContentRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateContentRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type UpdateContentRequest_Data struct {
	// resource name
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// unique id
	Id                   int64                      `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Attributes           *ExistingContentAttributes `protobuf:"bytes,3,opt,name=attributes" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *UpdateContentRequest_Data) Reset()         { *m = UpdateContentRequest_Data{} }
func (m *UpdateContentRequest_Data) String() string { return proto.CompactTextString(m) }
func (*UpdateContentRequest_Data) ProtoMessage()    {}
func (*UpdateContentRequest_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_cb789752e0a7ca97, []int{8, 0}
}
func (m *UpdateContentRequest_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateContentRequest_Data.Unmarshal(m, b)
}
func (m *UpdateContentRequest_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateContentRequest_Data.Marshal(b, m, deterministic)
}
func (dst *UpdateContentRequest_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateContentRequest_Data.Merge(dst, src)
}
func (m *UpdateContentRequest_Data) XXX_Size() int {
	return xxx_messageInfo_UpdateContentRequest_Data.Size(m)
}
func (m *UpdateContentRequest_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateContentRequest_Data.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateContentRequest_Data proto.InternalMessageInfo

func (m *UpdateContentRequest_Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UpdateContentRequest_Data) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateContentRequest_Data) GetAttributes() *ExistingContentAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterType((*Content)(nil), "dictybase.content.Content")
	proto.RegisterType((*ContentData)(nil), "dictybase.content.ContentData")
	proto.RegisterType((*ContentAttributes)(nil), "dictybase.content.ContentAttributes")
	proto.RegisterType((*ContentRequest)(nil), "dictybase.content.ContentRequest")
	proto.RegisterType((*ContentIdRequest)(nil), "dictybase.content.ContentIdRequest")
	proto.RegisterType((*NewContentAttributes)(nil), "dictybase.content.NewContentAttributes")
	proto.RegisterType((*StoreContentRequest)(nil), "dictybase.content.StoreContentRequest")
	proto.RegisterType((*StoreContentRequest_Data)(nil), "dictybase.content.StoreContentRequest.Data")
	proto.RegisterType((*ExistingContentAttributes)(nil), "dictybase.content.ExistingContentAttributes")
	proto.RegisterType((*UpdateContentRequest)(nil), "dictybase.content.UpdateContentRequest")
	proto.RegisterType((*UpdateContentRequest_Data)(nil), "dictybase.content.UpdateContentRequest.Data")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContentService service

type ContentServiceClient interface {
	// Gets the content of specified page(slug)
	GetContentBySlug(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*Content, error)
	GetContent(ctx context.Context, in *ContentIdRequest, opts ...grpc.CallOption) (*Content, error)
	// Store the content of a new page(slug)
	StoreContent(ctx context.Context, in *StoreContentRequest, opts ...grpc.CallOption) (*Content, error)
	// Update the content of an existing page
	UpdateContent(ctx context.Context, in *UpdateContentRequest, opts ...grpc.CallOption) (*Content, error)
	// Delete an existing page along with its content
	DeleteContent(ctx context.Context, in *ContentIdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Basic health check that always return success
	Healthz(ctx context.Context, in *jsonapi.HealthzIdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type contentServiceClient struct {
	cc *grpc.ClientConn
}

func NewContentServiceClient(cc *grpc.ClientConn) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) GetContentBySlug(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := grpc.Invoke(ctx, "/dictybase.content.ContentService/GetContentBySlug", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetContent(ctx context.Context, in *ContentIdRequest, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := grpc.Invoke(ctx, "/dictybase.content.ContentService/GetContent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) StoreContent(ctx context.Context, in *StoreContentRequest, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := grpc.Invoke(ctx, "/dictybase.content.ContentService/StoreContent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) UpdateContent(ctx context.Context, in *UpdateContentRequest, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := grpc.Invoke(ctx, "/dictybase.content.ContentService/UpdateContent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteContent(ctx context.Context, in *ContentIdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/dictybase.content.ContentService/DeleteContent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) Healthz(ctx context.Context, in *jsonapi.HealthzIdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/dictybase.content.ContentService/Healthz", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContentService service

type ContentServiceServer interface {
	// Gets the content of specified page(slug)
	GetContentBySlug(context.Context, *ContentRequest) (*Content, error)
	GetContent(context.Context, *ContentIdRequest) (*Content, error)
	// Store the content of a new page(slug)
	StoreContent(context.Context, *StoreContentRequest) (*Content, error)
	// Update the content of an existing page
	UpdateContent(context.Context, *UpdateContentRequest) (*Content, error)
	// Delete an existing page along with its content
	DeleteContent(context.Context, *ContentIdRequest) (*empty.Empty, error)
	// Basic health check that always return success
	Healthz(context.Context, *jsonapi.HealthzIdRequest) (*empty.Empty, error)
}

func RegisterContentServiceServer(s *grpc.Server, srv ContentServiceServer) {
	s.RegisterService(&_ContentService_serviceDesc, srv)
}

func _ContentService_GetContentBySlug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetContentBySlug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.content.ContentService/GetContentBySlug",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetContentBySlug(ctx, req.(*ContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.content.ContentService/GetContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetContent(ctx, req.(*ContentIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_StoreContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).StoreContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.content.ContentService/StoreContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).StoreContent(ctx, req.(*StoreContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_UpdateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).UpdateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.content.ContentService/UpdateContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).UpdateContent(ctx, req.(*UpdateContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.content.ContentService/DeleteContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteContent(ctx, req.(*ContentIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(jsonapi.HealthzIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.content.ContentService/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).Healthz(ctx, req.(*jsonapi.HealthzIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dictybase.content.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContentBySlug",
			Handler:    _ContentService_GetContentBySlug_Handler,
		},
		{
			MethodName: "GetContent",
			Handler:    _ContentService_GetContent_Handler,
		},
		{
			MethodName: "StoreContent",
			Handler:    _ContentService_StoreContent_Handler,
		},
		{
			MethodName: "UpdateContent",
			Handler:    _ContentService_UpdateContent_Handler,
		},
		{
			MethodName: "DeleteContent",
			Handler:    _ContentService_DeleteContent_Handler,
		},
		{
			MethodName: "Healthz",
			Handler:    _ContentService_Healthz_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content.proto",
}

func init() { proto.RegisterFile("content.proto", fileDescriptor_content_cb789752e0a7ca97) }

var fileDescriptor_content_cb789752e0a7ca97 = []byte{
	// 928 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x2e, 0x69, 0xf9, 0x47, 0xe3, 0x58, 0x50, 0x36, 0x4e, 0x2a, 0xab, 0x4e, 0xed, 0x32, 0x4e,
	0x13, 0x24, 0x16, 0x59, 0xa8, 0x40, 0x81, 0x34, 0x40, 0x5b, 0x29, 0x4e, 0x53, 0x03, 0x69, 0x61,
	0xd0, 0x29, 0xd0, 0x5f, 0x04, 0x2b, 0x71, 0x43, 0x6f, 0x2d, 0x71, 0x19, 0x71, 0x15, 0x87, 0x0e,
	0x72, 0xe9, 0xa9, 0xf7, 0x3e, 0x40, 0x81, 0x5e, 0x7a, 0xe8, 0xa1, 0x0f, 0xd0, 0xc7, 0xe8, 0x03,
	0x04, 0x08, 0xfa, 0x10, 0x3d, 0x16, 0xbb, 0x5c, 0xfe, 0x53, 0x8a, 0x2e, 0xe2, 0x72, 0xe7, 0x9b,
	0xf9, 0x66, 0xbe, 0x9d, 0x59, 0x11, 0x36, 0x86, 0xcc, 0xe3, 0xc4, 0xe3, 0xa6, 0x3f, 0x61, 0x9c,
	0xa1, 0x8b, 0x0e, 0x1d, 0xf2, 0x70, 0x80, 0x03, 0x62, 0x2a, 0x43, 0x7b, 0xdb, 0x65, 0xcc, 0x1d,
	0x11, 0x0b, 0xfb, 0xd4, 0xc2, 0x9e, 0xc7, 0x38, 0xe6, 0x94, 0x79, 0x41, 0xe4, 0xd0, 0xde, 0x55,
	0x56, 0xf9, 0x36, 0x98, 0x3e, 0xb1, 0x9e, 0x50, 0x32, 0x72, 0x1e, 0x8f, 0x71, 0x70, 0xaa, 0x10,
	0x3b, 0x45, 0x04, 0xa7, 0x63, 0x12, 0x70, 0x3c, 0xf6, 0x15, 0xe0, 0x9d, 0x22, 0x80, 0x8c, 0x7d,
	0x1e, 0x2a, 0xe3, 0x56, 0xd1, 0x88, 0xbd, 0xd8, 0x74, 0x2d, 0xc9, 0x55, 0xe6, 0xf6, 0x53, 0xc0,
	0x3c, 0xf1, 0xf4, 0x71, 0x38, 0x62, 0xd8, 0x99, 0x0f, 0x9a, 0x90, 0xa7, 0x53, 0x12, 0xa8, 0xaa,
	0xdb, 0x1f, 0xb9, 0x94, 0x9f, 0x4c, 0x07, 0xe6, 0x90, 0x8d, 0xad, 0xf1, 0x19, 0xe5, 0xa7, 0xec,
	0xcc, 0x72, 0x59, 0x47, 0x1a, 0x3b, 0xcf, 0xf0, 0x88, 0x3a, 0x98, 0xb3, 0x49, 0x60, 0x25, 0xcb,
	0xc8, 0xcf, 0x78, 0x0a, 0xab, 0xf7, 0x22, 0x95, 0x50, 0x17, 0x6a, 0x0e, 0xe6, 0xb8, 0xa5, 0xed,
	0x6a, 0x37, 0xd7, 0xbb, 0xef, 0x9a, 0x25, 0x1d, 0x4d, 0x85, 0x3c, 0xc0, 0x1c, 0xdb, 0x12, 0x8b,
	0xba, 0xb0, 0x3c, 0xa2, 0xde, 0x69, 0xd0, 0xd2, 0xa5, 0xd3, 0x76, 0xc6, 0x09, 0xfb, 0xd4, 0x54,
	0xb9, 0x9a, 0x0f, 0x05, 0xc6, 0x8e, 0xa0, 0xc6, 0x9f, 0x1a, 0xac, 0x67, 0x22, 0x21, 0x04, 0x35,
	0x1e, 0xfa, 0x44, 0xf2, 0xd6, 0x6d, 0xb9, 0x46, 0x0d, 0xd0, 0xa9, 0x23, 0x83, 0x2e, 0xd9, 0x3a,
	0x75, 0xd0, 0x01, 0x00, 0xe6, 0x7c, 0x42, 0x07, 0x53, 0x4e, 0x82, 0xd6, 0x92, 0x24, 0xdb, 0x9b,
	0x9d, 0x61, 0x2f, 0xc1, 0xda, 0x19, 0xbf, 0x34, 0xdb, 0xda, 0xe2, 0xd9, 0xfe, 0xa1, 0xc3, 0xc5,
	0x52, 0x54, 0x91, 0xb3, 0x87, 0xc7, 0x49, 0xce, 0x62, 0x2d, 0xf6, 0x82, 0xd1, 0xd4, 0x95, 0x59,
	0xd7, 0x6d, 0xb9, 0x46, 0x57, 0x01, 0x86, 0x13, 0x82, 0x39, 0x71, 0x1e, 0x0f, 0x42, 0x99, 0xf7,
	0x92, 0x5d, 0x57, 0x3b, 0xfd, 0x50, 0x98, 0xa7, 0xbe, 0x13, 0x9b, 0x6b, 0x91, 0x59, 0xed, 0xf4,
	0x43, 0x74, 0x27, 0xf5, 0xc6, 0xbc, 0xb5, 0x2c, 0x93, 0x6e, 0x9b, 0x51, 0x3b, 0x99, 0x71, 0x3b,
	0x99, 0x8f, 0xe2, 0x66, 0x4c, 0x22, 0xf7, 0xb8, 0x70, 0x8d, 0x23, 0x63, 0xde, 0x5a, 0x79, 0xb3,
	0xab, 0x42, 0xf7, 0x38, 0x6a, 0xc1, 0xaa, 0x92, 0xb3, 0xb5, 0x2a, 0x4b, 0x89, 0x5f, 0xd1, 0x36,
	0xd4, 0x45, 0xa5, 0x81, 0x8f, 0x87, 0xa4, 0xb5, 0x26, 0x6d, 0xe9, 0x86, 0xb1, 0x07, 0x0d, 0x25,
	0x94, 0x1d, 0xb5, 0x66, 0xa2, 0x88, 0x96, 0x2a, 0x62, 0x18, 0xd0, 0x54, 0xa8, 0x43, 0x27, 0xc6,
	0x45, 0xa7, 0xad, 0xc5, 0xa7, 0x6d, 0xfc, 0xa5, 0xc1, 0xe6, 0x57, 0xe4, 0xac, 0x2c, 0xfb, 0xcd,
	0xac, 0xec, 0xfd, 0xcd, 0xd7, 0xaf, 0x76, 0x9a, 0xd0, 0xf8, 0x1e, 0x77, 0xce, 0xf7, 0x7b, 0x9d,
	0xf3, 0xfd, 0x0f, 0x3a, 0x77, 0x7e, 0xbc, 0xad, 0x0e, 0xe3, 0x7a, 0x4e, 0x78, 0xd9, 0x48, 0xfd,
	0x95, 0xd7, 0xaf, 0x76, 0xf4, 0xe6, 0x5b, 0xd9, 0x03, 0xd8, 0x4d, 0x6b, 0x5d, 0x92, 0x31, 0x25,
	0xe6, 0x1b, 0x2d, 0xad, 0x79, 0x2f, 0x5b, 0x73, 0x2d, 0x87, 0xc9, 0xd4, 0xfe, 0xb7, 0x06, 0x97,
	0x8e, 0x39, 0x9b, 0x90, 0x82, 0x02, 0x9f, 0xe6, 0x66, 0xea, 0x76, 0x45, 0xc7, 0x56, 0x78, 0x99,
	0xe9, 0x80, 0xb5, 0x87, 0x50, 0x9b, 0x39, 0x24, 0x0f, 0x72, 0x43, 0x11, 0x4d, 0xe0, 0x8d, 0x0a,
	0x8a, 0x2a, 0x29, 0xb3, 0x73, 0x61, 0x38, 0xb0, 0x75, 0xff, 0x39, 0x0d, 0x38, 0xf5, 0xdc, 0xb2,
	0xe6, 0xd7, 0x73, 0x3d, 0xaa, 0xe5, 0x95, 0x4c, 0x7b, 0x35, 0xa3, 0xa4, 0x5e, 0xa9, 0xa4, 0xf1,
	0x9b, 0x0e, 0x9b, 0x5f, 0x4b, 0x7c, 0x41, 0xa4, 0xcf, 0x72, 0x22, 0xed, 0x57, 0x54, 0x50, 0xe5,
	0x96, 0x51, 0xa9, 0x74, 0x5d, 0xdc, 0x85, 0xf5, 0x28, 0x33, 0x79, 0x8b, 0xab, 0xfb, 0xa2, 0xdc,
	0xfe, 0x9f, 0x8b, 0x8b, 0xfe, 0x4b, 0x1c, 0x9c, 0xda, 0xaa, 0x44, 0xb1, 0x6e, 0x3f, 0x9f, 0x23,
	0x79, 0x91, 0xe8, 0x61, 0xc5, 0xbd, 0x54, 0x55, 0xc0, 0x4c, 0x79, 0xb3, 0xe7, 0xd0, 0xfd, 0x65,
	0x39, 0x19, 0xa1, 0x63, 0x32, 0x79, 0x46, 0x87, 0x04, 0xf9, 0xd0, 0x7c, 0x40, 0xb8, 0xda, 0xec,
	0x87, 0xc7, 0xe2, 0x52, 0x79, 0x6f, 0xf6, 0xc5, 0xa7, 0xb4, 0x69, 0xb7, 0x67, 0x43, 0x8c, 0xab,
	0x3f, 0xff, 0xf3, 0xef, 0xaf, 0xfa, 0xdb, 0xe8, 0xb2, 0xa5, 0x2c, 0x81, 0x25, 0x26, 0xd3, 0x7a,
	0x21, 0x7e, 0x5f, 0x22, 0x02, 0x90, 0x32, 0xa2, 0x6b, 0xb3, 0x03, 0x25, 0xf3, 0x3b, 0x97, 0xed,
	0x8a, 0x64, 0x6b, 0xa2, 0x46, 0xca, 0xf6, 0x82, 0x3a, 0x2f, 0xd1, 0x09, 0x5c, 0xc8, 0xb6, 0x3e,
	0x7a, 0x7f, 0xb1, 0xd9, 0x98, 0xcb, 0xb5, 0x29, 0xb9, 0x1a, 0x46, 0x3d, 0xe1, 0xfa, 0x58, 0xbb,
	0x85, 0x18, 0x6c, 0xe4, 0xfa, 0x07, 0xdd, 0x58, 0xb0, 0xc3, 0xe6, 0x72, 0x6d, 0x49, 0xae, 0x4b,
	0xdd, 0x42, 0x5d, 0x82, 0xd0, 0x81, 0x8d, 0x03, 0x32, 0x22, 0x29, 0xe1, 0x42, 0x22, 0x5e, 0x29,
	0xb5, 0xe7, 0x7d, 0xf1, 0x11, 0x11, 0x0b, 0x78, 0xab, 0x28, 0xe0, 0x0f, 0xb0, 0xfa, 0x05, 0xc1,
	0x23, 0x7e, 0x72, 0x9e, 0x2b, 0x28, 0xfb, 0x47, 0xa6, 0xec, 0x6f, 0xe6, 0x68, 0x4a, 0x0e, 0x40,
	0x6b, 0xd6, 0x49, 0xe4, 0xd2, 0x9f, 0xc2, 0x65, 0x36, 0x71, 0xcb, 0x79, 0xf7, 0x2f, 0xa8, 0xc4,
	0x8f, 0x44, 0x84, 0x23, 0xed, 0xbb, 0x4f, 0x32, 0x1f, 0x1e, 0x12, 0xdd, 0x17, 0x1f, 0x2a, 0x2e,
	0xeb, 0xb8, 0xc4, 0x93, 0x2c, 0x56, 0x12, 0x03, 0xfb, 0x34, 0x88, 0x8b, 0xb8, 0xab, 0x9e, 0xff,
	0x69, 0xda, 0xef, 0xfa, 0xda, 0xc1, 0xe1, 0xbd, 0x47, 0xdf, 0xf6, 0x8e, 0x0e, 0x07, 0x2b, 0xd2,
	0xe5, 0xc3, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x92, 0xf8, 0xc4, 0xd4, 0x09, 0x00, 0x00,
}
