package restapi

import (

  //nt/http to handle http request
  "net/http"
  //string conversion
  "strconv"
  //we need to import models so that ReturnedBook can be understood
  // "../models"

  "github.com/gin-gonic/gin"
  "github.com/asmitaneupane/api/models"
)

func getAllReturnedBooks(c *gin.Context) {
  var ReturnedBooks []models.ReturnedBook
  err := models.GetAllReturnedBooks(&ReturnedBooks)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, ReturnedBooks)
}

func getReturnedBook(c *gin.Context) {

  var ReturnedBook models.ReturnedBook
  id, _ := strconv.Atoi(c.Param("id"))
  err := models.GetReturnedBook(&ReturnedBook, uint(id))

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, ReturnedBook)

}

func addNewReturnedBook(c *gin.Context) {
  var ReturnedBook models.ReturnedBook
  if err := c.ShouldBindJSON(&ReturnedBook); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
    return
  }
  err := models.AddNewReturnedBook(&ReturnedBook)

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  err = models.GetReturnedBook(&ReturnedBook, ReturnedBook.ID)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, ReturnedBook)

}

func updateReturnedBook(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))

  var ReturnedBook models.ReturnedBook
  if err := c.ShouldBindJSON(&ReturnedBook); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
    return
  }
  if err := models.UpdateReturnedBook(&ReturnedBook, uint(id)); err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  if err := models.GetReturnedBook(&ReturnedBook, uint(id)); err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, &ReturnedBook)

}

func deleteReturnedBook(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  err := models.DeleteReturnedBook(uint(id))
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "msg": "ReturnedBook has been deleted",
  })
}