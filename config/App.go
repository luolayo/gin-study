package config

import "github.com/luolayo/gin-study/enum"

// AppConfig struct
type AppConfig struct {
	// Name of the application
	Name string
	// Port of the application
	Port string
	// Version of the application
	Version string
	// Mode of the application (development, release)
	Mode enum.AppModel
	// IsInstalled is the application installed
	Installed bool
}
