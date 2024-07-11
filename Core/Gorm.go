package Core

import (
	"github.com/luolayo/gin-study/Logger"
	"github.com/luolayo/gin-study/Model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var logger = Logger.NewLogger(Logger.ErrorLevel)

func GetGorm() *gorm.DB {
	dbConfig := GetDatabaseConfig()
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database %s", err)
	}
	sqlDb, err := db.DB()
	if err != nil || sqlDb == nil {
		logger.Error("Failed to get database connection %s", err)
		return nil
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour * 5)
	err = db.AutoMigrate(&Model.User{}, &Model.Meta{}, &Model.Content{}, &Model.Comment{}, &Model.Relationship{})
	if err != nil {
		logger.Error("Failed to migrate user model %s", err)
		return nil
	}
	return db
}
func CloseGorm(db *gorm.DB) {
	sqlDb, err := db.DB()
	if sqlDb == nil {
		return
	}
	if err != nil {
		logger.Error("Failed to get database connection %s", err)
	}
	err = sqlDb.Close()
	if err != nil {
		logger.Error("Failed to close database connection %s", err)
	}
}

func UserGorm() *gorm.DB {
	return GetGorm().Model(&Model.User{})
}

func MetaGorm() *gorm.DB {
	return GetGorm().Model(&Model.Meta{})
}
func ContentGorm() *gorm.DB {
	return GetGorm().Model(&Model.Content{})
}
func CommentGorm() *gorm.DB {
	return GetGorm().Model(&Model.Comment{})
}
func RelationshipGorm() *gorm.DB {
	return GetGorm().Model(&Model.Relationship{})
}
