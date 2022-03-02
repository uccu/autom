package middleware

import (
	"autom/conf"
	"autom/http_error"

	"github.com/gin-gonic/gin"
)

func HeaderAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !conf.Http.HeaderCheck {
			c.Next()
			return
		}

		token := c.GetHeader("X-Gitlab-Token")

		if token == "" {
			panic(http_error.NoXGitlabToken)
		}

		if token != conf.Http.Token {
			panic(http_error.XGitlabTokenNotMatch)
		}

		c.Next()
	}
}
