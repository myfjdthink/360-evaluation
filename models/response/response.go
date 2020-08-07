package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	x
	c.JSON(code, Response{data, msg})
}

func Err(msg string, c *gin.Context) {
	Result(http.StatusBadRequest, nil, msg, c)
}

func OK(data interface{}, c *gin.Context) {
	Result(http.StatusOK, data, "", c)
}
