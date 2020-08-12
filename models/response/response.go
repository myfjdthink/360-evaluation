package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	msg := ""
	switch err.(type) {
	case validator.ValidationErrors:
		msg = Translate(err.(validator.ValidationErrors))
	default:
		msg = err.Error()
	}
	Result(http.StatusBadRequest, nil, msg, c)
}

func OK(data interface{}, c *gin.Context) {
	Result(http.StatusOK, data, "", c)
}
