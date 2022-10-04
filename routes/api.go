package routes

import (
	"github.com/labstack/echo/v4"
	controllers "setup/controllers"
)

func Api() *echo.Echo {
	api := E.Group("/api")
	api.GET("/getUser", controllers.GetUsers)
	return E
}
