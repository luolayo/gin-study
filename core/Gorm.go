package core

import (
	"context"
	"database/sql"
	"github.com/luolayo/gin-study/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type GormClient struct {
	client *gorm.DB
	ctx    context.Context
}

// GetGorm get gorm
// use gorm to connect mysql
func GetGorm() *GormClient {
	// Get configuration
	dbConfig := config.GetDatabaseConfig()
	// Generate DSN link
	dsn := dbConfig.Username + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	// Connect to gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	ctx := context.Background()
	if err != nil {
		return &GormClient{
			client: nil,
			ctx:    ctx,
		}
	}
	// Get sql.DB -> mysql connection pool
	sqlDb, err := db.DB()
	if err != nil || sqlDb == nil {
		return &GormClient{
			client: nil,
			ctx:    ctx,
		}
	}
	// Set mysql connection pool configuration
	sqlDb.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDb.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Second * time.Duration(dbConfig.ConnMaxLifetime))
	return &GormClient{
		client: db,
		ctx:    ctx,
	}
}

// CloseGorm close gorm
// close gorm connection
func CloseGorm(db *sql.DB) {
	if db == nil {
		return
	}
	err := db.Close()
	if err != nil {
		panic(err)
		return
	}
}

func (r *GormClient) CheckGormConnection() bool {
	if r.client == nil {
		return false
	}
	db, err := r.client.DB()
	if err != nil {
		return false
	}
	defer CloseGorm(db)
	return true
}

func (r *GormClient) GetClient() *gorm.DB {
	return r.client
}
