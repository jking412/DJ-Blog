package core

import (
	"DJ-Blog/internal/http/request"
	"DJ-Blog/internal/model"
	"DJ-Blog/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type IArticleController interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Index(c *gin.Context)
	Update(c *gin.Context)

	ShowArticleDetail(c *gin.Context)
	ShowArticleByTime(c *gin.Context)
	ShowArticleByTag(c *gin.Context)
	ShowArticleByCategory(c *gin.Context)

	Search(c *gin.Context)
}

type ITagController interface {
	ShowTags(c *gin.Context)
}

type ICategoryController interface {
	ShowCategories(c *gin.Context)
}

func ArticleCreate(req *request.ArticleCreateReq) (interface{}, bool) {
	errs := make(map[string][]string)
	errs = request.ValidateArticleReq(req)

	if len(errs) > 0 {
		return errs, false
	}

	article := &service.Article{
		ArticleModel: model.ArticleModel{
			Title:         req.Title,
			OriginContent: req.OriginContent,
			ParseContent:  req.ParseContent,
			ImgUrl:        req.ImgUrl,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			UserId:        req.UserId,
		},
	}

	for _, tagId := range req.Tags {
		article.Tags = append(article.Tags, model.TagModel{Id: tagId})
	}

	for _, categoryId := range req.Categories {
		article.Categories = append(article.Categories, model.CategoryModel{Id: categoryId})
	}

	var serviceArticle *service.Article

	if serviceArticle = service.CreateArticle(article); serviceArticle.Id == 0 {
		logrus.Error("服务器创建文章失败")
		errs["server"] = []string{"服务器创建文章失败"}
		return errs, false
	}

	return serviceArticle, true
}

func ArticleUpdate(req *request.ArticleUpdateReq) (interface{}, bool) {
	errs := make(map[string][]string)
	errs = request.ValidateUpdateArticleReq(req)

	if len(errs) > 0 {
		return errs, false
	}

	serviceArticle := &service.Article{
		ArticleModel: model.ArticleModel{
			Id:            req.Id,
			Title:         req.Title,
			OriginContent: req.OriginContent,
			ParseContent:  req.ParseContent,
			ImgUrl:        req.ImgUrl,
			UpdatedAt:     time.Now(),
			UserId:        req.UserId,
		},
	}

	for _, tagId := range req.Tags {
		serviceArticle.Tags = append(serviceArticle.Tags, model.TagModel{Id: tagId})
	}

	for _, categoryId := range req.Categories {
		serviceArticle.Categories = append(serviceArticle.Categories, model.CategoryModel{Id: categoryId})
	}

	if !service.UpdateArticle(serviceArticle) {
		logrus.Error("服务器更新文章失败")
		errs["server"] = []string{"服务器更新文章失败"}
		return errs, false
	}

	return serviceArticle, true
}
