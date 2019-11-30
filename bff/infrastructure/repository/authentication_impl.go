package repository

import (
	"context"

	"github.com/abyssparanoia/catharsis/bff/domain/repository"
	"github.com/abyssparanoia/catharsis/pkg/log"
	pb "github.com/abyssparanoia/catharsis/proto"
	"go.uber.org/zap"
)

type authentication struct {
	authenticationClient pb.AuthenticationClient
}

func (r *authentication) SignIn(ctx context.Context, userID string, password string) (string, string, error) {

	message := &pb.SignInMessageRequest{UserId: userID, Password: password}
	res, err := r.authenticationClient.SignIn(ctx, message)
	if err != nil {
		log.Errorf(ctx, "r.authenticationClient.SignIn", zap.Error(err))
		return "", "", err
	}

	return res.AccessToken, res.RefreshToken, nil
}

// NewAuthentication ... get authentication repository
func NewAuthentication(authenticationClient pb.AuthenticationClient) repository.Authentication {
	return &authentication{authenticationClient}
}
