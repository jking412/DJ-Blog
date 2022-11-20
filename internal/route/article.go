package route

import (
	"DJ-Blog/internal/core"
	"github.com/gin-gonic/gin"
)

func registerArticleRoutes(r *gin.Engine, articleController core.IArticleController) {

	r.GET("/article", articleController.Index)
	r.POST("/article", articleController.Create)

	articleGroup := r.Group("/article")

	articleGroup.GET("/:id", articleController.ShowArticleDetail)
	articleGroup.GET("/time", articleController.ShowArticleByTime)
	articleGroup.GET("/tag/:id", articleController.ShowArticleByTag)
	articleGroup.GET("/category/:id", articleController.ShowByArticleCategory)

	articleGroup.POST("/search", articleController.Search)

	articleGroup.DELETE("/:id", articleController.Delete)

	articleGroup.PUT("/:id", articleController.Update)
}

func registerTagRoutes(r *gin.Engine, tagController core.ITagController) {

	r.GET("/tag", tagController.ShowTags)

}

func registerCategoryRoutes(r *gin.Engine, categoryController core.ICategoryController) {
	r.GET("/category", categoryController.ShowCategories)

	categoryGroup := r.Group("/category")

	categoryGroup.GET("/:id", categoryController.ShowSpecificCategory)
}
