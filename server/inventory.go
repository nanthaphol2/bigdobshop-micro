package server

import (
	"log"

	"github.com/nanthaphol2/bigdobshop-micro/modules/inventory/inventoryHandler"
	"github.com/nanthaphol2/bigdobshop-micro/modules/inventory/inventoryRepository"
	"github.com/nanthaphol2/bigdobshop-micro/modules/inventory/inventoryUsecase"
	"github.com/nanthaphol2/bigdobshop-micro/pkg/grpccon"

	inventoryPb "github.com/nanthaphol2/bigdobshop-micro/modules/inventory/inventoryPb"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.InventoryUrl)

		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Inventory gRPC server listening on %s", s.cfg.Grpc.InventoryUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	inventory.GET("", s.healthCheckService)
}
