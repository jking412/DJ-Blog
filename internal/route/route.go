package route

import (
	_ "DJ-Blog/docs"
	"DJ-Blog/internal/controller"
	"DJ-Blog/pkg/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

//@title DJ-Blog API文档
//@description DJ-Blog API文档
//@version 1.0

// host  localhost:8000

func RegisterRoutes(r *gin.Engine) {
	r.Use(sessions.Sessions("DJ-Blog", session.Store))

	r.Any("/ping", Ping)

	registerUserRoutes(r.Group("/user"), controller.NewUserController())

	// 处理swagger文档的接口配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(NoRouteHandler)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}

func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"msg": "页面不存在",
	})
}
