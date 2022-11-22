package main

import (
	database "setup/database"
	routes "setup/routes"
)

func main() {
	database.Init()
	routes.Init()
}
