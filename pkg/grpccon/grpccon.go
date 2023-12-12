package grpccon

import (
	"errors"
	"log"
	"net"

	"github.com/nanthaphol2/bigdobshop-micro/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authPb "github.com/nanthaphol2/bigdobshop-micro/modules/auth/authPb"
	cartPb "github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartPb"
	inventoryPb "github.com/nanthaphol2/bigdobshop-micro/modules/inventory/inventoryPb"
	profilePb "github.com/nanthaphol2/bigdobshop-micro/modules/profile/profilePb"
)

type (
	GrpcClientFactoryHandler interface {
		Auth() authPb.AuthGrpcServiceClient
		Profile() profilePb.ProfileGrpcServiceClient
		Cart() cartPb.CartGrpcServiceClient
		Inventory() inventoryPb.InventoryGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}

	grpcAuth struct {
	}
)

func (g *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Profile() profilePb.ProfileGrpcServiceClient {
	return profilePb.NewProfileGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Cart() cartPb.CartGrpcServiceClient {
	return cartPb.NewCartGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Inventory() inventoryPb.InventoryGrpcServiceClient {
	return inventoryPb.NewInventoryGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Printf("Error: Grpc client connection failed: %s", err.Error())
		return nil, errors.New("error: grpc client connection failed")
	}

	return &grpcClientFactory{
		client: clientConn,
	}, nil
}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)

	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}

	return grpcServer, lis
}
