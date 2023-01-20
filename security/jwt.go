package security

import (
	"backend/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const SECRET_KEY = "chanhxaucho"

type JwtCustomClaims struct {
	Id   string `json:"id,omitempty" db:"user_id, omitempty"`
	Role string `json:"role,omitempty" db:"role, omitempty"`
	jwt.RegisteredClaims
}

func GenToken(user model.User) (string, error) {
	claims := &JwtCustomClaims{
		user.Id,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return result, nil
}
