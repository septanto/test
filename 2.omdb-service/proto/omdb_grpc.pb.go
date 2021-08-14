// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package omdb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OmdbServiceClient is the client API for OmdbService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OmdbServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SeachReply, error)
	Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailReply, error)
}

type omdbServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOmdbServiceClient(cc grpc.ClientConnInterface) OmdbServiceClient {
	return &omdbServiceClient{cc}
}

func (c *omdbServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SeachReply, error) {
	out := new(SeachReply)
	err := c.cc.Invoke(ctx, "/omdb.OmdbService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omdbServiceClient) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailReply, error) {
	out := new(DetailReply)
	err := c.cc.Invoke(ctx, "/omdb.OmdbService/Detail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OmdbServiceServer is the server API for OmdbService service.
// All implementations must embed UnimplementedOmdbServiceServer
// for forward compatibility
type OmdbServiceServer interface {
	Search(context.Context, *SearchRequest) (*SeachReply, error)
	Detail(context.Context, *DetailRequest) (*DetailReply, error)
	mustEmbedUnimplementedOmdbServiceServer()
}

// UnimplementedOmdbServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOmdbServiceServer struct {
}

func (UnimplementedOmdbServiceServer) Search(context.Context, *SearchRequest) (*SeachReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedOmdbServiceServer) Detail(context.Context, *DetailRequest) (*DetailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedOmdbServiceServer) mustEmbedUnimplementedOmdbServiceServer() {}

// UnsafeOmdbServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OmdbServiceServer will
// result in compilation errors.
type UnsafeOmdbServiceServer interface {
	mustEmbedUnimplementedOmdbServiceServer()
}

func RegisterOmdbServiceServer(s grpc.ServiceRegistrar, srv OmdbServiceServer) {
	s.RegisterService(&OmdbService_ServiceDesc, srv)
}

func _OmdbService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmdbServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omdb.OmdbService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmdbServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmdbService_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmdbServiceServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omdb.OmdbService/Detail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmdbServiceServer).Detail(ctx, req.(*DetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OmdbService_ServiceDesc is the grpc.ServiceDesc for OmdbService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OmdbService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "omdb.OmdbService",
	HandlerType: (*OmdbServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _OmdbService_Search_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _OmdbService_Detail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/omdb.proto",
}
