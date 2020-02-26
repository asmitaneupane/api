package database

import (
  "github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
  var err error
  DB, err = gorm.Open("sqlite3", "test.db")
  if err != nil {
    panic(err)
  }
}
func ConnectDB() *gorm.DB {
  return DB
}