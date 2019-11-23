package model

// User ... user domain model
type User struct {
	ID        string `json:"id"`
	Password  string `json:"-"`
	CreatedAt int64  `json:"created_at"`
}
