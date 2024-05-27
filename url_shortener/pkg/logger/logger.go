package logger

import (
	"log/slog"
	"os"
)

// InitLogger initializes a new logger instance based on the provided
// mode string. It returns a pointer to the initialized slog.Logger.
//
// If mode is "dev", it configures the logger for debug level logging
// to stdout.
//
// If mode is "production", it configures the logger for info level
// logging to stdout.
func InitLogger(mode string) *slog.Logger {
	var newLogger *slog.Logger

	switch mode {
	case "dev":
		newLogger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "production":
		newLogger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return newLogger
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
