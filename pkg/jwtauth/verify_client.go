package jwtauth

import (
	"crypto/rsa"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

// VerifyClient ... jwt verify client
type VerifyClient struct {
	verifyKey *rsa.PublicKey
}

// NewVerifyClient ... get verify client
func NewVerifyClient(verifyKeyPath string) *VerifyClient {
	verifyBytes, err := ioutil.ReadFile(verifyKeyPath)
	if err != nil {
		panic(err)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}

	return &VerifyClient{verifyKey}
}
