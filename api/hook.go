package api

import (
	"autom/middleware"
	"autom/service/hook"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HookController struct {
}

func HookRegister(r *gin.RouterGroup) {
	c := HookController{}
	r.POST("/hook", c.hook)
}

func (*HookController) hook(c *gin.Context) {

	client := hook.NewHookClient(c)
	if client == nil {
		logrus.Warn("config not exist")
		c.AbortWithStatus(404)
		return
	}

	if !client.CheckRight() {
		logrus.Warn("permission validation failed")
		c.AbortWithStatus(404)
		return
	}

	if !client.Run() {
		c.AbortWithStatus(404)
		return
	}
	middleware.Success(c)
}
