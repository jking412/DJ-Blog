package main

import (
	"DJ-Blog/api"
	"DJ-Blog/boot"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	boot.Initialize()

	r := gin.Default()

	api.Register(r)

	if err := r.Run(":" + config.GetString("server.port")); err != nil {
		logrus.Error("Server start failed")
		panic(err)
	}
}
