// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SlbUkEntitiesClient is the client API for SlbUkEntities service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SlbUkEntitiesClient interface {
	GetEntities(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetEntityResponse, error)
	GetEntitiesStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SlbUkEntities_GetEntitiesStreamClient, error)
}

type slbUkEntitiesClient struct {
	cc grpc.ClientConnInterface
}

func NewSlbUkEntitiesClient(cc grpc.ClientConnInterface) SlbUkEntitiesClient {
	return &slbUkEntitiesClient{cc}
}

func (c *slbUkEntitiesClient) GetEntities(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetEntityResponse, error) {
	out := new(GetEntityResponse)
	err := c.cc.Invoke(ctx, "/slbuk.grpc.SlbUkEntities/GetEntities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slbUkEntitiesClient) GetEntitiesStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SlbUkEntities_GetEntitiesStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SlbUkEntities_serviceDesc.Streams[0], "/slbuk.grpc.SlbUkEntities/GetEntitiesStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &slbUkEntitiesGetEntitiesStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SlbUkEntities_GetEntitiesStreamClient interface {
	Recv() (*Entity, error)
	grpc.ClientStream
}

type slbUkEntitiesGetEntitiesStreamClient struct {
	grpc.ClientStream
}

func (x *slbUkEntitiesGetEntitiesStreamClient) Recv() (*Entity, error) {
	m := new(Entity)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SlbUkEntitiesServer is the server API for SlbUkEntities service.
// All implementations must embed UnimplementedSlbUkEntitiesServer
// for forward compatibility
type SlbUkEntitiesServer interface {
	GetEntities(context.Context, *Empty) (*GetEntityResponse, error)
	GetEntitiesStream(*Empty, SlbUkEntities_GetEntitiesStreamServer) error
	mustEmbedUnimplementedSlbUkEntitiesServer()
}

// UnimplementedSlbUkEntitiesServer must be embedded to have forward compatible implementations.
type UnimplementedSlbUkEntitiesServer struct {
}

func (UnimplementedSlbUkEntitiesServer) GetEntities(context.Context, *Empty) (*GetEntityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntities not implemented")
}
func (UnimplementedSlbUkEntitiesServer) GetEntitiesStream(*Empty, SlbUkEntities_GetEntitiesStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEntitiesStream not implemented")
}
func (UnimplementedSlbUkEntitiesServer) mustEmbedUnimplementedSlbUkEntitiesServer() {}

// UnsafeSlbUkEntitiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SlbUkEntitiesServer will
// result in compilation errors.
type UnsafeSlbUkEntitiesServer interface {
	mustEmbedUnimplementedSlbUkEntitiesServer()
}

func RegisterSlbUkEntitiesServer(s *grpc.Server, srv SlbUkEntitiesServer) {
	s.RegisterService(&_SlbUkEntities_serviceDesc, srv)
}

func _SlbUkEntities_GetEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlbUkEntitiesServer).GetEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/slbuk.grpc.SlbUkEntities/GetEntities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlbUkEntitiesServer).GetEntities(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SlbUkEntities_GetEntitiesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SlbUkEntitiesServer).GetEntitiesStream(m, &slbUkEntitiesGetEntitiesStreamServer{stream})
}

type SlbUkEntities_GetEntitiesStreamServer interface {
	Send(*Entity) error
	grpc.ServerStream
}

type slbUkEntitiesGetEntitiesStreamServer struct {
	grpc.ServerStream
}

func (x *slbUkEntitiesGetEntitiesStreamServer) Send(m *Entity) error {
	return x.ServerStream.SendMsg(m)
}

var _SlbUkEntities_serviceDesc = grpc.ServiceDesc{
	ServiceName: "slbuk.grpc.SlbUkEntities",
	HandlerType: (*SlbUkEntitiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntities",
			Handler:    _SlbUkEntities_GetEntities_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEntitiesStream",
			Handler:       _SlbUkEntities_GetEntitiesStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc.proto",
}