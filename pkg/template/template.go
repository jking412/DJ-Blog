package template

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"html/template"
)

func InitTemplate(r *gin.Engine, pattern ...string) {
	for _, p := range pattern {
		r.LoadHTMLGlob(p)
	}
	r.SetFuncMap(template.FuncMap{})
	logrus.Info("Template loaded")
}
