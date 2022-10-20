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

type PostCreateReq struct {
	Title   string
	Tag     string
	Content string
}

type PostDetailResp struct {
	request.PostBaseResp
	Content string
	Likes   uint64
	Stared  bool
}

type PostUpdateReq struct {
	Id      uint64
	Content string
	Likes   uint64
	Stared  bool
}

func (pc *PostController) Index(c *gin.Context) {
	post := &model.Post{}
	posts, _ := post.GetAllPosts()

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

func (pc *PostController) Create(c *gin.Context) {
	Title := c.PostForm("title")
	Tag := c.PostForm("tag")
	Content := c.PostForm("content")
	postCreateReq := &PostCreateReq{
		Title:   Title,
		Tag:     Tag,
		Content: Content,
	}
	post := &model.Post{
		Title:   postCreateReq.Title,
		Tag:     postCreateReq.Tag,
		Content: postCreateReq.Content,
	}
	post.Create()
	c.Redirect(http.StatusFound, "/")
}

func (pc *PostController) Detail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Warn("Request param id is not a number")
		c.HTML(http.StatusBadRequest, "error", gin.H{})
	}
	post := &model.Post{}
	post.Id = id
	post, _ = post.GetPostById()
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

func (pc *PostController) Update(c *gin.Context) {
	Id, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)
	Content := c.PostForm("content")
	Likes, _ := strconv.ParseUint(c.PostForm("likes"), 10, 64)
	Stared, _ := strconv.ParseBool(c.PostForm("stared"))
	postUpdateReq := &PostUpdateReq{
		Id:      Id,
		Content: Content,
		Likes:   Likes,
		Stared:  Stared,
	}
	post := &model.Post{
		BaseModel: model.BaseModel{
			Id: postUpdateReq.Id,
		},
		Content: postUpdateReq.Content,
		Likes:   postUpdateReq.Likes,
		Stared:  postUpdateReq.Stared,
	}
	post.Update()
	c.Redirect(http.StatusFound, "/post/"+strconv.FormatUint(post.Id, 10))
}

func (pc *PostController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Warn("Request param id is not a number")
		c.HTML(http.StatusBadRequest, "error", gin.H{})
	}
	post := &model.Post{}
	post.Id = id
	post.DeletePostById()
	c.Redirect(http.StatusFound, "/")
}
