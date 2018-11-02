// Code generated by protoc-gen-go. DO NOT EDIT.
// source: request.proto

package jsonapi

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A `GetRequest` defines various url and query parameters that could be passed
// in a HTTP **GET** request to a singular resource. Majority of the request
// parameters are identical or similar to [jsonapi](http://jsonapi.org).
type GetRequest struct {
	// An unique identifier, for example:
	// "/users/34"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// include query parameter to retrieve any particular or particular
	// combination of relationships. Multiple include values are delimited by
	// comma(,).
	//
	// For example,
	//      /{resource_name}/13?include=baz
	//      /{resource_name}/13?include=baz,bot
	Include string `protobuf:"bytes,2,opt,name=include,proto3" json:"include,omitempty"`
	// fields query parameter to retrieve any particular or any particular
	// combination of attributes. Multiple fields values are delimited by comma(,).
	//
	// For example
	//      /{resource_name}/29?fields=foo
	//      /{resource_name}/?fields=foo,bar
	Fields               string   `protobuf:"bytes,3,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetRequest) GetInclude() string {
	if m != nil {
		return m.Include
	}
	return ""
}

func (m *GetRequest) GetFields() string {
	if m != nil {
		return m.Fields
	}
	return ""
}

// A `GetEmailRequest` is identical to GetRequest except `email` id used as unique identifier.
type GetEmailRequest struct {
	// Email id
	// "/users/newman@seinfeld.org"
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// include query parameter to retrieve any particular or particular
	// combination of relationships. Multiple include values are delimited by
	// comma(,).
	//
	// For example,
	//          /{resource_name}/13?include=baz
	//          /{resource_name}/13?include=baz,bot
	Include string `protobuf:"bytes,2,opt,name=include,proto3" json:"include,omitempty"`
	// fields query parameter to retrieve any particular or any particular
	// combination of attributes. Multiple fields values are delimited by comma(,).
	//
	// For example,
	//          /{resource_name}/29?fields=foo
	//          /{resource_name}/?fields=foo,bar
	Fields               string   `protobuf:"bytes,3,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEmailRequest) Reset()         { *m = GetEmailRequest{} }
func (m *GetEmailRequest) String() string { return proto.CompactTextString(m) }
func (*GetEmailRequest) ProtoMessage()    {}
func (*GetEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{1}
}

func (m *GetEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEmailRequest.Unmarshal(m, b)
}
func (m *GetEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEmailRequest.Marshal(b, m, deterministic)
}
func (m *GetEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEmailRequest.Merge(m, src)
}
func (m *GetEmailRequest) XXX_Size() int {
	return xxx_messageInfo_GetEmailRequest.Size(m)
}
func (m *GetEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEmailRequest proto.InternalMessageInfo

func (m *GetEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *GetEmailRequest) GetInclude() string {
	if m != nil {
		return m.Include
	}
	return ""
}

func (m *GetEmailRequest) GetFields() string {
	if m != nil {
		return m.Fields
	}
	return ""
}

// A `GetRequestWithFields` is a subset of GetRequest which only allow the fields parameter.
type GetRequestWithFields struct {
	// An unique identifier, for example:
	// "/users/34"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// fields query parameter to retrieve any particular or any particular
	// combination of attributes. Multiple fields values are delimited by comma(,).
	//
	// For example
	// /{resource_name}/29?fields=foo
	// /{resource_name}/?fields=foo,bar
	Fields               string   `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequestWithFields) Reset()         { *m = GetRequestWithFields{} }
func (m *GetRequestWithFields) String() string { return proto.CompactTextString(m) }
func (*GetRequestWithFields) ProtoMessage()    {}
func (*GetRequestWithFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{2}
}

func (m *GetRequestWithFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequestWithFields.Unmarshal(m, b)
}
func (m *GetRequestWithFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequestWithFields.Marshal(b, m, deterministic)
}
func (m *GetRequestWithFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequestWithFields.Merge(m, src)
}
func (m *GetRequestWithFields) XXX_Size() int {
	return xxx_messageInfo_GetRequestWithFields.Size(m)
}
func (m *GetRequestWithFields) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequestWithFields.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequestWithFields proto.InternalMessageInfo

func (m *GetRequestWithFields) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetRequestWithFields) GetFields() string {
	if m != nil {
		return m.Fields
	}
	return ""
}

// A `RelationshipRequest` defines the url parameter for relationship resources
// that are given in the links field of relationship object
type RelationshipRequest struct {
	// An unique identifier, for example:
	// "/users/45/roles" or "/users/45/relationships/roles"
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationshipRequest) Reset()         { *m = RelationshipRequest{} }
func (m *RelationshipRequest) String() string { return proto.CompactTextString(m) }
func (*RelationshipRequest) ProtoMessage()    {}
func (*RelationshipRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{3}
}

func (m *RelationshipRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationshipRequest.Unmarshal(m, b)
}
func (m *RelationshipRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationshipRequest.Marshal(b, m, deterministic)
}
func (m *RelationshipRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationshipRequest.Merge(m, src)
}
func (m *RelationshipRequest) XXX_Size() int {
	return xxx_messageInfo_RelationshipRequest.Size(m)
}
func (m *RelationshipRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationshipRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RelationshipRequest proto.InternalMessageInfo

func (m *RelationshipRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// A `RelationshipRequestWithPagination` is a `RelationshipRequest` with pagination
type RelationshipRequestWithPagination struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pagenum              int64    `protobuf:"varint,4,opt,name=pagenum,proto3" json:"pagenum,omitempty"`
	Pagesize             int64    `protobuf:"varint,3,opt,name=pagesize,proto3" json:"pagesize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelationshipRequestWithPagination) Reset()         { *m = RelationshipRequestWithPagination{} }
func (m *RelationshipRequestWithPagination) String() string { return proto.CompactTextString(m) }
func (*RelationshipRequestWithPagination) ProtoMessage()    {}
func (*RelationshipRequestWithPagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{4}
}

func (m *RelationshipRequestWithPagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelationshipRequestWithPagination.Unmarshal(m, b)
}
func (m *RelationshipRequestWithPagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelationshipRequestWithPagination.Marshal(b, m, deterministic)
}
func (m *RelationshipRequestWithPagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelationshipRequestWithPagination.Merge(m, src)
}
func (m *RelationshipRequestWithPagination) XXX_Size() int {
	return xxx_messageInfo_RelationshipRequestWithPagination.Size(m)
}
func (m *RelationshipRequestWithPagination) XXX_DiscardUnknown() {
	xxx_messageInfo_RelationshipRequestWithPagination.DiscardUnknown(m)
}

var xxx_messageInfo_RelationshipRequestWithPagination proto.InternalMessageInfo

func (m *RelationshipRequestWithPagination) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RelationshipRequestWithPagination) GetPagenum() int64 {
	if m != nil {
		return m.Pagenum
	}
	return 0
}

func (m *RelationshipRequestWithPagination) GetPagesize() int64 {
	if m != nil {
		return m.Pagesize
	}
	return 0
}

// A `ListRequest` defines various url and query parameters that could be
// passed in a HTTP **GET** request to a collection resource. All collection
// resources are expected to support pagination. Majority of the request
// parameters are identical or similar to [jsonapi](http://jsonapi.org).
type ListRequest struct {
	// include query parameter to retrieve any particular or particular
	// combination of relationships. Multiple include values are delimited by
	// comma(,).
	//
	// For example,
	// /{resource_name}/13?include=baz
	// /{resource_name}/13?include=baz,bot
	Include string `protobuf:"bytes,1,opt,name=include,proto3" json:"include,omitempty"`
	// fields query parameter to retrieve any particular or any particular
	// combination of attributes. Multiple fields values are delimited by comma(,).
	//
	// For example
	// /{resource_name}/29?fields=foo
	// /{resource_name}/?fields=foo,bar
	Fields string `protobuf:"bytes,3,opt,name=fields,proto3" json:"fields,omitempty"`
	// The page number to fetch
	Pagenum int64 `protobuf:"varint,4,opt,name=pagenum,proto3" json:"pagenum,omitempty"`
	// Number of records per page
	Pagesize int64 `protobuf:"varint,5,opt,name=pagesize,proto3" json:"pagesize,omitempty"`
	// The `filter` query parameter restricts the data return by the
	// collection. To use it, supply an attribute to filter, followed by a
	// filter expression. It uses the following syntax...
	//        attribute operator expression
	// attribute - Any one of the valid attribute of the resource.
	// operator - Defines the type of filter match to use. It could be any of
	// the following four and all of them should be URL-encoded.
	//
	//              ==  Equals (URL encoding is %3D%3D)
	//              !=  Not equals
	//              =@  Contains substring
	//              !@  Not contains substring
	//
	// expression - The value that will be included or excluded from the
	// result. URL-reserved characters must be URL-encoded.
	// For example, the following filter returns all users with last name `Gag`.
	//           /users?filter=last_name%3D%3Dgag
	//
	// Filter can be combined using OR or AND boolean logic.
	//   * The OR is represented using a comma(,).
	//   * The AND is represented using a semi-colon(;).
	//   * AND and OR operators can be combined and AND takes precedence over OR.
	Filter               string   `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{5}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetInclude() string {
	if m != nil {
		return m.Include
	}
	return ""
}

func (m *ListRequest) GetFields() string {
	if m != nil {
		return m.Fields
	}
	return ""
}

func (m *ListRequest) GetPagenum() int64 {
	if m != nil {
		return m.Pagenum
	}
	return 0
}

func (m *ListRequest) GetPagesize() int64 {
	if m != nil {
		return m.Pagesize
	}
	return 0
}

func (m *ListRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// A `SimpleListRequest` is identical to `ListRequest` except it does not support
// pagination. The rest of the parameters are identical to `ListRequest` definition.
type SimpleListRequest struct {
	// include query parameter to retrieve any particular or particular
	// combination of relationships. Multiple include values are delimited by
	// comma(,).
	//
	// For example,
	// /{resource_name}/13?include=baz
	// /{resource_name}/13?include=baz,bot
	Include string `protobuf:"bytes,1,opt,name=include,proto3" json:"include,omitempty"`
	// fields query parameter to retrieve any particular or any particular
	// combination of attributes. Multiple fields values are delimited by comma(,).
	//
	// For example
	// /{resource_name}/29?fields=foo
	// /{resource_name}/?fields=foo,bar
	Fields string `protobuf:"bytes,2,opt,name=fields,proto3" json:"fields,omitempty"`
	// The `filter` query parameter restricts the data return by the
	// collection. To use it, supply an attribute to filter, followed by a
	// filter expression. It uses the following syntax...
	//        attribute operator expression
	// attribute - Any one of the valid attribute of the resource.
	// operator - Defines the type of filter match to use. It could be any of
	// the following four and all of them should be URL-encoded.
	//
	//              ==  Equals (URL encoding is %3D%3D)
	//              !=  Not equals
	//              =@  Contains substring
	//              !@  Not contains substring
	//
	// expression - The value that will be included or excluded from the
	// result. URL-reserved characters must be URL-encoded.
	// For example, the following filter returns all users with last name `Gag`.
	//           /users?filter=last_name%3D%3Dgag
	//
	// Filter can be combined using OR or AND boolean logic.
	//   * The OR is represented using a comma(,).
	//   * The AND is represented using a semi-colon(;).
	//   * AND and OR operators can be combined and AND takes precedence over OR.
	Filter               string   `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleListRequest) Reset()         { *m = SimpleListRequest{} }
func (m *SimpleListRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleListRequest) ProtoMessage()    {}
func (*SimpleListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{6}
}

func (m *SimpleListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleListRequest.Unmarshal(m, b)
}
func (m *SimpleListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleListRequest.Marshal(b, m, deterministic)
}
func (m *SimpleListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleListRequest.Merge(m, src)
}
func (m *SimpleListRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleListRequest.Size(m)
}
func (m *SimpleListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleListRequest proto.InternalMessageInfo

func (m *SimpleListRequest) GetInclude() string {
	if m != nil {
		return m.Include
	}
	return ""
}

func (m *SimpleListRequest) GetFields() string {
	if m != nil {
		return m.Fields
	}
	return ""
}

func (m *SimpleListRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// A `DeleteRequest` defines the url parameter that must be passed
// to remove a singular resource.
type DeleteRequest struct {
	// An unique identifier, for example:
	// "/users/34"
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type IdRequest struct {
	// An unique identifier
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdRequest) Reset()         { *m = IdRequest{} }
func (m *IdRequest) String() string { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()    {}
func (*IdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{8}
}

func (m *IdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdRequest.Unmarshal(m, b)
}
func (m *IdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdRequest.Marshal(b, m, deterministic)
}
func (m *IdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdRequest.Merge(m, src)
}
func (m *IdRequest) XXX_Size() int {
	return xxx_messageInfo_IdRequest.Size(m)
}
func (m *IdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IdRequest proto.InternalMessageInfo

func (m *IdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type IdResponse struct {
	// An unique identifier
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdResponse) Reset()         { *m = IdResponse{} }
func (m *IdResponse) String() string { return proto.CompactTextString(m) }
func (*IdResponse) ProtoMessage()    {}
func (*IdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{9}
}

func (m *IdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IdResponse.Unmarshal(m, b)
}
func (m *IdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IdResponse.Marshal(b, m, deterministic)
}
func (m *IdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdResponse.Merge(m, src)
}
func (m *IdResponse) XXX_Size() int {
	return xxx_messageInfo_IdResponse.Size(m)
}
func (m *IdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IdResponse proto.InternalMessageInfo

func (m *IdResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type HealthzIdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthzIdRequest) Reset()         { *m = HealthzIdRequest{} }
func (m *HealthzIdRequest) String() string { return proto.CompactTextString(m) }
func (*HealthzIdRequest) ProtoMessage()    {}
func (*HealthzIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{10}
}

func (m *HealthzIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthzIdRequest.Unmarshal(m, b)
}
func (m *HealthzIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthzIdRequest.Marshal(b, m, deterministic)
}
func (m *HealthzIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthzIdRequest.Merge(m, src)
}
func (m *HealthzIdRequest) XXX_Size() int {
	return xxx_messageInfo_HealthzIdRequest.Size(m)
}
func (m *HealthzIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthzIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthzIdRequest proto.InternalMessageInfo

func (m *HealthzIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// ExistResponse wraps a boolean response
type ExistResponse struct {
	// exist or non-existant
	Exist                bool     `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistResponse) Reset()         { *m = ExistResponse{} }
func (m *ExistResponse) String() string { return proto.CompactTextString(m) }
func (*ExistResponse) ProtoMessage()    {}
func (*ExistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f73548e33e655fe, []int{11}
}

func (m *ExistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistResponse.Unmarshal(m, b)
}
func (m *ExistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistResponse.Marshal(b, m, deterministic)
}
func (m *ExistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistResponse.Merge(m, src)
}
func (m *ExistResponse) XXX_Size() int {
	return xxx_messageInfo_ExistResponse.Size(m)
}
func (m *ExistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExistResponse proto.InternalMessageInfo

func (m *ExistResponse) GetExist() bool {
	if m != nil {
		return m.Exist
	}
	return false
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "dictybase.api.jsonapi.GetRequest")
	proto.RegisterType((*GetEmailRequest)(nil), "dictybase.api.jsonapi.GetEmailRequest")
	proto.RegisterType((*GetRequestWithFields)(nil), "dictybase.api.jsonapi.GetRequestWithFields")
	proto.RegisterType((*RelationshipRequest)(nil), "dictybase.api.jsonapi.RelationshipRequest")
	proto.RegisterType((*RelationshipRequestWithPagination)(nil), "dictybase.api.jsonapi.RelationshipRequestWithPagination")
	proto.RegisterType((*ListRequest)(nil), "dictybase.api.jsonapi.ListRequest")
	proto.RegisterType((*SimpleListRequest)(nil), "dictybase.api.jsonapi.SimpleListRequest")
	proto.RegisterType((*DeleteRequest)(nil), "dictybase.api.jsonapi.DeleteRequest")
	proto.RegisterType((*IdRequest)(nil), "dictybase.api.jsonapi.IdRequest")
	proto.RegisterType((*IdResponse)(nil), "dictybase.api.jsonapi.IdResponse")
	proto.RegisterType((*HealthzIdRequest)(nil), "dictybase.api.jsonapi.HealthzIdRequest")
	proto.RegisterType((*ExistResponse)(nil), "dictybase.api.jsonapi.ExistResponse")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor_7f73548e33e655fe) }

var fileDescriptor_7f73548e33e655fe = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0x4d, 0xb6, 0x5d, 0x77, 0x8f, 0xae, 0xd6, 0x58, 0x25, 0x56, 0xa1, 0x35, 0x50, 0xe8,
	0x85, 0x4d, 0x2e, 0x04, 0x6f, 0x04, 0xc1, 0xd8, 0x5a, 0x17, 0x44, 0x96, 0x28, 0x48, 0x05, 0x2f,
	0x66, 0x37, 0xc7, 0xec, 0xd1, 0x49, 0x66, 0xcc, 0xcc, 0x5a, 0x2d, 0xf8, 0x08, 0xbe, 0x84, 0x2f,
	0x26, 0xf8, 0x14, 0x5e, 0xca, 0x4c, 0xb2, 0x7f, 0x42, 0x77, 0x61, 0xf5, 0x2a, 0xf3, 0xcd, 0x39,
	0xe7, 0x77, 0xe6, 0x7c, 0xe4, 0x40, 0xaf, 0xc4, 0xcf, 0x13, 0x54, 0x3a, 0x94, 0xa5, 0xd0, 0xc2,
	0xbb, 0x95, 0xd2, 0x48, 0x7f, 0x1b, 0x32, 0x85, 0x21, 0x93, 0x14, 0x7e, 0x54, 0xa2, 0x60, 0x92,
	0x76, 0x1e, 0x65, 0xa4, 0xc7, 0x93, 0x61, 0x38, 0x12, 0x79, 0x94, 0x9f, 0x91, 0xfe, 0x24, 0xce,
	0xa2, 0x4c, 0x1c, 0xda, 0x9a, 0xc3, 0x2f, 0x8c, 0x53, 0xca, 0xb4, 0x28, 0x55, 0x34, 0x3b, 0x56,
	0xb8, 0xe0, 0x15, 0xc0, 0x09, 0xea, 0xa4, 0x6a, 0xe1, 0x5d, 0x03, 0x97, 0x52, 0xdf, 0xd9, 0x73,
	0x0e, 0x5a, 0x89, 0x4b, 0xa9, 0xe7, 0xc3, 0x65, 0x2a, 0x46, 0x7c, 0x92, 0xa2, 0xef, 0xee, 0x39,
	0x07, 0xdd, 0x64, 0x2a, 0xbd, 0xdb, 0xd0, 0xfe, 0x40, 0xc8, 0x53, 0xe5, 0xb7, 0x6c, 0xa0, 0x56,
	0xc1, 0x29, 0x5c, 0x3f, 0x41, 0x7d, 0x9c, 0x33, 0xe2, 0x53, 0xe8, 0x36, 0x6c, 0xa2, 0xd1, 0x96,
	0xdb, 0x4d, 0x2a, 0xf1, 0x1f, 0xe8, 0x27, 0xb0, 0x3d, 0x7f, 0xea, 0x5b, 0xd2, 0xe3, 0xe7, 0xf6,
	0xfe, 0xc2, 0xa3, 0xe7, 0xf5, 0x6e, 0xa3, 0x7e, 0x1f, 0x6e, 0x26, 0xc8, 0x99, 0x26, 0x51, 0xa8,
	0x31, 0xc9, 0x15, 0x33, 0x07, 0x04, 0xf7, 0x97, 0xa4, 0x99, 0x7e, 0x03, 0x96, 0x51, 0x61, 0x03,
	0xcb, 0x8c, 0x92, 0x2c, 0xc3, 0x62, 0x92, 0xfb, 0x1b, 0xf6, 0x72, 0x2a, 0xbd, 0x1d, 0xe8, 0x98,
	0xa3, 0xa2, 0x73, 0xb4, 0xf3, 0xb4, 0x92, 0x99, 0x0e, 0x7e, 0x38, 0x70, 0xe5, 0x25, 0xa9, 0x99,
	0xfd, 0x0b, 0x9e, 0x38, 0x6b, 0x79, 0xb2, 0x66, 0xdf, 0xcd, 0x66, 0xdf, 0x8a, 0xc6, 0x35, 0x96,
	0x7e, 0x7b, 0x4a, 0x33, 0x2a, 0x78, 0x0f, 0x37, 0x5e, 0x53, 0x2e, 0x39, 0xfe, 0xeb, 0xa3, 0x1a,
	0x46, 0x2f, 0xe0, 0x5b, 0x0d, 0xfc, 0x2e, 0xf4, 0x8e, 0x90, 0xa3, 0xc6, 0x55, 0xd6, 0xdf, 0x85,
	0x6e, 0x3f, 0x5d, 0x15, 0xbc, 0x07, 0x60, 0x82, 0x4a, 0x8a, 0x42, 0xe1, 0x85, 0xe8, 0x03, 0xd8,
	0x7a, 0x81, 0x8c, 0xeb, 0xf1, 0xf9, 0x9c, 0xe0, 0xcf, 0x73, 0xe2, 0xce, 0xef, 0x5f, 0xbb, 0x1b,
	0x5b, 0x97, 0x7c, 0xd7, 0x66, 0xef, 0x43, 0xef, 0xf8, 0xab, 0x9d, 0xb1, 0xc6, 0x99, 0x7f, 0xd4,
	0x5c, 0xd8, 0xec, 0x4e, 0x52, 0x89, 0xf8, 0x3b, 0xdc, 0x11, 0x65, 0x16, 0x2e, 0xdd, 0xb8, 0xf8,
	0x6a, 0xdd, 0x66, 0x60, 0xf6, 0x68, 0xe0, 0xbc, 0x8b, 0x17, 0x36, 0xd0, 0x56, 0xc4, 0x4c, 0xa1,
	0xd9, 0xc1, 0x0c, 0x0b, 0xbb, 0x6b, 0xd1, 0x8c, 0xc3, 0x24, 0xa9, 0x88, 0x49, 0x8a, 0x6a, 0xd6,
	0xe3, 0xfa, 0xfb, 0xc7, 0x71, 0x7e, 0xba, 0x9d, 0xa3, 0xfe, 0xb3, 0x37, 0xa7, 0x4f, 0x07, 0xfd,
	0x61, 0xdb, 0x96, 0x3d, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x65, 0xd3, 0x94, 0x8d, 0x02, 0x04,
	0x00, 0x00,
}
