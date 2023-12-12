package server

import (
	"log"

	"github.com/nanthaphol2/bigdobshop-micro/modules/auth/authHandler"
	authPb "github.com/nanthaphol2/bigdobshop-micro/modules/auth/authPb"
	"github.com/nanthaphol2/bigdobshop-micro/modules/auth/authRepository"
	"github.com/nanthaphol2/bigdobshop-micro/modules/auth/authUsecase"
	"github.com/nanthaphol2/bigdobshop-micro/pkg/grpccon"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	// gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Auth gRPC server listening on %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	auth.GET("", s.healthCheckService)
}
