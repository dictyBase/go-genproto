// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package identity

import (
	context "context"
	jsonapi "github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IdentityServiceClient is the client API for IdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityServiceClient interface {
	// Gets the specified identity
	GetIdentityFromProvider(ctx context.Context, in *IdentityProviderReq, opts ...grpc.CallOption) (*Identity, error)
	GetIdentity(ctx context.Context, in *jsonapi.IdRequest, opts ...grpc.CallOption) (*Identity, error)
	ExistProviderIdentity(ctx context.Context, in *IdentityProviderReq, opts ...grpc.CallOption) (*jsonapi.ExistResponse, error)
	// Create a new identity
	CreateIdentity(ctx context.Context, in *CreateIdentityReq, opts ...grpc.CallOption) (*Identity, error)
	// Delete an existing identity
	DeleteIdentity(ctx context.Context, in *jsonapi.IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Basic health check that always return success
	Healthz(ctx context.Context, in *jsonapi.HealthzIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type identityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityServiceClient(cc grpc.ClientConnInterface) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) GetIdentityFromProvider(ctx context.Context, in *IdentityProviderReq, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := c.cc.Invoke(ctx, "/dictybase.identity.IdentityService/GetIdentityFromProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) GetIdentity(ctx context.Context, in *jsonapi.IdRequest, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := c.cc.Invoke(ctx, "/dictybase.identity.IdentityService/GetIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) ExistProviderIdentity(ctx context.Context, in *IdentityProviderReq, opts ...grpc.CallOption) (*jsonapi.ExistResponse, error) {
	out := new(jsonapi.ExistResponse)
	err := c.cc.Invoke(ctx, "/dictybase.identity.IdentityService/ExistProviderIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) CreateIdentity(ctx context.Context, in *CreateIdentityReq, opts ...grpc.CallOption) (*Identity, error) {
	out := new(Identity)
	err := c.cc.Invoke(ctx, "/dictybase.identity.IdentityService/CreateIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) DeleteIdentity(ctx context.Context, in *jsonapi.IdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dictybase.identity.IdentityService/DeleteIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) Healthz(ctx context.Context, in *jsonapi.HealthzIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dictybase.identity.IdentityService/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServiceServer is the server API for IdentityService service.
// All implementations must embed UnimplementedIdentityServiceServer
// for forward compatibility
type IdentityServiceServer interface {
	// Gets the specified identity
	GetIdentityFromProvider(context.Context, *IdentityProviderReq) (*Identity, error)
	GetIdentity(context.Context, *jsonapi.IdRequest) (*Identity, error)
	ExistProviderIdentity(context.Context, *IdentityProviderReq) (*jsonapi.ExistResponse, error)
	// Create a new identity
	CreateIdentity(context.Context, *CreateIdentityReq) (*Identity, error)
	// Delete an existing identity
	DeleteIdentity(context.Context, *jsonapi.IdRequest) (*emptypb.Empty, error)
	// Basic health check that always return success
	Healthz(context.Context, *jsonapi.HealthzIdRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedIdentityServiceServer()
}

// UnimplementedIdentityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityServiceServer struct {
}

func (UnimplementedIdentityServiceServer) GetIdentityFromProvider(context.Context, *IdentityProviderReq) (*Identity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentityFromProvider not implemented")
}
func (UnimplementedIdentityServiceServer) GetIdentity(context.Context, *jsonapi.IdRequest) (*Identity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentity not implemented")
}
func (UnimplementedIdentityServiceServer) ExistProviderIdentity(context.Context, *IdentityProviderReq) (*jsonapi.ExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistProviderIdentity not implemented")
}
func (UnimplementedIdentityServiceServer) CreateIdentity(context.Context, *CreateIdentityReq) (*Identity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIdentity not implemented")
}
func (UnimplementedIdentityServiceServer) DeleteIdentity(context.Context, *jsonapi.IdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIdentity not implemented")
}
func (UnimplementedIdentityServiceServer) Healthz(context.Context, *jsonapi.HealthzIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthz not implemented")
}
func (UnimplementedIdentityServiceServer) mustEmbedUnimplementedIdentityServiceServer() {}

// UnsafeIdentityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServiceServer will
// result in compilation errors.
type UnsafeIdentityServiceServer interface {
	mustEmbedUnimplementedIdentityServiceServer()
}

func RegisterIdentityServiceServer(s grpc.ServiceRegistrar, srv IdentityServiceServer) {
	s.RegisterService(&IdentityService_ServiceDesc, srv)
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
	in := new(jsonapi.IdRequest)
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
		return srv.(IdentityServiceServer).GetIdentity(ctx, req.(*jsonapi.IdRequest))
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
	in := new(jsonapi.IdRequest)
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
		return srv.(IdentityServiceServer).DeleteIdentity(ctx, req.(*jsonapi.IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(jsonapi.HealthzIdRequest)
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
		return srv.(IdentityServiceServer).Healthz(ctx, req.(*jsonapi.HealthzIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityService_ServiceDesc is the grpc.ServiceDesc for IdentityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityService_ServiceDesc = grpc.ServiceDesc{
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
	Metadata: "dictybase/identity/identity.proto",
}