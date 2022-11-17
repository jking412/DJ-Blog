package controller

import "github.com/gin-gonic/gin"

type ArticleController struct {
}

func NewArticleController() *ArticleController {
	return new(ArticleController)
}

func (a *ArticleController) Create(c *gin.Context) {

}

func (a *ArticleController) Delete(c *gin.Context) {

}

func (a *ArticleController) Index(c *gin.Context) {

}

func (a *ArticleController) Update(c *gin.Context) {

}

func (a *ArticleController) ShowArticleDetail(c *gin.Context) {

}

func (a *ArticleController) ShowArticleByTime(c *gin.Context) {

}

func (a *ArticleController) ShowArticleByTag(c *gin.Context) {

}

func (a *ArticleController) ShowByArticleCategory(c *gin.Context) {

}

func (a *ArticleController) Search(c *gin.Context) {

}

func (a *ArticleController) ShowTags(c *gin.Context) {

}

func (a *ArticleController) ShowCategories(c *gin.Context) {

}

func (a *ArticleController) ShowSpecificCategory(c *gin.Context) {

}
