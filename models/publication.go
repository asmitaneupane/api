package models

import "github.com/asmitaneupane/api/database"

//GetAllPublications fetch all Publications
func GetAllPublications(Publication *[]Publication) (err error) {
  if err = database.ConnectDB().Find(Publication).Error; err != nil {
    return err
  }
  return nil
}

//GetPublication Endpoint
func GetPublication(Publication *Publication, id uint) (err error) {
  if err = database.ConnectDB().First(&Publication, "id=?", id).Error; err != nil {
    return err
  }
  return nil
}

//AddNewPublication endpoint hit
func AddNewPublication(Publication *Publication) (err error) {
  if err = database.ConnectDB().Create(&Publication).Error; err != nil {
    return err
  }
  return nil
}

//UpdatePublication endpoint models
func UpdatePublication(Publication *Publication, id uint) (err error) {

  if err = database.ConnectDB().Model(&Publication).Where("id=?", id).Updates(&Publication).Error; err != nil {
    return err
  }
  return nil

}

//DeletePublication Endpoint
func DeletePublication(id uint) (err error) {
  if err = database.ConnectDB().Where("id = ?", id).Delete(&Publication{}).Error; err != nil {
    return err
  }
  return nil
}