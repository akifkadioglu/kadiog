package database

import (
	"chatApp/env"
	models "chatApp/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error
var dbAddress = env.GoDotEnvVariable("DB_USERNAME") + ":" +
	env.GoDotEnvVariable("DB_PASSWORD") +
	"@tcp(" +
	env.GoDotEnvVariable("DB_HOST") +
	":" +
	env.GoDotEnvVariable("DB_PORT") +
	")/"

func Init() gorm.DB {
	dns := dbAddress + env.GoDotEnvVariable("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dns))
	db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
	return *db
}
func DBManager() gorm.DB {
	return *db
}
