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
	c.JSON(code, Response{data, msg})
}

func Err(err error, c *gin.Context) {
	Result(http.StatusBadRequest, nil, err.Error(), c)
}

func OK(data interface{}, c *gin.Context) {
	Result(http.StatusOK, data, "", c)
}
