package restapi

import (

  //nt/http to handle http request
  "net/http"
  //string conversion
  "strconv"
  //we need to import models so that User can be understood
  // "../models"

  "github.com/gin-gonic/gin"
)

func getAllUsers(c *gin.Context) {
  var user []models.User
  err := models.GetAllUsers(&user)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, user)
}

func getUser(c *gin.Context) {

  var user models.User
  id, _ := strconv.Atoi(c.Param("id"))
  err := models.GetUser(&user, uint(id))

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, user)

}

func addNewUser(c *gin.Context) {
  var user models.User
  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
    return
  }
  err := models.AddNewUser(&user)

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  err = models.GetUser(&user, user.ID)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, user)

}

func updateUser(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))

  var user models.User
  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
    return
  }
  if err := models.UpdateUser(&user, uint(id)); err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  if err := models.GetUser(&user, uint(id)); err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, &user)

}

func deleteUser(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  err := models.DeleteUser(uint(id))
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "msg": "user has been deleted",
  })
}