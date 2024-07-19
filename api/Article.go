package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/util"
)

// Article godoc
// @Summary Article
// @Description Article
// @Tags Article
// @Schemes http https
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} interceptor.ResponseSuccess[model.ArticleResponse]
// @Failure 400 {object} interceptor.ResponseError
// @router /article/:id [Get]
func Article(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		interceptor.BadRequest(c, "Invalid parameter", nil)
		return
	}
	article := model.Article{}
	global.GormDB.Where("id = ?", id).First(&article)
	if article.ID == 0 {
		interceptor.BadRequest(c, "Article does not exist", nil)
		return
	}
	interceptor.Success(c, "success", articleToArticleResponse(article))
}

// ArticleList godoc
// @Summary ArticleList
// @Description ArticleList
// @Tags Article
// @Schemes http https
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param user_id query string false "user_id"
// @Param limit query string true "limit"
// @Param offset query string true "offset"
// @Success 200 {object} interceptor.ResponseSuccess[[]model.ArticleResponse]
// @Failure 400 {object} interceptor.ResponseError
// @router /article [Get]
func ArticleList(c *gin.Context) {
	userId := c.Query("user_id")
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "1")
	limitInt := util.StringToInt(limit)
	offsetInt := util.StringToInt(offset)
	if limitInt == 0 {
		interceptor.BadRequest(c, "Invalid parameter", nil)
		return
	}
	if userId == "" {
		articles := articleAll(limitInt, offsetInt)
		interceptor.Success(c, "success", articlesToArticleResponses(articles))
		return
	}
	articles := articleByUserID(userId, limitInt, offsetInt)
	interceptor.Success(c, "success", articlesToArticleResponses(articles))
}

func articleAll(limt int, offset int) []model.Article {
	var articles []model.Article
	global.GormDB.Limit(limt).Offset(offset).Find(&articles)
	return articles
}
func articleByUserID(userId string, limt int, offset int) []model.Article {
	var articles []model.Article
	global.GormDB.Where("user_id = ?", userId).Limit(limt).Offset(offset).Find(&articles)
	return articles
}

func articleToArticleResponse(article model.Article) model.ArticleResponse {
	return model.ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Context:   article.Context,
		UserID:    article.UserID,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}

func articlesToArticleResponses(articles []model.Article) []model.ArticleResponse {
	var articleResponses []model.ArticleResponse
	for _, article := range articles {
		articleResponses = append(articleResponses, articleToArticleResponse(article))
	}
	return articleResponses
}
