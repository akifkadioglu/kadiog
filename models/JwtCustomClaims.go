package models

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Time   string `json:"time"`
	jwt.StandardClaims
}
