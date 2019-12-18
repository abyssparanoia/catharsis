package jwtauth

import "context"

// JwtauthSign ... sign service interface
type JwtauthSign interface {
	GenerateToken(ctx context.Context, claims *Claims) (string, string, error)
}
