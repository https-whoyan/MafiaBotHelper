package log

import (
	"os"
	"strings"
	"sync"

	"github.com/https-whoyan/MafiaCore/log"
)

type Logger = log.Logger

var appLogger Logger
var appLoggerOnce sync.Once

func SetLogger(logger Logger) {
	appLoggerOnce.Do(func() {
		appLogger = logger
	})
}

func GetLogger() Logger {
	return appLogger
}

func NewDefaultLogger() Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Ltime|log.Lshortfile)
}

func NewLoggerWithPredix(prefix string, fromParent bool) Logger {
	if appLogger == nil {
		panic("log not initialized")
	}
	if fromParent {
		prefix = appLogger.Prefix() + " " + prefix
	}
	prefix = strings.TrimSpace(prefix)
	return log.New(
		appLogger.Writer(),
		prefix,
		appLogger.Flags(),
	)
}
