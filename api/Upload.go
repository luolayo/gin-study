package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/interceptor"
	"path/filepath"
	"time"
)

// Upload godoc
// @Summary Upload file
// @Description Upload file
// @Tags Upload
// @Accept  mpfd
// @Produce  json
// @Param file formData file true "file"
// @Success 200 {object} interceptor.ResponseSuccess[string]
// @Router /upload [post]swg
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		interceptor.BadRequest(c, "file not found", interceptor.ValidateErr(err))
		return
	}
	// check file type
	if file.Header.Get("Content-Type") != "image/jpeg" {
		interceptor.BadRequest(c, "file type not allowed", nil)
		return
	}
	// Save file name as current time
	filename := time.Now().Format("2006-01-02-15-04-05.000") + filepath.Ext(file.Filename)
	interceptor.Success(c, "Upload file successful", filename)
}
