package controller

import (
	"DJ-Blog/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostController struct {
}

func (pc *PostController) Index(c *gin.Context) {
	post := &model.Post{}
	posts := post.GetAllPosts()

	c.HTML(http.StatusOK, "index", gin.H{
		"posts": posts,
	})
}
