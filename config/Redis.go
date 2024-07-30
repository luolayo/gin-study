package config

type RedisConfig struct {
	// Host redis host
	Host string `json:"host" binding:"required" form:"host" example:"localhost"`
	// Port redis port
	Port string `json:"port" binding:"required" form:"port" example:"6379"`
	// DB Which repository is it
	DB int `json:"db" binding:"required" form:"db" example:"1"`
	// DialTimeout redis dial timeout
	DialTimeout int `json:"dialTimeout" binding:"required" form:"dialTimeout" example:"10"`
	// ReadTimeout redis read timeout
	ReadTimeout int `json:"readTimeout" binding:"required" form:"readTimeout" example:"30"`
	// WriteTimeout redis write timeout
	WriteTimeout int `json:"writeTimeout" binding:"required" form:"writeTimeout" example:"30"`
	// PoolSize redis pool size
	PoolSize int `json:"poolSize" binding:"required" form:"poolSize" example:"10"`
	// PoolTimeout redis pool timeout
	PoolTimeout int `json:"poolTimeout" binding:"required" form:"poolTimeout" example:"30"`
}
