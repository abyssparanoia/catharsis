package repository

import "context"

// Authentication ... authentication reposiroty interface
type Authentication interface {
	SignIn(ctx context.Context, userID string, password string) (string, string, error)
}
