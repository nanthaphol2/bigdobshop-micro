package authHandler

import "github.com/nanthaphol2/bigdobshop-micro/modules/auth/authUsecase"

type (
	authGrpcHandlerService struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandlerService {
	return &authGrpcHandlerService{
		authUsecase: authUsecase,
	}
}
