package repository

import (
	"context"

	"github.com/abyssparanoia/catharsis/authentication/domain/model"
)

// User ... user repositoy interface
type User interface {
	Get(ctx context.Context, userID string) (*model.User, error)
	Create(ctx context.Context, payload struct {
		DisplayName         string
		IconImagePath       string
		BackgroundImagePath string
		Profile             string
		Email               string
		CreatedAt           int64
	}) (*model.User, error)
}
