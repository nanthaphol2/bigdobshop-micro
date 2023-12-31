// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: modules/profile/profilePb/profilePb.proto

package bigdobshop_micro

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

// ProfileGrpcServiceClient is the client API for ProfileGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileGrpcServiceClient interface {
	CredentialSearch(ctx context.Context, in *CredentialSearchReq, opts ...grpc.CallOption) (*ProfileProfile, error)
	FindOneProfileProfileToRefresh(ctx context.Context, in *FindOneProfileProfileToRefreshReq, opts ...grpc.CallOption) (*ProfileProfile, error)
	GetProfileSavingAccount(ctx context.Context, in *GetProfileSavingAccountReq, opts ...grpc.CallOption) (*GetProfileSavingAccountRes, error)
}

type profileGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileGrpcServiceClient(cc grpc.ClientConnInterface) ProfileGrpcServiceClient {
	return &profileGrpcServiceClient{cc}
}

func (c *profileGrpcServiceClient) CredentialSearch(ctx context.Context, in *CredentialSearchReq, opts ...grpc.CallOption) (*ProfileProfile, error) {
	out := new(ProfileProfile)
	err := c.cc.Invoke(ctx, "/ProfileGrpcService/CredentialSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileGrpcServiceClient) FindOneProfileProfileToRefresh(ctx context.Context, in *FindOneProfileProfileToRefreshReq, opts ...grpc.CallOption) (*ProfileProfile, error) {
	out := new(ProfileProfile)
	err := c.cc.Invoke(ctx, "/ProfileGrpcService/FindOneProfileProfileToRefresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileGrpcServiceClient) GetProfileSavingAccount(ctx context.Context, in *GetProfileSavingAccountReq, opts ...grpc.CallOption) (*GetProfileSavingAccountRes, error) {
	out := new(GetProfileSavingAccountRes)
	err := c.cc.Invoke(ctx, "/ProfileGrpcService/GetProfileSavingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileGrpcServiceServer is the server API for ProfileGrpcService service.
// All implementations must embed UnimplementedProfileGrpcServiceServer
// for forward compatibility
type ProfileGrpcServiceServer interface {
	CredentialSearch(context.Context, *CredentialSearchReq) (*ProfileProfile, error)
	FindOneProfileProfileToRefresh(context.Context, *FindOneProfileProfileToRefreshReq) (*ProfileProfile, error)
	GetProfileSavingAccount(context.Context, *GetProfileSavingAccountReq) (*GetProfileSavingAccountRes, error)
	mustEmbedUnimplementedProfileGrpcServiceServer()
}

// UnimplementedProfileGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileGrpcServiceServer struct {
}

func (UnimplementedProfileGrpcServiceServer) CredentialSearch(context.Context, *CredentialSearchReq) (*ProfileProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CredentialSearch not implemented")
}
func (UnimplementedProfileGrpcServiceServer) FindOneProfileProfileToRefresh(context.Context, *FindOneProfileProfileToRefreshReq) (*ProfileProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneProfileProfileToRefresh not implemented")
}
func (UnimplementedProfileGrpcServiceServer) GetProfileSavingAccount(context.Context, *GetProfileSavingAccountReq) (*GetProfileSavingAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileSavingAccount not implemented")
}
func (UnimplementedProfileGrpcServiceServer) mustEmbedUnimplementedProfileGrpcServiceServer() {}

// UnsafeProfileGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileGrpcServiceServer will
// result in compilation errors.
type UnsafeProfileGrpcServiceServer interface {
	mustEmbedUnimplementedProfileGrpcServiceServer()
}

func RegisterProfileGrpcServiceServer(s grpc.ServiceRegistrar, srv ProfileGrpcServiceServer) {
	s.RegisterService(&ProfileGrpcService_ServiceDesc, srv)
}

func _ProfileGrpcService_CredentialSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileGrpcServiceServer).CredentialSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileGrpcService/CredentialSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileGrpcServiceServer).CredentialSearch(ctx, req.(*CredentialSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileGrpcService_FindOneProfileProfileToRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneProfileProfileToRefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileGrpcServiceServer).FindOneProfileProfileToRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileGrpcService/FindOneProfileProfileToRefresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileGrpcServiceServer).FindOneProfileProfileToRefresh(ctx, req.(*FindOneProfileProfileToRefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileGrpcService_GetProfileSavingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileSavingAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileGrpcServiceServer).GetProfileSavingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProfileGrpcService/GetProfileSavingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileGrpcServiceServer).GetProfileSavingAccount(ctx, req.(*GetProfileSavingAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileGrpcService_ServiceDesc is the grpc.ServiceDesc for ProfileGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProfileGrpcService",
	HandlerType: (*ProfileGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CredentialSearch",
			Handler:    _ProfileGrpcService_CredentialSearch_Handler,
		},
		{
			MethodName: "FindOneProfileProfileToRefresh",
			Handler:    _ProfileGrpcService_FindOneProfileProfileToRefresh_Handler,
		},
		{
			MethodName: "GetProfileSavingAccount",
			Handler:    _ProfileGrpcService_GetProfileSavingAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/profile/profilePb/profilePb.proto",
}
