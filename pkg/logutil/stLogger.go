package logutil

import (
	"fmt"
	"os"
	"time"
)

type stLogger struct {
	EnableLevel LogLevel
}

func newstLogger(level string) *stLogger {
	return &stLogger{
		EnableLevel: parseLogLevel(level),
	}
}

func (l *stLogger) enable(logLevel LogLevel) bool {
	return logLevel >= l.EnableLevel
}

func (l *stLogger) output(logLevel LogLevel, msg string) {
	if l.enable(logLevel) {
		now := time.Now()
		time := now.Format("2006-01-02 15:04:05")
		fmt.Fprintf(os.Stdout, "[%s][%v] %s\n", time, logLevel, msg)
	}
}

func (l *stLogger) Debug(msg string) {
	l.output(DEBUG, msg)
}

func (l *stLogger) INFO(msg string) {
	l.output(INFO, msg)
}

func (l *stLogger) WARNING(msg string) {
	l.output(WARNING, msg)
}

func (l *stLogger) ERROR(msg string) {
	l.output(ERROR, msg)
}

func (l *stLogger) FATAL(msg string) {
	l.output(FATAL, msg)
}
