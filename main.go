package main

import (
	"360-evaluation/models"
	"360-evaluation/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	models.Setup()
}

func main() {
	r := gin.Default()
	routers.InitRouter(r)
	r.Run()
}
