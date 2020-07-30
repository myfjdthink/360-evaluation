package routers

import (
	v1 "360-evaluation/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	apiv1 := e.Group("/api/v1")
	v1.InitUserRouter(apiv1)
}
