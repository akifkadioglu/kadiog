package database

import (
	"setup/helpers"
	models "setup/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (d *MySql) main() *gorm.DB {
	var dns = helpers.GoDotEnvVariable("DB_USERNAME") + ":" + helpers.GoDotEnvVariable("DB_PASSWORD") + "@tcp(" + helpers.GoDotEnvVariable("DB_HOST") + ":" + helpers.GoDotEnvVariable("DB_PORT") + ")/" + helpers.GoDotEnvVariable("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"

	d.db, err = gorm.Open(mysql.Open(dns))
	d.db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err.Error())
	}
	return d.db
}
