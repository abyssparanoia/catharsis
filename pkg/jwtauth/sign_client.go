package jwtauth

import (
	"crypto/rsa"
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
)

// SignClient ... jwt client for sign
type SignClient struct {
	signKey *rsa.PrivateKey
}

// NewSignClient ... get sign client
func NewSignClient(signKeyPath string) *SignClient {
	signBytes, err := ioutil.ReadFile(signKeyPath)
	if err != nil {
		panic(err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	return &SignClient{signKey}
}
