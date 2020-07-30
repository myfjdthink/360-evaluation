package v1

import (
	"360-evaluation/models"
	"360-evaluation/service/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRouter(r *gin.RouterGroup) {
	userApi := r.Group("/user")
	userApi.GET("/hello", Hello)
	userApi.GET("/add", AddUser)
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, user_service.Hello("deo"))
}

func AddUser(c *gin.Context) {
	name := "deo"
	err := models.AddUser(name)
	if err != nil {
		_, errMsg := fmt.Printf("插入失败 %s", err.Error())
		c.JSON(http.StatusOK, errMsg)
	}
	c.JSON(http.StatusOK, "插入成功")
}
