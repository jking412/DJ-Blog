package core

import "github.com/gin-gonic/gin"

type IOSSCore interface {
	UploadFile(c *gin.Context)
	DownloadFile(c *gin.Context)
}
