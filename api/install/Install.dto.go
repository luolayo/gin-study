package install

type AdminConfig struct {
	Password   string `json:"password" binding:"required" form:"password" example:"123456"`
	Name       string `json:"name" binding:"required" form:"name" example:"admin"`
	Phone      string `json:"phone" binding:"required" form:"phone" example:"12345678901"`
	Url        string `json:"url" binding:"required" form:"url" example:"http://localhost"`
	ScreenName string `json:"screenName" binding:"required" form:"screenName" example:"admin"`
}
