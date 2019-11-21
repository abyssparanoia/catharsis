package jwtauth

import "context"

// GetClaimsFromContext ... get claims from context
func GetClaimsFromContext(ctx context.Context) Claims {
	return ctx.Value(claimsContextKey).(Claims)
}

// SetClaimsToContext ... set claims to context
func SetClaimsToContext(ctx context.Context, claims *Claims) context.Context {
	return context.WithValue(ctx, claimsContextKey, &claims)
}
