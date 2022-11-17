package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func init() {
	logrus.SetReportCaller(true)
	stdoutWriter := os.Stdout
	logrus.SetOutput(io.MultiWriter(stdoutWriter))
	logrus.Info("Logger initialized")
}
