package util

import (
	"github.com/luolayo/gin-study/core"
	"github.com/luolayo/gin-study/enum"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	core.InitViper(enum.ConfigDevelopmentPath)
	global.InitRedis()
}
func TestCreateToken(t *testing.T) {
	user := model.User{
		Uid:  1,
		Name: "luolayo",
	}
	token, err := CreateToken(user)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwiTmFtZSI6Imx1b2xheW8iLCJzdWIiOiJUb2tlbiIsImV4cCI6MTcyMjQ0Nzk5OSwiaWF0IjoxNzIyMzYxNTk5fQ.4CKDwlZ-jHA7W8DlBCnEJP9kvhEyXY9cKgpV9NXmdg0"
	claims, err := ParseToken(token)
	if err != nil {
		assert.Equal(t, "token expired", err.Error())
	}
	assert.Equal(t, 1, claims.ID)
}
