package service

import "context"

// Authentication ... authentication service interface
type Authentication interface {
	SignIn(ctx context.Context, userID string, password string) (string, string, error)
}
