package main

import (
	"log"
	"net"

	"github.com/abyssparanoia/catharsis/authentication/handler"
	"github.com/abyssparanoia/catharsis/authentication/infrastructure/repository"
	"github.com/abyssparanoia/catharsis/authentication/service"
	pb "github.com/abyssparanoia/catharsis/proto/authentication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func newAuthenticationServer(port string) *grpc.Server {

	userRepository := repository.NewUserMock()

	authenticationService := service.NewAuthentication(userRepository)

	authenticationHandler := handler.NewAuthenticationHandler(authenticationService)

	server := grpc.NewServer()

	pb.RegisterAuthenticationServer(server, authenticationHandler)
	reflection.Register(server)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	go func() {
		log.Printf("Listening grpc on %s:%s\n", "localhost", port)
		if err := server.Serve(listen); err != nil {
			log.Print(err)
		}
	}()

	return server
}
