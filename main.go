package main

import (
	"DJ-Blog/boot"
	"DJ-Blog/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	boot.Initialize()

	r := gin.Default()

	if err := r.Run(":" + config.LoadString("server.port")); err != nil {
		logrus.Panic("Server start failed")
	}
}
