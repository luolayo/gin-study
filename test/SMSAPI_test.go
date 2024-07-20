package test

import (
	"github.com/luolayo/gin-study/core"
	"github.com/luolayo/gin-study/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSentVerificationCode(t *testing.T) {
	core.InitGlobal()
	r := router.GetRouter()
	w := httptest.NewRecorder()
	phone := "18888888888"
	req := httptest.NewRequest("GET", "/SMS/send?phone_number="+phone, nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
