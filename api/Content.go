package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/util"
	"time"
)

// CreateContent godoc
// @Summary Create a new content
// @Description Create a new content
// @Tags Content
// @Accept x-www-form-urlencoded
// @Produce json
// @Param Authorization header string true " Authorization token" example({{token}})
// @Param data body model.ContentRequest true "Content data"
// @Success 200 {object} interceptor.ResponseSuccess[model.Content]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @Failure 500 {object} interceptor.ResponseError
// @Router /content [post]
func CreateContent(c *gin.Context) {
	contentRequest := model.ContentRequest{}
	content := model.Content{}
	if err := c.ShouldBind(&contentRequest); err != nil {
		interceptor.BadRequest(c, "Invalid request", interceptor.ValidateErr(err))
		return
	}
	claims, ok := c.Get("claims")
	if !ok {
		interceptor.Unauthorized(c, "Unauthorized")
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	user := model.User{}
	global.GormDB.Where("id = ?", jwtClaims.ID).First(&user)
	switch checkContentType(contentRequest.Type) {
	case model.TypePost:
		content = createContentTypeIsPost(&contentRequest, jwtClaims.ID)
	case model.TypePage:
		if user.Group != model.GroupAdmin {
			interceptor.Forbidden(c, "You are not allowed to create page content")
			return
		}
		content = createContentTypeIsPage(&contentRequest, jwtClaims.ID)
	case model.TypeAttachment:
		content = createContentTypeIsAttachment(&contentRequest, jwtClaims.ID)
	default:
		interceptor.BadRequest(c, "Invalid content type", nil)
	}
	tx := global.GormDB.Begin()
	if err := tx.Create(&content).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Create content failed")
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Create content failed")
		return
	}
	interceptor.Success(c, "Create content success", content)
}

// GetContent godoc
// @Summary Get content by id
// @Description Get content by id
// @Tags Content
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path string true "Content ID" example(1)
// @Success 200 {object} interceptor.ResponseSuccess[model.Content]
// @Failure 404 {object} interceptor.ResponseError
// @Router /content/{id} [get]
func GetContent(c *gin.Context) {
	id := c.Param("id")
	content := model.Content{}
	global.GormDB.Where("cid = ?", id).Where("status = ?", model.ViewStatusPublic).First(&content)
	if content.Cid == 0 {
		interceptor.NotFound(c, "Content not found", nil)
		return
	}
	interceptor.Success(c, "Get content success", content)
}

// GetPostContentList godoc
// @Summary Get post content list
// @Description Get post content list
// @Tags Content
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} interceptor.ResponseSuccess[[]model.Content]
// @Router /content/post [get]
func GetPostContentList(c *gin.Context) {
	var content []model.Content
	global.GormDB.Where(model.Content{Type: model.TypePost}).Find(&content)
	interceptor.Success(c, "Get content list success", content)
}

// GetPageContentList godoc
// @Summary Get page content list
// @Description page post content list
// @Tags Content
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} interceptor.ResponseSuccess[[]model.Content]
// @Router /content/page [get]
func GetPageContentList(c *gin.Context) {
	var content []model.Content
	global.GormDB.Where(model.Content{Type: model.TypePage}).Find(&content)
	interceptor.Success(c, "Get content list success", content)
}

// GetAttachmentContentList godoc
// @Summary Get attachment content list
// @Description Get attachment content list
// @Tags Content
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path string true "Content ID" example(1)
// @Success 200 {object} interceptor.ResponseSuccess[model.Content]
// @Router /content/attachment/{id} [get]
func GetAttachmentContentList(c *gin.Context) {
	id := c.Param("id")
	content := model.Content{}
	global.GormDB.Where("parent = ?", id).Where("type = ?", model.TypeAttachment).First(&content)
	interceptor.Success(c, "Get content list success", content)
}

// UpdateContent godoc
// @Summary Update content by id
// @Description Update content by id
// @Tags Content
// @Accept x-www-form-urlencoded
// @Produce json
// @Param Authorization header string true " Authorization token" example({{token}})
// @Param id path string true "Content ID" example(1)
// @Param data body model.ContentUpdate true "Content data"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @Failure 404 {object} interceptor.ResponseError
// @Failure 500 {object} interceptor.ResponseError
// @Router /content/{id} [put]
func UpdateContent(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		interceptor.Unauthorized(c, "Unauthorized")
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	id := c.Param("id")
	content := model.Content{}
	global.GormDB.Where("cid = ?", id).First(&content)
	if content.Type == model.TypeAttachment {
		interceptor.BadRequest(c, "Attachment content cannot be updated", nil)
		return
	}
	if content.Cid == 0 {
		interceptor.NotFound(c, "Content not found", nil)
		return
	}
	if int(content.AuthorId) != jwtClaims.ID {
		interceptor.BadRequest(c, "You are not the author of this content", nil)
		return
	}
	contentRequest := model.ContentUpdate{}
	if err := c.ShouldBind(&contentRequest); err != nil {
		interceptor.BadRequest(c, "Invalid request", interceptor.ValidateErr(err))
		return
	}
	if contentRequest.Title != "" {
		content.Title = contentRequest.Title
	}
	if contentRequest.Text != "" {
		content.Text = contentRequest.Text
	}
	if contentRequest.Order != 0 {
		content.Order = gconv.Uint(contentRequest.Order)
	}
	t := time.Now()
	content.Modified = &t
	tx := global.GormDB.Begin()
	if err := tx.Save(&content).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Update content failed")
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Update content failed")
		return
	}
	interceptor.Success(c, "Update content success", interceptor.Empty{})
}

// DeleteContent godoc
// @Summary Delete content by id
// @Description Delete content by id
// @Tags Content
// @Accept x-www-form-urlencoded
// @Produce json
// @Param Authorization header string true " Authorization token" example({{token}})
// @Param id path string true "Content ID" example(1)
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @Failure 404 {object} interceptor.ResponseError
// @Failure 500 {object} interceptor.ResponseError
func DeleteContent(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		interceptor.Unauthorized(c, "Unauthorized")
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	id := c.Param("id")
	content := model.Content{}
	global.GormDB.Where("cid = ?", id).First(&content)
	if content.Cid == 0 {
		interceptor.NotFound(c, "Content not found", nil)
		return
	}
	user := model.User{}
	global.GormDB.Where("id = ?", jwtClaims.ID).First(&user)
	if user.Group == model.GroupAdmin {
		if ok := deleteContent(&content); ok != nil {
			interceptor.ServerError(c, "Delete content failed")
			return
		}
		interceptor.Success(c, "Delete content success", interceptor.Empty{})
		return
	}
	if int(content.AuthorId) == jwtClaims.ID {
		if ok := deleteContent(&content); ok != nil {
			interceptor.ServerError(c, "Delete content failed")
			return
		}
		interceptor.Success(c, "Delete content success", interceptor.Empty{})
		return
	}
	interceptor.BadRequest(c, "You are not the author of this content", nil)
}

// ApproveRelease godoc
// @Summary Approve release by id
// @Description Approve release by id
// @Tags Content
// @Accept x-www-form-urlencoded
// @Produce json
// @Param Authorization header string true " Authorization token" example({{token}})
// @Param id path string true "Content ID" example(1)
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @Failure 404 {object} interceptor.ResponseError
// @Failure 500 {object} interceptor.ResponseError
func ApproveRelease(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		interceptor.BadRequest(c, "Invalid request", nil)
		return
	}
	content := model.Content{}
	global.GormDB.Where("cid = ?", id).First(&content)
	if content.Cid == 0 {
		interceptor.NotFound(c, "Content not found", nil)
		return
	}
	content.Status = model.ViewStatusPublic
	tx := global.GormDB.Begin()
	if err := tx.Save(&content).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Approve release failed")
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Approve release failed")
		return
	}
	interceptor.Success(c, "Approve release success", interceptor.Empty{})
}

// GetPostContentListPublic godoc
// @Summary Get post content list public
// @Description Get post content list public
// @Tags Content
// @Accept x-www-form-urlencoded
// @Produce json
// @Success 200 {object} interceptor.ResponseSuccess[[]model.Content]
// @Router /content/all [get]
func GetPostContentListPublic(c *gin.Context) {
	var content []model.Content
	global.GormDB.Where(model.Content{Type: model.TypePost, Status: model.ViewStatusPublic}).Find(&content)
	interceptor.Success(c, "Get content list success", content)
}

// checkContentType checks if the content type is valid.
func checkContentType(contentType model.Type) model.Type {
	switch contentType {
	case "post":
		return model.TypePost
	case "page":
		return model.TypePage
	case "attachment":
		return model.TypeAttachment
	default:
		return model.TypePost
	}
}

func createContentTypeIsPost(contentRequest *model.ContentRequest, uid int) model.Content {
	id := getLastId()
	newContent := model.Content{
		Cid:      id,
		Title:    contentRequest.Title,
		Slug:     gconv.String(id),
		AuthorId: gconv.Uint(uid),
		Text:     contentRequest.Text,
		Type:     model.TypePost,
	}
	return newContent
}

func createContentTypeIsPage(contentRequest *model.ContentRequest, uid int) model.Content {
	id := getLastId()
	if contentRequest.Slug == "" {
		contentRequest.Slug = gconv.String(id)
	}
	newContent := model.Content{
		Cid:      id,
		Title:    contentRequest.Title,
		Slug:     contentRequest.Slug,
		AuthorId: gconv.Uint(uid),
		Text:     contentRequest.Text,
		Type:     model.TypePage,
	}
	return newContent
}

func createContentTypeIsAttachment(contentRequest *model.ContentRequest, uid int) model.Content {
	id := getLastId()
	newContent := model.Content{
		Cid:      id,
		Title:    contentRequest.Title,
		Slug:     contentRequest.Title,
		AuthorId: gconv.Uint(uid),
		Text:     contentRequest.Text,
		Type:     model.TypeAttachment,
	}
	return newContent
}

func getLastId() uint {
	content := model.Content{}
	global.GormDB.Last(&content)
	return content.Cid
}

func deleteContent(content *model.Content) error {
	tx := global.GormDB.Begin()
	if err := tx.Delete(&content).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
