package profileUsecase

import "github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileRepository"

type (
	ProfileUsecaseService interface{}

	profileUsecase struct {
		profileRepository profileRepository.ProfileRepositoryService
	}
)

func NewProfileUsecase(profileRepository profileRepository.ProfileRepositoryService) ProfileUsecaseService {
	return &profileUsecase{profileRepository}
}
