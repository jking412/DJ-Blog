package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"strings"
)

func InitLogger() {
	logrus.SetReportCaller(true)
	logFile := strings.Split(config.GetString("log.path"), "/")
	if len(logFile) > 1 {
		_, err := os.Stat(logFile[0])
		if err != nil {
			err = os.Mkdir(logFile[0], os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	stdoutWriter := os.Stdout
	_, err := os.OpenFile(config.GetString("log.path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Print("Failed to open log file")
	}

	//logrus.SetOutput(io.MultiWriter(stdoutWriter, FileWriter))
	logrus.SetOutput(io.MultiWriter(stdoutWriter))
	logrus.Info("Logger initialized")
}
