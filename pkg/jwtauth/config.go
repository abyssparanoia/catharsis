package jwtauth

type contextKey string

const (
	// accessTokenExpiredHours ... access token expired hours
	accessTokenExpiredHours = 1

	// refreshTokenExpiredHours ... refresh token expired hours
	refreshTokenExpiredHours = 72

	// claimsContextKey ... jwt claim key
	claimsContextKey contextKey = "jwtauth:claims"
)
