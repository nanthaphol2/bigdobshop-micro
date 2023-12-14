package profileHandler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nanthaphol2/bigdobshop-micro/config"
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile"
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileUsecase"
	"github.com/nanthaphol2/bigdobshop-micro/pkg/request"
	"github.com/nanthaphol2/bigdobshop-micro/pkg/response"
)

type (
	ProfileHttpHandlerService interface {
		CreateProfile(c echo.Context) error
	}

	profileHttpHandler struct {
		cfg            *config.Config
		profileUsecase profileUsecase.ProfileUsecaseService
	}
)

func NewProfileHttpHandler(cfg *config.Config, profileUsecase profileUsecase.ProfileUsecaseService) ProfileHttpHandlerService {
	return &profileHttpHandler{cfg, profileUsecase}
}

func (h *profileHttpHandler) CreateProfile(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(profile.CreateProfileReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.profileUsecase.CreateProfile(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}
