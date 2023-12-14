package profileUsecase

import (
	"context"
	"errors"

	"github.com/nanthaphol2/bigdobshop-micro/modules/profile"
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileRepository"
	"github.com/nanthaphol2/bigdobshop-micro/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	ProfileUsecaseService interface {
		CreateProfile(pctx context.Context, req *profile.CreateProfileReq) (string, error)
	}

	profileUsecase struct {
		profileRepository profileRepository.ProfileRepositoryService
	}
)

func NewProfileUsecase(profileRepository profileRepository.ProfileRepositoryService) ProfileUsecaseService {
	return &profileUsecase{profileRepository: profileRepository}
}

func (u *profileUsecase) CreateProfile(pctx context.Context, req *profile.CreateProfileReq) (string, error) {
	if !u.profileRepository.IsUniqueProfile(pctx, req.Email, req.Username) {
		return "", errors.New("error: email or username already exist")
	}

	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error: failed to hash password")
	}

	// Insert one player
	profileId, err := u.profileRepository.InsertOneProfile(pctx, &profile.Profile{
		Email:     req.Email,
		Password:  string(hashedPassword),
		Username:  req.Username,
		CreatedAt: utils.LocalTime(),
		UpdatedAt: utils.LocalTime(),
		ProfileRoles: []profile.ProfileRole{
			{
				RoleTitle: "user",
				RoleCode:  0,
			},
		},
	})
	return profileId.Hex(), nil
	// return u.FindOnePlayerProfile(pctx, playerId.Hex())
}
