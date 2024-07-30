package core

import (
	"github.com/luolayo/gin-study/enum"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCheckConfigFileFail tests the CheckConfigFile function.
func TestCheckConfigFileFail(t *testing.T) {
	ok := CheckConfigFile(enum.ConfigDevelopmentPath)
	assert.Equal(t, ok, false)
}

// TestCheckConfigFile tests the CheckConfigFile function.
func TestCheckConfigFile(t *testing.T) {
	ok := CheckConfigFile(enum.ConfigDevelopmentPath)
	assert.Equal(t, ok, true)
}

// TestCreateConfigFile tests the CreateConfigFile function.
func TestCreateConfigFile(t *testing.T) {
	CreateConfigFile(enum.ConfigDevelopmentPath)
	ok := CheckConfigFile(enum.ConfigDevelopmentPath)
	assert.Equal(t, ok, true)
}

// TestWriteBaseConfig tests the WriteBaseConfig function.
func TestWriteBaseConfig(t *testing.T) {
	InitViper(enum.ConfigDevelopmentPath)
	WriteBaseConfig()
}

// TestViper tests the InitViper function.
func TestViper(t *testing.T) {
	InitViper(enum.ConfigDevelopmentPath)
	assert.Equal(t, enum.AppName, viper.GetString("app.name"))
	assert.Equal(t, "development", viper.GetString("app.mode"))
	assert.Equal(t, enum.AppPort, viper.GetString("app.port"))
	assert.Equal(t, enum.AppVersion, viper.GetString("app.version"))
}
