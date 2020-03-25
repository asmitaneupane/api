package models

import "github.com/asmitaneupane/api/database"

//GetAllCategories fetch all categories
func GetAllCategories(category *[]Category) (err error) {
	if err = database.ConnectDB().Preload("Book").Find(category).Error; err != nil {
		return err
	}
	return nil
}

//GetCategory Endpoint
func GetCategory(category *Category, id uint) (err error) {
	if err = database.ConnectDB().First(&category, "id=?", id).Error; err != nil {
		return err
	}
	return nil
}

//AddNewCategory endpoint hit
func AddNewCategory(category *Category) (err error) {
	if err = database.ConnectDB().Create(&category).Error; err != nil {
		return err
	}
	return nil
}

//UpdateCategory endpoint models
func UpdateCategory(category *Category, id uint) (err error) {

	if err = database.ConnectDB().Model(&category).Where("id=?", id).Updates(&category).Error; err != nil {
		return err
	}
	return nil

}

//DeleteCategory Endpoint
func DeleteCategory(id uint) (err error) {
	if err = database.ConnectDB().Where("id = ?", id).Delete(&Category{}).Error; err != nil {
		return err
	}
	return nil
}
