package route

import (
	"DJ-Blog/internal/core"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(r *gin.Engine, userController core.IUserController) {

	userGroup := r.Group("/user")

	userGroup.GET("/logout", userController.Logout)

	userGroup.POST("/login", userController.Login)
	userGroup.POST("/register", userController.Register)
}
