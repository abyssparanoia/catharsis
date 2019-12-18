package service

import (
	"context"

	"github.com/abyssparanoia/catharsis/authentication/domain/model"
)

// User ... user service interface
type User interface {
	Create(ctx context.Context, paylaod UserCreatePayload) (*UserCreateResult, error)
}

// UserCreatePayload ... paylaod for create user service
type UserCreatePayload struct {
	Password            string
	DisplayName         string
	IconImagePath       string
	BackgroundImagePath string
	Profile             *string
	Email               *string
}

// UserCreateResult ... result for create user service
type UserCreateResult struct {
	User *model.User
}
