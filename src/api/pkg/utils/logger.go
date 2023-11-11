package utils

import "github.com/sirupsen/logrus"

func NewLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02/01/2006-15:04:05",
	})

	logrus.SetLevel(logrus.DebugLevel)
}
