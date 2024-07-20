package test

import (
	"github.com/luolayo/gin-study/core"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/util"
	"testing"
)

func TestGetJWTToken(t *testing.T) {
	core.InitGlobal()
	user := model.User{
		Uid:      1,
		Name:     "test",
		Password: "test",
	}
	token, err := util.CreateToken(user)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("token is empty")
	}
}

func TestParseJWTToken(t *testing.T) {
	core.InitGlobal()
	user := model.User{
		Uid:      1,
		Name:     "test",
		Password: "test",
	}
	token, err := util.CreateToken(user)
	if err != nil {
		t.Error(err)
	}
	if token == "" {
		t.Error("token is empty")
	}
	user2, err := util.ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	if user2.ID != int(user.Uid) {
		t.Errorf("expected %d, but got %d", user.Uid, user2.ID)
	}
	if user2.Name != user.Name {
		t.Errorf("expected %s, but got %s", user.Name, user2.Name)
	}
}
