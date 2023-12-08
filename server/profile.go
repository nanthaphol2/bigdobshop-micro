package server

import (
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileHandler"
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileRepository"
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileUsecase"
)

func (s *server) profileService() {
	repo := profileRepository.NewProfileRepository(s.db)
	usecase := profileUsecase.NewProfileUsecase(repo)
	httpHandler := profileHandler.NewProfileHttpHandler(s.cfg, usecase)
	grpcHandler := profileHandler.NewProfileGrpcHandler(usecase)
	queueHandler := profileHandler.NewProfileQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	profile := s.app.Group("/profile_v1")

	profile.GET("", s.healthCheckService)
}
