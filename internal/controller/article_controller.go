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

// Create godoc
// @Summary 创建文章
// @Description 创建文章
// @Tags Article
// @Accept  json
// @Produce  json
// @Param ArticleCreateRequest body request.ArticleCreateReq true "文章信息"
// @Success 200 {object} response.Status{data=service.Article} "ok"
// @Failure 401 {object} response.Status "Unauthorized"
// @Failure 422 {object} response.Status "Unprocessable Entity"
// @Failure 500 {object} response.Status "Internal Server Error"
// @Router /article [post]
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

	serviceArticle := value.(*service.Article)

	response.EndWithOK(c, serviceArticle)
}

// Delete godoc
// @Summary 删除文章
// @Description 删除文章
// @Tags Article
// @Accept  json
// @Produce  json
// @Param id path int true "文章ID"
// @Success 200 {object} response.Status{data=uint32} "ok"
// @Failure 401 {object} response.Status "Unauthorized"
// @Failure 422 {object} response.Status "Unprocessable Entity"'
// @Failure 500 {object} response.Status "Internal Server Error"
// @Router /article/{id} [delete]
func (a *ArticleController) Delete(c *gin.Context) {
	articleId, err := strconv.ParseUint(c.Param("id"), 10, 32)

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

// Index godoc
// @Summary 获取文章列表
// @Description 获取文章列表
// @Tags Article
// @Param pageNum query int true "页码"
// @Param pageSize query int true "每页数量"
// @Success 200 {object} response.Status{data=[]service.Article} "ok"
// @Failure 400 {object} response.Status "Bad Request"
// @Failure 401 {object} response.Status "Unauthorized"
// @Failure 422 {object} response.Status "Unprocessable Entity"
// @Failure 500 {object} response.Status "Internal Server Error"
// @Router /article [get]
func (a *ArticleController) Index(c *gin.Context) {

	pageNum, err := strconv.ParseInt(c.Query("pageNum"), 10, 32)
	if err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithBadRequest(c, nil)
		return
	}

	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 32)
	if err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithBadRequest(c, nil)
		return
	}

	req := &request.PaginationReq{
		PageNum:  int(pageNum),
		PageSize: int(pageSize),
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

// Update godoc
// @Summary 更新文章
// @Description 更新文章
// @Tags Article
// @Accept  json
// @Produce  json
// @Param id path int true "文章ID"
// @Param ArticleUpdateRequest body request.ArticleUpdateReq true "文章信息"
// @Success 200 {object} response.Status{data=service.Article} "ok"
// @Failure 401 {object} response.Status "Unauthorized"
// @Failure 422 {object} response.Status "Unprocessable Entity"
// @Failure 500 {object} response.Status "Internal Server Error"
// @Router /article/{id} [put]
func (a *ArticleController) Update(c *gin.Context) {
	articleId, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

// ShowArticleDetail godoc
// @Summary 获取文章详情
// @Description 获取文章详情
// @Tags Article
// @Param id path int true "文章ID"
// @Success 200 {object} response.Status{data=service.Article} "ok"
// @Failure 400 {object} response.Status "Bad Request"
// @Failure 401 {object} response.Status "Unauthorized"
// @Failure 422 {object} response.Status "Unprocessable Entity"
// @Failure 500 {object} response.Status "Internal Server Error"
// @Router /article/{id} [get]
func (a *ArticleController) ShowArticleDetail(c *gin.Context) {

	articleId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		logrus.Info("文章请求参数错误", err)
		response.EndWithUnprocessableData(c, nil)
		return
	}

	if !service.IsExistArticleById(uint32(articleId)) {
		logrus.Info("文章不存在")
		response.EndWithBadRequest(c, nil)
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

// ShowArticleByTime godoc
// @Summary 根据时间获取文章列表
// @Description 根据文章创建时间由新到旧获取文章列表
// @Tags Article
// @Param pageNum query int true "页码"
// @Param pageSize query int true "每页数量"
// @Success 200 {object} response.Status{data=[]service.Article} "ok"
// @Failure 400 {object} response.Status "Bad Request"
// @Failure 401 {object} response.Status "Unauthorized"
// @Failure 422 {object} response.Status "Unprocessable Entity"
// @Failure 500 {object} response.Status "Internal Server Error"
// @Router /article/time [get]
func (a *ArticleController) ShowArticleByTime(c *gin.Context) {

	pageNum, err := strconv.ParseInt(c.Query("pageNum"), 10, 32)
	if err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithBadRequest(c, nil)
		return
	}

	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 32)
	if err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithBadRequest(c, nil)
		return
	}

	req := &request.PaginationReq{
		PageNum:  int(pageNum),
		PageSize: int(pageSize),
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

// ShowArticleByTag godoc
// @Summary 根据标签获取文章列表
// @Description 根据标签获取文章列表
// @Tags Article
// @Param id path int true "标签的ID"
// @Param pageNum query int true "页码"
// @Param pageSize query int true "每页数量"
// @Success 200 {object} response.Status{data=[]service.Article} "ok"
// @Failure 400 {object} response.Status "Bad Request"
// @Failure 401 {object} response.Status "Unauthorized"
// @Failure 422 {object} response.Status "Unprocessable Entity"
// @Failure 500 {object} response.Status "Internal Server Error"
// @Router /article/tag/{id} [get]
func (a *ArticleController) ShowArticleByTag(c *gin.Context) {

	pageNum, err := strconv.ParseInt(c.Query("pageNum"), 10, 32)
	if err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithBadRequest(c, nil)
		return
	}

	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 32)
	if err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithBadRequest(c, nil)
		return
	}

	req := &request.PaginationReq{
		PageNum:  int(pageNum),
		PageSize: int(pageSize),
	}

	tagId, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

// ShowArticleByCategory godoc
// @Summary 根据分类获取文章列表
// @Description 根据分类获取文章列表
// @Tags Article
// @Param id path int true "分类的ID"
// @Param pageNum query int true "页码"
// @Param pageSize query int true "每页数量"
// @Success 200 {object} response.Status{data=[]service.Article} "ok"
// @Failure 400 {object} response.Status "Bad Request"
// @Failure 401 {object} response.Status "Unauthorized"
// @Failure 422 {object} response.Status "Unprocessable Entity"
// @Failure 500 {object} response.Status "Internal Server Error"
// @Router /article/category/{id} [get]
func (a *ArticleController) ShowArticleByCategory(c *gin.Context) {

	pageNum, err := strconv.ParseInt(c.Query("pageNum"), 10, 32)
	if err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithBadRequest(c, nil)
		return
	}

	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 32)
	if err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithBadRequest(c, nil)
		return
	}

	req := &request.PaginationReq{
		PageNum:  int(pageNum),
		PageSize: int(pageSize),
	}

	categoryId, err := strconv.ParseUint(c.Param("id"), 10, 32)
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

// Search godoc
// @Summary 搜索文章
// @Description 通过关键词搜索文章
// @Tags Article
// @Param keyword query string true "关键词"
// @Param pageNum query int true "页码"
// @Param pageSize query int true "每页数量"
// @Success 200 {object} response.Status{data=[]service.Article} "ok"
// @Failure 400 {object} response.Status "Bad Request"
// @Failure 401 {object} response.Status "Unauthorized"
// @Failure 422 {object} response.Status "Unprocessable Entity"
// @Failure 500 {object} response.Status "Internal Server Error"
// @Router /article/search [get]
func (a *ArticleController) Search(c *gin.Context) {

	pageNum, err := strconv.ParseInt(c.Query("pageNum"), 10, 32)
	if err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithBadRequest(c, nil)
		return
	}

	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 32)
	if err != nil {
		logrus.Info("分页请求参数错误", err)
		response.EndWithBadRequest(c, nil)
		return
	}

	req := &request.PaginationReq{
		PageNum:  int(pageNum),
		PageSize: int(pageSize),
	}

	keyword := c.Query("keyword")

	if keyword == "" {
		logrus.Info("关键词不能为空")
		response.EndWithBadRequest(c, nil)
		return
	}

	articles, ok := service.SearchArticle(keyword)
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

// ShowTags godoc
// @Summary 获取所有标签
// @Description 获取所有标签
// @Tags Article
// @Success 200 {object} response.Status{data=[]model.TagModel} "ok"
// @Failure 500 {object} response.Status "Internal Server Error"
// @Router /article/tag [get]
func (a *ArticleController) ShowTags(c *gin.Context) {

	tags, ok := service.GetTags()
	if !ok {
		logrus.Info("获取标签列表失败")
		response.EndWithInternalServerError(c, nil)
		return
	}

	response.EndWithOK(c, tags)
}

// ShowCategories godoc
// @Summary 获取所有分类
// @Description 获取所有分类
// @Tags Article
// @Success 200 {object} response.Status{data=[]model.CategoryModel} "ok"
// @Failure 500 {object} response.Status "Internal Server Error"
// @Router /category [get]
func (a *ArticleController) ShowCategories(c *gin.Context) {

	categories, ok := service.GetCategories()
	if !ok {
		logrus.Info("获取分类列表失败")
		response.EndWithInternalServerError(c, nil)
		return
	}

	response.EndWithOK(c, categories)
}
