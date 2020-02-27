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

//Books Model
type Book struct {
	gorm.Model
	Title string
	Publisher string
	Author string
	Year int64
}