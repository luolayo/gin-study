package test

import (
	Core "github.com/luolayo/gin-study/Core"
	"github.com/luolayo/gin-study/Model"
	"testing"
)

func TestGorm(t *testing.T) {
	// Test code here
	db := Core.GetGorm()
	if db == nil {
		t.Error("Gorm is nil")
	} else {
		t.Log("Gorm is not nil")
	}
}

func TestGormSelect(t *testing.T) {
	db := Core.GetGorm()
	if db == nil {
		t.Error("Gorm is nil")
	} else {
		t.Log("Gorm is not nil")
	}
	err := db.AutoMigrate(&Model.Test{})
	if err != nil {
		t.Fatal(err)
	}
	db.Create(&Model.Test{})
}

func TestGormSelectAll(t *testing.T) {
	db := Core.GetGorm()
	if db == nil {
		t.Error("Gorm is nil")
	} else {
		t.Log("Gorm is not nil")
	}
	var tests []Model.Test
	db.Find(&tests)
	t.Log(tests)
}

func TestGormConnect(t *testing.T) {
	userGorm := Core.UserGorm()
	if userGorm == nil {
		t.Error("Gorm is nil")
		return
	} else {
		t.Log("Gorm is not nil")
	}
	user := Model.User{}
	userGorm.Find(&user)
	t.Log(user)
	defer Core.CloseGorm(userGorm)
}
