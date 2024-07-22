package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
)

func CheckInstalled() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the application is installed
		// If not installed, redirect to the installation page
		if c.Request.URL.Path == "/install" {
			if _, err := os.Stat("install/install.lock"); os.IsNotExist(err) {
				c.Redirect(302, "/")
				c.Abort()
			}
		}
		// If installed, continue to the next step
		c.Next()
	}
}
