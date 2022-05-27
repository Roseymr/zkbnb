// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package l2BlockMonitor

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

// L2BlockMonitorClient is the client API for L2BlockMonitor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type L2BlockMonitorClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type l2BlockMonitorClient struct {
	cc grpc.ClientConnInterface
}

func NewL2BlockMonitorClient(cc grpc.ClientConnInterface) L2BlockMonitorClient {
	return &l2BlockMonitorClient{cc}
}

func (c *l2BlockMonitorClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/l2BlockMonitor.L2BlockMonitor/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// L2BlockMonitorServer is the server API for L2BlockMonitor service.
// All implementations must embed UnimplementedL2BlockMonitorServer
// for forward compatibility
type L2BlockMonitorServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedL2BlockMonitorServer()
}

// UnimplementedL2BlockMonitorServer must be embedded to have forward compatible implementations.
type UnimplementedL2BlockMonitorServer struct {
}

func (UnimplementedL2BlockMonitorServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedL2BlockMonitorServer) mustEmbedUnimplementedL2BlockMonitorServer() {}

// UnsafeL2BlockMonitorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to L2BlockMonitorServer will
// result in compilation errors.
type UnsafeL2BlockMonitorServer interface {
	mustEmbedUnimplementedL2BlockMonitorServer()
}

func RegisterL2BlockMonitorServer(s grpc.ServiceRegistrar, srv L2BlockMonitorServer) {
	s.RegisterService(&L2BlockMonitor_ServiceDesc, srv)
}

func _L2BlockMonitor_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(L2BlockMonitorServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/l2BlockMonitor.L2BlockMonitor/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(L2BlockMonitorServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// L2BlockMonitor_ServiceDesc is the grpc.ServiceDesc for L2BlockMonitor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var L2BlockMonitor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "l2BlockMonitor.L2BlockMonitor",
	HandlerType: (*L2BlockMonitorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _L2BlockMonitor_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "l2BlockMonitor.proto",
}