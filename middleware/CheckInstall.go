package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/interceptor/res"
	"github.com/spf13/viper"
	"net/http"
)

// CheckInstall function
// This method is used to check if the application has been installed.
// If the application is installed, proceed to next. Otherwise, proceed to install
func CheckInstall() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use Viber to check if it is installed
		if !viper.GetBool("app.Installed") {
			// If the path is "install", no path under it will jump
			installPath := []string{"/install/check", "/install/step1", "/install/step2", "/install/step3"}
			for _, path := range installPath {
				if c.Request.URL.Path == path {
					c.Next()
					return
				}
			}
			// If not installed, return a 501 status code requesting completion of the installation program
			res.Error(c, http.StatusNotImplemented, "Please install the application first", nil)
			c.Abort()
			return
		}
		// If installed, proceed to the next step
		c.Next()

	}
}
