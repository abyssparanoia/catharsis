package model

// User ... user domain model
type User struct {
	ID        string `json:"id"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}
