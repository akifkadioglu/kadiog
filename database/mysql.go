package database

import (
	"setup/env"
	models "setup/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (d *Mysql) main() {
	var dns = env.Getenv(env.DB_USERNAME) + ":" + env.Getenv(env.DB_PASSWORD) + "@tcp(" + env.Getenv(env.DB_HOST) + ":" + env.Getenv(env.DB_PORT) + ")/" + env.Getenv(env.DB_DATABASE) + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dns))
	db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
}
