// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: system.compatibility.external.proto

package compatibility_grpc

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

// CompatibilityCheckerClient is the client API for CompatibilityChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompatibilityCheckerClient interface {
	//*
	//Получить уровень совместимости текущего клиента с сервером
	CheckCompatibility(ctx context.Context, in *CheckCompatibilityRequest, opts ...grpc.CallOption) (*CheckCompatibilityResponse, error)
}

type compatibilityCheckerClient struct {
	cc grpc.ClientConnInterface
}

func NewCompatibilityCheckerClient(cc grpc.ClientConnInterface) CompatibilityCheckerClient {
	return &compatibilityCheckerClient{cc}
}

func (c *compatibilityCheckerClient) CheckCompatibility(ctx context.Context, in *CheckCompatibilityRequest, opts ...grpc.CallOption) (*CheckCompatibilityResponse, error) {
	out := new(CheckCompatibilityResponse)
	err := c.cc.Invoke(ctx, "/cheetah.system.compatibility.CompatibilityChecker/CheckCompatibility", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompatibilityCheckerServer is the server API for CompatibilityChecker service.
// All implementations must embed UnimplementedCompatibilityCheckerServer
// for forward compatibility
type CompatibilityCheckerServer interface {
	//*
	//Получить уровень совместимости текущего клиента с сервером
	CheckCompatibility(context.Context, *CheckCompatibilityRequest) (*CheckCompatibilityResponse, error)
	mustEmbedUnimplementedCompatibilityCheckerServer()
}

// UnimplementedCompatibilityCheckerServer must be embedded to have forward compatible implementations.
type UnimplementedCompatibilityCheckerServer struct {
}

func (UnimplementedCompatibilityCheckerServer) CheckCompatibility(context.Context, *CheckCompatibilityRequest) (*CheckCompatibilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCompatibility not implemented")
}
func (UnimplementedCompatibilityCheckerServer) mustEmbedUnimplementedCompatibilityCheckerServer() {}

// UnsafeCompatibilityCheckerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompatibilityCheckerServer will
// result in compilation errors.
type UnsafeCompatibilityCheckerServer interface {
	mustEmbedUnimplementedCompatibilityCheckerServer()
}

func RegisterCompatibilityCheckerServer(s grpc.ServiceRegistrar, srv CompatibilityCheckerServer) {
	s.RegisterService(&CompatibilityChecker_ServiceDesc, srv)
}

func _CompatibilityChecker_CheckCompatibility_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckCompatibilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompatibilityCheckerServer).CheckCompatibility(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheetah.system.compatibility.CompatibilityChecker/CheckCompatibility",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompatibilityCheckerServer).CheckCompatibility(ctx, req.(*CheckCompatibilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompatibilityChecker_ServiceDesc is the grpc.ServiceDesc for CompatibilityChecker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompatibilityChecker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cheetah.system.compatibility.CompatibilityChecker",
	HandlerType: (*CompatibilityCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckCompatibility",
			Handler:    _CompatibilityChecker_CheckCompatibility_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system.compatibility.external.proto",
}
