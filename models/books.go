package models

import "github.com/asmitaneupane/api/database"

//GetAllBooks fetch all books
func GetAllBooks(book *[]Book) (err error) {
  if err = database.ConnectDB().Find(book).Error; err != nil {
    return err
  }
  return nil
}

//GetBook Endpoint
func GetBook(book *Book, id uint) (err error) {
  if err = database.ConnectDB().First(&book, "id=?", id).Error; err != nil {
    return err
  }
  return nil
}

//AddNewBook endpoint hit
func AddNewBook(book *Book) (err error) {
  if err = database.ConnectDB().Create(&book).Error; err != nil {
    return err
  }
  return nil
}

//UpdateBook endpoint models
func UpdateBook(book *Book, id uint) (err error) {

  if err = database.ConnectDB().Model(&book).Where("id=?", id).Updates(&book).Error; err != nil {
    return err
  }
  return nil

}

//DeleteBook Endpoint
func DeleteBook(id uint) (err error) {
  if err = database.ConnectDB().Where("id = ?", id).Delete(&Book{}).Error; err != nil {
    return err
  }
  return nil
}