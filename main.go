package main

import (
	"360-evaluation/config"
	"360-evaluation/models"
	"360-evaluation/routers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	models.Setup()
}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	routers.InitRouter(r)
	err := r.Run(fmt.Sprintf("127.0.0.1:%d", config.Config.Port))
	if err != nil {
		fmt.Printf("server start fail %e", err)
	}
}
