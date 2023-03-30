package internal

import (
	"gateway/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// 自定义响应
func Custom(c *gin.Context, status int, data any, message string) {
	c.JSON(http.StatusOK, Response{
		Status:  status,
		Data:    data,
		Message: message,
	})
	c.Abort()
}

// 成功默认响应
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Data:    data,
		Message: "请求成功",
	})
	c.Abort()
}

// 请求出错响应
func Error(c *gin.Context, err error) {
	c.JSON(http.StatusOK, Response{
		Status:  errno.GetCode(err.Error()),
		Message: err.Error(),
		Data:    nil,
	})
	c.Abort()
}

// 服务端出错响应
func ServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
		Status:  http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	})
}
