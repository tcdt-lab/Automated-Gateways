// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: iop.proto

package scripts

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

// IOPClient is the client API for IOP service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IOPClient interface {
	Invoke(ctx context.Context, in *IOPInfo, opts ...grpc.CallOption) (*Response, error)
	QueryMethods(ctx context.Context, in *OutsiderNetwork, opts ...grpc.CallOption) (IOP_QueryMethodsClient, error)
}

type iOPClient struct {
	cc grpc.ClientConnInterface
}

func NewIOPClient(cc grpc.ClientConnInterface) IOPClient {
	return &iOPClient{cc}
}

func (c *iOPClient) Invoke(ctx context.Context, in *IOPInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/IOP/invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iOPClient) QueryMethods(ctx context.Context, in *OutsiderNetwork, opts ...grpc.CallOption) (IOP_QueryMethodsClient, error) {
	stream, err := c.cc.NewStream(ctx, &IOP_ServiceDesc.Streams[0], "/IOP/queryMethods", opts...)
	if err != nil {
		return nil, err
	}
	x := &iOPQueryMethodsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IOP_QueryMethodsClient interface {
	Recv() (*IOPInfo, error)
	grpc.ClientStream
}

type iOPQueryMethodsClient struct {
	grpc.ClientStream
}

func (x *iOPQueryMethodsClient) Recv() (*IOPInfo, error) {
	m := new(IOPInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IOPServer is the server API for IOP service.
// All implementations must embed UnimplementedIOPServer
// for forward compatibility
type IOPServer interface {
	Invoke(context.Context, *IOPInfo) (*Response, error)
	QueryMethods(*OutsiderNetwork, IOP_QueryMethodsServer) error
	mustEmbedUnimplementedIOPServer()
}

// UnimplementedIOPServer must be embedded to have forward compatible implementations.
type UnimplementedIOPServer struct {
}

func (UnimplementedIOPServer) Invoke(context.Context, *IOPInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (UnimplementedIOPServer) QueryMethods(*OutsiderNetwork, IOP_QueryMethodsServer) error {
	return status.Errorf(codes.Unimplemented, "method QueryMethods not implemented")
}
func (UnimplementedIOPServer) mustEmbedUnimplementedIOPServer() {}

// UnsafeIOPServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IOPServer will
// result in compilation errors.
type UnsafeIOPServer interface {
	mustEmbedUnimplementedIOPServer()
}

func RegisterIOPServer(s grpc.ServiceRegistrar, srv IOPServer) {
	s.RegisterService(&IOP_ServiceDesc, srv)
}

func _IOP_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IOPInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IOPServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/IOP/invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IOPServer).Invoke(ctx, req.(*IOPInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _IOP_QueryMethods_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OutsiderNetwork)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IOPServer).QueryMethods(m, &iOPQueryMethodsServer{stream})
}

type IOP_QueryMethodsServer interface {
	Send(*IOPInfo) error
	grpc.ServerStream
}

type iOPQueryMethodsServer struct {
	grpc.ServerStream
}

func (x *iOPQueryMethodsServer) Send(m *IOPInfo) error {
	return x.ServerStream.SendMsg(m)
}

// IOP_ServiceDesc is the grpc.ServiceDesc for IOP service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IOP_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "IOP",
	HandlerType: (*IOPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "invoke",
			Handler:    _IOP_Invoke_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "queryMethods",
			Handler:       _IOP_QueryMethods_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "iop.proto",
}
