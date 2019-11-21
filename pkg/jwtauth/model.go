package jwtauth

import jwt "github.com/dgrijalva/jwt-go"

// Claims ... claims model
type Claims struct {
	UserID string
	jwt.StandardClaims
}

func (c *Claims) set(jc jwt.Claims) {
	mc := jc.(jwt.MapClaims)
	c.UserID = mc["UserID"].(string)
}
