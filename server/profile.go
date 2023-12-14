package server

import (
	"log"

	"github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileHandler"
	profilePb "github.com/nanthaphol2/bigdobshop-micro/modules/profile/profilePb"
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileRepository"
	"github.com/nanthaphol2/bigdobshop-micro/modules/profile/profileUsecase"
	"github.com/nanthaphol2/bigdobshop-micro/pkg/grpccon"
)

func (s *server) profileService() {
	repo := profileRepository.NewProfileRepository(s.db)
	usecase := profileUsecase.NewProfileUsecase(repo)
	httpHandler := profileHandler.NewProfileHttpHandler(s.cfg, usecase)
	grpcHandler := profileHandler.NewProfileGrpcHandler(usecase)
	queueHandler := profileHandler.NewProfileQueueHandler(s.cfg, usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ProfileUrl)

		profilePb.RegisterProfileGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Profile gRPC server listening on %s", s.cfg.Grpc.ProfileUrl)
		grpcServer.Serve(lis)
	}()

	_ = grpcHandler
	_ = queueHandler

	profile := s.app.Group("/profile_v1")

	profile.GET("", s.healthCheckService)
	profile.POST("/profile/register", httpHandler.CreateProfile)
}
