package model

import (
	"time"
)

// User ... user domain model
type User struct {
	ID                  string
	Password            string
	DisplayName         string
	IconImagePath       string
	BackgroundImagePath string
	Profile             *string
	Email               *string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time
}
