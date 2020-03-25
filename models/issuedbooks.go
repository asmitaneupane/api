package models

import "github.com/asmitaneupane/api/database"

//GetAllIssuedBooks fetch all issuedbooks
func GetAllIssuedBooks(issuedbook *[]IssuedBook) (err error) {
	if err = database.ConnectDB().Find(issuedbook).Error; err != nil {
		return err
	}
	return nil
}

//GetIssuedBook Endpoint
func GetIssuedBook(issuedbook *IssuedBook, id uint) (err error) {
	if err = database.ConnectDB().First(&issuedbook, "id=?", id).Error; err != nil {
		return err
	}
	return nil
}

//AddNewIssuedBook endpoint hit
func AddNewIssuedBook(issuedbook *IssuedBook) (err error) {
	if err = database.ConnectDB().Create(&issuedbook).Error; err != nil {
		return err
	}
	return nil
}

//UpdateIssuedBook endpoint models
func UpdateIssuedBook(issuedbook *IssuedBook, id uint) (err error) {

	if err = database.ConnectDB().Model(&issuedbook).Where("id=?", id).Updates(&issuedbook).Error; err != nil {
		return err
	}
	return nil

}

//DeleteIssuedBook Endpoint
func DeleteIssuedBook(id uint) (err error) {
	if err = database.ConnectDB().Where("id = ?", id).Delete(&IssuedBook{}).Error; err != nil {
		return err
	}
	return nil
}
