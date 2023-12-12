package cartHandler

import (
	"context"

	cartPb "github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartPb"
	"github.com/nanthaphol2/bigdobshop-micro/modules/cart/cartUsecase"
)

type (
	cartGrpcHandlerService struct {
		cartUsecase cartUsecase.CartUsecaseService
		cartPb.UnimplementedCartGrpcServiceServer
	}
)

func NewCartGrpcHandler(cartUsecase cartUsecase.CartUsecaseService) *cartGrpcHandlerService {
	return &cartGrpcHandlerService{
		cartUsecase: cartUsecase,
	}
}

func (g *cartGrpcHandlerService) FindCartsInIds(ctx context.Context, req *cartPb.FindCartsInIdsReq) (*cartPb.FindCartsInIdsRes, error) {
	return nil, nil
}
