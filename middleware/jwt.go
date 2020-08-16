package middleware

import (
	"360-evaluation/models/response"
	"360-evaluation/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		success := true
		if token == "" {
			success = false
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil || claims.ExpiresAt < time.Now().Unix() {
				success = false
			}
		}

		if !success {
			response.Result(http.StatusForbidden, nil, "请重新登陆", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
