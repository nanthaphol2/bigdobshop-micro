package cartUsecase

import "github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartRepository"

type (
	CartUsecaseService interface{}

	cartUsecase struct {
		cartRepository cartRepository.CartRepositoryService
	}
)

func NewCartUsecase(cartRepository cartRepository.CartRepositoryService) CartUsecaseService {
	return &cartUsecase{
		cartRepository: cartRepository,
	}
}
