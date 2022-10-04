package routes

import (
	controllers "setup/controllers"
)

func Api() {
	api := E.Group("/api")
	api.GET("/getUser", controllers.GetUsers)
}
