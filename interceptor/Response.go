package interceptor

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseSuccess[T any] struct {
	Code    int    `json:"code" description:"Status code" example:"200"`
	Message string `json:"message" description:"Response message" example:"Success"`
	Data    T      `json:"data"`
}

type ResponseError struct {
	Code    int      `json:"code" description:"Status code" example:"400"`
	Message string   `json:"message" description:"Response message" example:"Invalid request"`
	Err     []string `json:"err"`
}
type Empty struct{}

func Success[T any](c *gin.Context, msg string, data T) {
	c.JSON(http.StatusOK, ResponseSuccess[T]{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, msg string, err []string) {
	c.JSON(code, ResponseError{
		Code:    code,
		Message: msg,
		Err:     err,
	})
}

func NotFound(c *gin.Context, msg string, err []string) {
	Error(c, http.StatusNotFound, msg, err)
}

func BadRequest(c *gin.Context, msg string, err []string) {
	Error(c, http.StatusBadRequest, msg, err)
}

func Unauthorized(c *gin.Context, msg string) {
	Error(c, http.StatusUnauthorized, msg, nil)
}

func Forbidden(c *gin.Context, msg string) {
	Error(c, http.StatusForbidden, msg, nil)
}

func ServerError(c *gin.Context, msg string) {
	Error(c, http.StatusInternalServerError, msg, nil)
}
