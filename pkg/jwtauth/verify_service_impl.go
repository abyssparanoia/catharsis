package jwtauth

import (
	"context"

	"github.com/abyssparanoia/catharsis/pkg/log"
	jwt "github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type verifyService struct {
	jwtVerifyClient *VerifyClient
}

func (s *verifyService) Authenticate(ctx context.Context, token string) (*Claims, error) {

	parser := new(jwt.Parser)

	claims := &Claims{}

	_, err := parser.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {

		return s.jwtVerifyClient.verifyKey, nil
	})

	if err != nil {
		log.Errorf(ctx, "parser.ParseWithClaims: ", zap.Error(err))
		return nil, err
	}

	return claims, nil

}

// NewVerifyService ... get new verify service
func NewVerifyService(jwtVerifyClient *VerifyClient) VerifyService {
	return &verifyService{jwtVerifyClient}
}
