package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
)

func InitLogger() {
	logrus.SetReportCaller(true)
	stdoutWriter := os.Stdout
	FileWriter, err := os.OpenFile(viper.GetString("log.path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Print("Failed to open log file")
	}
	logrus.SetOutput(io.MultiWriter(stdoutWriter, FileWriter))
	logrus.Info("Logger initialized")
}
