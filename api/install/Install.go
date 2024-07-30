package install

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/config"
	"github.com/luolayo/gin-study/global"
	"github.com/luolayo/gin-study/interceptor/res"
	"github.com/luolayo/gin-study/interceptor/validate"
)

type Step struct {
	Step int `json:"step" example:"1"`
}

// CheckStep
// @Summary Check the installation step
// @Description Check which step the installation program has reached, return to the required steps, and if the installation is complete, return the 201 status code
// @Tags Install
// @Scheme http https
// @Produce json
// @Success 201 {object} res.Response[res.Empty]
// @Success 200 {object} res.Response[Step]
// @Router /install/check [get]
func CheckStep(c *gin.Context) {
	appConfig := config.GetAppConfig()
	if appConfig.Installed {
		res.Created(c)
		return
	}
	if !global.DB.CheckGormConnection() {
		res.Success(c, Step{
			Step: 1,
		})
		return
	}
	if !global.Redis.CheckRedisConnection() {
		res.Success(c, Step{
			Step: 2,
		})
		return
	}
	res.Success(c, Step{
		Step: 3,
	})
}

// Step1
// @Summary Initialize the database
// @Description Write the database configuration steps. If the database has already been configured, return a 201 status code
// @Tags Install
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data body config.DatabaseConfig true "Database configuration"
// @Success 201 {object} res.Response[res.Empty]
// @Success 200 {object} res.Response[res.Empty]
// @Failure 400 {object} res.ErrorRes[[]string]
// @Failure 500 {object} res.ErrorRes[nil]
// @Router /install/step1 [post]
func Step1(c *gin.Context) {
	if global.DB.CheckGormConnection() {
		res.Created(c)
		return
	}
	gormConfig := config.DatabaseConfig{}
	if err := c.ShouldBind(&gormConfig); err != nil {
		res.BadRequest(c, validate.Err(err))
		return
	}
	// Check if the database is connected. If the database connection is successful, modifications are not allowed
	if !initGorm(&gormConfig) {
		res.ServerError(c, "Writing configuration file failed")
		return
	}
	global.InitDB()
	if !global.DB.CheckGormConnection() {
		res.ServerError(c, "Failed to initialize gorm2")
		return
	}
	global.InitDB()
	err := global.AutoMigrate()
	if err != nil {
		res.ServerError(c, "Failed to migrate database")
		return
	}
	res.SuccessNoData(c)
}

// Step2
// @Summary Initialize the redis
// @Description Write the redis configuration steps. If the redis has already been configured, return a 201 status code
// @Tags Install
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data body config.RedisConfig true "Reids configuration"
// @Success 201 {object} res.Response[res.Empty]
// @Success 200 {object} res.Response[res.Empty]
// @Failure 400 {object} res.ErrorRes[[]string]
// @Failure 500 {object} res.ErrorRes[nil]
// @Router /install/step2 [post]
func Step2(c *gin.Context) {
	if global.Redis.CheckRedisConnection() {
		res.Created(c)
		return
	}
	redisConfig := config.RedisConfig{}
	if err := c.ShouldBind(&redisConfig); err != nil {
		res.BadRequest(c, validate.Err(err))
		return
	}
	if !initRedis(&redisConfig) {
		res.ServerError(c, "Writing configuration file failed")
		return
	}
	global.InitRedis()
	if !global.Redis.CheckRedisConnection() {
		res.ServerError(c, "Failed to initialize redis")
		return
	}
	res.SuccessNoData(c)
}

// Step3
// @Summary Create an admin account
// @Description Used to create a system administrator account, which can only have one and can only be created through this step during system initialization
// @Tags Install
// @Accept x-www-form-urlencoded
// @Produce json
// @Param data body AdminConfig true "Reids configuration"
// @Success 201 {object} res.Response[res.Empty]
// @Success 200 {object} res.Response[res.Empty]
// @Failure 400 {object} res.ErrorRes[[]string]
// @Failure 500 {object} res.ErrorRes[nil]
// @Router /install/step3 [post]
func Step3(c *gin.Context) {
	if testAdminIsExist() {
		res.Created(c)
	}
	adminUser := AdminConfig{}
	if err := c.ShouldBind(&adminUser); err != nil {
		res.BadRequest(c, validate.Err(err))
		return
	}
	if !initAdminUser(&adminUser) {
		res.ServerError(c, "Failed to create admin user")
		return
	}
	if !initSystemConfig() {
		res.ServerError(c, "Failed to initialize system configuration")
		return
	}
	res.SuccessNoData(c)
}
