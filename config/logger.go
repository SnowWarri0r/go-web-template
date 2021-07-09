package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func InitLogger() {
	log = logrus.New()
	if mode == "prod" {
		log.SetLevel(logrus.WarnLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
	file, err := os.OpenFile("./log/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func Log() *logrus.Logger {
	return log
}
