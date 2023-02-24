package database

import (
	"setup/env"
	models "setup/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
If you are using Windows
	- Download the GCC from here http://tdm-gcc.tdragon.net/download
	- Install GCC
*/

func (d *SQLite) main() {
	db, err = gorm.Open(sqlite.Open("./database/"+env.Getenv(env.DB_DATABASE)+".db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
}
