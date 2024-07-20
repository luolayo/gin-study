package test

import (
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/util"
	"testing"
)

func TestCryption(t *testing.T) {
	global.Init()
	str := "hello"
	encrypted, err := util.Encrypt(str)
	if err != nil {
		t.Error(err)
	}
	decrypted, err := util.Decrypt(encrypted)
	if err != nil {
		t.Error(err)
	}
	if decrypted != str {
		t.Errorf("expected %s, but got %s", str, decrypted)
	}
}
