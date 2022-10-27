package middleware

import (
	"DJ-Blog/pkg/sessionpkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := sessionpkg.GetUser(c)
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			c.Abort()
		} else {
			sessionpkg.FlushSession(c)
			c.Next()
		}
	}
}
