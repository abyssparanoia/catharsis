package firebaseauth

import (
	"context"
	"errors"

	"github.com/abyssparanoia/catharsis/pkg/log"
	"go.uber.org/zap"
)

type firebaseauth struct {
}

// SetCustomClaims ... set custom claims
func (s *firebaseauth) SetCustomClaims(ctx context.Context, userID string, claims *Claims) error {
	c, err := getAuthClient(ctx)
	if err != nil {
		log.Errorf(ctx, "getAuthClient", zap.Error(err))
		return err
	}

	err = c.SetCustomUserClaims(ctx, userID, claims.ToMap())
	if err != nil {
		log.Errorf(ctx, "c.SetCustomUserClaims", zap.Error(err))
		return err
	}

	return nil
}

// Authentication ... authenticate
func (s *firebaseauth) Authentication(ctx context.Context, ah string) (string, *Claims, error) {
	var userID string
	claims := &Claims{}

	c, err := getAuthClient(ctx)
	if err != nil {
		log.Warningf(ctx, "getAuthClient", zap.Error(err))
		return userID, claims, err
	}

	token := getTokenByAuthHeader(ah)
	if token == "" {
		err := errors.New("token empty error")
		log.Warningf(ctx, "token empty error", zap.Error(err))
		return userID, claims, err
	}

	t, err := c.VerifyIDToken(ctx, token)
	if err != nil {
		log.Warningf(ctx, "c.VerifyIDToken", zap.Error(err))
		return userID, claims, err
	}

	userID = t.UID
	claims.SetMap(t.Claims)

	return userID, claims, nil
}

// New ... get firebaseauth
func New() Firebaseauth {
	return &firebaseauth{}
}
