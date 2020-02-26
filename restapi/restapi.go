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
}
