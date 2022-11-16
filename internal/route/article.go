package route

import (
	"DJ-Blog/internal/core"
	"github.com/gin-gonic/gin"
)

func registerArticleRoutes(articleGroup *gin.RouterGroup, articleController core.IArticleController) {
	articleGroup.GET("/create", articleController.Create)
	articleGroup.GET("/update", articleController.Update)
	articleGroup.GET("/", articleController.Index)
	articleGroup.GET("/:id", articleController.ShowArticleDetail)
	articleGroup.GET("/time", articleController.ShowArticleByTime)
	articleGroup.GET("/:id/tag", articleController.ShowArticleByTag)
	articleGroup.GET("/:id/category", articleController.ShowByArticleCategory)

	articleGroup.POST("/", articleController.DoCreate)
	articleGroup.POST("/search", articleController.Search)

	articleGroup.DELETE("/:id", articleController.Delete)

	articleGroup.PUT("/:id", articleController.DoUpdate)
}

func registerTagRoutes(tagGroup *gin.RouterGroup, tagController core.ITagController) {
	tagGroup.GET("/", tagController.ShowTags)
}

func registerCategoryRoutes(categoryGroup *gin.RouterGroup, categoryController core.ICategoryController) {
	categoryGroup.GET("/", categoryController.ShowCategories)
	categoryGroup.GET("/:id", categoryController.ShowSpecificCategory)
}
