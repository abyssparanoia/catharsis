package service

import (
	"context"
	"errors"
	"log"

	"github.com/abyssparanoia/catharsis/authentication/domain/repository"
)

type authentication struct {
	userRepository repository.User
}

func (s *authentication) SignIn(ctx context.Context, userID string, password string) (accessToken string, refreshToken string, err error) {

	user, err := s.userRepository.Get(ctx, userID)
	if err != nil {
		log.Fatalf("s.userRepository.Get: %s", err)
		return accessToken, refreshToken, err
	}

	// TODO(abyssparanoia): hash check
	if user.Password != password {
		err = errors.New("invalid password")
		return accessToken, refreshToken, err
	}

	// TODO: generate token by usign jwt-go
	accessToken = "access_token"
	refreshToken = "refresh_token"

	return accessToken, refreshToken, nil
}

// NewAuthentication ... get authentication service
func NewAuthentication(userRepository repository.User) Authentication {
	return &authentication{userRepository}
}
