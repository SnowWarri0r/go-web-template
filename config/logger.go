package config

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var log *logrus.Logger

func InitLogger() {
	log = logrus.New()
	if mode == "prod" {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(logrus.DebugLevel)
	}
	log.SetReportCaller(config.Conf.ReportCaller)
	log.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	file, err := os.OpenFile("./log/"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func Log() *logrus.Logger {
	return log
}
