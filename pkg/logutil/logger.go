package logutil

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"

	"github.com/zhendong233/Books/pkg/envutil"
)

var Logger zerolog.Logger

func init() {
	Logger = newLogger()
}

func newLogger() zerolog.Logger {
	outputFile := envutil.GetBool("OUTPUT_LOG_FILE", false)
	if outputFile {
		now := time.Now().Format("2006-01-02")
		f, err := os.OpenFile(fmt.Sprintf("./log/log_%s.txt", now), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0o666)
		if err != nil {
			panic(err)
		}
		return zerolog.New(f).With().Timestamp().Logger()
	}
	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}
