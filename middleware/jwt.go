package middleware

import (
	"backend/model"
	"backend/security"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(model.JwtCustomClaims)
		},
		SigningKey: []byte(security.SECRET_KEY),
	}
	/*
		config := middleware.JWTConfig{
			Claims:     &model.JwtCustomClaims{},
			SigningKey: []byte(security.SECRET_KEY),
		} */
	return echojwt.WithConfig(config)
}
