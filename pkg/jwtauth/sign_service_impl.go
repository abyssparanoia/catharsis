package jwtauth

import (
	"context"
	"time"

	"github.com/abyssparanoia/catharsis/pkg/log"
	jwt "github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type signService struct {
	jwtSignCli *SignClient
}

func (s *signService) GenerateToken(ctx context.Context, claims *Claims) (string, string, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	accessTokenString, err := accessToken.SignedString(s.jwtSignCli.signKey)
	if err != nil {
		log.Errorf(ctx, "accessToken.SignedString: ", zap.Error(err))
		return "", "", err
	}

	refreshToken := jwt.New(jwt.SigningMethodRS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * refreshTokenExpiredHours).Unix()

	refreshTokenString, err := refreshToken.SignedString(s.jwtSignCli.signKey)
	if err != nil {
		log.Errorf(ctx, "refreshToken.SignedString: ", zap.Error(err))
	}

	return accessTokenString, refreshTokenString, nil
}

// NewSignService ... get new sign service
func NewSignService(jwtSignCli *SignClient) SignService {
	return &signService{jwtSignCli}
}
