package entity

// User ... user entity
type User struct {
	ID       string `db:"id"`
	Password string `db:"password"`
	BaseEntity
}
