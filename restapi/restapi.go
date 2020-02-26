package restapi

import(
	"github.com/gin-gonic/gin"
	"github.com/asmitaneupane/api/models"
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
