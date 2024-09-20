package test

import (
	"github.com/gin-gonic/gin"
	"github.com/luolayo/gin-study/core/ip"
	"github.com/luolayo/gin-study/interceptor/res"
	"log"
)

// GetIPAddredd godoc
// @Summary GetIPAddredd
// @Tags Test
// @Scheme http https
// @Produce json
// @Accept x-www-form-urlencoded
// @Success 200 {object} res.Response[ip.Address]
// @Router /test/IP [get]
func GetIPAddredd(c *gin.Context) {
	IP := ip.NewIp()
	IP.Init()
	address, err := IP.Find(c.ClientIP())
	if err != nil {
		log.Fatalf("ip type error %v", err)
	}
	res.Success(c, address)
}
