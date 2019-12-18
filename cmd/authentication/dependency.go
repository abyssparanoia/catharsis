package main

import (
	"fmt"
	"net"

	"github.com/abyssparanoia/catharsis/pkg/psql"

	"github.com/abyssparanoia/catharsis/pkg/jwtauth"

	"github.com/abyssparanoia/catharsis/authentication/handler"
	"github.com/abyssparanoia/catharsis/authentication/infrastructure/repository"
	"github.com/abyssparanoia/catharsis/authentication/service"
	pb "github.com/abyssparanoia/catharsis/proto"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

func recoveryFuncFactory(logger *zap.Logger) func(p interface{}) error {
	return func(p interface{}) error {
		logger.Error(fmt.Sprintf("p: %+v\n", p))
		return grpc.Errorf(codes.Internal, "Unexpected error")
	}
}

func newAuthenticationServer(logger *zap.Logger, env *environment) *grpc.Server {

	grpc_zap.ReplaceGrpcLoggerV2(logger)

	jwtSignClient := jwtauth.NewSignClient(env.SignKeyPath)

	psqlCfg := psql.NewConfig(env.DBHost, env.DBPort, env.DBUser, env.DBPassword, env.DBDatabase)

	psqlClient := psql.NewClient(psqlCfg)

	userRepository := repository.NewUser(psqlClient)

	jwtauthSign := jwtauth.NewSign(jwtSignClient)

	authenticationService := service.NewAuthentication(userRepository, jwtauthSign)
	userService := service.NewUser(userRepository)

	authenticationHandler := handler.NewAuthenticationHandler(authenticationService, userService)

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(recoveryFuncFactory(logger)),
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(opts...),
		),
	)

	pb.RegisterAuthenticationServer(server, authenticationHandler)
	reflection.Register(server)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", env.Port))
	if err != nil {
		panic(err)
	}

	go func() {
		logger.Info(fmt.Sprintf("Listening grpc on %s:%s", "localhost", env.Port))
		if err := server.Serve(listen); err != nil {
			logger.Error("server.Serve", zap.Error(err))
		}
	}()

	return server
}
