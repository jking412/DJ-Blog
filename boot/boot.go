package boot

import (
	"DJ-Blog/model"
	"DJ-Blog/pkg/database"
	"DJ-Blog/pkg/logger"
	"DJ-Blog/pkg/markdown"
	"DJ-Blog/pkg/search"
	"DJ-Blog/pkg/sessionpkg"
	"DJ-Blog/pkg/viperlib"
	"github.com/sirupsen/logrus"
	"time"
)

func Initialize() {
	viperlib.InitConfig()
	logger.InitLogger()

	for {
		err := database.InitDB()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	err := database.DB.AutoMigrate(&model.User{},
		&model.Post{})
	if err != nil {
		logrus.Error("Database migration failed")
		panic(err)
	} else {
		logrus.Info("Database migration success")
	}

	sessionpkg.InitSession()
	search.InitZincClient()

	createIndex()

	markdown.InitMarkDownParser()
}
