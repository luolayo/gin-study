package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/util"
)

// GetComments godoc
// @Summary Get comments
// @Description Get comments
// @Tags Comment
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path string true "Content ID"
// @Success 200 {object} interceptor.ResponseSuccess[[]model.Comment]
// @Router /comment/{id} [get]
func GetComments(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		interceptor.BadRequest(c, "Bad Request", nil)
		return
	}
	var comments []model.Comment
	global.GormDB.Where("cid = ?", id).Find(&comments)
	interceptor.Success(c, "Success", comments)
}

// CreateComment godoc
// @Summary Create comment
// @Description Create comment
// @Tags Comment
// @Accept x-www-form-urlencoded
// @Produce json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Param id path string true "Content ID"
// @Param data body model.CommentRequest true "Comment"
// @Success 200 {object} interceptor.ResponseSuccess[model.Comment]
// @Router /comment/{id} [post]
func CreateComment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		interceptor.BadRequest(c, "Bad Request", nil)
		return
	}
	commnetRequest := model.CommentRequest{}
	if err := c.ShouldBind(&commnetRequest); err != nil {
		interceptor.BadRequest(c, "Bad Request", interceptor.ValidateErr(err))
		return
	}
	claims, ok := c.Get("claims")
	if !ok {
		interceptor.Unauthorized(c, "Unauthorized")
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	user := model.User{}
	global.GormDB.Where("uid = ?", jwtClaims.ID).First(&user)
	if user.Uid == 0 {
		interceptor.Unauthorized(c, "Unauthorized")
		return
	}
	comment := model.Comment{
		Cid:        gconv.Uint(id),
		AuthorName: user.Name,
		AuthorId:   user.Uid,
		Url:        user.Url,
		Ip:         c.ClientIP(),
		Agent:      c.GetHeader("User-Agent"),
		Text:       commnetRequest.Text,
		Status:     model.Pending,
	}
	tx := global.GormDB.Begin()
	if err := tx.Create(&comment).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Internal Server Error")
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Internal Server Error")
		return
	}
	interceptor.Success(c, "Success", comment)
}

// DeleteComment godoc
// @Summary Delete comment
// @Description Delete comment
// @Tags Comment
// @Accept x-www-form-urlencoded
// @Produce json
// @Param Authorization header string true "Authorization token" example({{token}})
// @Param id path string true "Comment ID"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @Router /comment/{id} [delete]
func DeleteComment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		interceptor.BadRequest(c, "Bad Request", nil)
		return
	}
	commnetRequest := model.CommentRequest{}
	if err := c.ShouldBind(&commnetRequest); err != nil {
		interceptor.BadRequest(c, "Bad Request", interceptor.ValidateErr(err))
		return
	}
	claims, ok := c.Get("claims")
	if !ok {
		interceptor.Unauthorized(c, "Unauthorized")
		return
	}
	jwtClaims := claims.(util.JwtCustomClaims)
	user := model.User{}
	global.GormDB.Where("uid = ?", jwtClaims.ID).First(&user)
	if user.Uid == 0 {
		interceptor.Unauthorized(c, "Unauthorized")
		return
	}
	if checkAuthority(id, user) {
		interceptor.Forbidden(c, "Forbidden")
		return
	}
	tx := global.GormDB.Begin()
	if err := tx.Where("coid = ?", id).Delete(&model.Comment{}).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Internal Server Error")
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Internal Server Error")
		return
	}
	interceptor.Success(c, "Success", interceptor.Empty{})
}

// ApproveComment godoc
// @Summary Approve comment
// @Description Approve comment
// @Tags Comment
// @Accept x-www-form-urlencoded
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param id path string true "Comment ID"
// @Success 200 {object} interceptor.ResponseSuccess[interceptor.Empty]
// @Failure 400 {object} interceptor.ResponseError
// @Failure 401 {object} interceptor.ResponseError
// @Failure 403 {object} interceptor.ResponseError
// @Router /comment/approve/{id} [get]
func ApproveComment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		interceptor.BadRequest(c, "Bad Request", nil)
		return
	}
	tx := global.GormDB.Begin()
	if err := tx.Model(&model.Comment{}).Where("coid = ?", id).Update("status", model.Approved).Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Internal Server Error")
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		interceptor.ServerError(c, "Internal Server Error")
		return
	}
	interceptor.Success(c, "Success", interceptor.Empty{})

}

func checkAuthority(id string, user model.User) bool {
	if user.Group == model.GroupAdmin {
		return true
	}
	comment := model.Comment{}
	global.GormDB.Where("author_id = ?", id).First(&comment)
	if comment.AuthorId == user.Uid {
		return true
	}
	return false
}
