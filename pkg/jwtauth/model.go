package jwtauth

import jwt "github.com/dgrijalva/jwt-go"

// Claims ... claims model
type Claims struct {
	UserID int64
	jwt.StandardClaims
}

func (c *Claims) set(jc jwt.Claims) {
	mc := jc.(jwt.MapClaims)
	c.UserID = int64(mc["UserID"].(float64))
}
