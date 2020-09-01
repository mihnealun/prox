package container

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

type LogDecorator interface {
	WithFields(map[string]interface{}) LoggerOperations
}
type LoggerOperations interface {
	Info(...interface{})
	Debug(...interface{})
	Error(...interface{})
	Print(...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(string, ...interface{})
}
type Logger interface {
	LoggerOperations
	LogDecorator
}

func newStdLogger(ctx context.Context, c *Config) (Logger, error) {
	file, err := os.OpenFile(c.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	var level logrus.Level
	level, err = logrus.ParseLevel(c.LogLevel)
	if err != nil {
		return nil, err
	}

	l := logrus.New()
	formatter := logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
			logrus.FieldKeyMsg:  "message",
		},
		PrettyPrint: false,
	}
	if c.Env == "development" {
		formatter.PrettyPrint = true
	}
	l.SetFormatter(&formatter)
	l.SetLevel(level)
	l.SetOutput(file)

	return &stdLogger{
		logger:  l,
		context: ctx,
	}, nil
}

type stdLogger struct {
	logger  *logrus.Logger
	context context.Context
}

func (l *stdLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *stdLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *stdLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *stdLogger) Print(args ...interface{}) {
	l.logger.Print(args...)
}

func (l *stdLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *stdLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *stdLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *stdLogger) WithFields(f map[string]interface{}) LoggerOperations {
	return l.logger.WithFields(f)
}
