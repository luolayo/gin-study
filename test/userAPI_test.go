package test

import (
	"github.com/goccy/go-json"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegitserUser(t *testing.T) {
	global.Init()
	r := router.GetRouter()
	w := httptest.NewRecorder()
	phone := "18888888888"
	req, _ := http.NewRequest("GET", "/SMS/send?phone_number="+phone, nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	userRegisterInfo := model.UserRegister{
		Name:            "luolayo",
		Password:        "123456",
		ConfirmPassword: "123456",
		Phone:           phone,
		Code:            "123456",
		Url:             "https://www.luola.me",
		ScreenName:      "罗拉",
	}
	userRegisterInfoJson, _ := json.Marshal(userRegisterInfo)
	req, _ = http.NewRequest("POST", "/user/register", strings.NewReader(string(userRegisterInfoJson)))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	t.Log(w.Body.String())
}

func TestCheckUserPhone(t *testing.T) {
	global.Init()
	r := router.GetRouter()
	w := httptest.NewRecorder()
	phone := "18888888888"
	req, _ := http.NewRequest("GET", "/user/checkPhone?phone="+phone, nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	t.Log(w.Body.String())
}

func TestCheckUserName(t *testing.T) {
	global.Init()
	r := router.GetRouter()
	w := httptest.NewRecorder()
	name := "luolayo"
	req, _ := http.NewRequest("GET", "/user/checkName?name="+name, nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	t.Log(w.Body.String())
}

func TestUserLogin(t *testing.T) {
	global.Init()
	r := router.GetRouter()
	w := httptest.NewRecorder()
	userLoginInfo := model.UserLogin{
		Name:     "admin",
		Password: "123456",
	}
	userLoginInfoJson, _ := json.Marshal(userLoginInfo)
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(string(userLoginInfoJson)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
