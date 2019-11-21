package jwtauth

type contextKey string

const (
	// TokenExpiredDays ... expired days
	tokenExpiredDays = 7

	// claimsContextKey ... jwt claim key
	claimsContextKey contextKey = "jwtauth:claims"
)
