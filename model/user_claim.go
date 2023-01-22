package model

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	Id   string `json:"id,omitempty" db:"user_id, omitempty"`
	Role Role   `json:"role,omitempty" db:"role, omitempty"`
	jwt.StandardClaims
}
