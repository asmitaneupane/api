package restapi

import (

	//nt/http to handle http request
	"net/http"
	//string conversion
	"strconv"
	//we need to import models so that IssuedBook can be understood
	// "../models"

	"github.com/asmitaneupane/api/models"
	"github.com/gin-gonic/gin"
)

func getAllIssuedBooks(c *gin.Context) {
	var issuedbooks []models.IssuedBook
	err := models.GetAllIssuedBooks(&issuedbooks)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, issuedbooks)
}

func getIssuedBook(c *gin.Context) {

	var issuedbook models.IssuedBook
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetIssuedBook(&issuedbook, uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, issuedbook)

}

func addNewIssuedBook(c *gin.Context) {
	var issuedbook models.IssuedBook
	if err := c.ShouldBindJSON(&issuedbook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	err := models.AddNewIssuedBook(&issuedbook)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	err = models.GetIssuedBook(&issuedbook, issuedbook.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, issuedbook)

}

func updateIssuedBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var issuedbook models.IssuedBook
	if err := c.ShouldBindJSON(&issuedbook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if err := models.UpdateIssuedBook(&issuedbook, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := models.GetIssuedBook(&issuedbook, uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &issuedbook)

}

func deleteIssuedBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteIssuedBook(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "issuedbook has been deleted",
	})
}
