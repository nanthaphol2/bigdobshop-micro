package server

import (
	"github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartHandler"
	"github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartRepository"
	"github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartUsecase"
)

func (s *server) cartService() {
	repo := cartRepository.NewCartRepository(s.db)
	usecase := cartUsecase.NewCartUsecase(repo)
	httpHandler := cartHandler.NewCartHttpHandler(s.cfg, usecase)
	grpcHandler := cartHandler.NewCartGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	cart := s.app.Group("/cart_v1")

	cart.GET("", s.healthCheckService)
}
