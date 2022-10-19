package api

import (
	"DJ-Blog/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.Engine) {
	r.Any("/ping", Ping)

	pc := &controller.PostController{}

	r.GET("/", pc.Index)

	postGroup := r.Group("/post")
	{
		postGroup.GET("/:id", pc.Detail)
		postGroup.GET("/delete/:id", pc.Delete)
	}

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
