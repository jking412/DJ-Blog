package route

import (
	"DJ-Blog/internal/core"
	"DJ-Blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

func registerArticleRoutes(r *gin.Engine, articleController core.IArticleController) {

	r.GET("/article", middleware.Auth(), articleController.Index)
	r.POST("/article", middleware.Auth(), articleController.Create)

	articleGroup := r.Group("/article")
	articleGroup.Use(middleware.Auth())

	articleGroup.GET("/:id", articleController.ShowArticleDetail)
	articleGroup.GET("/time", articleController.ShowArticleByTime)
	articleGroup.GET("/tag/:id", articleController.ShowArticleByTag)
	articleGroup.GET("/category/:id", articleController.ShowArticleByCategory)

	articleGroup.GET("/search", articleController.Search)

	articleGroup.DELETE("/:id", articleController.Delete)

	articleGroup.PUT("/:id", articleController.Update)
}

func registerTagRoutes(r *gin.Engine, tagController core.ITagController) {

	r.GET("/tag", middleware.Auth(), tagController.ShowTags)

}

func registerCategoryRoutes(r *gin.Engine, categoryController core.ICategoryController) {

	r.GET("/category", middleware.Auth(), categoryController.ShowCategories)

}
