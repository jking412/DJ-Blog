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

	//postGroup := r.Group("/post")
	//{
	//}

}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}
