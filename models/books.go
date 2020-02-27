package models

import "github.com/asmitaneupane/api/database"

//getAllBooks fetch all books
func getAllBooks(books *[]Books) (err error) {
	if err = database.ConnectDB().Find(books).Error; err != nil {
		return err
	}
	return nil
}

//GetBooks Endpoint
func GetBooks(books *Books, id uint) (err error) {
	if err = database.ConnectDB().First(&books, "id=?", id).Error; err != nil {
		return err
	}
	return nil
}

//AddNewBook endpoint hit
func AddNewBook(books *Books) (err error) {
	if err = database.ConnectDB().Create(&books).Error; err != nil {
		return err
	}
	return nil
}

//UpdateBooks endpoint models
func UpdateBooks(books *Books, id uint) (err error) {

	if err = database.ConnectDB().Model(&books).Where("id=?", id).Updates(&books).Error; err != nil {
		return err
	}
	return nil

}

//DeleteBook Endpoint
func DeleteBook(id uint) (err error) {
	if err = database.ConnectDB().Where("id = ?", id).Delete(&Books{}).Error; err != nil {
		return err
	}
	return nil
}
