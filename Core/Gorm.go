package Core

import (
	"github.com/luolayo/gin-study/Logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func GetGorm() *gorm.DB {
	system := GetSystemConfig()
	logger := &Logger.Logger{}
	if system.Environment == "development" {
		logger = Logger.NewLogger(Logger.InfoLevel)
	} else {
		logger = Logger.NewLogger(Logger.ErrorLevel)
	}
	dbConfig := GetDatabaseConfig()
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database %s", err)
	}
	sqlDb, err := db.DB()
	if err != nil || sqlDb == nil {
		logger.Error("Failed to get database connection %s", err)
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour * 5)
	return db
}
