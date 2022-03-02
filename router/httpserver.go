package router

import (
	"autom/conf"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HttpServerRun() {

	gin.SetMode("release")
	r := InitRouter()

	go func() {
		if err := r.Run(":" + conf.Http.Port); err != nil {
			logrus.Errorf("HttpServerRun:%s err:%v", ":"+conf.Http.Port, err)
		}
	}()
}
