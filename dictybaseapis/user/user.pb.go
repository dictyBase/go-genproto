// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "google.golang.org/genproto/protobuf/field_mask"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
import dictybase_api_jsonapi "github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
import dictybase_api_jsonapi1 "github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UpdateUserRequest struct {
	Data *UpdateUserRequest_Data `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	// An optional mask specifying which fields to update.
	// Presence of this field allow partial updates.
	UpdateMask *google_protobuf1.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateUserRequest) Reset()                    { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()               {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *UpdateUserRequest) GetData() *UpdateUserRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UpdateUserRequest) GetUpdateMask() *google_protobuf1.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type UpdateUserRequest_Data struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Unique id, required
	Id            int64                      `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Attributes    *UserAttributes            `protobuf:"bytes,3,opt,name=attributes" json:"attributes,omitempty"`
	Relationships *ExistingUserRelationships `protobuf:"bytes,4,opt,name=relationships" json:"relationships,omitempty"`
}

func (m *UpdateUserRequest_Data) Reset()                    { *m = UpdateUserRequest_Data{} }
func (m *UpdateUserRequest_Data) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRequest_Data) ProtoMessage()               {}
func (*UpdateUserRequest_Data) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

func (m *UpdateUserRequest_Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UpdateUserRequest_Data) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateUserRequest_Data) GetAttributes() *UserAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *UpdateUserRequest_Data) GetRelationships() *ExistingUserRelationships {
	if m != nil {
		return m.Relationships
	}
	return nil
}

type CreateUserRequest struct {
	Data *CreateUserRequest_Data `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *CreateUserRequest) GetData() *CreateUserRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

// The payload for new user
type CreateUserRequest_Data struct {
	Type          string                `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Attributes    *UserAttributes       `protobuf:"bytes,2,opt,name=attributes" json:"attributes,omitempty"`
	Relationships *NewUserRelationships `protobuf:"bytes,3,opt,name=relationships" json:"relationships,omitempty"`
}

func (m *CreateUserRequest_Data) Reset()                    { *m = CreateUserRequest_Data{} }
func (m *CreateUserRequest_Data) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest_Data) ProtoMessage()               {}
func (*CreateUserRequest_Data) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

func (m *CreateUserRequest_Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateUserRequest_Data) GetAttributes() *UserAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *CreateUserRequest_Data) GetRelationships() *NewUserRelationships {
	if m != nil {
		return m.Relationships
	}
	return nil
}

// A user resource.
type User struct {
	Data  *UserData                    `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Links *dictybase_api_jsonapi.Links `protobuf:"bytes,2,opt,name=links" json:"links,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *User) GetData() *UserData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *User) GetLinks() *dictybase_api_jsonapi.Links {
	if m != nil {
		return m.Links
	}
	return nil
}

// A user collection resource.
type UserCollection struct {
	Data  []*UserData                            `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	Links *dictybase_api_jsonapi.PaginationLinks `protobuf:"bytes,2,opt,name=links" json:"links,omitempty"`
	Meta  *dictybase_api_jsonapi.Meta            `protobuf:"bytes,3,opt,name=meta" json:"meta,omitempty"`
}

func (m *UserCollection) Reset()                    { *m = UserCollection{} }
func (m *UserCollection) String() string            { return proto.CompactTextString(m) }
func (*UserCollection) ProtoMessage()               {}
func (*UserCollection) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *UserCollection) GetData() []*UserData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UserCollection) GetLinks() *dictybase_api_jsonapi.PaginationLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *UserCollection) GetMeta() *dictybase_api_jsonapi.Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

// A top level container for user data.
type UserData struct {
	// The resource name.
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Unique id.
	Id            int64                        `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Attributes    *UserAttributes              `protobuf:"bytes,3,opt,name=attributes" json:"attributes,omitempty"`
	Links         *dictybase_api_jsonapi.Links `protobuf:"bytes,4,opt,name=links" json:"links,omitempty"`
	Relationships *ExistingUserRelationships   `protobuf:"bytes,5,opt,name=relationships" json:"relationships,omitempty"`
}

func (m *UserData) Reset()                    { *m = UserData{} }
func (m *UserData) String() string            { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()               {}
func (*UserData) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *UserData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UserData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserData) GetAttributes() *UserAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *UserData) GetLinks() *dictybase_api_jsonapi.Links {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *UserData) GetRelationships() *ExistingUserRelationships {
	if m != nil {
		return m.Relationships
	}
	return nil
}

// A container for user fields.
type UserAttributes struct {
	// First name.
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	// Last name.
	LastName string `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	// Email.
	Email string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	// Organization in which the user belong.
	Organization string `protobuf:"bytes,4,opt,name=organization" json:"organization,omitempty"`
	// Group in which the user belong.
	Group string `protobuf:"bytes,5,opt,name=group" json:"group,omitempty"`
	// Address.
	FirstAddress string `protobuf:"bytes,6,opt,name=first_address,json=firstAddress" json:"first_address,omitempty"`
	// More address.
	SecondAddress string `protobuf:"bytes,7,opt,name=second_address,json=secondAddress" json:"second_address,omitempty"`
	// City.
	City string `protobuf:"bytes,8,opt,name=city" json:"city,omitempty"`
	// State.
	State string `protobuf:"bytes,9,opt,name=state" json:"state,omitempty"`
	// Zipcode.
	Zip string `protobuf:"bytes,10,opt,name=zip" json:"zip,omitempty"`
	// Country.
	Country string `protobuf:"bytes,11,opt,name=country" json:"country,omitempty"`
	// Phone no.
	Phone string `protobuf:"bytes,12,opt,name=phone" json:"phone,omitempty"`
	// Current status of user.
	IsActive bool `protobuf:"varint,13,opt,name=is_active,json=isActive" json:"is_active,omitempty"`
	// Timestamp for creation and update
	CreatedAt *google_protobuf3.Timestamp `protobuf:"bytes,14,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf3.Timestamp `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *UserAttributes) Reset()                    { *m = UserAttributes{} }
func (m *UserAttributes) String() string            { return proto.CompactTextString(m) }
func (*UserAttributes) ProtoMessage()               {}
func (*UserAttributes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *UserAttributes) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserAttributes) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserAttributes) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserAttributes) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func (m *UserAttributes) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *UserAttributes) GetFirstAddress() string {
	if m != nil {
		return m.FirstAddress
	}
	return ""
}

func (m *UserAttributes) GetSecondAddress() string {
	if m != nil {
		return m.SecondAddress
	}
	return ""
}

func (m *UserAttributes) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *UserAttributes) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *UserAttributes) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *UserAttributes) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *UserAttributes) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserAttributes) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *UserAttributes) GetCreatedAt() *google_protobuf3.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *UserAttributes) GetUpdatedAt() *google_protobuf3.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// The relationship definition for existing users.
type ExistingUserRelationships struct {
	Roles *ExistingUserRelationships_Roles `protobuf:"bytes,1,opt,name=roles" json:"roles,omitempty"`
}

func (m *ExistingUserRelationships) Reset()                    { *m = ExistingUserRelationships{} }
func (m *ExistingUserRelationships) String() string            { return proto.CompactTextString(m) }
func (*ExistingUserRelationships) ProtoMessage()               {}
func (*ExistingUserRelationships) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *ExistingUserRelationships) GetRoles() *ExistingUserRelationships_Roles {
	if m != nil {
		return m.Roles
	}
	return nil
}

// Relationships with role resource.
type ExistingUserRelationships_Roles struct {
	// Http links with role resource.
	Links *dictybase_api_jsonapi.Links `protobuf:"bytes,1,opt,name=links" json:"links,omitempty"`
}

func (m *ExistingUserRelationships_Roles) Reset()         { *m = ExistingUserRelationships_Roles{} }
func (m *ExistingUserRelationships_Roles) String() string { return proto.CompactTextString(m) }
func (*ExistingUserRelationships_Roles) ProtoMessage()    {}
func (*ExistingUserRelationships_Roles) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{6, 0}
}

func (m *ExistingUserRelationships_Roles) GetLinks() *dictybase_api_jsonapi.Links {
	if m != nil {
		return m.Links
	}
	return nil
}

// The relationship definition for creating new users.
type NewUserRelationships struct {
	Roles []*NewUserRelationships_Roles `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
}

func (m *NewUserRelationships) Reset()                    { *m = NewUserRelationships{} }
func (m *NewUserRelationships) String() string            { return proto.CompactTextString(m) }
func (*NewUserRelationships) ProtoMessage()               {}
func (*NewUserRelationships) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *NewUserRelationships) GetRoles() []*NewUserRelationships_Roles {
	if m != nil {
		return m.Roles
	}
	return nil
}

// Relationships with role resource.
type NewUserRelationships_Roles struct {
	// A role [resource identifier object](http://jsonapi.org/format/#document-resource-identifier-objects).
	Data *dictybase_api_jsonapi.Data `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *NewUserRelationships_Roles) Reset()                    { *m = NewUserRelationships_Roles{} }
func (m *NewUserRelationships_Roles) String() string            { return proto.CompactTextString(m) }
func (*NewUserRelationships_Roles) ProtoMessage()               {}
func (*NewUserRelationships_Roles) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7, 0} }

func (m *NewUserRelationships_Roles) GetData() *dictybase_api_jsonapi.Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateUserRequest)(nil), "dictybase.user.UpdateUserRequest")
	proto.RegisterType((*UpdateUserRequest_Data)(nil), "dictybase.user.UpdateUserRequest.Data")
	proto.RegisterType((*CreateUserRequest)(nil), "dictybase.user.CreateUserRequest")
	proto.RegisterType((*CreateUserRequest_Data)(nil), "dictybase.user.CreateUserRequest.Data")
	proto.RegisterType((*User)(nil), "dictybase.user.User")
	proto.RegisterType((*UserCollection)(nil), "dictybase.user.UserCollection")
	proto.RegisterType((*UserData)(nil), "dictybase.user.UserData")
	proto.RegisterType((*UserAttributes)(nil), "dictybase.user.UserAttributes")
	proto.RegisterType((*ExistingUserRelationships)(nil), "dictybase.user.ExistingUserRelationships")
	proto.RegisterType((*ExistingUserRelationships_Roles)(nil), "dictybase.user.ExistingUserRelationships.Roles")
	proto.RegisterType((*NewUserRelationships)(nil), "dictybase.user.NewUserRelationships")
	proto.RegisterType((*NewUserRelationships_Roles)(nil), "dictybase.user.NewUserRelationships.Roles")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	// Gets the specified user
	GetUser(ctx context.Context, in *dictybase_api_jsonapi1.GetRequest, opts ...grpc.CallOption) (*User, error)
	// List all users
	ListUsers(ctx context.Context, in *dictybase_api_jsonapi1.ListRequest, opts ...grpc.CallOption) (*UserCollection, error)
	// Create an user
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Update an user
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Delete an user
	DeleteUser(ctx context.Context, in *dictybase_api_jsonapi1.DeleteRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *dictybase_api_jsonapi1.GetRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/dictybase.user.UserService/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *dictybase_api_jsonapi1.ListRequest, opts ...grpc.CallOption) (*UserCollection, error) {
	out := new(UserCollection)
	err := grpc.Invoke(ctx, "/dictybase.user.UserService/ListUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/dictybase.user.UserService/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/dictybase.user.UserService/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *dictybase_api_jsonapi1.DeleteRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/dictybase.user.UserService/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	// Gets the specified user
	GetUser(context.Context, *dictybase_api_jsonapi1.GetRequest) (*User, error)
	// List all users
	ListUsers(context.Context, *dictybase_api_jsonapi1.ListRequest) (*UserCollection, error)
	// Create an user
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// Update an user
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	// Delete an user
	DeleteUser(context.Context, *dictybase_api_jsonapi1.DeleteRequest) (*google_protobuf2.Empty, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dictybase_api_jsonapi1.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.user.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*dictybase_api_jsonapi1.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dictybase_api_jsonapi1.ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.user.UserService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUsers(ctx, req.(*dictybase_api_jsonapi1.ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.user.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dictybase_api_jsonapi1.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.user.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*dictybase_api_jsonapi1.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dictybase.user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UserService_ListUsers_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 983 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcb, 0x8e, 0xdc, 0x44,
	0x17, 0x96, 0xfb, 0x32, 0xd3, 0x3e, 0x9d, 0xee, 0xff, 0x9f, 0x62, 0x84, 0x4c, 0x4f, 0x80, 0xc1,
	0x09, 0x51, 0x18, 0x81, 0x2d, 0x0d, 0x9b, 0x90, 0x41, 0x88, 0x9e, 0x0b, 0x51, 0x50, 0x12, 0x46,
	0x26, 0x11, 0x97, 0x4d, 0xab, 0xba, 0x5d, 0xe3, 0x29, 0xc6, 0xed, 0x32, 0xae, 0xea, 0x40, 0x07,
	0xb1, 0xe1, 0x11, 0xe0, 0x01, 0x90, 0x58, 0xb0, 0xe3, 0x09, 0x78, 0x03, 0xb6, 0x6c, 0x59, 0x22,
	0xf1, 0x00, 0x6c, 0xb2, 0x44, 0x75, 0xca, 0xee, 0xbb, 0x27, 0x1d, 0x21, 0xb1, 0xe9, 0x76, 0x9d,
	0xfa, 0xbe, 0xef, 0x7c, 0xe7, 0xb8, 0x7c, 0x6c, 0x80, 0x91, 0x64, 0x99, 0x97, 0x66, 0x42, 0x09,
	0xd2, 0x0e, 0xf9, 0x40, 0x8d, 0xfb, 0x54, 0x32, 0x4f, 0x47, 0x3b, 0x57, 0x23, 0x21, 0xa2, 0x98,
	0xf9, 0x34, 0xe5, 0x3e, 0x4d, 0x12, 0xa1, 0xa8, 0xe2, 0x22, 0x91, 0x06, 0xdd, 0xd9, 0xcd, 0x77,
	0x71, 0xd5, 0x1f, 0x9d, 0xf9, 0x67, 0x9c, 0xc5, 0x61, 0x6f, 0x48, 0xe5, 0x45, 0x8e, 0x78, 0x75,
	0x11, 0xa1, 0xf8, 0x90, 0x49, 0x45, 0x87, 0x69, 0x0e, 0xd8, 0x59, 0x04, 0xb0, 0x61, 0xaa, 0xc6,
	0xf9, 0xe6, 0xb5, 0x89, 0x1b, 0x34, 0xf0, 0x85, 0x14, 0x89, 0xfe, 0x4f, 0xe9, 0x38, 0x16, 0x34,
	0xbc, 0x1c, 0x94, 0xb1, 0x2f, 0x47, 0x4c, 0x2a, 0x03, 0x72, 0x7f, 0xab, 0xc0, 0xd6, 0xa3, 0x34,
	0xa4, 0x8a, 0x3d, 0x92, 0x2c, 0x0b, 0xcc, 0x1e, 0xb9, 0x0d, 0xb5, 0x90, 0x2a, 0xea, 0x58, 0xbb,
	0xd6, 0xcd, 0xe6, 0xfe, 0x0d, 0x6f, 0xbe, 0x78, 0x6f, 0x89, 0xe0, 0x1d, 0x53, 0x45, 0x03, 0xe4,
	0x90, 0x03, 0x68, 0x8e, 0x70, 0x1f, 0xcb, 0x75, 0x2a, 0x28, 0xd1, 0xf1, 0x4c, 0x39, 0x5e, 0x51,
	0x8e, 0xf7, 0x81, 0xee, 0xc8, 0x7d, 0x2a, 0x2f, 0x02, 0x30, 0x70, 0x7d, 0xdd, 0xf9, 0xd5, 0x82,
	0x9a, 0xd6, 0x22, 0x04, 0x6a, 0x6a, 0x9c, 0x32, 0x74, 0x60, 0x07, 0x78, 0x4d, 0xda, 0x50, 0xe1,
	0x21, 0x0a, 0x56, 0x83, 0x0a, 0x0f, 0xc9, 0x7b, 0x00, 0x54, 0xa9, 0x8c, 0xf7, 0x47, 0x8a, 0x49,
	0xa7, 0x8a, 0x89, 0x5e, 0x59, 0xf2, 0x2a, 0x59, 0xd6, 0x9d, 0xa0, 0x82, 0x19, 0x06, 0xf9, 0x08,
	0x5a, 0x19, 0x8b, 0xcd, 0x8d, 0x3b, 0xe7, 0xa9, 0x74, 0x6a, 0x28, 0xf1, 0xc6, 0xa2, 0xc4, 0xc9,
	0xd7, 0x5c, 0x2a, 0x9e, 0x44, 0xa6, 0xe0, 0x19, 0x42, 0x30, 0xcf, 0x77, 0x9f, 0x5a, 0xb0, 0x75,
	0x94, 0xb1, 0xe7, 0x6b, 0xe6, 0x12, 0x61, 0xa6, 0x99, 0x9d, 0x9f, 0x2f, 0xeb, 0xc7, 0x7c, 0xfd,
	0x95, 0xe7, 0xae, 0xff, 0xc3, 0xc5, 0xfa, 0x4d, 0x0b, 0xaf, 0x2f, 0x4a, 0x3c, 0x60, 0x5f, 0x3d,
	0xb3, 0xf4, 0x73, 0xa8, 0x69, 0x0c, 0x79, 0x73, 0xae, 0x58, 0x67, 0x95, 0x9b, 0x99, 0xb3, 0xb2,
	0x0f, 0xf5, 0x98, 0x27, 0x17, 0x85, 0xf9, 0xab, 0x33, 0x70, 0x9a, 0x72, 0x2f, 0x3f, 0xb2, 0xde,
	0x3d, 0x8d, 0x09, 0x0c, 0xd4, 0xfd, 0xc5, 0x82, 0xb6, 0x96, 0x39, 0x12, 0x71, 0xcc, 0x06, 0xda,
	0xc1, 0x4c, 0xd2, 0xea, 0x1a, 0x49, 0xdf, 0x9d, 0x4f, 0x7a, 0xa3, 0x24, 0xe9, 0x29, 0x8d, 0x78,
	0x82, 0x15, 0xce, 0xa6, 0x27, 0x3e, 0xd4, 0x86, 0x4c, 0xd1, 0xbc, 0x57, 0x3b, 0x25, 0xe4, 0xfb,
	0x4c, 0xa7, 0xd3, 0x40, 0xf7, 0x6f, 0x0b, 0x1a, 0x85, 0x83, 0xff, 0xe4, 0x58, 0x4f, 0x9a, 0x5a,
	0x5b, 0xbb, 0xa9, 0xcb, 0x8f, 0x42, 0xfd, 0x5f, 0x3e, 0x0a, 0x7f, 0x55, 0xcd, 0x5d, 0x9a, 0x7a,
	0x24, 0x2f, 0x03, 0x9c, 0xf1, 0x4c, 0xaa, 0x5e, 0x42, 0x87, 0x45, 0x07, 0x6c, 0x8c, 0x3c, 0xa0,
	0x43, 0x46, 0x76, 0xc0, 0x8e, 0x69, 0xb1, 0x5b, 0xc1, 0xdd, 0x86, 0x0e, 0xe0, 0xe6, 0x36, 0xd4,
	0xd9, 0x90, 0xf2, 0x18, 0xdb, 0x61, 0x07, 0x66, 0x41, 0x5c, 0xb8, 0x22, 0xb2, 0x88, 0x26, 0xfc,
	0x09, 0x66, 0xc6, 0x82, 0xed, 0x60, 0x2e, 0xa6, 0x99, 0x51, 0x26, 0x46, 0x29, 0x56, 0x64, 0x07,
	0x66, 0x41, 0xae, 0x41, 0xcb, 0x78, 0xa1, 0x61, 0x98, 0x31, 0x29, 0x9d, 0x0d, 0x43, 0xc5, 0x60,
	0xd7, 0xc4, 0xc8, 0xeb, 0xd0, 0x96, 0x6c, 0x20, 0x92, 0x70, 0x82, 0xda, 0x44, 0x54, 0xcb, 0x44,
	0x0b, 0x18, 0x81, 0xda, 0x80, 0xab, 0xb1, 0xd3, 0x30, 0xf7, 0x54, 0x5f, 0xeb, 0xac, 0x52, 0x51,
	0xc5, 0x1c, 0xdb, 0x64, 0xc5, 0x05, 0xf9, 0x3f, 0x54, 0x9f, 0xf0, 0xd4, 0x01, 0x8c, 0xe9, 0x4b,
	0xe2, 0xc0, 0xe6, 0x40, 0x8c, 0x12, 0x95, 0x8d, 0x9d, 0x26, 0x46, 0x8b, 0xa5, 0x56, 0x48, 0xcf,
	0x45, 0xc2, 0x9c, 0x2b, 0x46, 0x01, 0x17, 0xba, 0x49, 0x5c, 0xf6, 0xe8, 0x40, 0xf1, 0xc7, 0xcc,
	0x69, 0xed, 0x5a, 0x37, 0x1b, 0x41, 0x83, 0xcb, 0x2e, 0xae, 0xc9, 0x3b, 0x00, 0x03, 0x1c, 0x26,
	0x61, 0x8f, 0x2a, 0xa7, 0x5d, 0x32, 0x78, 0x1f, 0x16, 0x2f, 0x9a, 0xc0, 0xce, 0xd1, 0x5d, 0xa5,
	0xa9, 0x66, 0x0a, 0x23, 0xf5, 0x7f, 0xcf, 0xa6, 0xe6, 0xe8, 0xae, 0x72, 0x7f, 0xb4, 0xe0, 0xa5,
	0xd2, 0x63, 0x41, 0x4e, 0xa0, 0x9e, 0x89, 0x98, 0xc9, 0x7c, 0x20, 0xf8, 0x6b, 0x1f, 0x28, 0x2f,
	0xd0, 0xb4, 0xc0, 0xb0, 0x3b, 0x07, 0x50, 0xc7, 0xf5, 0xf4, 0x70, 0x5b, 0xeb, 0x4f, 0x8c, 0xef,
	0x2d, 0xd8, 0x5e, 0x35, 0xc3, 0xc8, 0xfb, 0x53, 0x73, 0x7a, 0x70, 0xec, 0xad, 0x33, 0xf8, 0xe6,
	0x7d, 0xdd, 0x2a, 0x7c, 0xf9, 0x73, 0x73, 0xaf, 0x6c, 0x2c, 0x4c, 0xa7, 0xd0, 0xfe, 0x1f, 0x55,
	0x68, 0x6a, 0xf1, 0x8f, 0x59, 0xf6, 0x98, 0x0f, 0x18, 0xf9, 0x04, 0x36, 0xef, 0x30, 0x85, 0x33,
	0xf4, 0xb5, 0x12, 0xf6, 0x1d, 0xa6, 0xf2, 0x57, 0x44, 0x67, 0x7b, 0xd5, 0x3c, 0x70, 0x5f, 0xf8,
	0xee, 0xf7, 0x3f, 0x7f, 0xa8, 0xb4, 0x48, 0xd3, 0xd7, 0x31, 0xe9, 0x7f, 0xc3, 0xc3, 0x6f, 0x49,
	0x0f, 0xec, 0x7b, 0x5c, 0xa2, 0xb2, 0x24, 0x6e, 0x69, 0xbf, 0xe4, 0x44, 0x7b, 0xe5, 0xac, 0x99,
	0x4e, 0x5b, 0xb7, 0x8d, 0x59, 0x1a, 0x64, 0xc3, 0x64, 0x21, 0x9f, 0x02, 0x4c, 0xdf, 0x61, 0x73,
	0xe6, 0x57, 0xbf, 0xdf, 0x4a, 0xcc, 0x6f, 0xa1, 0x6c, 0xd3, 0xcd, 0x65, 0x6f, 0x5b, 0x7b, 0x5a,
	0x79, 0xfa, 0xa9, 0xb1, 0xac, 0xbc, 0xf4, 0x19, 0x72, 0xb9, 0xf2, 0xfe, 0x8c, 0x72, 0x0f, 0xe0,
	0x98, 0xc5, 0x2c, 0x57, 0xbe, 0x5e, 0x76, 0xbb, 0x10, 0x52, 0x88, 0xbf, 0xb8, 0xf4, 0x3c, 0x9c,
	0xe8, 0x4f, 0xb2, 0xa2, 0xeb, 0x7b, 0xb3, 0x5d, 0x3f, 0xe4, 0x40, 0x44, 0x16, 0x2d, 0xd8, 0x39,
	0xb4, 0x75, 0xba, 0x53, 0x4d, 0x3f, 0xb5, 0x3e, 0xbf, 0x15, 0x71, 0x75, 0x3e, 0xea, 0x7b, 0x03,
	0x31, 0xf4, 0x11, 0x77, 0xa8, 0x3f, 0xd5, 0x22, 0xf1, 0x56, 0xc4, 0x12, 0x4c, 0xe1, 0x4f, 0xd8,
	0x34, 0xe5, 0x12, 0xd5, 0x0f, 0xf4, 0xcf, 0x53, 0xcb, 0xfa, 0xa9, 0xd2, 0x38, 0xbe, 0x7b, 0xf4,
	0xf0, 0xb3, 0xee, 0xe9, 0xdd, 0xfe, 0x06, 0x82, 0xdf, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0x60, 0x36, 0xa1, 0xaf, 0x0a, 0x00, 0x00,
}
