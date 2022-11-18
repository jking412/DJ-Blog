package controller

import (
	"DJ-Blog/internal/core"
	"DJ-Blog/internal/http/request"
	"DJ-Blog/internal/http/response"
	"DJ-Blog/internal/service"
	"DJ-Blog/pkg/paginations"
	"DJ-Blog/pkg/session"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ArticleController struct {
}

func NewArticleController() *ArticleController {
	return new(ArticleController)
}

func (a *ArticleController) Create(c *gin.Context) {

	req := &request.ArticleCreateReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Info("创建文章请求JSON格式错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	user := session.GetUser(c)
	req.UserId = user.Id

	value, ok := core.ArticleCreate(req)
	if !ok {
		logrus.Info("创建文章失败", value)
		response.EndWithUnsatisfiedRequest(c, value)
		return
	}

	userService := value.(*service.Article)

	response.EndWithOK(c, userService)
}

func (a *ArticleController) Delete(c *gin.Context) {
	articleId, err := strconv.ParseUint(c.Param("articleId"), 10, 32)

	if err != nil {
		logrus.Info("删除文章请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	ok := service.DeleteArticleById(uint32(articleId))
	if !ok {
		logrus.Info("删除文章失败")
		response.EndWithUnsatisfiedRequest(c, nil)
		return
	}

	response.EndWithOK(c, articleId)
}

func (a *ArticleController) Index(c *gin.Context) {

	req := request.PaginationReq{}

	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	articles, ok := service.GetArticleList()
	if !ok {
		logrus.Info("获取文章列表失败")
		response.EndWithInternalServerError(c, nil)
		return
	}

	articles, ok = paginations.ArticlePagination(articles, req.PageNum, req.PageSize)
	if !ok {
		logrus.Info("分页失败")
		response.EndWithUnsatisfiedRequest(c, nil)
		return
	}

	response.EndWithOK(c, articles)
}

func (a *ArticleController) Update(c *gin.Context) {
	articleId, err := strconv.ParseUint(c.Param("articleId"), 10, 32)
	req := &request.ArticleUpdateReq{}

	if err != nil {
		logrus.Info("删除文章请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	if err := c.ShouldBindJSON(req); err != nil {
		logrus.Info("创建文章请求JSON格式错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	user := session.GetUser(c)
	req.UserId = user.Id
	req.Id = uint32(articleId)

	value, ok := core.ArticleUpdate(req)
	if !ok {
		logrus.Info("更新文章失败", value)
		response.EndWithUnsatisfiedRequest(c, value)
		return
	}

	serviceArticle := value.(*service.Article)

	response.EndWithOK(c, serviceArticle)
}

func (a *ArticleController) ShowArticleDetail(c *gin.Context) {

	articleId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		logrus.Info("文章请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	article, ok := service.GetArticleById(uint32(articleId))
	if !ok {
		logrus.Info("获取文章失败")
		response.EndWithInternalServerError(c, nil)
		return
	}

	response.EndWithOK(c, article)
}

func (a *ArticleController) ShowArticleByTime(c *gin.Context) {

	req := request.PaginationReq{}

	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	articles, ok := service.GetArticleOrderByTime()
	if !ok {
		logrus.Info("获取文章列表失败")
		response.EndWithInternalServerError(c, nil)
		return
	}

	articles, ok = paginations.ArticlePagination(articles, req.PageNum, req.PageSize)
	if !ok {
		logrus.Info("分页失败")
		response.EndWithUnprocessableData(c, nil)
		return
	}

	response.EndWithOK(c, articles)
}

func (a *ArticleController) ShowArticleByTag(c *gin.Context) {

	req := request.PaginationReq{}

	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	tagId, err := strconv.ParseUint(c.Param("tagId"), 10, 32)
	if err != nil {
		logrus.Info("标签请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	articles, ok := service.GetArticleByTagId(uint32(tagId))
	if !ok {
		logrus.Info("获取文章列表失败")
		response.EndWithInternalServerError(c, nil)
		return
	}

	articles, ok = paginations.ArticlePagination(articles, req.PageNum, req.PageSize)
	if !ok {
		logrus.Info("分页失败")
		response.EndWithUnprocessableData(c, nil)
		return
	}

	response.EndWithOK(c, articles)
}

func (a *ArticleController) ShowByArticleCategory(c *gin.Context) {

	req := request.PaginationReq{}

	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	categoryId, err := strconv.ParseUint(c.Param("categoryId"), 10, 32)
	if err != nil {
		logrus.Info("分类请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	articles, ok := service.GetArticleByCategoryId(uint32(categoryId))
	if !ok {
		logrus.Info("获取文章列表失败")
		response.EndWithInternalServerError(c, nil)
		return
	}

	articles, ok = paginations.ArticlePagination(articles, req.PageNum, req.PageSize)
	if !ok {
		logrus.Info("分页失败")
		response.EndWithUnprocessableData(c, nil)
		return
	}

	response.EndWithOK(c, articles)
}

func (a *ArticleController) Search(c *gin.Context) {

}

func (a *ArticleController) ShowTags(c *gin.Context) {

	tags, ok := service.GetTags()
	if !ok {
		logrus.Info("获取标签列表失败")
		response.EndWithInternalServerError(c, nil)
		return
	}

	response.EndWithOK(c, tags)
}

func (a *ArticleController) ShowCategories(c *gin.Context) {

	categories, ok := service.GetCategories()
	if !ok {
		logrus.Info("获取分类列表失败")
		response.EndWithInternalServerError(c, nil)
		return
	}

	response.EndWithOK(c, categories)
}

func (a *ArticleController) ShowSpecificCategory(c *gin.Context) {

	req := request.PaginationReq{}

	if err := c.ShouldBindQuery(&req); err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	categoryId, err := strconv.ParseUint(c.Param("categoryId"), 10, 32)
	if err != nil {
		logrus.Info("分类请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	articles, ok := service.GetArticleByCategoryId(uint32(categoryId))
	if !ok {
		logrus.Info("获取文章列表失败")
		response.EndWithInternalServerError(c, nil)
		return
	}

	articles, ok = paginations.ArticlePagination(articles, req.PageNum, req.PageSize)
	if !ok {
		logrus.Info("分页失败")
		response.EndWithUnprocessableData(c, nil)
		return
	}

	response.EndWithOK(c, articles)
}
