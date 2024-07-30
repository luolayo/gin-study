package install

import (
	"github.com/luolayo/gin-study/config"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/model"
	"github.com/luolayo/gin-study/util"
	"github.com/spf13/viper"
)

func initGorm(gormConfig *config.DatabaseConfig) bool {
	viper.Set("database.host", gormConfig.Host)
	viper.Set("database.port", gormConfig.Port)
	viper.Set("database.username", gormConfig.Username)
	viper.Set("database.password", gormConfig.Password)
	viper.Set("database.database", gormConfig.Database)
	viper.Set("database.MaxIdleConns", gormConfig.MaxIdleConns)
	viper.Set("database.MaxOpenConns", gormConfig.MaxOpenConns)
	viper.Set("database.ConnMaxLifetime", gormConfig.ConnMaxLifetime)
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func initRedis(redisConfig *config.RedisConfig) bool {
	viper.Set("redis.host", redisConfig.Host)
	viper.Set("redis.port", redisConfig.Port)
	viper.Set("redis.DB", redisConfig.DB)
	viper.Set("redis.DialTimeout", redisConfig.DialTimeout)
	viper.Set("redis.ReadTimeout", redisConfig.ReadTimeout)
	viper.Set("redis.WriteTimeout", redisConfig.WriteTimeout)
	viper.Set("redis.PoolSize", redisConfig.PoolSize)
	viper.Set("redis.PoolTimeout", redisConfig.PoolTimeout)
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func initAdminUser(adminUser *AdminConfig) bool {
	enPassword, err := util.Encrypt(adminUser.Password)
	if err != nil {
		global.LOG.Error("Failed to encrypt the password %v", err)
		return false
	}
	user := model.User{
		Name:       adminUser.Name,
		Password:   enPassword,
		Phone:      adminUser.Phone,
		Url:        adminUser.Url,
		ScreenName: adminUser.ScreenName,
		Group:      model.GroupAdmin,
	}
	client := global.DB.GetClient()
	tx := client.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		global.LOG.Error("Failed to create admin user %v", err)
		return false
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		global.LOG.Error("Failed to create admin user %v", err)
		return false
	}
	return true
}

func initSystemConfig() bool {
	viper.Set("app.installed", true)
	viper.Set("aliyun.accessKeyId", "")
	viper.Set("aliyun.accessKeySecret", "")
	viper.Set("aliyun.signName", "")
	viper.Set("aliyun.templateCode", "")
	err := viper.WriteConfig()
	if err != nil {
		panic(err)
		return false
	}
	return true
}

func testAdminIsExist() bool {
	client := global.DB.GetClient()
	user := model.User{}
	client.Where(user.Group, model.GroupAdmin).First(&user)
	if user.Uid > 0 {
		return true
	}
	return false
}
