package database

import (
	"DJ-Blog/pkg/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dsn struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

var DB *gorm.DB

func Init() bool {
	d := dsn{
		Username: config.LoadString("database.username"),
		Password: config.LoadString("database.password"),
		Host:     config.LoadString("database.host"),
		Port:     config.LoadString("database.port"),
		DBName:   config.LoadString("database.dbname"),
	}
	var err error
	DB, err = gorm.Open(mysql.Open(d.String()), &gorm.Config{})
	if err != nil {
		logrus.Error("dsn ", d.String())
		return false
	}
	logrus.Info("Database connected")
	return true
}

func (d *dsn) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.DBName)
}
