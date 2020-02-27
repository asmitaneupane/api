package restapi

import (

  //nt/http to handle http request
  "net/http"
  //string conversion
  "strconv"
  //we need to import models so that Book can be understood
  // "../models"

  "github.com/gin-gonic/gin"
  "github.com/asmitaneupane/api/models"
)

func getAllBooks(c *gin.Context) {
  var book []models.Book
  err := models.GetAllBooks(&book)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, book)
}

func getBook(c *gin.Context) {

  var book models.Book
  id, _ := strconv.Atoi(c.Param("id"))
  err := models.GetBook(&book, uint(id))

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, book)

}

func addNewBook(c *gin.Context) {
  var book models.Book
  if err := c.ShouldBindJSON(&book); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
    return
  }
  err := models.AddNewBook(&book)

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  err = models.GetBook(&book, book.ID)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, book)

}

func updateBook(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))

  var book models.Book
  if err := c.ShouldBindJSON(&book); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
    return
  }
  if err := models.UpdateBook(&book, uint(id)); err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  if err := models.GetBook(&book, uint(id)); err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, &book)

}

func deleteBook(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  err := models.DeleteBook(uint(id))
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "msg": "book has been deleted",
  })
}