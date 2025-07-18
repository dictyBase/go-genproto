// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: dictybase/feature_annotation/feature_annotation.proto

package feature_annotation

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
	FeatureAnnotationService_CreateFeatureAnnotation_FullMethodName          = "/dictybase.feature_annotation.FeatureAnnotationService/CreateFeatureAnnotation"
	FeatureAnnotationService_GetFeatureAnnotation_FullMethodName             = "/dictybase.feature_annotation.FeatureAnnotationService/GetFeatureAnnotation"
	FeatureAnnotationService_GetFeatureAnnotationByName_FullMethodName       = "/dictybase.feature_annotation.FeatureAnnotationService/GetFeatureAnnotationByName"
	FeatureAnnotationService_UpdateFeatureAnnotation_FullMethodName          = "/dictybase.feature_annotation.FeatureAnnotationService/UpdateFeatureAnnotation"
	FeatureAnnotationService_DeleteFeatureAnnotation_FullMethodName          = "/dictybase.feature_annotation.FeatureAnnotationService/DeleteFeatureAnnotation"
	FeatureAnnotationService_AddTag_FullMethodName                           = "/dictybase.feature_annotation.FeatureAnnotationService/AddTag"
	FeatureAnnotationService_AddTags_FullMethodName                          = "/dictybase.feature_annotation.FeatureAnnotationService/AddTags"
	FeatureAnnotationService_SetTags_FullMethodName                          = "/dictybase.feature_annotation.FeatureAnnotationService/SetTags"
	FeatureAnnotationService_UpdateTag_FullMethodName                        = "/dictybase.feature_annotation.FeatureAnnotationService/UpdateTag"
	FeatureAnnotationService_RemoveTag_FullMethodName                        = "/dictybase.feature_annotation.FeatureAnnotationService/RemoveTag"
	FeatureAnnotationService_RemoveTags_FullMethodName                       = "/dictybase.feature_annotation.FeatureAnnotationService/RemoveTags"
	FeatureAnnotationService_ListFeatureAnnotationsByPubmedId_FullMethodName = "/dictybase.feature_annotation.FeatureAnnotationService/ListFeatureAnnotationsByPubmedId"
	FeatureAnnotationService_ListFeatureAnnotationsByDOI_FullMethodName      = "/dictybase.feature_annotation.FeatureAnnotationService/ListFeatureAnnotationsByDOI"
)

// FeatureAnnotationServiceClient is the client API for FeatureAnnotationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeatureAnnotationServiceClient interface {
	// Create a feature annotation
	CreateFeatureAnnotation(ctx context.Context, in *NewFeatureAnnotation, opts ...grpc.CallOption) (*FeatureAnnotation, error)
	// Retrieves the specified feature annotation
	GetFeatureAnnotation(ctx context.Context, in *FeatureAnnotationId, opts ...grpc.CallOption) (*FeatureAnnotation, error)
	// Retrieves the specified feature annotation by its name
	GetFeatureAnnotationByName(ctx context.Context, in *FeatureName, opts ...grpc.CallOption) (*FeatureAnnotation, error)
	// Update an existing feature annotation. Any included tags will replace all
	// existing tags (identical behavior to SetTags)
	UpdateFeatureAnnotation(ctx context.Context, in *FeatureAnnotationUpdate, opts ...grpc.CallOption) (*FeatureAnnotation, error)
	// Delete an existing feature annotation
	DeleteFeatureAnnotation(ctx context.Context, in *DeleteFeatureAnnotationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Add tag to an existing feature annotation
	AddTag(ctx context.Context, in *AddTagRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error)
	// Add multiple tags to an existing feature annotation
	AddTags(ctx context.Context, in *AddTagsRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error)
	// Replace all tags for a feature annotation (idempotent full-update)
	SetTags(ctx context.Context, in *SetTagsRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error)
	// Deprecated: Do not use.
	// Update an existing tag in a feature annotation
	UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error)
	// Deprecated: Do not use.
	// Remove a tag from a feature annotation
	RemoveTag(ctx context.Context, in *RemoveTagRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error)
	// Remove a tag from a feature annotation
	RemoveTags(ctx context.Context, in *RemoveTagsRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error)
	// Retrieves a list of feature annotations by PubMed ID
	ListFeatureAnnotationsByPubmedId(ctx context.Context, in *PubmedId, opts ...grpc.CallOption) (*FeatureAnnotationCollection, error)
	// Retrieves a list of feature annotations by DOI (Digital Object Identifier)
	ListFeatureAnnotationsByDOI(ctx context.Context, in *DOI, opts ...grpc.CallOption) (*FeatureAnnotationCollection, error)
}

type featureAnnotationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeatureAnnotationServiceClient(cc grpc.ClientConnInterface) FeatureAnnotationServiceClient {
	return &featureAnnotationServiceClient{cc}
}

func (c *featureAnnotationServiceClient) CreateFeatureAnnotation(ctx context.Context, in *NewFeatureAnnotation, opts ...grpc.CallOption) (*FeatureAnnotation, error) {
	out := new(FeatureAnnotation)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_CreateFeatureAnnotation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureAnnotationServiceClient) GetFeatureAnnotation(ctx context.Context, in *FeatureAnnotationId, opts ...grpc.CallOption) (*FeatureAnnotation, error) {
	out := new(FeatureAnnotation)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_GetFeatureAnnotation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureAnnotationServiceClient) GetFeatureAnnotationByName(ctx context.Context, in *FeatureName, opts ...grpc.CallOption) (*FeatureAnnotation, error) {
	out := new(FeatureAnnotation)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_GetFeatureAnnotationByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureAnnotationServiceClient) UpdateFeatureAnnotation(ctx context.Context, in *FeatureAnnotationUpdate, opts ...grpc.CallOption) (*FeatureAnnotation, error) {
	out := new(FeatureAnnotation)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_UpdateFeatureAnnotation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureAnnotationServiceClient) DeleteFeatureAnnotation(ctx context.Context, in *DeleteFeatureAnnotationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_DeleteFeatureAnnotation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureAnnotationServiceClient) AddTag(ctx context.Context, in *AddTagRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error) {
	out := new(FeatureAnnotation)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_AddTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureAnnotationServiceClient) AddTags(ctx context.Context, in *AddTagsRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error) {
	out := new(FeatureAnnotation)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_AddTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureAnnotationServiceClient) SetTags(ctx context.Context, in *SetTagsRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error) {
	out := new(FeatureAnnotation)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_SetTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *featureAnnotationServiceClient) UpdateTag(ctx context.Context, in *UpdateTagRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error) {
	out := new(FeatureAnnotation)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_UpdateTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *featureAnnotationServiceClient) RemoveTag(ctx context.Context, in *RemoveTagRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error) {
	out := new(FeatureAnnotation)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_RemoveTag_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureAnnotationServiceClient) RemoveTags(ctx context.Context, in *RemoveTagsRequest, opts ...grpc.CallOption) (*FeatureAnnotation, error) {
	out := new(FeatureAnnotation)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_RemoveTags_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureAnnotationServiceClient) ListFeatureAnnotationsByPubmedId(ctx context.Context, in *PubmedId, opts ...grpc.CallOption) (*FeatureAnnotationCollection, error) {
	out := new(FeatureAnnotationCollection)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_ListFeatureAnnotationsByPubmedId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *featureAnnotationServiceClient) ListFeatureAnnotationsByDOI(ctx context.Context, in *DOI, opts ...grpc.CallOption) (*FeatureAnnotationCollection, error) {
	out := new(FeatureAnnotationCollection)
	err := c.cc.Invoke(ctx, FeatureAnnotationService_ListFeatureAnnotationsByDOI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeatureAnnotationServiceServer is the server API for FeatureAnnotationService service.
// All implementations must embed UnimplementedFeatureAnnotationServiceServer
// for forward compatibility
type FeatureAnnotationServiceServer interface {
	// Create a feature annotation
	CreateFeatureAnnotation(context.Context, *NewFeatureAnnotation) (*FeatureAnnotation, error)
	// Retrieves the specified feature annotation
	GetFeatureAnnotation(context.Context, *FeatureAnnotationId) (*FeatureAnnotation, error)
	// Retrieves the specified feature annotation by its name
	GetFeatureAnnotationByName(context.Context, *FeatureName) (*FeatureAnnotation, error)
	// Update an existing feature annotation. Any included tags will replace all
	// existing tags (identical behavior to SetTags)
	UpdateFeatureAnnotation(context.Context, *FeatureAnnotationUpdate) (*FeatureAnnotation, error)
	// Delete an existing feature annotation
	DeleteFeatureAnnotation(context.Context, *DeleteFeatureAnnotationRequest) (*emptypb.Empty, error)
	// Add tag to an existing feature annotation
	AddTag(context.Context, *AddTagRequest) (*FeatureAnnotation, error)
	// Add multiple tags to an existing feature annotation
	AddTags(context.Context, *AddTagsRequest) (*FeatureAnnotation, error)
	// Replace all tags for a feature annotation (idempotent full-update)
	SetTags(context.Context, *SetTagsRequest) (*FeatureAnnotation, error)
	// Deprecated: Do not use.
	// Update an existing tag in a feature annotation
	UpdateTag(context.Context, *UpdateTagRequest) (*FeatureAnnotation, error)
	// Deprecated: Do not use.
	// Remove a tag from a feature annotation
	RemoveTag(context.Context, *RemoveTagRequest) (*FeatureAnnotation, error)
	// Remove a tag from a feature annotation
	RemoveTags(context.Context, *RemoveTagsRequest) (*FeatureAnnotation, error)
	// Retrieves a list of feature annotations by PubMed ID
	ListFeatureAnnotationsByPubmedId(context.Context, *PubmedId) (*FeatureAnnotationCollection, error)
	// Retrieves a list of feature annotations by DOI (Digital Object Identifier)
	ListFeatureAnnotationsByDOI(context.Context, *DOI) (*FeatureAnnotationCollection, error)
	mustEmbedUnimplementedFeatureAnnotationServiceServer()
}

// UnimplementedFeatureAnnotationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeatureAnnotationServiceServer struct {
}

func (UnimplementedFeatureAnnotationServiceServer) CreateFeatureAnnotation(context.Context, *NewFeatureAnnotation) (*FeatureAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeatureAnnotation not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) GetFeatureAnnotation(context.Context, *FeatureAnnotationId) (*FeatureAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeatureAnnotation not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) GetFeatureAnnotationByName(context.Context, *FeatureName) (*FeatureAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeatureAnnotationByName not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) UpdateFeatureAnnotation(context.Context, *FeatureAnnotationUpdate) (*FeatureAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFeatureAnnotation not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) DeleteFeatureAnnotation(context.Context, *DeleteFeatureAnnotationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeatureAnnotation not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) AddTag(context.Context, *AddTagRequest) (*FeatureAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTag not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) AddTags(context.Context, *AddTagsRequest) (*FeatureAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTags not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) SetTags(context.Context, *SetTagsRequest) (*FeatureAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTags not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) UpdateTag(context.Context, *UpdateTagRequest) (*FeatureAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTag not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) RemoveTag(context.Context, *RemoveTagRequest) (*FeatureAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTag not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) RemoveTags(context.Context, *RemoveTagsRequest) (*FeatureAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTags not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) ListFeatureAnnotationsByPubmedId(context.Context, *PubmedId) (*FeatureAnnotationCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeatureAnnotationsByPubmedId not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) ListFeatureAnnotationsByDOI(context.Context, *DOI) (*FeatureAnnotationCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeatureAnnotationsByDOI not implemented")
}
func (UnimplementedFeatureAnnotationServiceServer) mustEmbedUnimplementedFeatureAnnotationServiceServer() {
}

// UnsafeFeatureAnnotationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeatureAnnotationServiceServer will
// result in compilation errors.
type UnsafeFeatureAnnotationServiceServer interface {
	mustEmbedUnimplementedFeatureAnnotationServiceServer()
}

func RegisterFeatureAnnotationServiceServer(s grpc.ServiceRegistrar, srv FeatureAnnotationServiceServer) {
	s.RegisterService(&FeatureAnnotationService_ServiceDesc, srv)
}

func _FeatureAnnotationService_CreateFeatureAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewFeatureAnnotation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).CreateFeatureAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_CreateFeatureAnnotation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).CreateFeatureAnnotation(ctx, req.(*NewFeatureAnnotation))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_GetFeatureAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeatureAnnotationId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).GetFeatureAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_GetFeatureAnnotation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).GetFeatureAnnotation(ctx, req.(*FeatureAnnotationId))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_GetFeatureAnnotationByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeatureName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).GetFeatureAnnotationByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_GetFeatureAnnotationByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).GetFeatureAnnotationByName(ctx, req.(*FeatureName))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_UpdateFeatureAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeatureAnnotationUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).UpdateFeatureAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_UpdateFeatureAnnotation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).UpdateFeatureAnnotation(ctx, req.(*FeatureAnnotationUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_DeleteFeatureAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeatureAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).DeleteFeatureAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_DeleteFeatureAnnotation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).DeleteFeatureAnnotation(ctx, req.(*DeleteFeatureAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_AddTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).AddTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_AddTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).AddTag(ctx, req.(*AddTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_AddTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).AddTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_AddTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).AddTags(ctx, req.(*AddTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_SetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).SetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_SetTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).SetTags(ctx, req.(*SetTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_UpdateTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).UpdateTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_UpdateTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).UpdateTag(ctx, req.(*UpdateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_RemoveTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).RemoveTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_RemoveTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).RemoveTag(ctx, req.(*RemoveTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_RemoveTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).RemoveTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_RemoveTags_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).RemoveTags(ctx, req.(*RemoveTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_ListFeatureAnnotationsByPubmedId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubmedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).ListFeatureAnnotationsByPubmedId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_ListFeatureAnnotationsByPubmedId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).ListFeatureAnnotationsByPubmedId(ctx, req.(*PubmedId))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeatureAnnotationService_ListFeatureAnnotationsByDOI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DOI)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeatureAnnotationServiceServer).ListFeatureAnnotationsByDOI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FeatureAnnotationService_ListFeatureAnnotationsByDOI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeatureAnnotationServiceServer).ListFeatureAnnotationsByDOI(ctx, req.(*DOI))
	}
	return interceptor(ctx, in, info, handler)
}

// FeatureAnnotationService_ServiceDesc is the grpc.ServiceDesc for FeatureAnnotationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeatureAnnotationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dictybase.feature_annotation.FeatureAnnotationService",
	HandlerType: (*FeatureAnnotationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeatureAnnotation",
			Handler:    _FeatureAnnotationService_CreateFeatureAnnotation_Handler,
		},
		{
			MethodName: "GetFeatureAnnotation",
			Handler:    _FeatureAnnotationService_GetFeatureAnnotation_Handler,
		},
		{
			MethodName: "GetFeatureAnnotationByName",
			Handler:    _FeatureAnnotationService_GetFeatureAnnotationByName_Handler,
		},
		{
			MethodName: "UpdateFeatureAnnotation",
			Handler:    _FeatureAnnotationService_UpdateFeatureAnnotation_Handler,
		},
		{
			MethodName: "DeleteFeatureAnnotation",
			Handler:    _FeatureAnnotationService_DeleteFeatureAnnotation_Handler,
		},
		{
			MethodName: "AddTag",
			Handler:    _FeatureAnnotationService_AddTag_Handler,
		},
		{
			MethodName: "AddTags",
			Handler:    _FeatureAnnotationService_AddTags_Handler,
		},
		{
			MethodName: "SetTags",
			Handler:    _FeatureAnnotationService_SetTags_Handler,
		},
		{
			MethodName: "UpdateTag",
			Handler:    _FeatureAnnotationService_UpdateTag_Handler,
		},
		{
			MethodName: "RemoveTag",
			Handler:    _FeatureAnnotationService_RemoveTag_Handler,
		},
		{
			MethodName: "RemoveTags",
			Handler:    _FeatureAnnotationService_RemoveTags_Handler,
		},
		{
			MethodName: "ListFeatureAnnotationsByPubmedId",
			Handler:    _FeatureAnnotationService_ListFeatureAnnotationsByPubmedId_Handler,
		},
		{
			MethodName: "ListFeatureAnnotationsByDOI",
			Handler:    _FeatureAnnotationService_ListFeatureAnnotationsByDOI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dictybase/feature_annotation/feature_annotation.proto",
}
