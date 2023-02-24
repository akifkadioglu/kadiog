package main

import (
	"setup/database"
	"setup/env"
	"setup/localization"
	"setup/routes"
)

func main() {
	localization.Init()
	env.Init()
	database.Init()
	routes.Init()
}
