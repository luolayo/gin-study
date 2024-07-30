package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt_Decrypt(t *testing.T) {
	// Encrypt("123456") => "TGFT4yb-JZrBj_GHfJ1IIu2KS04mtw=="

	passwd := "123456"
	encrypted, err := Encrypt(passwd)
	if err != nil {
		t.Error(err)
	}
	decrpted, err := Decrypt(encrypted)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, passwd, decrpted)
}

func TestCompare(t *testing.T) {
	passwd := "123456"
	encrypted, err := Encrypt(passwd)
	if err != nil {
		t.Error(err)
	}
	flag, err := Compare(encrypted, passwd)
	if err != nil {
		t.Error(err)
	}
	assert.True(t, flag)
}
