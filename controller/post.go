package controller

import (
	"DJ-Blog/controller/request"
	"DJ-Blog/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type PostController struct {
}

type PostIndexResp struct {
	request.PostBaseResp
}

type PostDetailResp struct {
	request.PostBaseResp
	Content string
	Likes   uint64
	Stared  bool
}

func (pc *PostController) Index(c *gin.Context) {
	post := &model.Post{}
	posts := post.GetAllPosts()

	var postIndexResp []PostIndexResp

	for _, p := range posts {
		postIndexResp = append(postIndexResp, PostIndexResp{
			PostBaseResp: request.PostBaseResp{
				Id:        p.Id,
				CreatedAt: p.CreatedAt,
				UpdatedAt: p.UpdatedAt,
				Title:     p.Title,
				Tag:       p.Tag,
			},
		})
	}

	c.HTML(http.StatusOK, "index", gin.H{
		"posts": postIndexResp,
	})
}

func (pc *PostController) Detail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Warn("Request param id is not a number")
		c.HTML(http.StatusBadRequest, "error", gin.H{})
	}
	post := &model.Post{}
	post.Id = id
	post = post.GetPostById()
	postDetailResp := &PostDetailResp{
		PostBaseResp: request.PostBaseResp{
			Id:        post.Id,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
			Title:     post.Title,
			Tag:       post.Tag,
		},
		Content: post.Content,
		Likes:   post.Likes,
		Stared:  post.Stared,
	}
	c.HTML(http.StatusOK, "detail", gin.H{
		"post": postDetailResp,
	})
}
