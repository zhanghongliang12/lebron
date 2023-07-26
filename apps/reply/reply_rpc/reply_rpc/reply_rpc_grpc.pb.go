// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: reply_rpc.proto

package reply_rpc

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
	ReplyRpc_Ping_FullMethodName = "/reply_rpc.Reply_rpc/Ping"
)

// ReplyRpcClient is the client API for ReplyRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplyRpcClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type replyRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewReplyRpcClient(cc grpc.ClientConnInterface) ReplyRpcClient {
	return &replyRpcClient{cc}
}

func (c *replyRpcClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ReplyRpc_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplyRpcServer is the server API for ReplyRpc service.
// All implementations must embed UnimplementedReplyRpcServer
// for forward compatibility
type ReplyRpcServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedReplyRpcServer()
}

// UnimplementedReplyRpcServer must be embedded to have forward compatible implementations.
type UnimplementedReplyRpcServer struct {
}

func (UnimplementedReplyRpcServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedReplyRpcServer) mustEmbedUnimplementedReplyRpcServer() {}

// UnsafeReplyRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplyRpcServer will
// result in compilation errors.
type UnsafeReplyRpcServer interface {
	mustEmbedUnimplementedReplyRpcServer()
}

func RegisterReplyRpcServer(s grpc.ServiceRegistrar, srv ReplyRpcServer) {
	s.RegisterService(&ReplyRpc_ServiceDesc, srv)
}

func _ReplyRpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplyRpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReplyRpc_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplyRpcServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ReplyRpc_ServiceDesc is the grpc.ServiceDesc for ReplyRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReplyRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reply_rpc.Reply_rpc",
	HandlerType: (*ReplyRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ReplyRpc_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reply_rpc.proto",
}
