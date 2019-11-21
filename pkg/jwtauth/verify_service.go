package jwtauth

import (
	"context"
)

// VerifyService ... jwtauth verify service interface
type VerifyService interface {
	Authenticate(ctx context.Context, token string) (*Claims, error)
}
