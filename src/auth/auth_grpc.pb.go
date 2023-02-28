// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: auth.proto

package auth

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

// AuthsClient is the client API for Auths service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthsClient interface {
	IsAuth(ctx context.Context, in *GetAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type authsClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthsClient(cc grpc.ClientConnInterface) AuthsClient {
	return &authsClient{cc}
}

func (c *authsClient) IsAuth(ctx context.Context, in *GetAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/Auths/IsAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthsServer is the server API for Auths service.
// All implementations must embed UnimplementedAuthsServer
// for forward compatibility
type AuthsServer interface {
	IsAuth(context.Context, *GetAuthRequest) (*AuthResponse, error)
	mustEmbedUnimplementedAuthsServer()
}

// UnimplementedAuthsServer must be embedded to have forward compatible implementations.
type UnimplementedAuthsServer struct {
}

func (UnimplementedAuthsServer) IsAuth(context.Context, *GetAuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuth not implemented")
}
func (UnimplementedAuthsServer) mustEmbedUnimplementedAuthsServer() {}

// UnsafeAuthsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthsServer will
// result in compilation errors.
type UnsafeAuthsServer interface {
	mustEmbedUnimplementedAuthsServer()
}

func RegisterAuthsServer(s grpc.ServiceRegistrar, srv AuthsServer) {
	s.RegisterService(&Auths_ServiceDesc, srv)
}

func _Auths_IsAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthsServer).IsAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auths/IsAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthsServer).IsAuth(ctx, req.(*GetAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auths_ServiceDesc is the grpc.ServiceDesc for Auths service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auths_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Auths",
	HandlerType: (*AuthsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAuth",
			Handler:    _Auths_IsAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}