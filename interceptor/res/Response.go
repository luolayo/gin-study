package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response[T any] struct {
	// Return status Code
	Code int `json:"code" example:"200"`
	// Return Message
	Message string `json:"message" example:"ok"`
	// Return specific Data
	Data T `json:"data"`
}

type ErrorRes struct {
	// Return status Code
	Code int `json:"code" example:"200"`
	// Return Message
	Message string `json:"message" example:"ok"`
	// Return specific Data
	Errors []string `json:"errors"`
}
type Empty struct{}

func Base[T any](c *gin.Context, code int, msg string, data T) {
	c.JSON(http.StatusOK, Response[T]{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func Success[T any](c *gin.Context, data T) {
	Base(c, http.StatusOK, http.StatusText(http.StatusOK), data)
}

func SuccessNoData(c *gin.Context) {
	Base(c, http.StatusOK, http.StatusText(http.StatusOK), Empty{})
}
func Error(c *gin.Context, code int, msg string, err []string) {
	c.JSON(http.StatusOK, ErrorRes{
		Code:    code,
		Message: msg,
		Errors:  err,
	})
}

func Created(c *gin.Context) {
	Base(c, http.StatusCreated, http.StatusText(http.StatusCreated), Empty{})
}
func NotFound(c *gin.Context, err []string) {
	Error(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), err)
}

func BadRequest(c *gin.Context, err []string) {
	Error(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
}

func Unauthorized(c *gin.Context) {
	Error(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
}

func Forbidden(c *gin.Context) {
	Error(c, http.StatusForbidden, http.StatusText(http.StatusForbidden), nil)
}

func ServerError(c *gin.Context, msg string) {
	Error(c, http.StatusInternalServerError, msg, nil)
}
