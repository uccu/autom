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
		logrus.Warn("配置不存在")
		c.AbortWithStatus(404)
		return
	}

	client.CheckRight()

	if !client.CheckRight() {
		logrus.Warn("权限验证失败")
		c.AbortWithStatus(404)
		return
	}

	client.Run()
	middleware.Success(c)
}
