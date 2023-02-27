package asynq

import (
	"github.com/hibiken/asynq"
	"go-micro.dev/v4/logger"
)

func newAsynqLogger(log logger.Logger) asynq.Logger {
	return &asynqLogger{
		log: logger.NewHelper(log),
	}
}

type asynqLogger struct {
	asynq.Logger
	log *logger.Helper
}

func (l *asynqLogger) Debug(args ...interface{}) {
	l.log.Debug(args)
}

func (l *asynqLogger) Info(args ...interface{}) {
	l.log.Info(args)
}

func (l *asynqLogger) Warn(args ...interface{}) {
	l.log.Warn(args)
}

func (l *asynqLogger) Error(args ...interface{}) {
	l.log.Error(args)
}

func (l *asynqLogger) Fatal(args ...interface{}) {
	l.log.Fatal(args)
}
