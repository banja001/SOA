// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: stakeholder-service/stakeholder-service.proto

package stakeholder_service

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
	StakeholderService_Login_FullMethodName = "/StakeholderService/login"
)

// StakeholderServiceClient is the client API for StakeholderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StakeholderServiceClient interface {
	// Login
	Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*AuthenticationTokensResponse, error)
}

type stakeholderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStakeholderServiceClient(cc grpc.ClientConnInterface) StakeholderServiceClient {
	return &stakeholderServiceClient{cc}
}

func (c *stakeholderServiceClient) Login(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*AuthenticationTokensResponse, error) {
	out := new(AuthenticationTokensResponse)
	err := c.cc.Invoke(ctx, StakeholderService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StakeholderServiceServer is the server API for StakeholderService service.
// All implementations must embed UnimplementedStakeholderServiceServer
// for forward compatibility
type StakeholderServiceServer interface {
	// Login
	Login(context.Context, *Credentials) (*AuthenticationTokensResponse, error)
	mustEmbedUnimplementedStakeholderServiceServer()
}

// UnimplementedStakeholderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStakeholderServiceServer struct {
}

func (UnimplementedStakeholderServiceServer) Login(context.Context, *Credentials) (*AuthenticationTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedStakeholderServiceServer) mustEmbedUnimplementedStakeholderServiceServer() {}

// UnsafeStakeholderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StakeholderServiceServer will
// result in compilation errors.
type UnsafeStakeholderServiceServer interface {
	mustEmbedUnimplementedStakeholderServiceServer()
}

func RegisterStakeholderServiceServer(s grpc.ServiceRegistrar, srv StakeholderServiceServer) {
	s.RegisterService(&StakeholderService_ServiceDesc, srv)
}

func _StakeholderService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StakeholderServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StakeholderService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StakeholderServiceServer).Login(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

// StakeholderService_ServiceDesc is the grpc.ServiceDesc for StakeholderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StakeholderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StakeholderService",
	HandlerType: (*StakeholderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "login",
			Handler:    _StakeholderService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stakeholder-service/stakeholder-service.proto",
}
