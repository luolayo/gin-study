package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/logger"
	"strconv"
	"time"
)

// ANSI color codes for terminal output
const (
	Reset    = "\033[0m"
	RedBG    = "\033[41m"
	GreenBG  = "\033[42m"
	YellowBG = "\033[43m"
	CyanBG   = "\033[46m"
	White    = "\033[97m"
)

// colorForStatus returns a background color based on the status code
func colorForStatus(code int) string {
	switch {
	case code >= 200 && code < 300:
		return GreenBG
	case code >= 300 && code < 400:
		return CyanBG
	case code >= 400 && code < 500:
		return YellowBG
	default:
		return RedBG
	}
}

// colorForMethod returns a background color based on the HTTP method
func colorForMethod(method string) string {
	switch method {
	case "GET":
		return CyanBG
	case "POST":
		return GreenBG
	case "PUT":
		return YellowBG
	case "DELETE":
		return RedBG
	default:
		return Reset
	}
}

// LoggerMiddleware creates a new middleware for logging
func LoggerMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate the latency
		latency := time.Since(startTime)

		// Get status code and apply background color
		statusCode := c.Writer.Status()
		statusColor := colorForStatus(statusCode)
		coloredStatusCode := statusColor + White + " " + strconv.Itoa(c.Writer.Status()) + " " + Reset

		// Get method and apply background color
		method := c.Request.Method
		methodColor := colorForMethod(method)
		coloredMethod := methodColor + White + " " + method + " " + Reset

		// Get path and apply foreground color
		path := c.Request.URL.Path
		coloredPath := White + path + Reset

		// Get client IP
		clientIP := c.ClientIP()

		// Log request details
		log.Info("Request details status %v method %v path %v clientIP %v latency %v",
			coloredStatusCode,
			coloredMethod,
			coloredPath,
			clientIP,
			latency,
		)
	}
}
