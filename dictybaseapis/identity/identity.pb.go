// Code generated by protoc-gen-go. DO NOT EDIT.
// source: identity.proto

/*
Package identity is a generated protocol buffer package.

It is generated from these files:
	identity.proto

It has these top-level messages:
	IdentityAttributes
	IdentityData
	Identity
	IdentityProviderReq
	CreateIdentityReq
	NewIdentityAttributes
*/
package identity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Definition for various fields
type IdentityAttributes struct {
	// An unique identifier provided by the third party.
	// Generally it's an email id, however it could be something else specifically
	// provided by an provider.
	Identifier string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// Name of the provider, for example, orcid, google, facebook etc.
	Provider string `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	// The id of the user to which this identity is connected.
	// This id could be used to fetch a complete user response from the user service
	UserId int64 `protobuf:"varint,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Timestamp for creation and update
	CreatedAt *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *IdentityAttributes) Reset()                    { *m = IdentityAttributes{} }
func (m *IdentityAttributes) String() string            { return proto.CompactTextString(m) }
func (*IdentityAttributes) ProtoMessage()               {}
func (*IdentityAttributes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IdentityAttributes) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *IdentityAttributes) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *IdentityAttributes) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *IdentityAttributes) GetCreatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *IdentityAttributes) GetUpdatedAt() *google_protobuf1.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type IdentityData struct {
	// The resource name
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Unique id
	Id         int64                        `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Attributes *IdentityAttributes          `protobuf:"bytes,3,opt,name=attributes" json:"attributes,omitempty"`
	Links      *dictybase_api_jsonapi.Links `protobuf:"bytes,4,opt,name=links" json:"links,omitempty"`
}

func (m *IdentityData) Reset()                    { *m = IdentityData{} }
func (m *IdentityData) String() string            { return proto.CompactTextString(m) }
func (*IdentityData) ProtoMessage()               {}
func (*IdentityData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IdentityData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *IdentityData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IdentityData) GetAttributes() *IdentityAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *IdentityData) GetLinks() *dictybase_api_jsonapi.Links {
	if m != nil {
		return m.Links
	}
	return nil
}

type Identity struct {
	Data  *IdentityData                `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	Links *dictybase_api_jsonapi.Links `protobuf:"bytes,4,opt,name=links" json:"links,omitempty"`
}

func (m *Identity) Reset()                    { *m = Identity{} }
func (m *Identity) String() string            { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()               {}
func (*Identity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Identity) GetData() *IdentityData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Identity) GetLinks() *dictybase_api_jsonapi.Links {
	if m != nil {
		return m.Links
	}
	return nil
}

type IdentityProviderReq struct {
	// An unique identifier provided by the third party.
	// Generally it's an email id, however it could be something else specifically
	// provided by an provider.
	Identifier string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// Name of the provider, for example, orcid, google, facebook etc.
	Provider string `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
}

func (m *IdentityProviderReq) Reset()                    { *m = IdentityProviderReq{} }
func (m *IdentityProviderReq) String() string            { return proto.CompactTextString(m) }
func (*IdentityProviderReq) ProtoMessage()               {}
func (*IdentityProviderReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IdentityProviderReq) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *IdentityProviderReq) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

type CreateIdentityReq struct {
	Data *CreateIdentityReq_Data `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *CreateIdentityReq) Reset()                    { *m = CreateIdentityReq{} }
func (m *CreateIdentityReq) String() string            { return proto.CompactTextString(m) }
func (*CreateIdentityReq) ProtoMessage()               {}
func (*CreateIdentityReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateIdentityReq) GetData() *CreateIdentityReq_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateIdentityReq_Data struct {
	// resource name
	Type       string                 `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Attributes *NewIdentityAttributes `protobuf:"bytes,2,opt,name=attributes" json:"attributes,omitempty"`
}

func (m *CreateIdentityReq_Data) Reset()                    { *m = CreateIdentityReq_Data{} }
func (m *CreateIdentityReq_Data) String() string            { return proto.CompactTextString(m) }
func (*CreateIdentityReq_Data) ProtoMessage()               {}
func (*CreateIdentityReq_Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *CreateIdentityReq_Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateIdentityReq_Data) GetAttributes() *NewIdentityAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type NewIdentityAttributes struct {
	// An unique identifier provided by the third party.
	// Generally it's an email id, however it could be something else specifically
	// provided by an provider.
	Identifier string `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	// Name of the provider, for example, orcid, google, facebook etc.
	Provider string `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	// The id of the user to which this identity is connected.
	// This id could be used to fetch a complete user response from the user service
	UserId int64 `protobuf:"varint,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *NewIdentityAttributes) Reset()                    { *m = NewIdentityAttributes{} }
func (m *NewIdentityAttributes) String() string            { return proto.CompactTextString(m) }
func (*NewIdentityAttributes) ProtoMessage()               {}
func (*NewIdentityAttributes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *NewIdentityAttributes) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *NewIdentityAttributes) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *NewIdentityAttributes) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*IdentityAttributes)(nil), "dictybase.identity.IdentityAttributes")
	proto.RegisterType((*IdentityData)(nil), "dictybase.identity.IdentityData")
	proto.RegisterType((*Identity)(nil), "dictybase.identity.Identity")
	proto.RegisterType((*IdentityProviderReq)(nil), "dictybase.identity.IdentityProviderReq")
	proto.RegisterType((*CreateIdentityReq)(nil), "dictybase.identity.CreateIdentityReq")
	proto.RegisterType((*CreateIdentityReq_Data)(nil), "dictybase.identity.CreateIdentityReq.Data")
	proto.RegisterType((*NewIdentityAttributes)(nil), "dictybase.identity.NewIdentityAttributes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for IdentityService service

type IdentityServiceClient interface {
	// Gets the specified identity
	GetIdentityFromProvider(ctx context.Context, in *IdentityProviderReq, opts ...grpc.CallOption) (*Identity, error)
	GetIdentity(ctx context.Context, in *dictybase_api_jsonapi1.IdRequest, opts ...grpc.CallOption) (*Identity, error)
	ExistProviderIdentity(ctx context.Context, in *IdentityProviderReq, opts ...grpc.CallOption) (*dictybase_api_jsonapi1.ExistResponse, error)
	// Create a new identity
	CreateIdentity(ctx context.Context, in *CreateIdentityReq, opts ...grpc.CallOption) (*Identity, error)
	// Delete an existing identity
	DeleteIdentity(ctx context.Context, in *dictybase_api_jsonapi1.IdRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	// Basic health check that always return success
	Healthz(ctx context.Context, in *dictybase_api_jsonapi1.HealthzIdRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type identityServiceClient struct {
	cc *grpc.ClientConn
}

func NewIdentityServiceClient(cc *grpc.ClientConn) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) GetIdentityFromProvider(ctx context.Context, in *IdentityProviderReq, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := grpc.Invoke(ctx, "/dictybase.identity.IdentityService/GetIdentityFromProvider", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) GetIdentity(ctx context.Context, in *dictybase_api_jsonapi1.IdRequest, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := grpc.Invoke(ctx, "/dictybase.identity.IdentityService/GetIdentity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) ExistProviderIdentity(ctx context.Context, in *IdentityProviderReq, opts ...grpc.CallOption) (*dictybase_api_jsonapi1.ExistResponse, error) {
	out := new(dictybase_api_jsonapi1.ExistResponse)
	err := grpc.Invoke(ctx, "/dictybase.identity.IdentityService/ExistProviderIdentity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) CreateIdentity(ctx context.Context, in *CreateIdentityReq, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := grpc.Invoke(ctx, "/dictybase.identity.IdentityService/CreateIdentity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) DeleteIdentity(ctx context.Context, in *dictybase_api_jsonapi1.IdRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/dictybase.identity.IdentityService/DeleteIdentity", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) Healthz(ctx context.Context, in *dictybase_api_jsonapi1.HealthzIdRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/dictybase.identity.IdentityService/Healthz", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IdentityService service

type IdentityServiceServer interface {
	// Gets the specified identity
	GetIdentityFromProvider(context.Context, *IdentityProviderReq) (*Identity, error)
	GetIdentity(context.Context, *dictybase_api_jsonapi1.IdRequest) (*Identity, error)
	ExistProviderIdentity(context.Context, *IdentityProviderReq) (*dictybase_api_jsonapi1.ExistResponse, error)
	// Create a new identity
	CreateIdentity(context.Context, *CreateIdentityReq) (*Identity, error)
	// Delete an existing identity
	DeleteIdentity(context.Context, *dictybase_api_jsonapi1.IdRequest) (*google_protobuf2.Empty, error)
	// Basic health check that always return success
	Healthz(context.Context, *dictybase_api_jsonapi1.HealthzIdRequest) (*google_protobuf2.Empty, error)
}

func RegisterIdentityServiceServer(s *grpc.Server, srv IdentityServiceServer) {
	s.RegisterService(&_IdentityService_serviceDesc, srv)
}

func _IdentityService_GetIdentityFromProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityProviderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).GetIdentityFromProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.identity.IdentityService/GetIdentityFromProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).GetIdentityFromProvider(ctx, req.(*IdentityProviderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_GetIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dictybase_api_jsonapi1.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).GetIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.identity.IdentityService/GetIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).GetIdentity(ctx, req.(*dictybase_api_jsonapi1.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_ExistProviderIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdentityProviderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).ExistProviderIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.identity.IdentityService/ExistProviderIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).ExistProviderIdentity(ctx, req.(*IdentityProviderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_CreateIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIdentityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).CreateIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.identity.IdentityService/CreateIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).CreateIdentity(ctx, req.(*CreateIdentityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_DeleteIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dictybase_api_jsonapi1.IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).DeleteIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.identity.IdentityService/DeleteIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).DeleteIdentity(ctx, req.(*dictybase_api_jsonapi1.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dictybase_api_jsonapi1.HealthzIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.identity.IdentityService/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).Healthz(ctx, req.(*dictybase_api_jsonapi1.HealthzIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IdentityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dictybase.identity.IdentityService",
	HandlerType: (*IdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIdentityFromProvider",
			Handler:    _IdentityService_GetIdentityFromProvider_Handler,
		},
		{
			MethodName: "GetIdentity",
			Handler:    _IdentityService_GetIdentity_Handler,
		},
		{
			MethodName: "ExistProviderIdentity",
			Handler:    _IdentityService_ExistProviderIdentity_Handler,
		},
		{
			MethodName: "CreateIdentity",
			Handler:    _IdentityService_CreateIdentity_Handler,
		},
		{
			MethodName: "DeleteIdentity",
			Handler:    _IdentityService_DeleteIdentity_Handler,
		},
		{
			MethodName: "Healthz",
			Handler:    _IdentityService_Healthz_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity.proto",
}

func init() { proto.RegisterFile("identity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x69, 0xda, 0xa6, 0x13, 0x08, 0x65, 0x51, 0xdb, 0x28, 0x54, 0x10, 0xb9, 0x14, 0x4a,
	0x11, 0xb6, 0x14, 0xb8, 0x00, 0x12, 0x22, 0xfd, 0x83, 0x48, 0x08, 0x15, 0xd3, 0x0b, 0x08, 0xa9,
	0xda, 0x64, 0xa7, 0xce, 0x42, 0xe2, 0x75, 0xed, 0x4d, 0x21, 0x54, 0xbd, 0x70, 0xe7, 0xc4, 0x1b,
	0x70, 0xe1, 0xc6, 0x03, 0xf0, 0x1a, 0x3c, 0x00, 0x17, 0x1e, 0x82, 0x23, 0xf2, 0xda, 0x9b, 0x3a,
	0x24, 0x4d, 0x2a, 0x10, 0xa7, 0xd8, 0xde, 0x6f, 0xbe, 0x6f, 0xbe, 0xf9, 0x26, 0x36, 0x14, 0x38,
	0x43, 0x4f, 0x72, 0xd9, 0xb5, 0xfc, 0x40, 0x48, 0x41, 0x08, 0xe3, 0x0d, 0xd9, 0xad, 0xd3, 0x10,
	0x2d, 0x7d, 0x52, 0x5a, 0x74, 0x85, 0x70, 0x5b, 0x68, 0x53, 0x9f, 0xdb, 0xd4, 0xf3, 0x84, 0xa4,
	0x92, 0x0b, 0x2f, 0x8c, 0x2b, 0x4a, 0x57, 0x92, 0x53, 0x75, 0x57, 0xef, 0xec, 0xd9, 0x92, 0xb7,
	0x31, 0x94, 0xb4, 0xed, 0x27, 0x80, 0x4b, 0x7f, 0x02, 0xb0, 0xed, 0x6b, 0xbd, 0xd2, 0x52, 0x4f,
	0x4f, 0xd1, 0xbf, 0x0e, 0x85, 0x17, 0xfd, 0xfa, 0xb4, 0xdb, 0x12, 0x94, 0x8d, 0x06, 0x05, 0xb8,
	0xdf, 0xc1, 0x50, 0xc6, 0x20, 0xf3, 0x87, 0x01, 0xa4, 0x96, 0xb4, 0x5c, 0x95, 0x32, 0xe0, 0xf5,
	0x8e, 0xc4, 0x90, 0x5c, 0x06, 0x88, 0x8d, 0xec, 0x71, 0x0c, 0x8a, 0x46, 0xd9, 0x58, 0x99, 0x71,
	0x52, 0x4f, 0x48, 0x09, 0x72, 0x7e, 0x20, 0x0e, 0x38, 0xc3, 0xa0, 0x98, 0x51, 0xa7, 0xbd, 0x7b,
	0xb2, 0x00, 0xd3, 0x9d, 0x10, 0x83, 0x5d, 0xce, 0x8a, 0x13, 0x65, 0x63, 0x65, 0xc2, 0x99, 0x8a,
	0x6e, 0x6b, 0x8c, 0xdc, 0x05, 0x68, 0x04, 0x48, 0x25, 0xb2, 0x5d, 0x2a, 0x8b, 0x93, 0x65, 0x63,
	0x25, 0x5f, 0x29, 0x59, 0xb1, 0x4f, 0x4b, 0xfb, 0xb4, 0x76, 0xf4, 0x20, 0x9c, 0x99, 0x04, 0x5d,
	0x95, 0x51, 0x69, 0xc7, 0x67, 0xba, 0x74, 0x6a, 0x7c, 0x69, 0x82, 0xae, 0x4a, 0xf3, 0xab, 0x01,
	0x67, 0xb5, 0xc3, 0x0d, 0x2a, 0x29, 0x21, 0x90, 0x95, 0x5d, 0x1f, 0x13, 0x57, 0xea, 0x9a, 0x14,
	0x20, 0xc3, 0x99, 0x72, 0x32, 0xe1, 0x64, 0x38, 0x23, 0x5b, 0x00, 0xb4, 0x37, 0x0d, 0x65, 0x23,
	0x5f, 0xb9, 0x66, 0x0d, 0xa6, 0x6c, 0x0d, 0xce, 0xce, 0x49, 0x55, 0x92, 0x0a, 0x4c, 0xb6, 0xb8,
	0xf7, 0x26, 0x2c, 0x66, 0x15, 0xc5, 0x62, 0x8a, 0x82, 0xfa, 0xdc, 0x4a, 0x32, 0xb1, 0x9e, 0x44,
	0x18, 0x27, 0x86, 0x9a, 0x12, 0x72, 0x9a, 0x95, 0xdc, 0x81, 0x2c, 0xa3, 0x92, 0xaa, 0x5e, 0xf3,
	0x95, 0xf2, 0xa8, 0x0e, 0x22, 0x6f, 0x8e, 0x42, 0xff, 0x95, 0xea, 0x33, 0xb8, 0xa8, 0x99, 0xb6,
	0x93, 0x24, 0x1d, 0xdc, 0xff, 0x97, 0x45, 0x30, 0xbf, 0x19, 0x70, 0x61, 0x5d, 0x45, 0xa8, 0x99,
	0x23, 0xc6, 0x07, 0x7d, 0x96, 0x56, 0x87, 0x59, 0x1a, 0x28, 0xb2, 0x8e, 0xcd, 0x95, 0x10, 0xb2,
	0x27, 0xc6, 0x58, 0xeb, 0x8b, 0x2d, 0xa3, 0x14, 0x6e, 0x0c, 0x53, 0x78, 0x8a, 0x6f, 0x47, 0x27,
	0x67, 0xb6, 0x60, 0x6e, 0x28, 0xe8, 0xbf, 0xfc, 0x35, 0x2a, 0x5f, 0x26, 0xe1, 0xbc, 0xd6, 0x7a,
	0x8e, 0xc1, 0x01, 0x6f, 0x20, 0xf9, 0x68, 0xc0, 0xc2, 0x23, 0x94, 0xfa, 0xf1, 0x56, 0x20, 0xda,
	0x3a, 0x19, 0x72, 0x7d, 0xd4, 0x26, 0xa4, 0xf2, 0x2b, 0x2d, 0x8e, 0x02, 0x9a, 0x37, 0x3f, 0x7c,
	0xff, 0xf9, 0x29, 0xb3, 0x4c, 0x96, 0xec, 0xe4, 0x8c, 0x63, 0x68, 0x1f, 0xea, 0x76, 0x8f, 0xec,
	0xc3, 0x63, 0x5f, 0x47, 0xc4, 0x85, 0x7c, 0xaa, 0x1d, 0x52, 0x3e, 0x61, 0xab, 0x6a, 0xcc, 0x89,
	0xdf, 0x30, 0x63, 0xb4, 0x8b, 0x4a, 0x9b, 0x90, 0xd9, 0x3e, 0x6d, 0xce, 0x8e, 0x48, 0x13, 0xe6,
	0x36, 0xdf, 0xf1, 0x50, 0x6a, 0x1f, 0x3d, 0xc9, 0x53, 0xbb, 0xbe, 0x7a, 0x42, 0x6f, 0x8a, 0xd6,
	0xc1, 0xd0, 0x17, 0x5e, 0x88, 0xe6, 0x19, 0xd2, 0x86, 0x42, 0xff, 0xae, 0x91, 0xe5, 0x53, 0xed,
	0xe3, 0x18, 0x6b, 0xf3, 0xca, 0xda, 0xac, 0x99, 0x4f, 0x59, 0xbb, 0x67, 0xac, 0x12, 0x06, 0x85,
	0x0d, 0x6c, 0x61, 0x4a, 0x6e, 0xfc, 0x10, 0xe7, 0x07, 0xde, 0x72, 0x9b, 0xd1, 0x87, 0x40, 0x8f,
	0x6f, 0x75, 0x70, 0x7c, 0xaf, 0x60, 0xfa, 0x31, 0xd2, 0x96, 0x6c, 0xbe, 0xef, 0x1b, 0x58, 0x9a,
	0x3e, 0x39, 0x1f, 0xaf, 0x32, 0xab, 0x54, 0x80, 0xe4, 0xec, 0x66, 0x5c, 0xb2, 0xd6, 0x85, 0x79,
	0x11, 0xb8, 0x43, 0xec, 0xaf, 0x9d, 0x4b, 0x25, 0x21, 0xc5, 0xb6, 0xf1, 0xf2, 0xa1, 0xcb, 0x65,
	0xb3, 0x53, 0xb7, 0x1a, 0xa2, 0x6d, 0x2b, 0xfc, 0x5a, 0xf4, 0x2d, 0x72, 0xc5, 0x2d, 0x17, 0x3d,
	0xa5, 0x63, 0xf7, 0x58, 0xa8, 0xcf, 0x43, 0x6d, 0xa4, 0x7b, 0x5f, 0x5f, 0xfc, 0x32, 0x8c, 0xcf,
	0x99, 0xdc, 0x46, 0x6d, 0x7d, 0xe7, 0x45, 0x75, 0xbb, 0x56, 0x9f, 0x52, 0x45, 0xb7, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x93, 0x4d, 0x79, 0xee, 0x7e, 0x07, 0x00, 0x00,
}
