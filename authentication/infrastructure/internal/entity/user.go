package entity

import "github.com/abyssparanoia/catharsis/authentication/domain/model"

// User ... user entity
type User struct {
	ID       string `db:"id"`
	Password string `db:"password"`
	BaseEntity
}

// BuildFromModel ... build model from entity
func (e *User) BuildFromModel(m *model.User) {
	e.ID = m.ID
	e.Password = m.Password
}

// OutputModel ... output model from entity
func (e *User) OutputModel() *model.User {
	return &model.User{
		ID:       e.ID,
		Password: e.Password,
	}
}
