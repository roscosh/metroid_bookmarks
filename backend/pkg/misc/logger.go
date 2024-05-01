package misc

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"sync"
)

type LogLevel string

const (
	Debug   LogLevel = "DEBUG"
	Info    LogLevel = "INFO"
	Warning LogLevel = "WARNING"
	Error   LogLevel = "ERROR"
	Fatal   LogLevel = "FATAL"
)

type Logger struct {
	logger *logrus.Logger
}

var logger *Logger
var loggerOnce sync.Once

func (l *Logger) SetParams(logLevel string) {
	var level logrus.Level

	switch LogLevel(logLevel) {
	case Debug:
		level = logrus.DebugLevel
	case Info:
		level = logrus.InfoLevel
	case Warning:
		level = logrus.WarnLevel
	case Error:
		level = logrus.ErrorLevel
	case Fatal:
		level = logrus.FatalLevel
	default:
		level = logrus.InfoLevel
	}
	l.logger = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{FullTimestamp: true},
		Level:     level,
		ExitFunc:  os.Exit,
	}
}

func GetLogger() *Logger {
	loggerOnce.Do(func() {
		log := logrus.New()
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

func (l *Logger) Warn(message string) {
	l.logger.Warn(l.formatMessage(message))
}

func (l *Logger) Error(message string) {
	l.logger.Error(l.formatMessage(message))
}

func (l *Logger) Fatal(message string) {
	l.logger.Fatal(l.formatMessage(message))
}

func (l *Logger) Debugf(message string, args ...interface{}) {
	l.logger.Debugf(l.formatMessage(message), args)
}

func (l *Logger) Infof(message string, args ...interface{}) {
	l.logger.Infof(l.formatMessage(message), args)
}

func (l *Logger) Warnf(message string, args ...interface{}) {
	l.logger.Warnf(l.formatMessage(message), args)
}

func (l *Logger) Errorf(message string, args ...interface{}) {
	l.logger.Errorf(l.formatMessage(message), args)
}

func (l *Logger) Fatalf(message string, args ...interface{}) {
	l.logger.Fatalf(l.formatMessage(message), args)
}

func (l *Logger) formatMessage(message string) string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%v | %s", file, line, message)
}
