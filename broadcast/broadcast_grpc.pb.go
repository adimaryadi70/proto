// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.9.1
// source: broadcast.proto

package broadcast

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

// BroadcastServiceClient is the client API for BroadcastService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BroadcastServiceClient interface {
	BroadcastMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type broadcastServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBroadcastServiceClient(cc grpc.ClientConnInterface) BroadcastServiceClient {
	return &broadcastServiceClient{cc}
}

func (c *broadcastServiceClient) BroadcastMessage(ctx context.Context, in *Message, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/broadcast.BroadcastService/BroadcastMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastServiceServer is the server API for BroadcastService service.
// All implementations must embed UnimplementedBroadcastServiceServer
// for forward compatibility
type BroadcastServiceServer interface {
	BroadcastMessage(context.Context, *Message) (*EmptyResponse, error)
	mustEmbedUnimplementedBroadcastServiceServer()
}

// UnimplementedBroadcastServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBroadcastServiceServer struct {
}

func (UnimplementedBroadcastServiceServer) BroadcastMessage(context.Context, *Message) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastMessage not implemented")
}
func (UnimplementedBroadcastServiceServer) mustEmbedUnimplementedBroadcastServiceServer() {}

// UnsafeBroadcastServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BroadcastServiceServer will
// result in compilation errors.
type UnsafeBroadcastServiceServer interface {
	mustEmbedUnimplementedBroadcastServiceServer()
}

func RegisterBroadcastServiceServer(s grpc.ServiceRegistrar, srv BroadcastServiceServer) {
	s.RegisterService(&BroadcastService_ServiceDesc, srv)
}

func _BroadcastService_BroadcastMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServiceServer).BroadcastMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/broadcast.BroadcastService/BroadcastMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServiceServer).BroadcastMessage(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// BroadcastService_ServiceDesc is the grpc.ServiceDesc for BroadcastService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BroadcastService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "broadcast.BroadcastService",
	HandlerType: (*BroadcastServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadcastMessage",
			Handler:    _BroadcastService_BroadcastMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "broadcast.proto",
}
