package route

import (
	"DJ-Blog/internal/core"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(userGroup *gin.RouterGroup, userController core.IUserController) {
	userGroup.GET("/logout", userController.Logout)

	userGroup.POST("/login", userController.Login)
	userGroup.POST("/register", userController.Register)
}
