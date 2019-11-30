package repository

// Authentication ... authentication reposiroty interface
type Authentication interface {
	SignIn(ctx, userID string, password string) (string, string, error)
}
