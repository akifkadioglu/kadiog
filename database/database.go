package database

import (
	"setup/helpers"
	models "setup/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var dbAddress = helpers.GoDotEnvVariable("DB_USERNAME") + ":" +
	helpers.GoDotEnvVariable("DB_PASSWORD") +
	"@tcp(" +
	helpers.GoDotEnvVariable("DB_HOST") +
	":" +
	helpers.GoDotEnvVariable("DB_PORT") +
	")/" +
	helpers.GoDotEnvVariable("DB_DATABASE")

func Init() {
	dns := dbAddress + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dns))
	db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
}
func DBManager() gorm.DB {
	return *db
}
