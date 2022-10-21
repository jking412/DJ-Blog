package logger

import (
	"DJ-Blog/pkg/viperlib"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"strings"
)

func InitLogger() {
	logrus.SetReportCaller(true)
	logFile := strings.Split(viperlib.GetString("log.path"), "/")
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
	FileWriter, err := os.OpenFile(viperlib.GetString("log.path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Print("Failed to open log file")
	}
	logrus.SetOutput(io.MultiWriter(stdoutWriter, FileWriter))
	logrus.Info("Logger initialized")
}
