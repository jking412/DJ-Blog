package database

import (
	"DJ-Blog/pkg/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type dsn struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

var DB *gorm.DB
var loadCount = 0

func init() {
	d := dsn{
		Username: config.LoadString("database.username"),
		Password: config.LoadString("database.password"),
		Host:     config.LoadString("database.host"),
		Port:     config.LoadString("database.port"),
		DBName:   config.LoadString("database.dbname"),
	}
	var err error
	DB, err = gorm.Open(mysql.Open(d.String()), &gorm.Config{})
	for {
		// 在docker部署中由于数据库的初始化需要一定的时间(即便配置了数据库比Go程序先运行)
		// 过快的连接会导致失败，所以可以多尝试几次
		if err != nil && loadCount < 5 {
			d.Host = "mysql"
			loadCount++
			time.Sleep(time.Second)
			DB, err = gorm.Open(mysql.Open(d.String()), &gorm.Config{})
		} else {
			break
		}
	}
	logrus.Info("Database connected")
}

func (d *dsn) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.Username,
		d.Password,
		d.Host,
		d.Port,
		d.DBName)
}
