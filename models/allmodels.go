package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

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
	Year       string
	CategoryID int64
}

//Category Model
type Category struct {
	gorm.Model
	CName       string
	Description string
	Book        []Book `gorm:"foreignkey:CategoryID"`
}

//Issue Model
type IssuedBook struct {
	gorm.Model
	UserID    uint
	BookID    uint
	IssueDate time.Time
}

//ReturnedBook Model
type ReturnedBook struct {
	gorm.Model
	BookID     uint
	UserID     uint
	ReturnDate time.Time
}

//Publication Model
type Publication struct {
	gorm.Model
	PName    string
	PAddress string
	Phone    string
	Email    string
}
