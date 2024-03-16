package misc

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"sync"
)

type Logger struct {
	logger *logrus.Logger
}

var logger *Logger
var loggerOnce sync.Once

func GetLogger() *Logger {
	loggerOnce.Do(func() {
		log := &logrus.Logger{
			Out:       os.Stdout,
			Formatter: &logrus.TextFormatter{FullTimestamp: true},
			Level:     logrus.InfoLevel,
			ExitFunc:  os.Exit,
		}
		logger = &Logger{logger: log}

	})
	return logger
}

func (l *Logger) Debug(message string) {
	l.logger.Debug(l.formatMessage(message))
}

func (l *Logger) Info(message string) {
	l.logger.Info(l.formatMessage(message))
}

func (l *Logger) Error(message string) {
	l.logger.Error(l.formatMessage(message))
}

func (l *Logger) Debugf(message string, args ...interface{}) {
	l.logger.Debugf(l.formatMessage(message), args)
}

func (l *Logger) Infof(message string, args ...interface{}) {
	l.logger.Infof(l.formatMessage(message), args)
}

func (l *Logger) Errorf(message string, args ...interface{}) {
	l.logger.Errorf(l.formatMessage(message), args)
}

func (l *Logger) formatMessage(message string) string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("| %s:%v | %s", file, line, message)
}
