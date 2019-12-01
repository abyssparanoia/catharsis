package repository

import (
	"context"
	"database/sql"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"go.uber.org/zap"

	"github.com/abyssparanoia/catharsis/authentication/domain/model"
	"github.com/abyssparanoia/catharsis/authentication/domain/repository"
	"github.com/abyssparanoia/catharsis/authentication/infrastructure/internal/entity"
	dbmodels "github.com/abyssparanoia/catharsis/pkg/dbmodels/authentication"
	"github.com/abyssparanoia/catharsis/pkg/log"
)

type user struct {
	db *sql.DB
}

func (m *user) Get(ctx context.Context, userID string) (*model.User, error) {
	user := entity.User{
		dbmodels.User{
			ID:       "user_id",
			Password: "password",
		},
	}
	return user.OutputModel(), nil
}

func (m *user) Create(ctx context.Context, payload struct {
	Password            string
	DisplayName         string
	IconImagePath       string
	BackgroundImagePath string
	Profile             *string
	Email               *string
}) (*model.User, error) {

	user := &entity.User{
		dbmodels.User{
			Password:            payload.Password,
			DisplayName:         payload.DisplayName,
			IconImagePath:       payload.IconImagePath,
			BackgroundImagePath: payload.BackgroundImagePath,
			Profile:             null.StringFromPtr(payload.Profile),
			Email:               null.StringFromPtr(payload.Email),
		},
	}

	if err := user.Insert(ctx, m.db, boil.Infer()); err != nil {
		log.Errorf(ctx, "user.Insert", zap.Error(err))
		return nil, err
	}

	return user.OutputModel(), nil
}

// NewUser ... get user repository
func NewUser(db *sql.DB) repository.User {
	return &user{db}
}
