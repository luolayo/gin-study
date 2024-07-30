package config

import (
	"github.com/spf13/viper"
)

// GetAppConfig Get configuration items for app fields
func GetAppConfig() *AppConfig {
	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	appConfig := &AppConfig{}
	err = viper.UnmarshalKey("app", appConfig)
	if err != nil {
		panic(err)
	}
	return appConfig
}

// GetDatabaseConfig Get configuration items for database fields
func GetDatabaseConfig() *DatabaseConfig {
	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	databaseConfig := DatabaseConfig{}
	err = viper.UnmarshalKey("database", &databaseConfig)
	if err != nil {
		panic(err)
	}
	return &databaseConfig
}

// GetRedisConfig Get configuration items for redis fields
func GetRedisConfig() *RedisConfig {
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	redisConfig := RedisConfig{}
	err = viper.UnmarshalKey("redis", &redisConfig)
	if err != nil {
		panic(err)
	}
	return &redisConfig
}
