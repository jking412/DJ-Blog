package route

import (
	"DJ-Blog/pkg/session"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(sessions.Sessions("DJ-Blog", session.Store))

	r.Any("/ping", Ping)

	//userGroup := r.Group("/user")
	//{
	//
	//}

	//uc := &controller.UserController{}
	//
	//r.GET("/", pc.Index)
	//
	//userGroup := r.Group("/user")
	//{
	//	userGroup.GET("/register", uc.ShowRegister)
	//	userGroup.POST("/register", uc.RegisterRoutes)
	//	userGroup.GET("/login", uc.ShowLogin)
	//	userGroup.POST("/login", uc.Login)
	//
	//	userGroup.GET("/github/login", uc.GithubLogin)
	//	userGroup.GET("/github/oauth2", uc.GithubLoginCallback)
	//
	//	userGroup.GET("/logout", uc.Logout)
	//}
	//
	//postGroup := r.Group("/post")
	//{
	//	postGroup.GET("/:id", pc.Detail)
	//
	//	postGroup.GET("/search", pc.ShowSearch)
	//	postGroup.POST("/search", pc.Search)
	//
	//	postGroup.Use(middleware.Auth())
	//
	//	postGroup.GET("/create", pc.Create)
	//	postGroup.POST("/store", pc.Store)
	//
	//	postGroup.GET("/update/:id", pc.ShowUpdate)
	//	postGroup.POST("/update", pc.Update)
	//
	//	postGroup.GET("/delete/:id", pc.Delete)
	//
	//}

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
