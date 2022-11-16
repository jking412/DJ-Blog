package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := session.GetUser(c)
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			c.Abort()
		} else {
			session.FlushSession(c)
			c.Next()
		}
	}
}
