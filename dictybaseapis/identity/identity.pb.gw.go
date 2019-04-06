// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: identity.proto

/*
Package identity is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package identity

import (
	"context"
	"io"
	"net/http"

	"github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_IdentityService_GetIdentityFromProvider_0(ctx context.Context, marshaler runtime.Marshaler, client IdentityServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq IdentityProviderReq
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["provider"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "provider")
	}

	protoReq.Provider, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "provider", err)
	}

	val, ok = pathParams["identifier"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "identifier")
	}

	protoReq.Identifier, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "identifier", err)
	}

	msg, err := client.GetIdentityFromProvider(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_IdentityService_GetIdentity_0(ctx context.Context, marshaler runtime.Marshaler, client IdentityServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq jsonapi.IdRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.Int64(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.GetIdentity(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_IdentityService_CreateIdentity_0(ctx context.Context, marshaler runtime.Marshaler, client IdentityServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateIdentityReq
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateIdentity(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_IdentityService_DeleteIdentity_0(ctx context.Context, marshaler runtime.Marshaler, client IdentityServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq jsonapi.IdRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.Int64(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.DeleteIdentity(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_IdentityService_Healthz_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_IdentityService_Healthz_0(ctx context.Context, marshaler runtime.Marshaler, client IdentityServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq jsonapi.HealthzIdRequest
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_IdentityService_Healthz_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Healthz(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterIdentityServiceHandlerFromEndpoint is same as RegisterIdentityServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterIdentityServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterIdentityServiceHandler(ctx, mux, conn)
}

// RegisterIdentityServiceHandler registers the http handlers for service IdentityService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterIdentityServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterIdentityServiceHandlerClient(ctx, mux, NewIdentityServiceClient(conn))
}

// RegisterIdentityServiceHandlerClient registers the http handlers for service IdentityService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "IdentityServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "IdentityServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "IdentityServiceClient" to call the correct interceptors.
func RegisterIdentityServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client IdentityServiceClient) error {

	mux.Handle("GET", pattern_IdentityService_GetIdentityFromProvider_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_IdentityService_GetIdentityFromProvider_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_IdentityService_GetIdentityFromProvider_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_IdentityService_GetIdentity_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_IdentityService_GetIdentity_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_IdentityService_GetIdentity_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_IdentityService_CreateIdentity_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_IdentityService_CreateIdentity_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_IdentityService_CreateIdentity_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_IdentityService_DeleteIdentity_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_IdentityService_DeleteIdentity_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_IdentityService_DeleteIdentity_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_IdentityService_Healthz_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_IdentityService_Healthz_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_IdentityService_Healthz_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_IdentityService_GetIdentityFromProvider_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1, 1, 0, 4, 1, 5, 2}, []string{"identities", "provider", "identifier"}, ""))

	pattern_IdentityService_GetIdentity_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1}, []string{"identities", "id"}, ""))

	pattern_IdentityService_CreateIdentity_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"identities"}, ""))

	pattern_IdentityService_DeleteIdentity_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1}, []string{"identities", "id"}, ""))

	pattern_IdentityService_Healthz_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"healthz"}, ""))
)

var (
	forward_IdentityService_GetIdentityFromProvider_0 = runtime.ForwardResponseMessage

	forward_IdentityService_GetIdentity_0 = runtime.ForwardResponseMessage

	forward_IdentityService_CreateIdentity_0 = runtime.ForwardResponseMessage

	forward_IdentityService_DeleteIdentity_0 = runtime.ForwardResponseMessage

	forward_IdentityService_Healthz_0 = runtime.ForwardResponseMessage
)
