package models

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Name string `json:"name"`
	Time string `json:"time"`
	jwt.StandardClaims
}
