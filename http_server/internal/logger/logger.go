package logger

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"todo-list/internal/utils"
)

func SetupLogger() (*slog.Logger, *os.File) {
	var logFilename string

	if utils.IsProd() {
		logFilename = "logProd.log"
	} else {
		logFilename = "logDev.log"
	}

	logDir := "./logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0755) // Create the directory with permission 0755
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}
	logFilePath := filepath.Join(logDir, logFilename)

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	opts := &slog.HandlerOptions{
		ReplaceAttr: ReplaceAttr,
	}
	var handler slog.Handler
	if utils.IsProd() {
		handler = slog.NewJSONHandler(file, opts)
	} else {
		handler = slog.NewTextHandler(file, opts)
	}
	handler = &ContextHandler{handler}
	logger := slog.New(handler)

	logger.Info("bla-bla-bla_123")

	return logger, file
}
