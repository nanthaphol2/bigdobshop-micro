package cart

import "github.com/nanthaphol2/bigdobshop-micro/modules/models"

type (
	CreateCartReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
	}

	CartShowCase struct {
		CartId   string  `json:"cart_id"`
		Title    string  `json:"title"`
		Price    float64 `json:"price"`
		ImageUrl string  `json:"image_url"`
	}

	CartSearchReq struct {
		Title string `query:"title" validate:"max=64"`
		models.PaginateReq
	}

	CartUpdateReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
	}

	EnableOrDisableCartReq struct {
		UsageStatus bool `json:"usage_status"`
	}
)
