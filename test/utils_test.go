package test

import (
	"github.com/luolayo/gin-study/util"
	"testing"
)

func TestStringtoInt(t *testing.T) {
	if util.StringToInt("123") != 123 {
		t.Error("Expected 123 ")
	}
	if util.StringToInt("abc") != 0 {
		t.Error("Expected 0 ")
	}
}
