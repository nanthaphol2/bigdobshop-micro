package inventoryUsecase

import "github.com/nanthaphol2/bigdobshop-micro/modules/inventory/inventoryRepository"

type (
	InventoryUsecaseService interface {
	}

	inventoryUsecase struct {
		inventoryRepository inventoryRepository.InventoryRepositoryService
	}
)

func NewInventoryUsecase(inventoryRepository inventoryRepository.InventoryRepositoryService) InventoryUsecaseService {
	return &inventoryUsecase{
		inventoryRepository: inventoryRepository,
	}
}
