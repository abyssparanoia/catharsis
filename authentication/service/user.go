package service

import (
	"context"

	"github.com/abyssparanoia/catharsis/authentication/domain/model"
)

// User ... user service interface
type User interface {
	Create(ctx context.Context, payload struct {
		Password            string
		DisplayName         string
		IconImagePath       string
		BackgroundImagePath string
		Profile             *string
		Email               *string
	}) (*model.User, error)
}
