package database

import (
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type Idatabase interface {
	main()
}

type Mysql struct{}
type SQLite struct{}

func Init() {
	var database Mysql //If you want to use SQLite, write SQLite instead of Mysql.
	Idatabase.main(&database)
}

func DBManager() gorm.DB {
	return *db
}
