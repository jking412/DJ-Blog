package route

import (
	"DJ-Blog/internal/core"
	"github.com/gin-gonic/gin"
)

func registerGithubRoutes(githubGroup *gin.RouterGroup, githubController core.IGithubUserController) {
	githubGroup.GET("/login", githubController.GithubLogin)
	githubGroup.GET("/oauth2", githubController.GithubLoginCallback)
}
