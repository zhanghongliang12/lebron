// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: cart_rpc.proto

package cart_rpc

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

const (
	CartRpc_Ping_FullMethodName = "/cart_rpc.Cart_rpc/Ping"
)

// CartRpcClient is the client API for CartRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartRpcClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type cartRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewCartRpcClient(cc grpc.ClientConnInterface) CartRpcClient {
	return &cartRpcClient{cc}
}

func (c *cartRpcClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, CartRpc_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartRpcServer is the server API for CartRpc service.
// All implementations must embed UnimplementedCartRpcServer
// for forward compatibility
type CartRpcServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedCartRpcServer()
}

// UnimplementedCartRpcServer must be embedded to have forward compatible implementations.
type UnimplementedCartRpcServer struct {
}

func (UnimplementedCartRpcServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCartRpcServer) mustEmbedUnimplementedCartRpcServer() {}

// UnsafeCartRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartRpcServer will
// result in compilation errors.
type UnsafeCartRpcServer interface {
	mustEmbedUnimplementedCartRpcServer()
}

func RegisterCartRpcServer(s grpc.ServiceRegistrar, srv CartRpcServer) {
	s.RegisterService(&CartRpc_ServiceDesc, srv)
}

func _CartRpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartRpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartRpc_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartRpcServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// CartRpc_ServiceDesc is the grpc.ServiceDesc for CartRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart_rpc.Cart_rpc",
	HandlerType: (*CartRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CartRpc_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart_rpc.proto",
}
