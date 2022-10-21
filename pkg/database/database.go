package database

import (
	"DJ-Blog/pkg/viperlib"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	databaseType := viperlib.GetString("database.type")
	switch databaseType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			viperlib.GetString("database.mysql.username"),
			viperlib.GetString("database.mysql.password"),
			//viperlib.GetString("database.mysql.container_name"),
			viperlib.GetString("database.mysql.host"),
			viperlib.GetString("database.mysql.port"),
			viperlib.GetString("database.mysql.dbname"))
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			logrus.Error("Database connection failed. Database name: " + viperlib.GetString("database.mysql.dbname"))
			logrus.Error("dsn " + dsn)
			return err
		}
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(viperlib.GetString("database.sqlite.dbname")), &gorm.Config{})
		if err != nil {
			logrus.Error("Database connection failed. Database path: " + viperlib.GetString("database.sqlite.path"))
			return err
		}
	}
	logrus.Info("Database connected")
	return nil
}
