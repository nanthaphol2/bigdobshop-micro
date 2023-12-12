package inventoryHandler

import (
	"context"

	inventoryPb "github.com/nanthaphol2/bigdobshop-micro/modules/inventory/inventoryPb"
	"github.com/nanthaphol2/bigdobshop-micro/modules/inventory/inventoryUsecase"
)

type (
	inventoryGrpcHandlerService struct {
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
		inventoryPb.UnimplementedInventoryGrpcServiceServer
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandlerService {
	return &inventoryGrpcHandlerService{
		inventoryUsecase: inventoryUsecase,
	}
}

func (g *inventoryGrpcHandlerService) IsAvaliableToSell(ctx context.Context, req *inventoryPb.IsAvaliableToSellReq) (*inventoryPb.IsAvaliableToSellRes, error) {
	return nil, nil
}
