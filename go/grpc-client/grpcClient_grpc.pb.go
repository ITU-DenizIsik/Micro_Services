// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_Client

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

// GrpcClientClient is the client API for GrpcClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcClientClient interface {
	CourseToString(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error)
	ReplyCourseToString(ctx context.Context, in *Course, opts ...grpc.CallOption) (*CourseReply, error)
}

type grpcClientClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcClientClient(cc grpc.ClientConnInterface) GrpcClientClient {
	return &grpcClientClient{cc}
}

func (c *grpcClientClient) CourseToString(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/grpcClient/CourseToString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcClientClient) ReplyCourseToString(ctx context.Context, in *Course, opts ...grpc.CallOption) (*CourseReply, error) {
	out := new(CourseReply)
	err := c.cc.Invoke(ctx, "/grpcClient/ReplyCourseToString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcClientServer is the server API for GrpcClient service.
// All implementations must embed UnimplementedGrpcClientServer
// for forward compatibility
type GrpcClientServer interface {
	CourseToString(context.Context, *Course) (*Course, error)
	ReplyCourseToString(context.Context, *Course) (*CourseReply, error)
	mustEmbedUnimplementedGrpcClientServer()
}

// UnimplementedGrpcClientServer must be embedded to have forward compatible implementations.
type UnimplementedGrpcClientServer struct {
}

func (UnimplementedGrpcClientServer) CourseToString(context.Context, *Course) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CourseToString not implemented")
}
func (UnimplementedGrpcClientServer) ReplyCourseToString(context.Context, *Course) (*CourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplyCourseToString not implemented")
}
func (UnimplementedGrpcClientServer) mustEmbedUnimplementedGrpcClientServer() {}

// UnsafeGrpcClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcClientServer will
// result in compilation errors.
type UnsafeGrpcClientServer interface {
	mustEmbedUnimplementedGrpcClientServer()
}

func RegisterGrpcClientServer(s grpc.ServiceRegistrar, srv GrpcClientServer) {
	s.RegisterService(&GrpcClient_ServiceDesc, srv)
}

func _GrpcClient_CourseToString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcClientServer).CourseToString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcClient/CourseToString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcClientServer).CourseToString(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcClient_ReplyCourseToString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcClientServer).ReplyCourseToString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcClient/ReplyCourseToString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcClientServer).ReplyCourseToString(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

// GrpcClient_ServiceDesc is the grpc.ServiceDesc for GrpcClient service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GrpcClient_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcClient",
	HandlerType: (*GrpcClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CourseToString",
			Handler:    _GrpcClient_CourseToString_Handler,
		},
		{
			MethodName: "ReplyCourseToString",
			Handler:    _GrpcClient_ReplyCourseToString_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcClient.proto",
}
