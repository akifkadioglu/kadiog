package routes

import (
	"setup/adapter"
	controller "setup/controllers"
	authcontroller "setup/controllers/AuthController"
)

func Api() {
	api := Network.Group("/api")

	//Authenticate
	api.POST("/register", authcontroller.Register)
	api.POST("/login", authcontroller.Login)

	//Auth Users
	auth := api.Group("")
	auth.Use(adapter.AuthAdapter())

	auth.GET("/index", controller.Index)
}
