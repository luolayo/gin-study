package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/logger"
	"net/http"
	"runtime/debug"
)

// RecoveryMiddleware creates a new middleware for recovering from panics
func RecoveryMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log the panic
				log.Error("Panic recovered: %v\n%s", r, debug.Stack())

				// Return 500 Internal Server Error
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
