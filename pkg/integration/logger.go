package integration

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger zerolog.Logger

// InitLogger returns a Logger
func InitLogger() *Logger {
	newZerolog := zerolog.New(os.Stderr)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if os.Getenv("DEBUG") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	zerolog.GlobalLevel()
	logger := Logger(newZerolog)

	return &logger
}

// GetLevel returns the current Level.
func (t *Logger) getGlobalLevel() zerolog.Level {
	return zerolog.GlobalLevel()
}

// Debug sends a log event using debug level.
func (t *Logger) Debug(args ...interface{}) {
	log.Debug().Msg(fmt.Sprint(args...))
}

// Info sends a log event using info level.
func (t *Logger) Info(args ...interface{}) {
	log.Info().Msg(fmt.Sprint(args...))
}

// Error sends a log event using error level.
func (t *Logger) Error(err error, args ...interface{}) {
	log.Error().Err(err).Msg(fmt.Sprint(args...))
}
