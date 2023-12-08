package inventoryHandler

import "github.com/nanthaphol2/bigdobshop-micro/modules/inventory/inventoryUsecase"

type (
	inventoryGrpcHandlerService struct {
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandlerService {
	return &inventoryGrpcHandlerService{
		inventoryUsecase: inventoryUsecase,
	}
}
