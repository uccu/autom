package middleware

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		path := c.Request.URL.Path
		start := time.Now()
		method := c.Request.Method

		defer func() {
			var user string
			user = c.ClientIP()
			latency := time.Since(start)
			logrus.Infof("%s|%s %s:%s, latency:%v", color.CyanString("%3s", "api"), color.WhiteString("%s", user), color.YellowString("%s", method), path, latency)
		}()

		c.Next()

	}
}
