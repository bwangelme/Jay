package logger

import (
	"github.com/sirupsen/logrus"
)

var (
	logger    = logrus.New()
	errLogger = logrus.New()
)

func setLogger(log *logrus.Logger) {
	log.SetFormatter(&logrus.JSONFormatter{})
}

func InitLogger() {
	setLogger(logger)
	setLogger(logger)
	logger.Infof("Init Logger Success")
}

func WithField(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warningf(format string, args ...interface{}) {
	errLogger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	errLogger.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	errLogger.Panicf(format, args...)
}
