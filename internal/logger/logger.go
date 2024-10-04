package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func New() *Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		NoColor:    true,
	}

	logger := zerolog.New(output).With().Timestamp().Caller().Logger()
	return &Logger{logger}
}

func (l *Logger) SetLevel(level int) {
	l.Logger = l.Logger.Level(zerolog.Level(level))
}
