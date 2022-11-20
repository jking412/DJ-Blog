package middleware

import (
	"DJ-Blog/pkg/session"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := session.GetUser(c)
		fmt.Println(user)
		c.Next()
	}
}
