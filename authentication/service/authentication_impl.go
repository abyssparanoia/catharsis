package service

import (
	"context"
	"errors"

	"github.com/abyssparanoia/catharsis/pkg/jwtauth"

	"go.uber.org/zap"

	"github.com/abyssparanoia/catharsis/authentication/domain/repository"
	"github.com/abyssparanoia/catharsis/pkg/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authentication struct {
	userRepository repository.User
	jwtSignSvc     jwtauth.SignService
}

func (s *authentication) SignIn(ctx context.Context, userID string, password string) (accessToken string, refreshToken string, err error) {

	user, err := s.userRepository.Get(ctx, userID)
	if err != nil {
		log.Errorf(ctx, "s.userRepository.Get: %s", zap.Error(err))
		return accessToken, refreshToken, status.Errorf(codes.Internal, "s.userRepository.Get:%s", err)
	}

	// TODO(abyssparanoia): hash check
	if user.Password != password {
		err = errors.New("invalid password")
		log.Errorf(ctx, "invalid password", zap.Error(err))
		return accessToken, refreshToken, status.Errorf(codes.Internal, "invalid password")
	}

	claims := &jwtauth.Claims{}
	claims.Set(userID)

	accessToken, err = s.jwtSignSvc.GenerateAccessToken(ctx, claims)
	if err != nil {
		log.Errorf(ctx, "s.jwtSignSvc.GenerateAccessToken: %s", zap.Error(err))
		return accessToken, refreshToken, status.Errorf(codes.Internal, "s.jwtSignSvc.GenerateAccessToken:%s", err)
	}

	// TODO: generate token by usign jwt-go
	refreshToken = "refresh_token"

	return accessToken, refreshToken, nil
}

// NewAuthentication ... get authentication service
func NewAuthentication(userRepository repository.User, jwtSignSvc jwtauth.SignService) Authentication {
	return &authentication{userRepository, jwtSignSvc}
}
