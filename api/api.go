package api

import (
	"DJ-Blog/controller"
	"DJ-Blog/pkg/sessionpkg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.Engine) {
	r.Use(sessions.Sessions("DJ-Blog", sessionpkg.Store))

	r.Any("/ping", Ping)

	pc := &controller.PostController{}
	uc := &controller.UserController{}

	userGroup := r.Group("/user")
	{
		userGroup.GET("/register", uc.ShowRegister)
		userGroup.POST("/register", uc.Register)
		userGroup.GET("/login", uc.ShowLogin)
		userGroup.POST("/login", uc.Login)

		userGroup.GET("/github/login", uc.GithubLogin)
		userGroup.GET("/github/oauth2", uc.GithubLoginCallback)

		userGroup.GET("/logout", uc.Logout)
	}

	postGroup := r.Group("/post")
	//postGroup.Use(middleware.Auth())
	{
		postGroup.GET("/:id", pc.Detail)

		postGroup.GET("/create", pc.Create)
		postGroup.POST("/store", pc.Store)

		postGroup.GET("/update/:id", pc.ShowUpdate)
		postGroup.POST("/update", pc.Update)

		postGroup.GET("/delete/:id", pc.Delete)
	}

	//r.Use(middleware.Auth())
	r.GET("/", pc.Index)

	r.NoRoute(NoRouteHandler)

}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}

func NoRouteHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error", gin.H{})
}
