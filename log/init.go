package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func SetUp() {
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)

	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:    false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		QuoteEmptyFields: true,
		ForceColors:      true,
	})
}

func Info(msg ...interface{}) {
	log.Info(msg...)
}

func Warn(msg ...interface{}) {
	log.Warn(msg...)
}

func Debug(msg ...interface{}) {
	log.Debug(msg...)
}

func Error(msg ...interface{}) {
	log.Error(msg...)
}

func Fatal(msg ...interface{}) {
	log.Fatal(msg...)
}
