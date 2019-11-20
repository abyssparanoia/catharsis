package service

import "context"

// User ... user service interface
type User interface {
	SignIn(ctx context.Context, userID string, password string) (string, string, error)
}
