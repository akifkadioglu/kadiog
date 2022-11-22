package adapter

import (
	"setup/helpers"
	"setup/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthAdapter() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte(helpers.GoDotEnvVariable("APP_KEY")),
	}
	return middleware.JWTWithConfig(config)
}
