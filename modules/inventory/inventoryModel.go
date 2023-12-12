package inventory

import (
	"github.com/nanthaphol2/bigdobshop-micro/modules/cart"
	"github.com/nanthaphol2/bigdobshop-micro/modules/models"
)

type (
	UpdateInventoryReq struct {
		ProfileId string `json:"profile_id" validate:"required,max=64"`
		CartId    string `json:"cart_id" validate:"required,max=64"`
	}

	CartInInventory struct {
		InventoryId string `json:"inventory_id"`
		ProfileId   string `json:"profile_id"`
		*cart.CartShowCase
	}

	InventorySearchReq struct {
		models.PaginateReq
	}

	RollbackCartInventoryReq struct {
		InventoryId string `json:"inventory_id"`
		ProfileId   string `json:"profile_id"`
		CartId      string `json:"cart_id"`
	}
)
