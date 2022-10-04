package routes

import (
	controllers "chatApp/controllers"
	"github.com/labstack/echo/v4"
)

func Api()*echo.Echo {
	api := E.Group("/api")
	api.GET("/getUser", controllers.GetUsers)
	return E
}
