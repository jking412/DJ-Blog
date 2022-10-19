package boot

import (
	"DJ-Blog/pkg/database"
	"DJ-Blog/pkg/logger"
	"DJ-Blog/pkg/viperlib"
	"github.com/sirupsen/logrus"
)

func Initialize() {
	viperlib.InitConfig()
	logger.InitLogger()
	database.InitDB()

	err := database.DB.AutoMigrate()
	if err != nil {
		logrus.Error("Database migration failed")
		panic(err)
	} else {
		logrus.Info("Database migration success")
	}
}
