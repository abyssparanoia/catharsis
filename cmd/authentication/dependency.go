package main

import (
	"fmt"
	"net"

	"github.com/abyssparanoia/catharsis/authentication/handler"
	"github.com/abyssparanoia/catharsis/authentication/infrastructure/repository"
	"github.com/abyssparanoia/catharsis/authentication/service"
	pb "github.com/abyssparanoia/catharsis/proto/authentication"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func newAuthenticationServer(logger *zap.Logger, port string) *grpc.Server {

	grpc_zap.ReplaceGrpcLoggerV2(logger)

	userRepository := repository.NewUserMock()

	authenticationService := service.NewAuthentication(userRepository)

	authenticationHandler := handler.NewAuthenticationHandler(authenticationService)

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(logger),
		),
	)

	pb.RegisterAuthenticationServer(server, authenticationHandler)
	reflection.Register(server)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	go func() {
		logger.Info(fmt.Sprintf("Listening grpc on %s:%s", "localhost", port))
		if err := server.Serve(listen); err != nil {
			logger.Error("server.Serve", zap.Error(err))
		}
	}()

	return server
}
