// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: modules/inventory/inventoryPb/inventoryPb.proto

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

// InventoryGrpcServiceClient is the client API for InventoryGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryGrpcServiceClient interface {
	IsAvaliableToSell(ctx context.Context, in *IsAvaliableToSellReq, opts ...grpc.CallOption) (*IsAvaliableToSellRes, error)
}

type inventoryGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryGrpcServiceClient(cc grpc.ClientConnInterface) InventoryGrpcServiceClient {
	return &inventoryGrpcServiceClient{cc}
}

func (c *inventoryGrpcServiceClient) IsAvaliableToSell(ctx context.Context, in *IsAvaliableToSellReq, opts ...grpc.CallOption) (*IsAvaliableToSellRes, error) {
	out := new(IsAvaliableToSellRes)
	err := c.cc.Invoke(ctx, "/InventoryGrpcService/IsAvaliableToSell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryGrpcServiceServer is the server API for InventoryGrpcService service.
// All implementations must embed UnimplementedInventoryGrpcServiceServer
// for forward compatibility
type InventoryGrpcServiceServer interface {
	IsAvaliableToSell(context.Context, *IsAvaliableToSellReq) (*IsAvaliableToSellRes, error)
	mustEmbedUnimplementedInventoryGrpcServiceServer()
}

// UnimplementedInventoryGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryGrpcServiceServer struct {
}

func (UnimplementedInventoryGrpcServiceServer) IsAvaliableToSell(context.Context, *IsAvaliableToSellReq) (*IsAvaliableToSellRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAvaliableToSell not implemented")
}
func (UnimplementedInventoryGrpcServiceServer) mustEmbedUnimplementedInventoryGrpcServiceServer() {}

// UnsafeInventoryGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryGrpcServiceServer will
// result in compilation errors.
type UnsafeInventoryGrpcServiceServer interface {
	mustEmbedUnimplementedInventoryGrpcServiceServer()
}

func RegisterInventoryGrpcServiceServer(s grpc.ServiceRegistrar, srv InventoryGrpcServiceServer) {
	s.RegisterService(&InventoryGrpcService_ServiceDesc, srv)
}

func _InventoryGrpcService_IsAvaliableToSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAvaliableToSellReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryGrpcServiceServer).IsAvaliableToSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InventoryGrpcService/IsAvaliableToSell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryGrpcServiceServer).IsAvaliableToSell(ctx, req.(*IsAvaliableToSellReq))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryGrpcService_ServiceDesc is the grpc.ServiceDesc for InventoryGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InventoryGrpcService",
	HandlerType: (*InventoryGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsAvaliableToSell",
			Handler:    _InventoryGrpcService_IsAvaliableToSell_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/inventory/inventoryPb/inventoryPb.proto",
}