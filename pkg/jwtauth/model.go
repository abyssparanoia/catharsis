package jwtauth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Claims ... claims model
type Claims struct {
	UserID string
	jwt.StandardClaims
}

// Set ... set claims
func (c *Claims) Set(userID string) {
	c.UserID = userID
}
