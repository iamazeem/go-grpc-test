// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: dns/dns.proto

package dns

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

// ResolverClient is the client API for Resolver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResolverClient interface {
	Resolve(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type resolverClient struct {
	cc grpc.ClientConnInterface
}

func NewResolverClient(cc grpc.ClientConnInterface) ResolverClient {
	return &resolverClient{cc}
}

func (c *resolverClient) Resolve(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/dns.Resolver/Resolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResolverServer is the server API for Resolver service.
// All implementations must embed UnimplementedResolverServer
// for forward compatibility
type ResolverServer interface {
	Resolve(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedResolverServer()
}

// UnimplementedResolverServer must be embedded to have forward compatible implementations.
type UnimplementedResolverServer struct {
}

func (UnimplementedResolverServer) Resolve(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resolve not implemented")
}
func (UnimplementedResolverServer) mustEmbedUnimplementedResolverServer() {}

// UnsafeResolverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResolverServer will
// result in compilation errors.
type UnsafeResolverServer interface {
	mustEmbedUnimplementedResolverServer()
}

func RegisterResolverServer(s grpc.ServiceRegistrar, srv ResolverServer) {
	s.RegisterService(&Resolver_ServiceDesc, srv)
}

func _Resolver_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResolverServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dns.Resolver/Resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResolverServer).Resolve(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Resolver_ServiceDesc is the grpc.ServiceDesc for Resolver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resolver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dns.Resolver",
	HandlerType: (*ResolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resolve",
			Handler:    _Resolver_Resolve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dns/dns.proto",
}
