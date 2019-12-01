package model

// User ... user domain model
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
