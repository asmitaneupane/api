package main

import (
	"fmt"

	"github.com/asmitaneupane/api/database"
	"github.com/asmitaneupane/api/models"
	"github.com/asmitaneupane/api/restapi"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Connecting to DB")
	db := database.ConnectDB()

	fmt.Println("Automigrating User")
	db.AutoMigrate(&models.User{}, &models.Books{})

	//defer execute that function at last
	fmt.Println("Defer Close DB Connection")
	defer db.Close()

	router := gin.Default()
	r := router.Group("/v1")
	restapi.InitializeRoutes(r)

	fmt.Println("Running API")
	router.Run(":8080")
}
