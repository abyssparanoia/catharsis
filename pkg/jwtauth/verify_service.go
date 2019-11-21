package jwtauth

import (
	"context"
)

// VerifyService ... jwtauth verify service interface
type VerifyService interface {
	Validate(ctx context.Context, token string) (*Claims, error)
}
