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

// NewClient ... get sign client
func NewClient(signKeyPath string) *SignClient {
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
