package router

import (
	"github.com/uccu/autom/conf"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/uccu/go-stringify"
)

func HttpServerRun() {

	gin.SetMode("release")
	r := InitRouter()

	go func() {
		if err := r.Run(":" + stringify.ToString(conf.Http.Port)); err != nil {
			logrus.Errorf("HttpServerRun:%s err:%v", ":"+stringify.ToString(conf.Http.Port), err)
		}
	}()
}
