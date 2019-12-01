package service

import (
	"context"

	"github.com/abyssparanoia/catharsis/authentication/domain/model"
	"github.com/abyssparanoia/catharsis/authentication/domain/repository"
	"github.com/abyssparanoia/catharsis/pkg/log"
	"go.uber.org/zap"
)

type user struct {
	userRepository repository.User
}

func (s *user) Create(ctx context.Context, payload struct {
	Password            string
	DisplayName         string
	IconImagePath       string
	BackgroundImagePath string
	Profile             *string
	Email               *string
}) (*model.User, error) {

	user, err := s.userRepository.Create(ctx, struct {
		Password            string
		DisplayName         string
		IconImagePath       string
		BackgroundImagePath string
		Profile             *string
		Email               *string
	}{
		Password:            payload.Password,
		DisplayName:         payload.DisplayName,
		IconImagePath:       payload.IconImagePath,
		BackgroundImagePath: payload.BackgroundImagePath,
		Profile:             payload.Profile,
		Email:               payload.Email,
	})
	if err != nil {
		log.Errorf(ctx, "s.userRepository.Create", zap.Error(err))
		return nil, err
	}

	return user, nil
}

// NewUser ... get user service
func NewUser(userRepository repository.User) User {
	return &user{userRepository}
}
