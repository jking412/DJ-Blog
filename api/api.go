package api

import (
	"DJ-Blog/controller"
	"DJ-Blog/controller/middleware"
	"DJ-Blog/pkg/sessionpkg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func Register(r *gin.Engine) {
	r.Use(sessions.Sessions("DJ-Blog", sessionpkg.Store))

	r.Any("/ping", Ping)

	//r.GET("/test", Test)

	pc := &controller.PostController{}
	uc := &controller.UserController{}

	r.GET("/", pc.Index)

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
	{
		postGroup.GET("/:id", pc.Detail)

		postGroup.GET("/search", pc.ShowSearch)
		postGroup.POST("/search", pc.Search)

		postGroup.Use(middleware.Auth())

		postGroup.GET("/create", pc.Create)
		postGroup.POST("/store", pc.Store)

		postGroup.GET("/update/:id", pc.ShowUpdate)
		postGroup.POST("/update", pc.Update)

		postGroup.GET("/delete/:id", pc.Delete)

	}

	r.NoRoute(NoRouteHandler)

}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}

type Habit struct {
	Name    string
	content string
	time    int
}

func Test(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	//msg := rand.Int()
	time.Sleep(2 * time.Second)
	habits := []Habit{
		{
			"1",
			"2",
			3,
		},
		{
			"134",
			"234",
			3,
		},
		{
			"213",
			"2341",
			3,
		},
	}
	c.HTML(http.StatusOK, "test", gin.H{
		"habits": habits,
	})
}

func NoRouteHandler(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error", gin.H{})
}
