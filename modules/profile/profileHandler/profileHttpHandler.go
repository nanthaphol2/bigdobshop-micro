package profileHandler

import (
	"github.com/nanthaphol2/bigdobshop-micro/config"
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileUsecase"
)

type (
	ProfileHttpHandlerService interface{}

	profileHttpHandler struct {
		cfg            *config.Config
		profileUsecase profileUsecase.ProfileUsecaseService
	}
)

func NewProfileHttpHandler(cfg *config.Config, profileUsecase profileUsecase.ProfileUsecaseService) ProfileHttpHandlerService {
	return &profileHttpHandler{cfg, profileUsecase}
}
