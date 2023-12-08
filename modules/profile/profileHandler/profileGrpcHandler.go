package profileHandler

import "github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileUsecase"

type (
	profileGrpcHandlerService struct {
		profileUsecase profileUsecase.ProfileUsecaseService
	}
)

func NewProfileGrpcHandler(profileUsecase profileUsecase.ProfileUsecaseService) *profileGrpcHandlerService {
	return &profileGrpcHandlerService{
		profileUsecase: profileUsecase,
	}
}
