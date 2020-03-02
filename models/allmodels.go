package models

import "github.com/jinzhu/gorm"

//User Model
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Address   string
	UserName  string
	Password  string
}

//Books Model
type Book struct {
	gorm.Model
	Title      string
	Publisher  string
	Author     string
	Year       int64
	CategoryID int64
}

//Category Model
type Category struct {
	gorm.Model
	CName       string
	Description string
	Book        []Book `gorm:"foreignkey:CategoryID"`
}

//Publication Model
type Publication struct {
	gorm.Model
	PName    string
	PAddress string
	Phone    string
	Email    string
}
