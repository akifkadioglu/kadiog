package database

import (
	"setup/environment"
	"setup/helpers"
	models "setup/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (d *Mysql) main() {
	var dns = helpers.GoDotEnvVariable(environment.DB_USERNAME) + ":" + helpers.GoDotEnvVariable(environment.DB_PASSWORD) + "@tcp(" + helpers.GoDotEnvVariable(environment.DB_HOST) + ":" + helpers.GoDotEnvVariable(environment.DB_PORT) + ")/" + helpers.GoDotEnvVariable(environment.DB_DATABASE) + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dns))
	db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
}
