package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

// init initialize a logger
func init() {
	log = logrus.New()

	// set the logger level
	log.SetLevel(logrus.InfoLevel)

	// set the log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		Fatal("failed to set up log file: ", err)
	}
}

// Fatal is a wrapper of the logrus fatal method
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Warn is a wrapper of the logrus warn method
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Info is a wrapper of the logrus info method
func Info(args ...interface{}) {
	log.Info(args...)
}
