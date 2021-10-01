// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stock

import (
	context "context"
	upload "github.com/dictyBase/go-genproto/dictybaseapis/api/upload"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StockServiceClient is the client API for StockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockServiceClient interface {
	// Retrieves strain by ID
	GetStrain(ctx context.Context, in *StockId, opts ...grpc.CallOption) (*Strain, error)
	// Retrieves stock by ID
	GetPlasmid(ctx context.Context, in *StockId, opts ...grpc.CallOption) (*Plasmid, error)
	// Create a new strain
	CreateStrain(ctx context.Context, in *NewStrain, opts ...grpc.CallOption) (*Strain, error)
	// Create a new plasmid
	CreatePlasmid(ctx context.Context, in *NewPlasmid, opts ...grpc.CallOption) (*Plasmid, error)
	// Update an existing strain
	UpdateStrain(ctx context.Context, in *StrainUpdate, opts ...grpc.CallOption) (*Strain, error)
	// Update an existing plasmid
	UpdatePlasmid(ctx context.Context, in *PlasmidUpdate, opts ...grpc.CallOption) (*Plasmid, error)
	// Remove an existing stock
	RemoveStock(ctx context.Context, in *StockId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// List strains using pagination, ten entries are retrieved by default
	ListStrains(ctx context.Context, in *StockParameters, opts ...grpc.CallOption) (*StrainCollection, error)
	// List strains using strain id without any pagination
	ListStrainsByIds(ctx context.Context, in *StockIdList, opts ...grpc.CallOption) (*StrainList, error)
	// List plasmids using pagination, ten entries are retrieved by default
	ListPlasmids(ctx context.Context, in *StockParameters, opts ...grpc.CallOption) (*PlasmidCollection, error)
	// Load existing strain
	LoadStrain(ctx context.Context, in *ExistingStrain, opts ...grpc.CallOption) (*Strain, error)
	// Load existing plasmid
	LoadPlasmid(ctx context.Context, in *ExistingPlasmid, opts ...grpc.CallOption) (*Plasmid, error)
	// Upload obojson formatted file through client side streaming
	OboJSONFileUpload(ctx context.Context, opts ...grpc.CallOption) (StockService_OboJSONFileUploadClient, error)
}

type stockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStockServiceClient(cc grpc.ClientConnInterface) StockServiceClient {
	return &stockServiceClient{cc}
}

func (c *stockServiceClient) GetStrain(ctx context.Context, in *StockId, opts ...grpc.CallOption) (*Strain, error) {
	out := new(Strain)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/GetStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) GetPlasmid(ctx context.Context, in *StockId, opts ...grpc.CallOption) (*Plasmid, error) {
	out := new(Plasmid)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/GetPlasmid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) CreateStrain(ctx context.Context, in *NewStrain, opts ...grpc.CallOption) (*Strain, error) {
	out := new(Strain)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/CreateStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) CreatePlasmid(ctx context.Context, in *NewPlasmid, opts ...grpc.CallOption) (*Plasmid, error) {
	out := new(Plasmid)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/CreatePlasmid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) UpdateStrain(ctx context.Context, in *StrainUpdate, opts ...grpc.CallOption) (*Strain, error) {
	out := new(Strain)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/UpdateStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) UpdatePlasmid(ctx context.Context, in *PlasmidUpdate, opts ...grpc.CallOption) (*Plasmid, error) {
	out := new(Plasmid)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/UpdatePlasmid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) RemoveStock(ctx context.Context, in *StockId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/RemoveStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) ListStrains(ctx context.Context, in *StockParameters, opts ...grpc.CallOption) (*StrainCollection, error) {
	out := new(StrainCollection)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/ListStrains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) ListStrainsByIds(ctx context.Context, in *StockIdList, opts ...grpc.CallOption) (*StrainList, error) {
	out := new(StrainList)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/ListStrainsByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) ListPlasmids(ctx context.Context, in *StockParameters, opts ...grpc.CallOption) (*PlasmidCollection, error) {
	out := new(PlasmidCollection)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/ListPlasmids", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) LoadStrain(ctx context.Context, in *ExistingStrain, opts ...grpc.CallOption) (*Strain, error) {
	out := new(Strain)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/LoadStrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) LoadPlasmid(ctx context.Context, in *ExistingPlasmid, opts ...grpc.CallOption) (*Plasmid, error) {
	out := new(Plasmid)
	err := c.cc.Invoke(ctx, "/dictybase.stock.StockService/LoadPlasmid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockServiceClient) OboJSONFileUpload(ctx context.Context, opts ...grpc.CallOption) (StockService_OboJSONFileUploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &StockService_ServiceDesc.Streams[0], "/dictybase.stock.StockService/OboJSONFileUpload", opts...)
	if err != nil {
		return nil, err
	}
	x := &stockServiceOboJSONFileUploadClient{stream}
	return x, nil
}

type StockService_OboJSONFileUploadClient interface {
	Send(*upload.FileUploadRequest) error
	CloseAndRecv() (*upload.FileUploadResponse, error)
	grpc.ClientStream
}

type stockServiceOboJSONFileUploadClient struct {
	grpc.ClientStream
}

func (x *stockServiceOboJSONFileUploadClient) Send(m *upload.FileUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stockServiceOboJSONFileUploadClient) CloseAndRecv() (*upload.FileUploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(upload.FileUploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StockServiceServer is the server API for StockService service.
// All implementations must embed UnimplementedStockServiceServer
// for forward compatibility
type StockServiceServer interface {
	// Retrieves strain by ID
	GetStrain(context.Context, *StockId) (*Strain, error)
	// Retrieves stock by ID
	GetPlasmid(context.Context, *StockId) (*Plasmid, error)
	// Create a new strain
	CreateStrain(context.Context, *NewStrain) (*Strain, error)
	// Create a new plasmid
	CreatePlasmid(context.Context, *NewPlasmid) (*Plasmid, error)
	// Update an existing strain
	UpdateStrain(context.Context, *StrainUpdate) (*Strain, error)
	// Update an existing plasmid
	UpdatePlasmid(context.Context, *PlasmidUpdate) (*Plasmid, error)
	// Remove an existing stock
	RemoveStock(context.Context, *StockId) (*emptypb.Empty, error)
	// List strains using pagination, ten entries are retrieved by default
	ListStrains(context.Context, *StockParameters) (*StrainCollection, error)
	// List strains using strain id without any pagination
	ListStrainsByIds(context.Context, *StockIdList) (*StrainList, error)
	// List plasmids using pagination, ten entries are retrieved by default
	ListPlasmids(context.Context, *StockParameters) (*PlasmidCollection, error)
	// Load existing strain
	LoadStrain(context.Context, *ExistingStrain) (*Strain, error)
	// Load existing plasmid
	LoadPlasmid(context.Context, *ExistingPlasmid) (*Plasmid, error)
	// Upload obojson formatted file through client side streaming
	OboJSONFileUpload(StockService_OboJSONFileUploadServer) error
	mustEmbedUnimplementedStockServiceServer()
}

// UnimplementedStockServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStockServiceServer struct {
}

func (UnimplementedStockServiceServer) GetStrain(context.Context, *StockId) (*Strain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStrain not implemented")
}
func (UnimplementedStockServiceServer) GetPlasmid(context.Context, *StockId) (*Plasmid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlasmid not implemented")
}
func (UnimplementedStockServiceServer) CreateStrain(context.Context, *NewStrain) (*Strain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStrain not implemented")
}
func (UnimplementedStockServiceServer) CreatePlasmid(context.Context, *NewPlasmid) (*Plasmid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlasmid not implemented")
}
func (UnimplementedStockServiceServer) UpdateStrain(context.Context, *StrainUpdate) (*Strain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStrain not implemented")
}
func (UnimplementedStockServiceServer) UpdatePlasmid(context.Context, *PlasmidUpdate) (*Plasmid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlasmid not implemented")
}
func (UnimplementedStockServiceServer) RemoveStock(context.Context, *StockId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveStock not implemented")
}
func (UnimplementedStockServiceServer) ListStrains(context.Context, *StockParameters) (*StrainCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStrains not implemented")
}
func (UnimplementedStockServiceServer) ListStrainsByIds(context.Context, *StockIdList) (*StrainList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStrainsByIds not implemented")
}
func (UnimplementedStockServiceServer) ListPlasmids(context.Context, *StockParameters) (*PlasmidCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlasmids not implemented")
}
func (UnimplementedStockServiceServer) LoadStrain(context.Context, *ExistingStrain) (*Strain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadStrain not implemented")
}
func (UnimplementedStockServiceServer) LoadPlasmid(context.Context, *ExistingPlasmid) (*Plasmid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadPlasmid not implemented")
}
func (UnimplementedStockServiceServer) OboJSONFileUpload(StockService_OboJSONFileUploadServer) error {
	return status.Errorf(codes.Unimplemented, "method OboJSONFileUpload not implemented")
}
func (UnimplementedStockServiceServer) mustEmbedUnimplementedStockServiceServer() {}

// UnsafeStockServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockServiceServer will
// result in compilation errors.
type UnsafeStockServiceServer interface {
	mustEmbedUnimplementedStockServiceServer()
}

func RegisterStockServiceServer(s grpc.ServiceRegistrar, srv StockServiceServer) {
	s.RegisterService(&StockService_ServiceDesc, srv)
}

func _StockService_GetStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).GetStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/GetStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).GetStrain(ctx, req.(*StockId))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_GetPlasmid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).GetPlasmid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/GetPlasmid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).GetPlasmid(ctx, req.(*StockId))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_CreateStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewStrain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).CreateStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/CreateStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).CreateStrain(ctx, req.(*NewStrain))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_CreatePlasmid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPlasmid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).CreatePlasmid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/CreatePlasmid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).CreatePlasmid(ctx, req.(*NewPlasmid))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_UpdateStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StrainUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).UpdateStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/UpdateStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).UpdateStrain(ctx, req.(*StrainUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_UpdatePlasmid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlasmidUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).UpdatePlasmid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/UpdatePlasmid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).UpdatePlasmid(ctx, req.(*PlasmidUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_RemoveStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).RemoveStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/RemoveStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).RemoveStock(ctx, req.(*StockId))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_ListStrains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).ListStrains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/ListStrains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).ListStrains(ctx, req.(*StockParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_ListStrainsByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockIdList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).ListStrainsByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/ListStrainsByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).ListStrainsByIds(ctx, req.(*StockIdList))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_ListPlasmids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).ListPlasmids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/ListPlasmids",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).ListPlasmids(ctx, req.(*StockParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_LoadStrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistingStrain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).LoadStrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/LoadStrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).LoadStrain(ctx, req.(*ExistingStrain))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_LoadPlasmid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistingPlasmid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServiceServer).LoadPlasmid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dictybase.stock.StockService/LoadPlasmid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServiceServer).LoadPlasmid(ctx, req.(*ExistingPlasmid))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockService_OboJSONFileUpload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StockServiceServer).OboJSONFileUpload(&stockServiceOboJSONFileUploadServer{stream})
}

type StockService_OboJSONFileUploadServer interface {
	SendAndClose(*upload.FileUploadResponse) error
	Recv() (*upload.FileUploadRequest, error)
	grpc.ServerStream
}

type stockServiceOboJSONFileUploadServer struct {
	grpc.ServerStream
}

func (x *stockServiceOboJSONFileUploadServer) SendAndClose(m *upload.FileUploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stockServiceOboJSONFileUploadServer) Recv() (*upload.FileUploadRequest, error) {
	m := new(upload.FileUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StockService_ServiceDesc is the grpc.ServiceDesc for StockService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dictybase.stock.StockService",
	HandlerType: (*StockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStrain",
			Handler:    _StockService_GetStrain_Handler,
		},
		{
			MethodName: "GetPlasmid",
			Handler:    _StockService_GetPlasmid_Handler,
		},
		{
			MethodName: "CreateStrain",
			Handler:    _StockService_CreateStrain_Handler,
		},
		{
			MethodName: "CreatePlasmid",
			Handler:    _StockService_CreatePlasmid_Handler,
		},
		{
			MethodName: "UpdateStrain",
			Handler:    _StockService_UpdateStrain_Handler,
		},
		{
			MethodName: "UpdatePlasmid",
			Handler:    _StockService_UpdatePlasmid_Handler,
		},
		{
			MethodName: "RemoveStock",
			Handler:    _StockService_RemoveStock_Handler,
		},
		{
			MethodName: "ListStrains",
			Handler:    _StockService_ListStrains_Handler,
		},
		{
			MethodName: "ListStrainsByIds",
			Handler:    _StockService_ListStrainsByIds_Handler,
		},
		{
			MethodName: "ListPlasmids",
			Handler:    _StockService_ListPlasmids_Handler,
		},
		{
			MethodName: "LoadStrain",
			Handler:    _StockService_LoadStrain_Handler,
		},
		{
			MethodName: "LoadPlasmid",
			Handler:    _StockService_LoadPlasmid_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OboJSONFileUpload",
			Handler:       _StockService_OboJSONFileUpload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "dictybase/stock/stock.proto",
}
