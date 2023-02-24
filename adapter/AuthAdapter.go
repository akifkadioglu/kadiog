package adapter

import (
	"setup/env"
	"setup/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthAdapter() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte(env.Getenv("APP_KEY")),
	}
	return middleware.JWTWithConfig(config)
}
