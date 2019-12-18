package test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/abyssparanoia/catharsis/authentication/domain/model"
	mock_repository "github.com/abyssparanoia/catharsis/authentication/domain/repository/mock"
	"github.com/abyssparanoia/catharsis/authentication/service"

	"github.com/abyssparanoia/catharsis/pkg/jwtauth"
	mock_jwt_auth "github.com/abyssparanoia/catharsis/pkg/jwtauth/mock"
)

type mockExpectedUserGet struct {
	userID string
	user   *model.User
	err    error
}

type mockExpectedJwtAuthGenerateToken struct {
	claims       *jwtauth.Claims
	accessToken  string
	refreshToken string
	err          error
}

func Test_authentication_sign_in(t *testing.T) {
	type args struct {
		ctx      context.Context
		userID   string
		password string
	}

	type mock struct {
		getUser       mockExpectedUserGet
		generateToken mockExpectedJwtAuthGenerateToken
	}

	type result struct {
		accessToken  string
		refreshToken string
		err          error
	}

	tests := []struct {
		name    string
		args    args
		mock    mock
		want    result
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				userID:   "user_id",
				password: "password",
			},
			mock: mock{
				getUser: mockExpectedUserGet{
					userID: "user_id",
					user: &model.User{
						ID:        "user_id",
						Password:  "password",
						CreatedAt: time.Now(),
					},
					err: nil,
				},
				generateToken: mockExpectedJwtAuthGenerateToken{
					claims: &jwtauth.Claims{
						UserID: "user_id",
					},
					accessToken:  "access_token",
					refreshToken: "refresh_token",
					err:          nil,
				},
			},
			want: result{
				accessToken:  "access_token",
				refreshToken: "refresh_token",
				err:          nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			userReposiotry := mock_repository.NewMockUser(mc)
			userReposiotry.EXPECT().
				Get(gomock.Any(), tt.mock.getUser.userID).
				Return(tt.mock.getUser.user, tt.mock.getUser.err)

			jwtauthSign := mock_jwt_auth.NewMockJwtauthSign(mc)
			jwtauthSign.EXPECT().
				GenerateToken(gomock.Any(), tt.mock.generateToken.claims).
				Return(tt.mock.generateToken.accessToken, tt.mock.generateToken.refreshToken, tt.mock.generateToken.err)

			s := service.NewAuthentication(userReposiotry, jwtauthSign)

			accessToken, refreshToken, err := s.SignIn(tt.args.ctx, tt.args.userID, tt.args.password)

			if (err != nil) != tt.wantErr {
				t.Errorf("s.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if assert.NoError(t, err) {
				assert.Equal(t, tt.want.accessToken, accessToken)
				assert.Equal(t, tt.want.refreshToken, refreshToken)
			}

		})
	}

}
