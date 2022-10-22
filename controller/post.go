package controller

import (
	"DJ-Blog/controller/request"
	"DJ-Blog/model"
	"DJ-Blog/pkg/search"
	"DJ-Blog/pkg/sessionpkg"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type PostController struct {
}

type PostIndexResp struct {
	request.PostBaseResp
	AvatarUrl string
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
	Title   string
	Tag     string
	Content string
}

type PostUpdateResp struct {
	Id      uint64
	Title   string
	Tag     string
	Content string
}

type PostSearchResp struct {
	Id    uint64
	Title string
}

func (pc *PostController) Index(c *gin.Context) {
	post := &model.Post{}
	posts, _ := post.GetAllPosts()

	var postIndexResp []PostIndexResp

	for _, p := range posts {
		user, _ := (&model.User{
			BaseModel: model.BaseModel{
				Id: p.UserId,
			},
		}).GetUserById()
		postIndexResp = append(postIndexResp, PostIndexResp{
			PostBaseResp: request.PostBaseResp{
				Id:        p.Id,
				CreatedAt: p.CreatedAt,
				UpdatedAt: p.UpdatedAt,
				Title:     p.Title,
				Tag:       p.Tag,
				Author:    p.Author,
			},
			AvatarUrl: user.AvatarUrl,
		})
	}

	c.HTML(http.StatusOK, "index", gin.H{
		"posts": postIndexResp,
	})
}

func (pc *PostController) Create(c *gin.Context) {
	c.HTML(http.StatusOK, "create", gin.H{})
}

func (pc *PostController) Store(c *gin.Context) {
	Title := c.PostForm("title")
	Tag := c.PostForm("tag")
	Content := c.PostForm("content")
	postCreateReq := &PostCreateReq{
		Title:   Title,
		Tag:     Tag,
		Content: Content,
	}
	user := sessionpkg.GetUser(c)
	post := &model.Post{
		Title:   postCreateReq.Title,
		Tag:     postCreateReq.Tag,
		Content: postCreateReq.Content,
		UserId:  user.Id,
		Author:  user.Username,
	}
	post.Create()
	c.Redirect(http.StatusFound, "/")
}

func (pc *PostController) Detail(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Warn("Request param id is not a number")
		c.HTML(http.StatusBadRequest, "error", gin.H{})
		return
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
			Author:    post.Author,
		},
		Content: post.Content,
		Likes:   post.Likes,
		Stared:  post.Stared,
	}
	c.HTML(http.StatusOK, "detail", gin.H{
		"post": postDetailResp,
	})
}

func (pc *PostController) ShowUpdate(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logrus.Warn("Request param id is not a number")
		c.HTML(http.StatusBadRequest, "error", gin.H{})
		return
	}
	post := &model.Post{}
	post.Id = id
	post, _ = post.GetPostById()
	if post.UserId != sessionpkg.GetUser(c).Id {
		c.HTML(http.StatusUnauthorized, "error", gin.H{})
		return
	}
	postUpdateResp := &PostUpdateResp{
		Id:      post.Id,
		Title:   post.Title,
		Tag:     post.Tag,
		Content: post.Content,
	}
	c.HTML(http.StatusOK, "update", gin.H{
		"post": postUpdateResp,
	})
}

func (pc *PostController) Update(c *gin.Context) {
	id, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)
	userId, _ := (&model.Post{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}).GetUserId()
	if userId != sessionpkg.GetUser(c).Id {
		c.HTML(http.StatusUnauthorized, "error", gin.H{})
		return
	}
	Title := c.PostForm("title")
	Tag := c.PostForm("tag")
	Content := c.PostForm("content")
	postUpdateReq := &PostUpdateReq{
		Id:      id,
		Title:   Title,
		Tag:     Tag,
		Content: Content,
	}
	post := &model.Post{
		BaseModel: model.BaseModel{
			Id: postUpdateReq.Id,
		},
		Title:   postUpdateReq.Title,
		Tag:     postUpdateReq.Tag,
		Content: postUpdateReq.Content,
	}
	post.Update()
	c.Redirect(http.StatusFound, "/post/"+strconv.FormatUint(post.Id, 10))
}

func (pc *PostController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	userId, _ := (&model.Post{
		BaseModel: model.BaseModel{
			Id: id,
		},
	}).GetUserId()
	if userId != sessionpkg.GetUser(c).Id {
		c.HTML(http.StatusUnauthorized, "error", gin.H{})
		return
	}
	if err != nil {
		logrus.Warn("Request param id is not a number")
		c.HTML(http.StatusBadRequest, "error", gin.H{})
		return
	}
	post := &model.Post{}
	post.Id = id
	post.DeletePostById()
	c.Redirect(http.StatusFound, "/")
}

func (pc *PostController) ShowSearch(c *gin.Context) {
	c.HTML(http.StatusOK, "search", gin.H{})
}

func (pc *PostController) Search(c *gin.Context) {
	keyword := c.PostForm("keyword")
	rs, _ := search.ZincCli.Query("posts", keyword)
	var postSearchResp []PostSearchResp
	for _, r := range rs.Hits.HitItems {
		source := r.Source.(map[string]interface{})
		postSearchResp = append(postSearchResp, PostSearchResp{
			Id:    uint64(source["id"].(float64)),
			Title: source["title"].(string),
		})
	}
	c.HTML(http.StatusOK, "search", gin.H{
		"posts": postSearchResp,
	})
}
