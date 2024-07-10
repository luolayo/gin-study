package test

import (
	"github.com/luolayo/gin-study/Router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	router := Router.GetRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/test/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
