package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger is a wrapper around logrus.Logger
type Logger struct {
	log *logrus.Logger
}

// NewLogger creates a new logger instance with default settings
func NewLogger() *Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{}) // or TextFormatter{}

	return &Logger{
		log: log,
	}
}

// Info logs an info message
func (l *Logger) Info(message string) {
	l.log.Info(message)
}

// InfoWithFields logs an info message with fields
func (l *Logger) InfoWithFields(message string, fields map[string]interface{}) {
	if fields == nil {
		l.Info(message)
		return
	}
	l.log.WithFields(logrus.Fields(fields)).Info(message)
}

// Error logs an error message
func (l *Logger) Error(message string) {
	l.log.Error(message)
}

// ErrorWithFields logs an error message with fields
func (l *Logger) ErrorWithFields(message string, fields map[string]interface{}) {
	if fields == nil {
		l.Error(message)
		return
	}
	l.log.WithFields(logrus.Fields(fields)).Error(message)
}

// Warn logs a warning message
func (l *Logger) Warn(message string) {
	l.log.Warn(message)
}

// WarnWithFields logs a warning message with fields
func (l *Logger) WarnWithFields(message string, fields map[string]interface{}) {
	if fields == nil {
		l.Warn(message)
		return
	}
	l.log.WithFields(logrus.Fields(fields)).Warn(message)
}

// Debug logs a debug message
func (l *Logger) Debug(message string) {
	l.log.Debug(message)
}

// DebugWithFields logs a debug message with fields
func (l *Logger) DebugWithFields(message string, fields map[string]interface{}) {
	if fields == nil {
		l.Debug(message)
		return
	}
	l.log.WithFields(logrus.Fields(fields)).Debug(message)
}

// Custom method to set log level
func (l *Logger) SetLogLevel(level logrus.Level) {
	l.log.SetLevel(level)
}
