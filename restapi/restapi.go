package restapi

import(
	"github.com/gin-gonic/gin"
)

//InitializeRoutes for the API
func InitializeRoutes(r*gin.RouterGroup) {
	//users endpoint
	r.GET("/users", getAllUsers)
  	r.GET("/users/:id", getUser)
  	r.POST("/users", addNewUser)
  	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.GET("/books", getAllBooks)
	r.GET("/books/:id", getBook)
	r.POST("/books", addNewBook)
	r.PUT("/books/:id", updateBooks)
	r.DELETE("/books/:id", deleteBook)

	r.GET("/categories", getAllBooks)
	r.GET("/categories/:id", getBook)
	r.POST("/categories", addNewBook)
	r.PUT("/categories/:id", updateBooks)
	r.DELETE("/categories/:id", deleteBook)

}
