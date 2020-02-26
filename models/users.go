package models

import "github.com/asmitaneupane/api/database"

//GetAllUsers fetch all users
func GetAllUsers(user *[]User) (err error) {
  if err = database.ConnectDB().Find(user).Error; err != nil {
    return err
  }
  return nil
}

//GetUser Endpoint
func GetUser(user *User, id uint) (err error) {
  if err = database.ConnectDB().First(&user, "id=?", id).Error; err != nil {
    return err
  }
  return nil
}

//AddNewUser endpoint hit
func AddNewUser(user *User) (err error) {
  if err = database.ConnectDB().Create(&user).Error; err != nil {
    return err
  }
  return nil
}

//UpdateUser endpoint models
func UpdateUser(user *User, id uint) (err error) {

  if err = database.ConnectDB().Model(&user).Where("id=?", id).Updates(&user).Error; err != nil {
    return err
  }
  return nil

}

//DeleteUser Endpoint
func DeleteUser(id uint) (err error) {
  if err = database.ConnectDB().Where("id = ?", id).Delete(&User{}).Error; err != nil {
    return err
  }
  return nil
}