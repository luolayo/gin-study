package install

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/util"
	"os"
)

type Data struct {
	Name       string `json:"name" binding:"required" form:"name"`
	Phone      string `json:"phone" binding:"required" form:"phone"`
	Passsword  string `json:"password" binding:"required" form:"password"`
	ScreenName string `json:"screenName" binding:"required" form:"screenName"`
	Url        string `json:"url" binding:"required" form:"url"`
}
type Site struct {
	Title       string `json:"title" binding:"required" form:"title"`
	Description string `json:"description" binding:"required" form:"description"`
	Keywords    string `json:"keywords" binding:"required" form:"keywords"`
	Url         string `json:"url" binding:"required" form:"url"`
}

func ApplicationInitialization(c *gin.Context) {
	// check mysql and redis connection
	c.HTML(200, "install.html", gin.H{})
}

func CreateAdminUser(c *gin.Context) {
	// create admin user
	data := Data{}
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// create admin user
	encryptPasswd, _ := util.Encrypt(data.Passsword)
	user := model.User{
		Name:       data.Name,
		Phone:      data.Phone,
		Password:   encryptPasswd,
		ScreenName: data.ScreenName,
		Url:        data.Url,
		Group:      model.GroupAdmin,
	}
	if err := global.GormDB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// create install.lock
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		err := os.Mkdir("log", 0755)
		if err != nil {
			panic(err)
		} // Create log directory
	}
	file, err := os.OpenFile("install/install.lock", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	// 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	c.JSON(200, gin.H{"message": "success"})
}
