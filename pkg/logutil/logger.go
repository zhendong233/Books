package logutil

type LogLevel uint64

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var LevelMap = map[string]LogLevel{
	"DEBUG":   DEBUG,
	"INFO":    INFO,
	"WARNING": WARNING,
	"ERROR":   ERROR,
	"FATAL":   FATAL,
}

func parseLogLevel(msg string) LogLevel {
	logLevel, ok := LevelMap[msg]
	if !ok {
		return ERROR
	}
	return logLevel
}

type Logger interface {
	Debug(msg string)
	INFO(msg string)
	WARNING(msg string)
	ERROR(msg string)
	FATAL(msg string)
}

func NewLogger(level string) Logger {
	return newstLogger(level)
}
