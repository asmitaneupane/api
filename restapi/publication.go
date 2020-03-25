package restapi

import (

  //nt/http to handle http request
  "net/http"
  //string conversion
  "strconv"
  //we need to import models so that Publication can be understood
  // "../models"

  "github.com/gin-gonic/gin"
  "github.com/asmitaneupane/api/models"
)

func getAllPublications(c *gin.Context) {
  var Publications []models.Publication
  err := models.GetAllPublications(&Publications)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, Publications)
}

func getPublication(c *gin.Context) {

  var Publication models.Publication
  id, _ := strconv.Atoi(c.Param("id"))
  err := models.GetPublication(&Publication, uint(id))

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, Publication)

}

func addNewPublication(c *gin.Context) {
  var Publication models.Publication
  if err := c.ShouldBindJSON(&Publication); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
    return
  }
  err := models.AddNewPublication(&Publication)

  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  err = models.GetPublication(&Publication, Publication.ID)
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, Publication)

}

func updatePublication(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))

  var Publication models.Publication
  if err := c.ShouldBindJSON(&Publication); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
    return
  }
  if err := models.UpdatePublication(&Publication, uint(id)); err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }

  if err := models.GetPublication(&Publication, uint(id)); err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, &Publication)

}

func deletePublication(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))
  err := models.DeletePublication(uint(id))
  if err != nil {
    c.JSON(http.StatusNotFound, gin.H{
      "msg": err.Error(),
    })
    return
  }
  c.JSON(http.StatusOK, gin.H{
    "msg": "Publication has been deleted",
  })
}