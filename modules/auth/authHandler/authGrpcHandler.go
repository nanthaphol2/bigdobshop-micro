package authHandler

import (
	"context"

	authPb "github.com/nanthaphol2/bigdobshop-micro/modules/auth/authPb"
	"github.com/nanthaphol2/bigdobshop-micro/modules/auth/authUsecase"
)

type (
	authGrpcHandlerService struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandlerService {
	return &authGrpcHandlerService{
		authUsecase: authUsecase,
	}
}

func (g *authGrpcHandlerService) AccessTokenSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	return nil, nil
}

func (g *authGrpcHandlerService) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}
