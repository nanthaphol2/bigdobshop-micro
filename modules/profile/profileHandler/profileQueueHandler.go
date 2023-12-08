package profileHandler

import (
	"github.com/nanthaphol2/bigdobshop-micro/config"
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileUsecase"
)

type (
	ProfileQueueHandlerService interface{}

	profileQueueHandler struct {
		cfg            *config.Config
		profileUsecase profileUsecase.ProfileUsecaseService
	}
)

func NewProfileQueueHandler(cfg *config.Config, profileUsecase profileUsecase.ProfileUsecaseService) ProfileQueueHandlerService {
	return &profileQueueHandler{
		cfg:            cfg,
		profileUsecase: profileUsecase,
	}
}
