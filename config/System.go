package config

type System struct {
	AppName     string // The AppName of the application
	AppVersion  string // The AppVersion of the application
	Host        string // 	The Host the application will run on
	Port        string // The Port the application will run on
	Environment string // The Environment the application is running in
	CryPtKey    string // The CryPtKey of the application
	JwtSecret   string // The JWTKey of the application
}
