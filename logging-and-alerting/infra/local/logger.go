package local

import (
	"context"
	"log"

	laa "github.com/ryutah/sandbox/logging-and-alerting"
)

type Logger struct {
}

var _ laa.Logger = new(Logger)

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	log.Printf(format, v...)
}
