package jwtauth

import "context"

// SignService ... sign service interface
type SignService interface {
	GenerateToken(ctx context.Context, claims *Claims) (string, string, error)
}
