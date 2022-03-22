package api

import (
	"autom/middleware"
	"autom/service/hook"
	"encoding/json"
	"fmt"
	"io/ioutil"

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

	byt, _ := ioutil.ReadAll(c.Request.Body)

	c.String(200, string(byt))

	fmt.Println(string(byt))
	fmt.Println()
	b, _ := json.Marshal(c.Request.Header)
	fmt.Println(string(b))

	return

	client := hook.NewHookClient(c)
	if client == nil {
		logrus.Warn("config not exist")
		c.AbortWithStatus(404)
		return
	}

	go hook.Run(client)
	middleware.Success(c)
}
