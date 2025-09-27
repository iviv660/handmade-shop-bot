package logger

import (
	"github.com/rs/zerolog"
	"os"
)

type Logger struct {
	log zerolog.Logger
}

func New() *Logger {
	return &Logger{
		log: zerolog.New(os.Stdout).With().Timestamp().Logger(),
	}
}

func (l *Logger) Info(msg string, fields map[string]any) {
	l.log.Info().Fields(fields).Msg(msg)
}

func (l *Logger) Error(msg string, fields map[string]any) {
	l.log.Error().Fields(fields).Msg(msg)
}
