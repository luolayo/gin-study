package config

import (
	"github.com/luolayo/gin-study/enum"
	"github.com/spf13/viper"
	"testing"
)

// Some initialization done before testing
func init() {
	// Since the config package cannot call the core package, manual initialization is required here
	// Set the configuration file name, type, and path for viper
	viper.SetConfigName(enum.ConfigFileName)
	viper.SetConfigType(enum.ConfigFileType)
	viper.AddConfigPath(enum.ConfigDevelopmentPath)
	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
}

// TestGetAppConfig test GetAppConfig function.If there are no errors, it is correct
func TestGetAppConfig(t *testing.T) {
	appConfig := GetAppConfig()
	t.Log(appConfig)
}

// TestGetDatabaseConfig test GetDatabaseConfig function.If there are no errors, it is correct
func TestGetDatabaseConfig(t *testing.T) {
	databaseConfig := GetDatabaseConfig()
	t.Log(databaseConfig)
}

// TestGetRedisConfig test GetRedisConfig function.If there are no errors, it is correct
func TestGetRedisConfig(t *testing.T) {
	redisConfig := GetRedisConfig()
	t.Log(redisConfig)
}
