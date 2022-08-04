// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: proto/service.proto

package pb

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

// MemoviaServiceClient is the client API for MemoviaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemoviaServiceClient interface {
	CreateMemovia(ctx context.Context, in *CreateMemoviaRequest, opts ...grpc.CallOption) (*Memovia, error)
}

type memoviaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemoviaServiceClient(cc grpc.ClientConnInterface) MemoviaServiceClient {
	return &memoviaServiceClient{cc}
}

func (c *memoviaServiceClient) CreateMemovia(ctx context.Context, in *CreateMemoviaRequest, opts ...grpc.CallOption) (*Memovia, error) {
	out := new(Memovia)
	err := c.cc.Invoke(ctx, "/infrastracture.MemoviaService/CreateMemovia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemoviaServiceServer is the server API for MemoviaService service.
// All implementations must embed UnimplementedMemoviaServiceServer
// for forward compatibility
type MemoviaServiceServer interface {
	CreateMemovia(context.Context, *CreateMemoviaRequest) (*Memovia, error)
	mustEmbedUnimplementedMemoviaServiceServer()
}

// UnimplementedMemoviaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMemoviaServiceServer struct {
}

func (UnimplementedMemoviaServiceServer) CreateMemovia(context.Context, *CreateMemoviaRequest) (*Memovia, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMemovia not implemented")
}
func (UnimplementedMemoviaServiceServer) mustEmbedUnimplementedMemoviaServiceServer() {}

// UnsafeMemoviaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemoviaServiceServer will
// result in compilation errors.
type UnsafeMemoviaServiceServer interface {
	mustEmbedUnimplementedMemoviaServiceServer()
}

func RegisterMemoviaServiceServer(s grpc.ServiceRegistrar, srv MemoviaServiceServer) {
	s.RegisterService(&MemoviaService_ServiceDesc, srv)
}

func _MemoviaService_CreateMemovia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemoviaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoviaServiceServer).CreateMemovia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infrastracture.MemoviaService/CreateMemovia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoviaServiceServer).CreateMemovia(ctx, req.(*CreateMemoviaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MemoviaService_ServiceDesc is the grpc.ServiceDesc for MemoviaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemoviaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infrastracture.MemoviaService",
	HandlerType: (*MemoviaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMemovia",
			Handler:    _MemoviaService_CreateMemovia_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}