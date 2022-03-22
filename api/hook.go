package api

import (
	"github.com/uccu/autom/middleware"
	"github.com/uccu/autom/service/hook"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HookController struct {
}

func HookRegister(r *gin.RouterGroup) {
	c := HookController{}
	r.Any("/hook", c.hook)
}

func (*HookController) hook(c *gin.Context) {

	client := hook.NewHookClient(c)
	if client == nil {
		logrus.Warn("config not exist")
		c.AbortWithStatus(404)
		return
	}

	go hook.Run(client)
	middleware.Success(c)
}
