package entity

import "github.com/abyssparanoia/catharsis/authentication/domain/model"

// User ... user entity
type User struct {
	ID                  string
	Password            string
	DisplayName         string
	IconImagePath       string
	BackgroundImagePath string
	Profile             string
	Email               string
	CreatedAt           int64
	UpdatedAt           int64
	DeletedAt           int64
}

// BuildFromModel ... build model from entity
func (e *User) BuildFromModel(m *model.User) {
	e.ID = m.ID
	e.Password = m.Password
	e.DisplayName = m.DisplayName
	e.IconImagePath = m.IconImagePath
	e.BackgroundImagePath = m.BackgroundImagePath
	e.Profile = m.Profile
	e.Email = m.Email
	e.CreatedAt = m.CreatedAt
	e.UpdatedAt = m.UpdatedAt
	e.DeletedAt = m.DeletedAt
}

// OutputModel ... output model from entity
func (e *User) OutputModel() *model.User {
	return &model.User{
		ID:                  e.ID,
		Password:            e.Password,
		DisplayName:         e.DisplayName,
		IconImagePath:       e.IconImagePath,
		BackgroundImagePath: e.BackgroundImagePath,
		Profile:             e.Profile,
		Email:               e.Email,
		CreatedAt:           e.CreatedAt,
		UpdatedAt:           e.UpdatedAt,
		DeletedAt:           e.DeletedAt,
	}
}
