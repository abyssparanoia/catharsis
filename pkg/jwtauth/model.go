package jwtauth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Claims ... claims model
type Claims struct {
	UserID string
	jwt.StandardClaims
}

func (c *Claims) set(userID string) {
	c.UserID = userID
	c.ExpiresAt = time.Now().Add(time.Hour * tokenExpiredHours).Unix()
}
