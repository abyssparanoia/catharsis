package service

import (
	"context"

	"github.com/abyssparanoia/catharsis/bff/domain/repository"
	"github.com/abyssparanoia/catharsis/pkg/log"
	"go.uber.org/zap"
)

type authentication struct {
	authenticationRepository repository.Authentication
}

func (s authentication) SignIn(ctx context.Context, userID string, password string) (string, string, error) {

	accessToken, refreshToken, err := s.authenticationRepository.SignIn(ctx, userID, password)
	if err != nil {
		log.Errorf(ctx, "s.authenticationRepository.SignIn", zap.Error(err))
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// NewAuthentication ... get authentication service
func NewAuthentication(authenticationRepository repository.Authentication) Authentication {
	return &authentication{authenticationRepository}
}
