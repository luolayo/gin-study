package core

import (
	"database/sql"
	"github.com/luolayo/gin-study/config"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// GetGorm get gorm
// use gorm to connect mysql
func GetGorm() *gorm.DB {
	databaseConfig := config.GetDatabaseConfig()
	dsn := databaseConfig.Username + ":" + databaseConfig.Password + "@tcp(" + databaseConfig.Host + ":" + databaseConfig.Port + ")/" + databaseConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	sqlDb, err := db.DB()
	if err != nil || sqlDb == nil {
		global.LOG.Error("GetGorm", err)
		return nil
	}
	err = db.AutoMigrate(&model.Test{}, &model.User{})
	if err != nil {
		global.LOG.Error("AutoMigrate %v", err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour * 5)
	return db
}

// CloseGorm close gorm
// close gorm connection
func CloseGorm(db *sql.DB) {
	if db == nil {
		return
	}
	err := db.Close()
	if err != nil {
		global.LOG.Error("CloseGorm", err)
		return
	}
}
