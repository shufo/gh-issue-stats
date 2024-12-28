package utils

import (
	"fmt"
	"os"

	"golang.org/x/exp/slog"
)

var logger *slog.Logger

func SetupLogger(debug bool) {
	opts := &slog.HandlerOptions{}
	if debug {
		opts.Level = slog.LevelDebug
	} else {
		opts.Level = slog.LevelInfo
	}
	handler := slog.NewJSONHandler(os.Stderr, opts)
	logger = slog.New(handler)
}

func DebugPrintf(format string, a ...interface{}) {
	if debug {
		logger.Debug(fmt.Sprintf(format, a...))

	}
}
