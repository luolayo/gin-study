package config

type Redis struct {
	Host         string
	Port         string
	DB           string
	DialTimeout  string
	ReadTimeout  string
	WriteTimeout string
	PoolSize     string
	PoolTimeout  string
}
