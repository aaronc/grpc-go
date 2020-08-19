// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_lookup_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RouteLookupServiceClient is the client API for RouteLookupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteLookupServiceClient interface {
	// Lookup returns a target for a single key.
	RouteLookup(ctx context.Context, in *RouteLookupRequest, opts ...grpc.CallOption) (*RouteLookupResponse, error)
}

type routeLookupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteLookupServiceClient(cc grpc.ClientConnInterface) RouteLookupServiceClient {
	return &routeLookupServiceClient{cc}
}

func (c *routeLookupServiceClient) RouteLookup(ctx context.Context, in *RouteLookupRequest, opts ...grpc.CallOption) (*RouteLookupResponse, error) {
	out := new(RouteLookupResponse)
	err := c.cc.Invoke(ctx, "/grpc.lookup.v1.RouteLookupService/RouteLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteLookupServiceServer is the server API for RouteLookupService service.
// All implementations should embed UnimplementedRouteLookupServiceServer
// for forward compatibility
type RouteLookupServiceServer interface {
	// Lookup returns a target for a single key.
	RouteLookup(context.Context, *RouteLookupRequest) (*RouteLookupResponse, error)
}

// UnimplementedRouteLookupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRouteLookupServiceServer struct {
}

func (*UnimplementedRouteLookupServiceServer) RouteLookup(context.Context, *RouteLookupRequest) (*RouteLookupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteLookup not implemented")
}

func RegisterRouteLookupServiceServer(s grpc.ServiceRegistrar, srv RouteLookupServiceServer) {
	s.RegisterService(&_RouteLookupService_serviceDesc, srv)
}

func _RouteLookupService_RouteLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteLookupServiceServer).RouteLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.lookup.v1.RouteLookupService/RouteLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteLookupServiceServer).RouteLookup(ctx, req.(*RouteLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteLookupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.lookup.v1.RouteLookupService",
	HandlerType: (*RouteLookupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RouteLookup",
			Handler:    _RouteLookupService_RouteLookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/lookup/v1/rls.proto",
}
