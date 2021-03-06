package middleware

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/uccu/autom/conf"
	"github.com/uccu/autom/http_error"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		defer func() {
			if err := recover(); err != nil {
				latency := time.Since(start)
				if e, ok := err.(http_error.HttpError); ok {
					logrus.Infof("response: [%d,%v] %s", e.ErrorCode, latency, e.ErrorMsg)
					Error(c, ResponseCode(e.ErrorCode), e.ErrorMsg)
					return
				}

				//先做一下日志记录
				logrus.Warnf("", map[string]interface{}{
					"error": fmt.Sprint(err),
					"stack": string(debug.Stack()),
				})

				if !conf.DebugMode {
					Error(c, 500, "内部错误")
					return
				} else {
					Error(c, 500, fmt.Sprint(err))
					return
				}
			}
		}()
		c.Next()
	}
}
