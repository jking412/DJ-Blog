//package controller
//
//import (
//	"DJ-Blog/internal/http/request"
//	"DJ-Blog/model"
//	"DJ-Blog/pkg/search"
//	"github.com/gin-gonic/gin"
//	"github.com/sirupsen/logrus"
//	"net/http"
//	"strconv"
//)
//
//type PostController struct {
//}
//
//func (pc *PostController) Index(c *gin.Context) {
//	post := &model.Post{}
//	posts, _ := post.GetAllPosts()
//
//	var postIndexResp []request.PostIndexResp
//
//	for _, p := range posts {
//		user, _ := (&model.User{
//			BaseModel: model.BaseModel{
//				Id: p.UserId,
//			},
//		}).GetUserById()
//		postIndexResp = append(postIndexResp, request.PostIndexResp{
//			PostBaseResp: request.PostBaseResp{
//				Id:        p.Id,
//				CreatedAt: p.CreatedAt,
//				UpdatedAt: p.UpdatedAt,
//				Title:     p.Title,
//				Tag:       p.Tag,
//				Author:    p.Author,
//			},
//			AvatarUrl: user.AvatarUrl,
//		})
//	}
//
//	c.HTML(http.StatusOK, "index", gin.H{
//		"posts": postIndexResp,
//	})
//}
//
//func (pc *PostController) Create(c *gin.Context) {
//	c.HTML(http.StatusOK, "create", gin.H{})
//}
//
//func (pc *PostController) Store(c *gin.Context) {
//	Title := c.PostForm("title")
//	Tag := c.PostForm("tag")
//	Content := c.PostForm("content")
//	postCreateReq := &request.PostCreateReq{
//		Title:   Title,
//		Tag:     Tag,
//		Content: Content,
//	}
//	errs := request.ValidatePost(postCreateReq)
//	if len(errs) > 0 {
//		logrus.Warn("Request param is not valid", errs)
//		c.HTML(http.StatusBadRequest, "error", gin.H{
//			"errors": errs,
//		})
//		return
//	}
//	user := session.GetUser(c)
//	post := &model.Post{
//		Title:   postCreateReq.Title,
//		Tag:     postCreateReq.Tag,
//		Content: postCreateReq.Content,
//		UserId:  user.Id,
//		Author:  user.Username,
//	}
//	post.Create()
//	c.Redirect(http.StatusFound, "/")
//}
//
//func (pc *PostController) Detail(c *gin.Context) {
//	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
//	if err != nil {
//		logrus.Warn("Request param id is not a number")
//		c.HTML(http.StatusBadRequest, "error", gin.H{})
//		return
//	}
//	post := &model.Post{}
//	post.Id = id
//	post, _ = post.GetPostById()
//	user, _ := (&model.User{
//		BaseModel: model.BaseModel{
//			Id: post.UserId,
//		},
//	}).GetUserById()
//	postDetailResp := &request.PostDetailResp{
//		PostBaseResp: request.PostBaseResp{
//			Id:        post.Id,
//			CreatedAt: post.CreatedAt,
//			UpdatedAt: post.UpdatedAt,
//			Title:     post.Title,
//			Tag:       post.Tag,
//			Author:    post.Author,
//		},
//		AvatarUrl: user.AvatarUrl,
//		Content:   post.Content,
//		Likes:     post.Likes,
//		Stared:    post.Stared,
//	}
//	c.HTML(http.StatusOK, "detail", gin.H{
//		"post": postDetailResp,
//	})
//}
//
//func (pc *PostController) ShowUpdate(c *gin.Context) {
//	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
//	if err != nil {
//		logrus.Warn("Request param id is not a number")
//		c.HTML(http.StatusBadRequest, "error", gin.H{})
//		return
//	}
//	post := &model.Post{}
//	post.Id = id
//	post, _ = post.GetPostById()
//	if post.UserId != session.GetUser(c).Id {
//		c.HTML(http.StatusUnauthorized, "error", gin.H{})
//		return
//	}
//	postUpdateResp := &request.PostUpdateResp{
//		Id:      post.Id,
//		Title:   post.Title,
//		Tag:     post.Tag,
//		Content: post.Content,
//	}
//	c.HTML(http.StatusOK, "update", gin.H{
//		"post": postUpdateResp,
//	})
//}
//
//func (pc *PostController) Update(c *gin.Context) {
//	id, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)
//	userId, _ := (&model.Post{
//		BaseModel: model.BaseModel{
//			Id: id,
//		},
//	}).GetUserId()
//	if userId != session.GetUser(c).Id {
//		c.HTML(http.StatusUnauthorized, "error", gin.H{})
//		return
//	}
//	Title := c.PostForm("title")
//	Tag := c.PostForm("tag")
//	Content := c.PostForm("content")
//	postUpdateReq := &request.PostUpdateReq{
//		Id:      id,
//		Title:   Title,
//		Tag:     Tag,
//		Content: Content,
//	}
//	errs := request.ValidatePost(postUpdateReq)
//	if len(errs) > 0 {
//		logrus.Warn("Request param is not valid", errs)
//		c.HTML(http.StatusBadRequest, "error", gin.H{
//			"errors": errs,
//		})
//		return
//	}
//	post := &model.Post{
//		BaseModel: model.BaseModel{
//			Id: postUpdateReq.Id,
//		},
//		Title:   postUpdateReq.Title,
//		Tag:     postUpdateReq.Tag,
//		Content: postUpdateReq.Content,
//	}
//	post.Update()
//	c.Redirect(http.StatusFound, "/post/"+strconv.FormatUint(post.Id, 10))
//}
//
//func (pc *PostController) Delete(c *gin.Context) {
//	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
//	userId, _ := (&model.Post{
//		BaseModel: model.BaseModel{
//			Id: id,
//		},
//	}).GetUserId()
//	if userId != session.GetUser(c).Id {
//		c.HTML(http.StatusUnauthorized, "error", gin.H{})
//		return
//	}
//	if err != nil {
//		logrus.Warn("Request param id is not a number")
//		c.HTML(http.StatusBadRequest, "error", gin.H{})
//		return
//	}
//	post := &model.Post{}
//	post.Id = id
//	post.DeletePostById()
//	c.Redirect(http.StatusFound, "/")
//}
//
//func (pc *PostController) ShowSearch(c *gin.Context) {
//	c.HTML(http.StatusOK, "search", gin.H{})
//}
//
//func (pc *PostController) Search(c *gin.Context) {
//	keyword := c.PostForm("keyword")
//	rs, _ := search.ZincCli.Query("posts", keyword)
//	var postSearchResp []request.PostSearchResp
//	for _, r := range rs.Hits.HitItems {
//		source := r.Source.(map[string]interface{})
//		postSearchResp = append(postSearchResp, request.PostSearchResp{
//			Id:    uint64(source["id"].(float64)),
//			Title: source["title"].(string),
//		})
//	}
//	c.HTML(http.StatusOK, "search", gin.H{
//		"posts": postSearchResp,
//	})
//}
