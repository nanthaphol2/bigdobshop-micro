package cartHandler

import "github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartUsecase"

type (
	cartGrpcHandler struct {
		cartUsecase cartUsecase.CartUsecaseService
	}
)

func NewCartGrpcHandler(cartUsecase cartUsecase.CartUsecaseService) *cartGrpcHandler {
	return &cartGrpcHandler{
		cartUsecase: cartUsecase,
	}
}
