package security

import (
	"backend/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const SECRET_KEY = "chanhxaucho"

type JwtCustomClaims struct {
	Id   string     `json:"id,omitempty" db:"user_id, omitempty"`
	Role model.Role `json:"role,omitempty" db:"role, omitempty"`
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

func RemoveToken(tokena string) error {
	claims := &JwtCustomClaims{}
	fmt.Println(tokena)
	_, err := jwt.ParseWithClaims(tokena, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return err
	}
	fmt.Println(claims.ExpiresAt)
	claims.ExpiresAt = jwt.NewNumericDate(time.Now())
	fmt.Println(claims.ExpiresAt)

	return nil
}
