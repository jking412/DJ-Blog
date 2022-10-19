package template

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitTemplate(r *gin.Engine, pattern ...string) {
	for _, p := range pattern {
		r.LoadHTMLGlob(p)
	}
	logrus.Info("Template loaded")
}
