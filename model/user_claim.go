package model

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	Id   string `json:"id" db:"user_id, omitempty"`
	Role Role   `json:"role" db:"role, omitempty"`
	jwt.StandardClaims
}
