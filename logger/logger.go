package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

// Init initializes the global logger
func Init() {
	var cfg zap.Config
	cfg = zap.NewDevelopmentConfig()
	logger, err := cfg.Build()
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	log = logger
}

// Info logs an info message
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

// Error logs an error message
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

// Debug logs a debug message
func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

// Sync flushes any buffered log entries
func Sync() {
	_ = log.Sync()
}
