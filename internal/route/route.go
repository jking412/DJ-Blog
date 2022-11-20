package route

import (
	_ "DJ-Blog/docs"
	"DJ-Blog/front"
	"DJ-Blog/internal/controller"
	"DJ-Blog/internal/middleware"
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

	r.Use(middleware.Cors())

	r.Any("/ping", Ping)
	r.GET("/", Index)

	registerUserRoutes(r, controller.NewUserController())
	registerArticleRoutes(r, controller.NewArticleController())
	registerTagRoutes(r, controller.NewArticleController())
	registerCategoryRoutes(r, controller.NewArticleController())

	// 处理前端静态资源路由
	r.StaticFS("/static", http.FS(&front.UI{
		FS:         front.Build,
		StaticPath: "build/static/",
	}))

	// 处理swagger文档的接口配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(NoRouteHandler)
}

func Index(c *gin.Context) {

	indexHTML, err := front.Build.ReadFile("build/index.html")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "页面不存在",
		})
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", indexHTML)
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
