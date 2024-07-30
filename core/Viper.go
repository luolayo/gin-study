package core

import (
	"github.com/luolayo/gin-study/enum"
	"github.com/spf13/viper"
	"os"
)

// InitViper set viper conifg and read config file
func InitViper(path string) {
	viper.SetConfigName(enum.ConfigFileName)
	viper.SetConfigType(enum.ConfigFileType)
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
}

// CheckConfigFile checks if the configuration file exists
func CheckConfigFile(path string) bool {
	if _, err := os.Stat(path + "config.json"); os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateConfigFile creates a configuration file
func CreateConfigFile(path string) {
	// Create configuration file
	file, err := os.Create(path + "config.json")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	writeString, err := file.WriteString(`{}`)
	if err != nil {
		panic(err)
	}
	if writeString == 0 {
		panic("Failed to write configuration file")
	}
}

// WriteBaseConfig writes the basic configuration
func WriteBaseConfig() {
	viper.SetDefault("app.name", enum.AppName)
	viper.SetDefault("app.mode", enum.AppMode)
	viper.SetDefault("app.port", enum.AppPort)
	viper.SetDefault("app.version", enum.AppVersion)
	viper.SetDefault("app.Installed", false)
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
	}
}
