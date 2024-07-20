package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

// GetSystemConfig get system config
func GetSystemConfig() *System {
	return &System{
		AppName:     os.Getenv("APP_NAME"),
		AppVersion:  os.Getenv("APP_VERSION"),
		Port:        os.Getenv("PORT"),
		Environment: os.Getenv("ENVIRONMENT"),
		CryPtKey:    os.Getenv("CRYPT_KEY"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}
}

// GetDatabaseConfig get database config
func GetDatabaseConfig() *Database {
	return &Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
}

// GetAliYunConfig get aliyun config
func GetAliYunConfig() *Aliyun {
	return &Aliyun{
		AccessKeyID:     os.Getenv("ALIYUN_ACCESS_KEY_ID"),
		AccessKeySecret: os.Getenv("ALIYUN_ACCESS_KEY_SECRET"),
		SignName:        os.Getenv("ALIYUN_SIGN_NAME"),
		TemplateCode:    os.Getenv("ALIYUN_TEMPLATE_CODE"),
	}
}

// GetRedisConfig get redis config
func GetRedisConfig() *Redis {
	return &Redis{
		Host:         os.Getenv("REDIS_HOST"),
		Port:         os.Getenv("REDIS_PORT"),
		DB:           os.Getenv("REDIS_DB"),
		DialTimeout:  os.Getenv("REDIS_DIAL_TIMEOUT"),
		ReadTimeout:  os.Getenv("REDIS_READ_TIMEOUT"),
		WriteTimeout: os.Getenv("REDIS_WRITE_TIMEOUT"),
		PoolSize:     os.Getenv("REDIS_POOL_SIZE"),
		PoolTimeout:  os.Getenv("REDIS_POOL_TIMEOUT"),
	}
}
