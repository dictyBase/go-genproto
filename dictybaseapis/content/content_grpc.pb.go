// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: dictybase/content/content.proto

package content

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ContentService_GetContentBySlug_FullMethodName = "/dictybase.content.ContentService/GetContentBySlug"
	ContentService_GetContent_FullMethodName       = "/dictybase.content.ContentService/GetContent"
	ContentService_StoreContent_FullMethodName     = "/dictybase.content.ContentService/StoreContent"
	ContentService_UpdateContent_FullMethodName    = "/dictybase.content.ContentService/UpdateContent"
	ContentService_DeleteContent_FullMethodName    = "/dictybase.content.ContentService/DeleteContent"
	ContentService_ListContents_FullMethodName     = "/dictybase.content.ContentService/ListContents"
)

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentServiceClient interface {
	// Gets the content of specified page(slug)
	GetContentBySlug(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*Content, error)
	GetContent(ctx context.Context, in *ContentIdRequest, opts ...grpc.CallOption) (*Content, error)
	// Store the content of a new page(slug)
	StoreContent(ctx context.Context, in *StoreContentRequest, opts ...grpc.CallOption) (*Content, error)
	// Update the content of an existing page
	UpdateContent(ctx context.Context, in *UpdateContentRequest, opts ...grpc.CallOption) (*Content, error)
	// Delete an existing page along with its content
	DeleteContent(ctx context.Context, in *ContentIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List contents using pagination, ten entries are retrieved by default
	ListContents(ctx context.Context, in *ListParameters, opts ...grpc.CallOption) (*ContentCollection, error)
}

type contentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentServiceClient(cc grpc.ClientConnInterface) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) GetContentBySlug(ctx context.Context, in *ContentRequest, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, ContentService_GetContentBySlug_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetContent(ctx context.Context, in *ContentIdRequest, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, ContentService_GetContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) StoreContent(ctx context.Context, in *StoreContentRequest, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, ContentService_StoreContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) UpdateContent(ctx context.Context, in *UpdateContentRequest, opts ...grpc.CallOption) (*Content, error) {
	out := new(Content)
	err := c.cc.Invoke(ctx, ContentService_UpdateContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteContent(ctx context.Context, in *ContentIdRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ContentService_DeleteContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) ListContents(ctx context.Context, in *ListParameters, opts ...grpc.CallOption) (*ContentCollection, error) {
	out := new(ContentCollection)
	err := c.cc.Invoke(ctx, ContentService_ListContents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
// All implementations must embed UnimplementedContentServiceServer
// for forward compatibility
type ContentServiceServer interface {
	// Gets the content of specified page(slug)
	GetContentBySlug(context.Context, *ContentRequest) (*Content, error)
	GetContent(context.Context, *ContentIdRequest) (*Content, error)
	// Store the content of a new page(slug)
	StoreContent(context.Context, *StoreContentRequest) (*Content, error)
	// Update the content of an existing page
	UpdateContent(context.Context, *UpdateContentRequest) (*Content, error)
	// Delete an existing page along with its content
	DeleteContent(context.Context, *ContentIdRequest) (*emptypb.Empty, error)
	// List contents using pagination, ten entries are retrieved by default
	ListContents(context.Context, *ListParameters) (*ContentCollection, error)
	mustEmbedUnimplementedContentServiceServer()
}

// UnimplementedContentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentServiceServer struct {
}

func (UnimplementedContentServiceServer) GetContentBySlug(context.Context, *ContentRequest) (*Content, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContentBySlug not implemented")
}
func (UnimplementedContentServiceServer) GetContent(context.Context, *ContentIdRequest) (*Content, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContent not implemented")
}
func (UnimplementedContentServiceServer) StoreContent(context.Context, *StoreContentRequest) (*Content, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreContent not implemented")
}
func (UnimplementedContentServiceServer) UpdateContent(context.Context, *UpdateContentRequest) (*Content, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContent not implemented")
}
func (UnimplementedContentServiceServer) DeleteContent(context.Context, *ContentIdRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContent not implemented")
}
func (UnimplementedContentServiceServer) ListContents(context.Context, *ListParameters) (*ContentCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContents not implemented")
}
func (UnimplementedContentServiceServer) mustEmbedUnimplementedContentServiceServer() {}

// UnsafeContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServiceServer will
// result in compilation errors.
type UnsafeContentServiceServer interface {
	mustEmbedUnimplementedContentServiceServer()
}

func RegisterContentServiceServer(s grpc.ServiceRegistrar, srv ContentServiceServer) {
	s.RegisterService(&ContentService_ServiceDesc, srv)
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
		FullMethod: ContentService_GetContentBySlug_FullMethodName,
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
		FullMethod: ContentService_GetContent_FullMethodName,
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
		FullMethod: ContentService_StoreContent_FullMethodName,
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
		FullMethod: ContentService_UpdateContent_FullMethodName,
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
		FullMethod: ContentService_DeleteContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteContent(ctx, req.(*ContentIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_ListContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).ListContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_ListContents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).ListContents(ctx, req.(*ListParameters))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentService_ServiceDesc is the grpc.ServiceDesc for ContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentService_ServiceDesc = grpc.ServiceDesc{
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
			MethodName: "ListContents",
			Handler:    _ContentService_ListContents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dictybase/content/content.proto",
}
