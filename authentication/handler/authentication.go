package handler

import (
	"context"

	"github.com/abyssparanoia/catharsis/authentication/service"
	"github.com/abyssparanoia/catharsis/pkg/log"
	pb "github.com/abyssparanoia/catharsis/proto"
	"go.uber.org/zap"
)

// AuthenticationHandler ... authentication handler struct
type AuthenticationHandler struct {
	authenticationServive service.Authentication
}

// SignIn ... sign in handler
func (h *AuthenticationHandler) SignIn(ctx context.Context, req *pb.SignInMessageRequest) (*pb.SignInMessageResponse, error) {

	// TODO: addtinal validate or call other outer service

	accessToken, refreshToken, err := h.authenticationServive.SignIn(ctx, req.UserId, req.Password)
	if err != nil {
		log.Errorf(ctx, "h.authenticationServive.SignIn: ", zap.Error(err))
		return nil, err
	}

	return &pb.SignInMessageResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

// CreateUser ... create user handler
func (h *AuthenticationHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}

// NewAuthenticationHandler ... get authentication handler
func NewAuthenticationHandler(authenticationServive service.Authentication) pb.AuthenticationServer {
	return &AuthenticationHandler{authenticationServive}
}
