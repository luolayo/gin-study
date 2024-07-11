package Core

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/luolayo/gin-study/Config"
	"os"
)

func GetSystemConfig() Config.System {
	return Config.System{
		AppName:     os.Getenv("APP_NAME"),
		AppVersion:  os.Getenv("APP_VERSION"),
		Port:        os.Getenv("PORT"),
		Environment: os.Getenv("ENVIRONMENT"),
	}
}

func GetDatabaseConfig() Config.Database {
	return Config.Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
}
