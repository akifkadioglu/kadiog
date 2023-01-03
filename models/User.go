package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"not null; size:60"`
	Email    string `json:"email" gorm:"not null; size:60"`
	Username string `json:"username" gorm:"not null; size:60"`
	Password string `json:"-" gorm:"not null"`
}

/*
	Hooks are functions that are called before or after performing an action.
	For example, when deleting a user, you can save that user in a table called "deleted_users".
	Notice "(u *Users)" according to the model you created Modify this struct.
*/

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("AfterFind Function called")

	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate Function called")

	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate Function called")

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeUpdate Function called")

	return
}

func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Println("AfterUpdate Function called")

	return
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("BeforeDelete Function called")

	return
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("AfterDelete Function called")

	return
}
