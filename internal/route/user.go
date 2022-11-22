package route

import (
	"DJ-Blog/internal/controller"
	"DJ-Blog/internal/core"
	"DJ-Blog/internal/middleware"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(r *gin.Engine, userController core.IUserController) {

	userGroup := r.Group("/user")
	registerAvatarRoutes(userGroup, controller.NewUserController())

	userGroup.GET("/logout", userController.Logout)

	userGroup.POST("/login", userController.Login)
	userGroup.POST("/register", userController.Register)
}

func registerAvatarRoutes(userGroup *gin.RouterGroup, ossController core.IOSSController) {
	userGroup.GET("/avatar", middleware.Auth(), ossController.GetAvatarImg)
	userGroup.POST("/avatar", middleware.Auth(), ossController.UploadAvatarImg)
}
