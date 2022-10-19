package main

import (
	"DJ-Blog/api"
	"DJ-Blog/boot"
	"DJ-Blog/pkg/template"
	"DJ-Blog/pkg/viperlib"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	boot.Initialize()

	r := gin.Default()

	template.InitTemplate(r, viperlib.GetStrings("template.path")...)

	api.Register(r)

	if err := r.Run(":" + viperlib.GetString("server.port")); err != nil {
		logrus.Error("Server start failed")
		panic(err)
	}
}
