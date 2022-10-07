package adapter

import (
	"setup/env"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthAdapter() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		SigningKey: []byte(env.GoDotEnvVariable("APP_KEY")),
	}
	return middleware.JWTWithConfig(config)
}
