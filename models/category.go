package models

import "github.com/asmitaneupane/api/database"

//GetAllCategories fetch all categories
func GetAllCategories(category *[]Book) (err error) {
  if err = database.ConnectDB().Preload("Product").Find(category).Error; err != nil {
    return err
  }
  return nil
}

//GetCategories Endpoint
func GetCategories(category *Book, id uint) (err error) {
  if err = database.ConnectDB().First(&category, "id=?", id).Error; err != nil {
    return err
  }
  return nil
}

//AddNewCategories endpoint hit
func AddNewCategories(category *Book) (err error) {
  if err = database.ConnectDB().Create(&category).Error; err != nil {
    return err
  }
  return nil
}

//UpdateCategories endpoint models
func UpdateCategories(category *Book, id uint) (err error) {

  if err = database.ConnectDB().Model(&category).Where("id=?", id).Updates(&category).Error; err != nil {
    return err
  }
  return nil

}

//DeleteCategories Endpoint
func DeleteCategories(id uint) (err error) {
  if err = database.ConnectDB().Where("id = ?", id).Delete(&Book{}).Error; err != nil {
    return err
  }
  return nil
}