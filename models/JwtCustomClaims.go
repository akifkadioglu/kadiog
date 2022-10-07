package models

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}
