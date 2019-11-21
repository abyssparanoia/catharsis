package jwtauth

import "context"

// SignService ... sign service interface
type SignService interface {
	GenerateAccessToken(ctx context.Context, claims *Claims) (string, error)
}
