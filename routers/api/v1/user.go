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
	_, err := models.GetUserByName(name)
	if err != nil {
		c.JSON(http.StatusOK, fmt.Sprintf("已存在用户:%s", name))
		return
	}

	err = models.AddUser(name)
	if err != nil {
		c.JSON(http.StatusOK, fmt.Sprintf("插入失败:%s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, "插入成功")
}
