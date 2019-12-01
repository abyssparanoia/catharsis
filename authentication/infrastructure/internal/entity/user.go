package entity

import (
	"github.com/abyssparanoia/catharsis/authentication/domain/model"
	dbmodels "github.com/abyssparanoia/catharsis/pkg/dbmodels/authentication"
	"github.com/volatiletech/null"
)

// User ... user entity
type User struct {
	dbmodels.User
}

// BuildFromModel ... build model from entity
func (e *User) BuildFromModel(m *model.User) {
	e.ID = m.ID
	e.Password = m.Password
	e.DisplayName = m.DisplayName
	e.IconImagePath = m.IconImagePath
	e.BackgroundImagePath = m.BackgroundImagePath
	if m.Profile != nil {
		e.Profile = null.StringFromPtr(m.Profile)
	}
	if m.Email != nil {
		e.Email = null.StringFromPtr(m.Email)
	}
	e.CreatedAt = m.CreatedAt
	e.UpdatedAt = m.UpdatedAt
	if m.DeletedAt != nil {
		e.DeletedAt = null.TimeFromPtr(m.DeletedAt)
	}
}

// OutputModel ... output model from entity
func (e *User) OutputModel() *model.User {
	return &model.User{
		ID:                  e.ID,
		Password:            e.Password,
		DisplayName:         e.DisplayName,
		IconImagePath:       e.IconImagePath,
		BackgroundImagePath: e.BackgroundImagePath,
		Profile:             e.Profile.Ptr(),
		Email:               e.Email.Ptr(),
		CreatedAt:           e.CreatedAt,
		UpdatedAt:           e.UpdatedAt,
		DeletedAt:           e.DeletedAt.Ptr(),
	}
}
