package jwtauth

import (
	"context"

	"github.com/abyssparanoia/catharsis/pkg/log"
	jwt "github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type signService struct {
	jwtSignCli *SignClient
}

func (s *signService) GenerateAccessToken(ctx context.Context, claims *Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(s.jwtSignCli.signKey)
	if err != nil {
		log.Errorf(ctx, "token.SignedString: ", zap.Error(err))
		return "", err
	}

	return tokenString, nil
}

// NewSignService ... get new sign service
func NewSignService(jwtSignCli *SignClient) SignService {
	return &signService{jwtSignCli}
}
