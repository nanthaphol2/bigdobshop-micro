package cartHandler

import (
	"github.com/nanthaphol2/bigdobshop-micro/config"
	"github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartUsecase"
)

type (
	CartHttpHandlerService interface {
	}

	cartHttpHandler struct {
		cfg         *config.Config
		cartUsecase cartUsecase.CartUsecaseService
	}
)

func NewCartHttpHandler(cfg *config.Config, cartUsecase cartUsecase.CartUsecaseService) CartHttpHandlerService {
	return &cartHttpHandler{
		cfg:         cfg,
		cartUsecase: cartUsecase,
	}
}
