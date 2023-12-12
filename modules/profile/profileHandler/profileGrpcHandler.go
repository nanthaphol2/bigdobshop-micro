package profileHandler

import (
	"context"

	profilePb "github.com/nanthaphol2/bigdobshop-micro/modules/profile/profilePb"
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileUsecase"
)

type (
	profileGrpcHandlerService struct {
		profileUsecase profileUsecase.ProfileUsecaseService
		profilePb.UnimplementedProfileGrpcServiceServer
	}
)

func NewProfileGrpcHandler(profileUsecase profileUsecase.ProfileUsecaseService) *profileGrpcHandlerService {
	return &profileGrpcHandlerService{
		profileUsecase: profileUsecase,
	}
}

func (g *profileGrpcHandlerService) CredentialSearch(ctx context.Context, req *profilePb.CredentialSearchReq) (*profilePb.ProfileProfile, error) {
	return nil, nil
}

func (g *profileGrpcHandlerService) FindOneProfileProfileToRefresh(ctx context.Context, req *profilePb.FindOneProfileProfileToRefreshReq) (*profilePb.ProfileProfile, error) {
	return nil, nil
}

func (g *profileGrpcHandlerService) GetProfileSavingAccount(ctx context.Context, req *profilePb.GetProfileSavingAccountReq) (*profilePb.GetProfileSavingAccountRes, error) {
	return nil, nil
}
