package jwtauth

import (
	"context"
)

// JwtauthVerify ... jwtauth verify service interface
type JwtauthVerify interface {
	Validate(ctx context.Context, token string) (*Claims, error)
}
