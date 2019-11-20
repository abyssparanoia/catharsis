package repositoy

import (
	"context"

	"github.com/abyssparanoia/catharsis/authentication/domain/model"
)

// User ... user repositoy interface
type User interface {
	Get(ctx context.Context, userID string) (*model.User, error)
}
