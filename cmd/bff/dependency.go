package main

import (
	"github.com/abyssparanoia/catharsis/bff/handler"
	"github.com/abyssparanoia/catharsis/bff/infrastructure/repository"
	"github.com/abyssparanoia/catharsis/bff/service"
	pb "github.com/abyssparanoia/catharsis/proto"
	"google.golang.org/grpc"
)

type dependency struct {
	AuthenticationHandler *handler.Authentication
}

func (d *dependency) Inject(e *environment) {

	authenticationConn, err := grpc.Dial(e.AuthenticationURL, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	authenticationClient := pb.NewAuthenticationClient(authenticationConn)

	authenticationRepository := repository.NewAuthentication(authenticationClient)

	authenticationService := service.NewAuthentication(authenticationRepository)

	d.AuthenticationHandler = handler.NewAuthentication(authenticationService)
}
