package user

type RegisterDTO struct {
	// User name
	Name string `json:"name" form:"name" binding:"required" example:"admin"`
	// User password
	Password string `json:"password" form:"password" binding:"required" example:"123456"`
	// Confirm password is the same as password
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" binding:"required"  example:"123456"`
	// User phone number
	Phone string `json:"phone" form:"phone" binding:"required"  example:"18888888888"`
	// User avatar
	Url string `json:"url" form:"url"  example:"https://www.luola.me"`
	// User nickname
	ScreenName string `json:"screenName" form:"screenName"  example:"罗拉"`
	// Verification code
	Code string `json:"code" form:"code" binding:"required"  example:"123456"`
}

type LoginDTO struct {
	// User name
	Name string `json:"name" form:"name" binding:"required" example:"admin"`
	// User password
	Password string `json:"password" form:"password" binding:"required" example:"123456"`
}

type UpdateDTO struct {
	// User url
	Url string `json:"url" form:"url" example:"https://www.luola.me"`
	// User nickname
	ScreenName string `json:"screenName" example:"罗拉"`
}
