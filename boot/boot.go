package boot

import (
	"DJ-Blog/pkg/config"
	"DJ-Blog/pkg/database"
	"DJ-Blog/pkg/logger"
	"DJ-Blog/pkg/session"
	"github.com/sirupsen/logrus"
)

func Initialize() {
	config.Init("config")
	logger.Init()
	if !database.Init() {
		logrus.Panic("Database init failed")
	}
	session.Init()
}
