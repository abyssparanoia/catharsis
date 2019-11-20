package main

import (
	"context"
	"log"
	"net"

	pb "github.com/abyssparanoia/catharsis/proto/authentication"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SignIn(ctx context.Context, req *pb.SignInMessageRequest) (*pb.SignInMessageResponse, error) {
	log.Printf("Received: %v", req.UserId)
	return &pb.SignInMessageResponse{AccessToken: "access_token", RefreshToken: "refresh_token"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthenticationServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
