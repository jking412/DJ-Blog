package route

import (
	"DJ-Blog/internal/core"
	"github.com/gin-gonic/gin"
)

func registerUserRoutes(userGroup *gin.RouterGroup, userController core.IUserController) {
	userGroup.GET("/login", userController.Login)
	userGroup.GET("/register", userController.Register)
	userGroup.GET("/logout", userController.Logout)

	userGroup.POST("/login", userController.DoLogin)
	userGroup.POST("/register", userController.DoRegister)
}
