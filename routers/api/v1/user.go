package v1

import (
	"360-evaluation/models/response"
	"360-evaluation/service/user_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRouter(r *gin.RouterGroup) {
	userApi := r.Group("/user")
	userApi.GET("/hello", Hello)
	userApi.POST("/", AddUser)
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, user_service.Hello("deo"))
}

func AddUser(c *gin.Context) {
	name := "deo"
	err, user := user_service.AddUser(name)
	if err != nil {
		response.Err(err.Error(), c)
		return
	}
	response.OK(response.UserResponse{User: user}, c)
}
