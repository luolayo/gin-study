package test

import (
	"encoding/json"
	"github.com/luolayo/gin-study/Model"
	"github.com/luolayo/gin-study/Router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestPing(t *testing.T) {
	router := Router.GetRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestPong(t *testing.T) {
	r := Router.GetRouter()
	ts := httptest.NewRecorder()
	test := Model.Test{
		Msg: "hello",
	}
	data := url.Values{}
	data.Add("msg", test.Msg)
	reqBody := strings.NewReader(data.Encode())
	req, _ := http.NewRequest(http.MethodPost, "/test", reqBody)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.ServeHTTP(ts, req)
	assert.Equal(t, http.StatusOK, ts.Code)
	res := Model.Test{}
	err := json.Unmarshal(ts.Body.Bytes(), &res)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, test.Msg, res.Msg)
}
