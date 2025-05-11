package helper

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func NewLoggerConfigure(filePath string, logLevel logrus.Level, message string, level logrus.Level) {
	log := logrus.New()

	log.SetLevel(logLevel)

	file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(file)

	timeStr := time.Now().UTC()

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: timeStr.String(),
	})

	log.WithFields(logrus.Fields{
		"message": message,
	}).Log(level, message)
}
