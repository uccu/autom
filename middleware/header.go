package middleware

import (
	"autom/conf"

	"github.com/gin-gonic/gin"
)

func HeaderAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !conf.Http.HeaderCheck {
			c.Next()
			return
		}

		c.Next()
	}
}
