package jwtauth

type contextKey string

const (
	// TokenExpiredHours ... expired hours
	tokenExpiredHours = 1

	// claimsContextKey ... jwt claim key
	claimsContextKey contextKey = "jwtauth:claims"
)
