package main

import (
	database "chatApp/database"
	routes "chatApp/routes"
)

func main() {
	database.Init()
	routes.Init()
}
