package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null; size:60"`
	Email    string `gorm:"not null; size:60"`
	Username string `gorm:"not null; size:60"`
	Biograpy string `gorm:""`
	Password string `gorm:"not null"`
}
