package repository

import (
	"context"

	"github.com/abyssparanoia/catharsis/authentication/domain/model"
	"github.com/abyssparanoia/catharsis/authentication/domain/repository"
	"github.com/abyssparanoia/catharsis/authentication/infrastructure/internal/entity"
)

type userMock struct{}

func (m *userMock) Get(ctx context.Context, userID string) (*model.User, error) {
	user := entity.User{
		ID:        "user_id",
		Password:  "password",
		CreatedAt: 10000000,
		UpdatedAt: 10000000,
	}
	return user.OutputModel(), nil
}

// NewUserMock ... get user mock
func NewUserMock() repository.User {
	return &userMock{}
}
