package models

import "github.com/jinzhu/gorm"

//User Model
type User struct {
	gorm.Model
	FirstName string
	LastName string
	Address string
	UserName string
	Password string
}