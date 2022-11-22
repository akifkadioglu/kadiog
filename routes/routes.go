package routes

import (
	"setup/adapter"
	authcontroller "setup/controllers/AuthController"
	indexcontroller "setup/controllers/IndexController"

	"github.com/labstack/echo/v4"
)

/*
The "network" variable is used for publicly accessible urls.
The "auth" variable is urls that can only be used by users who own the token.
*/

func Web(Network *echo.Group) {
	auth := Network.Group("")
	auth.Use(adapter.AuthAdapter())

	//user
	Network.POST("/login", authcontroller.Login)
	Network.POST("/register", authcontroller.Register)

	Network.GET("/", indexcontroller.Index)
}
