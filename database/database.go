package database

import (
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type Idatabase interface {
	main() *gorm.DB
}
type MySql struct {
	db *gorm.DB
}
type Sqlite struct {
	db *gorm.DB
}

func Init() {
	var dbBuilder Idatabase = &MySql{db: db}
	db = dbBuilder.main()
}

func DBManager() gorm.DB {
	return *db
}
