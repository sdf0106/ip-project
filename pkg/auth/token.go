package auth

import "github.com/golang-jwt/jwt"

type TokenWithClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id""`
}
