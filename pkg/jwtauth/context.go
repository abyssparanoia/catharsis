package jwtauth

import "context"

// GetClaims ... get claims from context
func GetClaims(ctx context.Context) Claims {
	return ctx.Value(claimsContextKey).(Claims)
}

func setClaims(ctx context.Context, claims Claims) context.Context {
	return context.WithValue(ctx, claimsContextKey, claims)
}
