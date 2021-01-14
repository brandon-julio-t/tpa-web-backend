package models

import "github.com/dgrijalva/jwt-go"

type UserJwtClaims struct {
	jwt.StandardClaims
	UserID int64
}
