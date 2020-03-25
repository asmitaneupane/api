package models

import "github.com/asmitaneupane/api/database"

//GetAllReturnedBooks fetch all ReturnedBooks
func GetAllReturnedBooks(ReturnedBook *[]ReturnedBook) (err error) {
  if err = database.ConnectDB().Find(ReturnedBook).Error; err != nil {
    return err
  }
  return nil
}

//GetReturnedBook Endpoint
func GetReturnedBook(ReturnedBook *ReturnedBook, id uint) (err error) {
  if err = database.ConnectDB().First(&ReturnedBook, "id=?", id).Error; err != nil {
    return err
  }
  return nil
}

//AddNewReturnedBook endpoint hit
func AddNewReturnedBook(ReturnedBook *ReturnedBook) (err error) {
  if err = database.ConnectDB().Create(&ReturnedBook).Error; err != nil {
    return err
  }
  return nil
}

//UpdateReturnedBook endpoint models
func UpdateReturnedBook(ReturnedBook *ReturnedBook, id uint) (err error) {

  if err = database.ConnectDB().Model(&ReturnedBook).Where("id=?", id).Updates(&ReturnedBook).Error; err != nil {
    return err
  }
  return nil

}

//DeleteReturnedBook Endpoint
func DeleteReturnedBook(id uint) (err error) {
  if err = database.ConnectDB().Where("id = ?", id).Delete(&ReturnedBook{}).Error; err != nil {
    return err
  }
  return nil
}