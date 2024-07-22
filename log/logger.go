package log

import (
	"context"
	"log"
	"os"
)

type Logger interface {
	Log(ctx context.Context, message string, args ...any)
}

type logger struct {
	log *log.Logger
}

func (l *logger) Log(ctx context.Context, message string, args ...any) {
	select {
	case <-ctx.Done():
		return
	default:
		l.log.Printf(message, args...)
	}
}

func NewLogger() Logger {
	return &logger{
		log: log.New(os.Stdout, "INFO\t", log.LstdFlags),
	}
}
