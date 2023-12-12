package payment

type (
	CartServiceReq struct {
		Carts []*CartServiceReqDatum `json:"carts" validate:"required"`
	}

	CartServiceReqDatum struct {
		CartId string  `json:"cart_id" validate:"required,max=64"`
		Price  float64 `json:"price"`
	}
)
