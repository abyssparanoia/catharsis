package intercepter

import (
	"context"

	"github.com/abyssparanoia/catharsis/pkg/jwtauth"
	"github.com/abyssparanoia/catharsis/pkg/log"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Authentication ... authentication intercepter
type Authentication struct {
	jwtauthVerify jwtauth.JwtauthVerify
}

// Func ... intercepter function
func (i *Authentication) Func(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		log.Errorf(ctx, "grpc_auth.AuthFromMD", zap.Error(err))
		return nil, status.Errorf(codes.Unauthenticated, "could not read auth token: %v", err)
	}

	claims, err := i.jwtauthVerify.Validate(ctx, token)
	if err != nil {
		log.Errorf(ctx, "i.jwtauthVerify.Validate", zap.Error(err))
		return nil, nil
	}

	ctx = jwtauth.SetClaimsToContext(ctx, claims)

	return ctx, nil
}

// NewAuthentication ... get new authentication intercepter
func NewAuthentication(jwtauthVerify jwtauth.JwtauthVerify) *Authentication {
	return &Authentication{jwtauthVerify}
}
