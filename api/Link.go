package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor"
	"github.com/luolayo/gin-study/model"
)

// GetLinks godoc
// @Summary Get all approved links
// @Description Get all links
// @Tags Link
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} interceptor.ResponseSuccess[[]model.Link]
// @Router /link [get]
func GetLinks(c *gin.Context) {
	var links []model.Link
	global.GormDB.Where("stutas = ?", 1).Find(&links)
	interceptor.Success(c, "Success", links)
}

// CreateLink godoc
// @Summary Create a link
// @Description Create a link
// @Tags Link
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data body model.LinkRequest true "Link"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Router /link [post]
func CreateLink(c *gin.Context) {
	var link model.LinkRequest
	if err := c.ShouldBindJSON(&link); err != nil {
		interceptor.BadRequest(c, "Invalid request", interceptor.ValidateErr(err))
		return
	}
	tx := global.GormDB.Begin()
	if err := tx.Create(&model.Link{Name: link.Name, Image: link.Image, URL: link.URL}).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to create link")
		return
	}
	if err := tx.Commit().Error; err != nil {
		interceptor.ServerError(c, "Failed to create link")
		return
	}
	interceptor.Success(c, "Success", interceptor.Empty{})
}

// UpdateLink godoc
// @Summary Update a link
// @Description Update a link
// @Tags Link
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path string true "Link ID"
// @Param data body model.LinkUpdate true "Link"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Router /link/{id} [put]
func UpdateLink(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		interceptor.BadRequest(c, "Invalid request", []string{"Invalid id"})
		return
	}
	var linkUpdate = model.LinkUpdate{}
	if err := c.ShouldBindJSON(&linkUpdate); err != nil {
		interceptor.BadRequest(c, "Invalid request", interceptor.ValidateErr(err))
		return
	}
	tx := global.GormDB.Begin()
	link := model.Link{}
	if err := tx.Where("id = ?", id).First(&link).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to update link")
		return
	}
	if linkUpdate.Name != "" {
		link.Name = linkUpdate.Name
	}
	if linkUpdate.Image != "" {
		link.Image = linkUpdate.Image
	}
	if linkUpdate.URL != "" {
		link.URL = linkUpdate.URL
	}
	if linkUpdate.Sort != 0 {
		link.Sort = linkUpdate.Sort
	}
	if err := tx.Save(&link).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to update link")
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to update link")
		return
	}
	interceptor.Success(c, "Success", interceptor.Empty{})
}

// DeleteLink godoc
// @Summary Delete a link
// @Description Delete a link
// @Tags Link
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path string true "Link ID"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Router /link/{id} [delete]
func DeleteLink(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		interceptor.BadRequest(c, "Invalid request", []string{"Invalid id"})
		return
	}
	tx := global.GormDB.Begin()
	if err := tx.Where("id = ?", id).Delete(&model.Link{}).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to delete link")
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to delete link")
		return
	}
	interceptor.Success(c, "Success", interceptor.Empty{})
}

// ApproveLink godoc
// @Summary Approve a link
// @Description Approve a link
// @Tags Link
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path string true "Link ID"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Router /link/approve/{id} [patch]
func ApproveLink(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		interceptor.BadRequest(c, "Invalid request", []string{"Invalid id"})
		return
	}
	tx := global.GormDB.Begin()
	if err := tx.Model(&model.Link{}).Where("id = ?", id).Update("stutas", 1).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to approve link")
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Failed to approve link")
		return
	}
	interceptor.Success(c, "Success", interceptor.Empty{})
}

// GetLink godoc
// @Summary admin Get all links
// @Description Get all links
// @Tags Link
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} interceptor.ResponseSuccess[[]model.Link]
// @Router /link/all [get]
func GetLink(c *gin.Context) {
	var link model.Link
	global.GormDB.First(&link)
	interceptor.Success(c, "Success", link)
}

// GetLinkById godoc
// @Summary admin Get link by id
// @Description Get link by id
// @Tags Link
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path string true "Link ID"
// @Success 200 {object} interceptor.ResponseSuccess[model.Link]
func GetLinkById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		interceptor.BadRequest(c, "Invalid request", []string{"Invalid id"})
		return
	}
	var link model.Link
	global.GormDB.Where("id = ?", id).First(&link)
	interceptor.Success(c, "Success", link)
}
