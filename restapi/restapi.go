package restapi

import (
	"github.com/gin-gonic/gin"
)

//InitializeRoutes for the API
func InitializeRoutes(r *gin.RouterGroup) {
	//users endpoint
	r.GET("/users", getAllUsers)
	r.GET("/users/:id", getUser)
	r.POST("/users", addNewUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.GET("/books", getAllBooks)
	r.GET("/books/:id", getBook)
	r.POST("/books", addNewBook)
	r.PUT("/books/:id", updateBook)
	r.DELETE("/books/:id", deleteBook)

	r.GET("/categories", getAllCategories)
	r.GET("/categories/:id", getCategory)
	r.POST("/categories", addNewCategory)
	r.PUT("/categories/:id", updateCategory)
	r.DELETE("/categories/:id", deleteCategory)

	r.GET("/issuedbooks", getAllIssuedBooks)
	r.GET("/issuedbooks/:id", getIssuedBook)
	r.POST("/issuedbooks", addNewIssuedBook)
	r.PUT("/issuedbooks/:id", updateIssuedBook)
	r.DELETE("/issuedbooks/:id", deleteIssuedBook)

	r.GET("/publications", getAllPublications)
	r.GET("/publications/:id", getPublication)
	r.POST("/publications", addNewPublication)
	r.PUT("/publications/:id", updatePublication)
	r.DELETE("/publications/:id", deletePublication)

	r.GET("/returnedbooks", getAllReturnedBooks)
	r.GET("/returnedbooks/:id", getReturnedBook)
	r.POST("/returnedbooks", addNewReturnedBook)
	r.PUT("/returnedbooks/:id", updateReturnedBook)
	r.DELETE("/returnedbooks/:id", deleteReturnedBook)
}
