package middleware

import (
	"DJ-Blog/internal/http/response"
	"DJ-Blog/pkg/session"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := session.GetUser(c)
		if user == nil {
			logrus.Info("用户未登录")
			response.EndWithUnauthorized(c, nil)
			c.Abort()
		}
		c.Next()
	}
}
