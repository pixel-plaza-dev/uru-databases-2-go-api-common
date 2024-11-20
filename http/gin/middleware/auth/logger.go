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

// MethodNotSupported logs a MethodNotSupportedError
func (l Logger) MethodNotSupported(method string) {
	l.logger.LogMessageWithDetails("Method not supported", method)
}

// BaseUriIsLongerThanFullPath logs a BaseUriIsLongerThanFullPathError
func (l Logger) BaseUriIsLongerThanFullPath(fullPath string) {
	l.logger.LogMessageWithDetails(
		"Base URI is longer than full path",
		fullPath,
	)
}

// MissingRESTMapping logs a MissingRESTMappingError
func (l Logger) MissingRESTMapping(fullPath string) {
	l.logger.LogMessageWithDetails(fullPath, MissingRESTMappingError.Error())
}

// MissingGRPCMethod logs a MissingGRPCMethodError
func (l Logger) MissingGRPCMethod(fullPath string) {
	l.logger.LogMessageWithDetails(fullPath, MissingGRPCMethodError.Error())
}
