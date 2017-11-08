// Code generated by protoc-gen-go. DO NOT EDIT.
// source: permission.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	permission.proto
	role.proto
	user.proto

It has these top-level messages:
	UpdatePermissionRequest
	CreatePermissionRequest
	Permission
	PermissionCollection
	PermissionData
	PermissionAttributes
	UpdateRoleRequest
	CreateRoleRequest
	Role
	RoleCollection
	RoleData
	RoleAttributes
	ExistingRoleRelationships
	NewRoleRelationships
	UpdateUserRequest
	CreateUserRequest
	User
	UserCollection
	UserData
	UserAttributes
	ExistingUserRelationships
	NewUserRelationships
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "google.golang.org/genproto/protobuf/field_mask"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
import dictybase_api_jsonapi "github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
import dictybase_api_jsonapi1 "github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"

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

type UpdatePermissionRequest struct {
	Data *UpdatePermissionRequest_Data `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	// An optional mask specifying which fields to update.
	// Presence of this field allow partial updates.
	UpdateMask *google_protobuf1.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdatePermissionRequest) Reset()                    { *m = UpdatePermissionRequest{} }
func (m *UpdatePermissionRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatePermissionRequest) ProtoMessage()               {}
func (*UpdatePermissionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdatePermissionRequest) GetData() *UpdatePermissionRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UpdatePermissionRequest) GetUpdateMask() *google_protobuf1.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type UpdatePermissionRequest_Data struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Unique id, required
	Id         int64                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Attributes *PermissionAttributes `protobuf:"bytes,3,opt,name=attributes" json:"attributes,omitempty"`
}

func (m *UpdatePermissionRequest_Data) Reset()                    { *m = UpdatePermissionRequest_Data{} }
func (m *UpdatePermissionRequest_Data) String() string            { return proto.CompactTextString(m) }
func (*UpdatePermissionRequest_Data) ProtoMessage()               {}
func (*UpdatePermissionRequest_Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *UpdatePermissionRequest_Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UpdatePermissionRequest_Data) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdatePermissionRequest_Data) GetAttributes() *PermissionAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type CreatePermissionRequest struct {
	Data *CreatePermissionRequest_Data `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *CreatePermissionRequest) Reset()                    { *m = CreatePermissionRequest{} }
func (m *CreatePermissionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePermissionRequest) ProtoMessage()               {}
func (*CreatePermissionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreatePermissionRequest) GetData() *CreatePermissionRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

// The payload for new user
type CreatePermissionRequest_Data struct {
	Type       string                `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Attributes *PermissionAttributes `protobuf:"bytes,2,opt,name=attributes" json:"attributes,omitempty"`
}

func (m *CreatePermissionRequest_Data) Reset()                    { *m = CreatePermissionRequest_Data{} }
func (m *CreatePermissionRequest_Data) String() string            { return proto.CompactTextString(m) }
func (*CreatePermissionRequest_Data) ProtoMessage()               {}
func (*CreatePermissionRequest_Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *CreatePermissionRequest_Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreatePermissionRequest_Data) GetAttributes() *PermissionAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// A resource for managing user permission.
type Permission struct {
	Data *PermissionData `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *Permission) Reset()                    { *m = Permission{} }
func (m *Permission) String() string            { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()               {}
func (*Permission) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Permission) GetData() *PermissionData {
	if m != nil {
		return m.Data
	}
	return nil
}

// A permission collection resource.
type PermissionCollection struct {
	Data  []*PermissionData            `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	Links *dictybase_api_jsonapi.Links `protobuf:"bytes,2,opt,name=links" json:"links,omitempty"`
}

func (m *PermissionCollection) Reset()                    { *m = PermissionCollection{} }
func (m *PermissionCollection) String() string            { return proto.CompactTextString(m) }
func (*PermissionCollection) ProtoMessage()               {}
func (*PermissionCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PermissionCollection) GetData() []*PermissionData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PermissionCollection) GetLinks() *dictybase_api_jsonapi.Links {
	if m != nil {
		return m.Links
	}
	return nil
}

// A top level container for permission data.
type PermissionData struct {
	// The resource name.
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Unique id.
	Id         int64                        `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Attributes *PermissionAttributes        `protobuf:"bytes,3,opt,name=attributes" json:"attributes,omitempty"`
	Links      *dictybase_api_jsonapi.Links `protobuf:"bytes,4,opt,name=links" json:"links,omitempty"`
}

func (m *PermissionData) Reset()                    { *m = PermissionData{} }
func (m *PermissionData) String() string            { return proto.CompactTextString(m) }
func (*PermissionData) ProtoMessage()               {}
func (*PermissionData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PermissionData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PermissionData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PermissionData) GetAttributes() *PermissionAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *PermissionData) GetLinks() *dictybase_api_jsonapi.Links {
	if m != nil {
		return m.Links
	}
	return nil
}

// A container for permission fields.
type PermissionAttributes struct {
	// Kind of permission.
	Permission string `protobuf:"bytes,1,opt,name=permission" json:"permission,omitempty"`
	// Brief description of the type of permission.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// Timestamp for creation and update
	CreatedAt *google_protobuf3.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf3.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *PermissionAttributes) Reset()                    { *m = PermissionAttributes{} }
func (m *PermissionAttributes) String() string            { return proto.CompactTextString(m) }
func (*PermissionAttributes) ProtoMessage()               {}
func (*PermissionAttributes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PermissionAttributes) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

func (m *PermissionAttributes) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PermissionAttributes) GetCreatedAt() *google_protobuf3.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *PermissionAttributes) GetUpdatedAt() *google_protobuf3.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdatePermissionRequest)(nil), "dictybase.user.UpdatePermissionRequest")
	proto.RegisterType((*UpdatePermissionRequest_Data)(nil), "dictybase.user.UpdatePermissionRequest.Data")
	proto.RegisterType((*CreatePermissionRequest)(nil), "dictybase.user.CreatePermissionRequest")
	proto.RegisterType((*CreatePermissionRequest_Data)(nil), "dictybase.user.CreatePermissionRequest.Data")
	proto.RegisterType((*Permission)(nil), "dictybase.user.Permission")
	proto.RegisterType((*PermissionCollection)(nil), "dictybase.user.PermissionCollection")
	proto.RegisterType((*PermissionData)(nil), "dictybase.user.PermissionData")
	proto.RegisterType((*PermissionAttributes)(nil), "dictybase.user.PermissionAttributes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PermissionService service

type PermissionServiceClient interface {
	// Gets the specified permission
	GetPermission(ctx context.Context, in *dictybase_api_jsonapi1.GetRequest, opts ...grpc.CallOption) (*Permission, error)
	// List all permissions
	ListPermissions(ctx context.Context, in *dictybase_api_jsonapi1.ListRequest, opts ...grpc.CallOption) (*PermissionCollection, error)
	// Create an permission
	CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*Permission, error)
	// Update an permission
	UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*Permission, error)
	// Delete an permission
	DeletePermission(ctx context.Context, in *dictybase_api_jsonapi1.DeleteRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type permissionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPermissionServiceClient(cc *grpc.ClientConn) PermissionServiceClient {
	return &permissionServiceClient{cc}
}

func (c *permissionServiceClient) GetPermission(ctx context.Context, in *dictybase_api_jsonapi1.GetRequest, opts ...grpc.CallOption) (*Permission, error) {
	out := new(Permission)
	err := grpc.Invoke(ctx, "/dictybase.user.PermissionService/GetPermission", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) ListPermissions(ctx context.Context, in *dictybase_api_jsonapi1.ListRequest, opts ...grpc.CallOption) (*PermissionCollection, error) {
	out := new(PermissionCollection)
	err := grpc.Invoke(ctx, "/dictybase.user.PermissionService/ListPermissions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*Permission, error) {
	out := new(Permission)
	err := grpc.Invoke(ctx, "/dictybase.user.PermissionService/CreatePermission", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...grpc.CallOption) (*Permission, error) {
	out := new(Permission)
	err := grpc.Invoke(ctx, "/dictybase.user.PermissionService/UpdatePermission", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionServiceClient) DeletePermission(ctx context.Context, in *dictybase_api_jsonapi1.DeleteRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/dictybase.user.PermissionService/DeletePermission", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PermissionService service

type PermissionServiceServer interface {
	// Gets the specified permission
	GetPermission(context.Context, *dictybase_api_jsonapi1.GetRequest) (*Permission, error)
	// List all permissions
	ListPermissions(context.Context, *dictybase_api_jsonapi1.ListRequest) (*PermissionCollection, error)
	// Create an permission
	CreatePermission(context.Context, *CreatePermissionRequest) (*Permission, error)
	// Update an permission
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*Permission, error)
	// Delete an permission
	DeletePermission(context.Context, *dictybase_api_jsonapi1.DeleteRequest) (*google_protobuf2.Empty, error)
}

func RegisterPermissionServiceServer(s *grpc.Server, srv PermissionServiceServer) {
	s.RegisterService(&_PermissionService_serviceDesc, srv)
}

func _PermissionService_GetPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dictybase_api_jsonapi1.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).GetPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.user.PermissionService/GetPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).GetPermission(ctx, req.(*dictybase_api_jsonapi1.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_ListPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dictybase_api_jsonapi1.ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).ListPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.user.PermissionService/ListPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).ListPermissions(ctx, req.(*dictybase_api_jsonapi1.ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.user.PermissionService/CreatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).CreatePermission(ctx, req.(*CreatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_UpdatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).UpdatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.user.PermissionService/UpdatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).UpdatePermission(ctx, req.(*UpdatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PermissionService_DeletePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dictybase_api_jsonapi1.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PermissionServiceServer).DeletePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.user.PermissionService/DeletePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PermissionServiceServer).DeletePermission(ctx, req.(*dictybase_api_jsonapi1.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PermissionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dictybase.user.PermissionService",
	HandlerType: (*PermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPermission",
			Handler:    _PermissionService_GetPermission_Handler,
		},
		{
			MethodName: "ListPermissions",
			Handler:    _PermissionService_ListPermissions_Handler,
		},
		{
			MethodName: "CreatePermission",
			Handler:    _PermissionService_CreatePermission_Handler,
		},
		{
			MethodName: "UpdatePermission",
			Handler:    _PermissionService_UpdatePermission_Handler,
		},
		{
			MethodName: "DeletePermission",
			Handler:    _PermissionService_DeletePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "permission.proto",
}

func init() { proto.RegisterFile("permission.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4d, 0x6e, 0xd3, 0x40,
	0x18, 0x95, 0xdd, 0x80, 0xe8, 0x57, 0x48, 0xdb, 0xa1, 0xa2, 0xc5, 0x54, 0xa5, 0x98, 0x4a, 0xa0,
	0x0a, 0x6c, 0x29, 0x6c, 0x80, 0x6e, 0x9a, 0x36, 0x50, 0x55, 0x2a, 0x52, 0x64, 0xca, 0x02, 0x36,
	0x65, 0x12, 0x4f, 0xc3, 0x10, 0xc7, 0x33, 0xf5, 0x8c, 0x91, 0x22, 0x04, 0x0b, 0x96, 0x6c, 0xb9,
	0x01, 0x07, 0xe0, 0x06, 0x6c, 0x38, 0x00, 0x1b, 0xae, 0xc0, 0x21, 0x58, 0xa2, 0x19, 0x3b, 0xf1,
	0x4f, 0xe4, 0x10, 0x75, 0xc1, 0xa6, 0x75, 0xc6, 0xef, 0x7b, 0xef, 0x7d, 0x7f, 0x63, 0x58, 0xe2,
	0x24, 0x1a, 0x50, 0x21, 0x28, 0x0b, 0x1d, 0x1e, 0x31, 0xc9, 0x50, 0xdd, 0xa7, 0x5d, 0x39, 0xec,
	0x60, 0x41, 0x9c, 0x58, 0x90, 0xc8, 0x5a, 0xef, 0x31, 0xd6, 0x0b, 0x88, 0x8b, 0x39, 0x75, 0x71,
	0x18, 0x32, 0x89, 0x25, 0x65, 0xa1, 0x48, 0xd0, 0xd6, 0x66, 0xfa, 0x56, 0xff, 0xea, 0xc4, 0xa7,
	0xee, 0x29, 0x25, 0x81, 0x7f, 0x32, 0xc0, 0xa2, 0x9f, 0x22, 0x6e, 0x94, 0x11, 0x64, 0xc0, 0xe5,
	0x30, 0x7d, 0x79, 0x7b, 0x2c, 0xa6, 0xf9, 0xdf, 0x0a, 0x16, 0xaa, 0xff, 0x1c, 0x0f, 0x03, 0x86,
	0xfd, 0xe9, 0xa0, 0x88, 0x9c, 0xc5, 0x44, 0xc8, 0x14, 0x74, 0xb3, 0x2c, 0x23, 0xe9, 0x80, 0x08,
	0x89, 0x07, 0x3c, 0x01, 0xd8, 0x9f, 0x4d, 0x58, 0x7d, 0xc1, 0x7d, 0x2c, 0x49, 0x7b, 0x9c, 0xb2,
	0x97, 0x50, 0xa0, 0x5d, 0xa8, 0xf9, 0x58, 0xe2, 0x35, 0x63, 0xd3, 0xb8, 0xbb, 0xd0, 0xb8, 0xe7,
	0x14, 0x4b, 0xe0, 0x54, 0x84, 0x39, 0x2d, 0x2c, 0xb1, 0xa7, 0x23, 0xd1, 0x0e, 0x2c, 0xc4, 0x1a,
	0xa5, 0x53, 0x5f, 0x33, 0x35, 0x91, 0xe5, 0x24, 0xa6, 0x9c, 0x91, 0x29, 0xe7, 0xa9, 0xaa, 0xce,
	0x33, 0x2c, 0xfa, 0x1e, 0x24, 0x70, 0xf5, 0x6c, 0x71, 0xa8, 0x29, 0x2a, 0x84, 0xa0, 0x26, 0x87,
	0x9c, 0x68, 0x1b, 0xf3, 0x9e, 0x7e, 0x46, 0x75, 0x30, 0xa9, 0xaf, 0xf9, 0xe6, 0x3c, 0x93, 0xfa,
	0xa8, 0x05, 0x80, 0xa5, 0x8c, 0x68, 0x27, 0x96, 0x44, 0xac, 0xcd, 0x69, 0x9d, 0xad, 0xb2, 0xe1,
	0xcc, 0x6a, 0x73, 0x8c, 0xf5, 0x72, 0x71, 0xf6, 0x77, 0x03, 0x56, 0xf7, 0x23, 0x72, 0x9e, 0x62,
	0x54, 0x84, 0xe5, 0x8a, 0x61, 0xbd, 0x9e, 0x92, 0x4f, 0xd1, 0xbf, 0x79, 0x4e, 0xff, 0xbb, 0x00,
	0x19, 0x06, 0x35, 0x0a, 0x8e, 0x37, 0xaa, 0xd9, 0x32, 0x8f, 0xf6, 0x47, 0x58, 0xc9, 0xce, 0xf7,
	0x59, 0x10, 0x90, 0xae, 0x2c, 0x72, 0xcd, 0xcd, 0xca, 0x85, 0x1a, 0x70, 0x21, 0xa0, 0x61, 0x7f,
	0x94, 0xce, 0x7a, 0x2e, 0x08, 0x73, 0xea, 0xa4, 0x03, 0xeb, 0x1c, 0x29, 0x8c, 0x97, 0x40, 0xed,
	0x6f, 0x06, 0xd4, 0x8b, 0x64, 0xff, 0xaf, 0xfd, 0x99, 0xe1, 0xda, 0xec, 0x86, 0x7f, 0x1a, 0xf9,
	0x8a, 0x65, 0xc4, 0x68, 0x03, 0x20, 0xbb, 0x44, 0x52, 0xf3, 0xb9, 0x13, 0xb4, 0x09, 0x0b, 0x3e,
	0x11, 0xdd, 0x88, 0x72, 0x55, 0x60, 0x9d, 0xcb, 0xbc, 0x97, 0x3f, 0x42, 0x8f, 0x00, 0xba, 0x7a,
	0xaa, 0xfc, 0x13, 0x2c, 0xd3, 0xa4, 0x26, 0x77, 0xe7, 0x78, 0xb4, 0xd0, 0xde, 0x7c, 0x8a, 0x6e,
	0x4a, 0x15, 0x9a, 0x2c, 0x92, 0x0e, 0xad, 0xfd, 0x3b, 0x34, 0x45, 0x37, 0x65, 0xe3, 0x47, 0x0d,
	0x96, 0xb3, 0x84, 0x9e, 0x93, 0xe8, 0x1d, 0xed, 0x12, 0x44, 0xe1, 0xca, 0x01, 0x91, 0xb9, 0xe1,
	0xba, 0x55, 0x51, 0x9c, 0x03, 0x22, 0xd3, 0xd1, 0xb7, 0xac, 0xea, 0x06, 0xd8, 0xd7, 0x3f, 0xfd,
	0xfa, 0xfd, 0xc5, 0xbc, 0x8a, 0x96, 0xdd, 0xac, 0x24, 0xc2, 0x7d, 0x4f, 0xfd, 0x0f, 0xe8, 0x0c,
	0x16, 0x8f, 0xa8, 0xc8, 0x69, 0x09, 0x64, 0x57, 0x76, 0x42, 0x8c, 0xd5, 0xa6, 0xb4, 0x3b, 0x9b,
	0x63, 0x7b, 0x45, 0xeb, 0xd6, 0xd1, 0xe5, 0xbc, 0x2e, 0xe2, 0xb0, 0x54, 0xde, 0x5f, 0x74, 0x67,
	0xc6, 0x0d, 0x9f, 0x9a, 0xe6, 0xaa, 0x96, 0x5b, 0xb6, 0x0b, 0x72, 0x8f, 0x8d, 0x6d, 0xa5, 0x58,
	0xbe, 0x3e, 0x27, 0x15, 0x2b, 0x2e, 0xd8, 0x59, 0x14, 0x1b, 0x13, 0x8a, 0x7d, 0x58, 0x6a, 0x91,
	0x80, 0x14, 0x14, 0xb7, 0x2a, 0xea, 0x9a, 0x00, 0x47, 0x72, 0xd7, 0x26, 0x06, 0xe7, 0x89, 0xfa,
	0x56, 0x8d, 0x7a, 0xb8, 0x3d, 0xd9, 0xc3, 0x3d, 0x06, 0x88, 0x45, 0xbd, 0x92, 0xcd, 0xbd, 0xc5,
	0x4c, 0xba, 0xad, 0xa8, 0xda, 0xc6, 0xab, 0x87, 0x3d, 0x2a, 0xdf, 0xc4, 0x1d, 0xa7, 0xcb, 0x06,
	0xae, 0x46, 0xef, 0xa9, 0xef, 0x59, 0x8f, 0xdd, 0xef, 0x91, 0x50, 0xcb, 0xb9, 0x63, 0x0e, 0xcc,
	0xa9, 0x70, 0x15, 0xcf, 0x8e, 0xfa, 0xf3, 0xc7, 0x30, 0xbe, 0x9a, 0x97, 0x5a, 0x87, 0xfb, 0xc7,
	0x2f, 0x9b, 0xed, 0xc3, 0xce, 0x45, 0x0d, 0x7e, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xb7,
	0xad, 0xa4, 0xb9, 0x07, 0x00, 0x00,
}
