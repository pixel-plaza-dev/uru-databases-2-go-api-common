package auth

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

// Logger is the logger for the auth middleware
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger is the logger for the auth middleware
func NewLogger(logger commonlogger.Logger) Logger {
	return Logger{logger: logger}
}

// MissingRESTMapping logs a MissingRESTMappingError
func (l Logger) MissingRESTMapping(fullPath string) {
	l.logger.LogMessageWithDetails(fullPath, MissingRESTMappingError.Error())
}

// MissingGRPCMethod logs a MissingGRPCMethodError
func (l Logger) MissingGRPCMethod(fullPath string) {
	l.logger.LogMessageWithDetails(fullPath, MissingGRPCMethodError.Error())
}
