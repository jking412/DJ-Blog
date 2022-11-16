package boot

import (
	"DJ-Blog/pkg/config"
	"DJ-Blog/pkg/database"
	"DJ-Blog/pkg/logger"
	"DJ-Blog/pkg/session"
	"github.com/sirupsen/logrus"
)

func Initialize() {
	config.InitConfig("config")
	logger.InitLogger()
	if !database.InitDB() {
		logrus.Panic("Database init failed")
	}
	session.InitSession()
}
