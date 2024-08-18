package config

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

func Logger() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
	}
	return logger
}
