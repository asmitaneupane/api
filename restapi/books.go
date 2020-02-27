package restapi

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/asmitaneupane/api/models"
)

func getAllBooks(c *gin.Context)  {
	var books []models.Books
	err :=models.getAllBooks(&books)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, books)
}

func getBook(c *gin.Context) {

	var books models.Books
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.getBook(&books, uint(id))
  
	if err != nil {
	  c.JSON(http.StatusNotFound, gin.H{
		"msg": err.Error(),
	  })
	  return
	}
  
	c.JSON(http.StatusOK, books)
  
  }
  
  func addNewBook(c *gin.Context) {
	var books models.Books
	if err := c.ShouldBindJSON(&books); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	  return
	}
	err := models.addNewBook(&books)
  
	if err != nil {
	  c.JSON(http.StatusNotFound, gin.H{
		"msg": err.Error(),
	  })
	  return
	}
  
	err = models.getBook(&books, books.ID)
	if err != nil {
	  c.JSON(http.StatusNotFound, gin.H{
		"msg": err.Error(),
	  })
	  return
	}
  
	c.JSON(http.StatusOK, books)
  
  }
  
  func updateBooks(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
  
	var books models.Books
	if err := c.ShouldBindJSON(&books); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	  return
	}
	if err := models.updateBooks(&books, uint(id)); err != nil {
	  c.JSON(http.StatusNotFound, gin.H{
		"msg": err.Error(),
	  })
	  return
	}
  
	if err := models.getBook(&books, uint(id)); err != nil {
	  c.JSON(http.StatusNotFound, gin.H{
		"msg": err.Error(),
	  })
	  return
	}
	c.JSON(http.StatusOK, &books)
  
  }
  
  func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.deleteBook(uint(id))
	if err != nil {
	  c.JSON(http.StatusNotFound, gin.H{
		"msg": err.Error(),
	  })
	  return
	}
	c.JSON(http.StatusOK, gin.H{
	  "msg": "Book has been deleted",
	})
  }