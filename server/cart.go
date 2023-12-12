package server

import (
	"log"

	"github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartHandler"
	cartPb "github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartPb"
	"github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartRepository"
	"github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartUsecase"
	"github.com/nanthaphol2/bigdobshop-micro/pkg/grpccon"
)

func (s *server) cartService() {
	repo := cartRepository.NewCartRepository(s.db)
	usecase := cartUsecase.NewCartUsecase(repo)
	httpHandler := cartHandler.NewCartHttpHandler(s.cfg, usecase)
	grpcHandler := cartHandler.NewCartGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.CartUrl)

		cartPb.RegisterCartGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Cart gRPC server listening on %s", s.cfg.Grpc.CartUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	cart := s.app.Group("/cart_v1")

	cart.GET("", s.healthCheckService)
}
