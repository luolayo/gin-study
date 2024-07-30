package config

// DatabaseConfig struct
type DatabaseConfig struct {
	// Host Database address
	Host string `json:"host" binding:"required" form:"host" example:"localhost"`
	// Port Database port
	Port string `json:"port" binding:"required" form:"port" example:"3306"`
	// Username Database username
	Username string `json:"username" binding:"required" form:"username" example:"root"`
	// Password Database password
	Password string `json:"password" binding:"required" form:"password" example:"123456"`
	// Database Database name
	Database string `json:"database" binding:"required" form:"database" example:"gin_study"`
	// MaxIdleConns Maximum number of idle connections
	MaxIdleConns int `json:"maxIdleConns" binding:"required" form:"maxIdleConns" example:"10"`
	// MaxOpenConns Maximum number of open connections
	MaxOpenConns int `json:"maxOpenConns" binding:"required" form:"maxOpenConns" example:"30"`
	// ConnMaxLifetime Maximum connection lifetime
	ConnMaxLifetime int `json:"connMaxLifetime" binding:"required" form:"connMaxLifetime" example:"60"`
}
