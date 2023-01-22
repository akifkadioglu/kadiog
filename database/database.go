package database

import (
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type Idatabase interface {
	main()
}
type MySql struct {
	db *gorm.DB
}
type Sqlite struct {
	db *gorm.DB
}

func Init() {
	Idatabase.main(&MySql{db: db})
}

func DBManager() gorm.DB {
	return *db
}
