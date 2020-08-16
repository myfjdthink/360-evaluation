package v1

import (
	"360-evaluation/middleware"
	"360-evaluation/models/request"
	"360-evaluation/models/response"
	"360-evaluation/service"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	userApi := r.Group("/user")
	userApi.GET("/hello", middleware.Auth(), Hello)
	userApi.POST("/", AddUser)
	userApi.POST("/login", LoginByPassword)
}

func Hello(c *gin.Context) {
	response.OK(service.Hello("deo"), c)
}

func AddUser(c *gin.Context) {
	var p request.AddUser
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.Err(err, c)
		return
	}

	err, user := service.AddUser(p.Name)
	if err != nil {
		response.Err(err, c)
		return
	}
	response.OK(response.UserResponse{Name: user.Name}, c)
}

func LoginByPassword(c *gin.Context) {
	var p request.LoginByPassword
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.Err(err, c)
		return
	}
	err, token := service.LoginByPassword(p.Name, p.Password)
	if err != nil {
		response.Err(err, c)
		return
	}
	response.OK(response.LoginByPasswordRes{Token: token}, c)
}
