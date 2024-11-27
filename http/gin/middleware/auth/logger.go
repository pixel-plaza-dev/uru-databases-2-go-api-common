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
	l.logger.LogMessage(commonlogger.NewLogMessage("Method not supported", commonlogger.StatusWarning, method))
}

// BaseUriIsLongerThanFullPath logs a BaseUriIsLongerThanFullPathError
func (l Logger) BaseUriIsLongerThanFullPath(fullPath string) {
	l.logger.LogMessage(commonlogger.NewLogMessage("Base URI is longer than full path", commonlogger.StatusWarning, fullPath))
}

// MissingRESTMapping logs a MissingRESTMappingError
func (l Logger) MissingRESTMapping(fullPath string) {
	l.logger.LogMessage(commonlogger.NewLogMessage("Missing REST endpoint mapping", commonlogger.StatusWarning, fullPath))
}

// MissingGRPCMethod logs a MissingGRPCMethodError
func (l Logger) MissingGRPCMethod(fullPath string) {
	l.logger.LogMessage(commonlogger.NewLogMessage("Missing gRPC method", commonlogger.StatusWarning, fullPath))
}
