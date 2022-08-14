package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("IjsfhrnaabHBJfklawjn676hBHBdkbvleKf")

type JWTclaim struct {
	Username string
	jwt.RegisteredClaims
}