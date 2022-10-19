package boot

import (
	"DJ-Blog/model"
	"DJ-Blog/pkg/database"
	"DJ-Blog/pkg/logger"
	"DJ-Blog/pkg/markdown"
	"DJ-Blog/pkg/viperlib"
	"github.com/sirupsen/logrus"
)

func Initialize() {
	viperlib.InitConfig()
	logger.InitLogger()
	database.InitDB()

	err := database.DB.AutoMigrate(&model.Post{})
	if err != nil {
		logrus.Error("Database migration failed")
		panic(err)
	} else {
		logrus.Info("Database migration success")
	}

	markdown.InitMarkDownParser()
}
