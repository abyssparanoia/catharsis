package firebaseauth

import (
	"context"

	"github.com/abyssparanoia/catharsis/pkg/log"
	"go.uber.org/zap"

	"errors"
)

type firebaseauthDebug struct {
}

// Authentication ... authenticate
func (s *firebaseauthDebug) Authentication(ctx context.Context, ah string) (string, *Claims, error) {
	var userID string
	claims := &Claims{}

	// ユーザーを取得できたらデバッグリクエストと判定する
	if user := getUserByAuthHeader(ah); user != "" {
		claims = newDummyClaims()
		return user, claims, nil
	}

	// 通常の認証を行う
	token := getTokenByAuthHeader(ah)
	if token == "" {
		err := errors.New("token empty error")
		log.Warningf(ctx, "token empty error", zap.Error(err))
		return userID, claims, err
	}

	c, err := getAuthClient(ctx)
	if err != nil {
		log.Warningf(ctx, "getAuthClient", zap.Error(err))
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

// SetCustomClaims ... set custom claim
func (s *firebaseauthDebug) SetCustomClaims(ctx context.Context, userID string, claims *Claims) error {
	c, err := getAuthClient(ctx)
	if err != nil {
		log.Errorf(ctx, "getAuthClient", zap.Error(err))
		return err
	}

	ah := getAuthHeader(ctx)
	if getUserByAuthHeader(ah) == "" {
		err = c.SetCustomUserClaims(ctx, userID, claims.ToMap())
		if err != nil {
			log.Errorf(ctx, "c.SetCustomUserClaims", zap.Error(err))
			return err
		}
	}
	return nil
}

// NewDebug ... DebugFirebaseauthを作成する
func NewDebug() Firebaseauth {
	return &firebaseauthDebug{}
}
