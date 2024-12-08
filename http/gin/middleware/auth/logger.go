package auth

import (
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
)

// Logger is the logger for the auth middleware
type Logger struct {
	logger commonlogger.Logger
}

// NewLogger is the logger for the auth middleware
func NewLogger(logger commonlogger.Logger) (*Logger, error) {
	// Check if the logger is nil
	if logger == nil {
		return nil, commonlogger.NilLoggerError
	}

	return &Logger{logger: logger}, nil
}

// MethodNotSupported logs that the method is not supported
func (l *Logger) MethodNotSupported(method string) {
	l.logger.LogMessage(commonlogger.NewLogMessage("Method not supported", commonlogger.StatusWarning, method))
}

// BaseUriIsLongerThanFullPath logs that the base URI is longer than the full path
func (l *Logger) BaseUriIsLongerThanFullPath(fullPath string) {
	l.logger.LogMessage(
		commonlogger.NewLogMessage(
			"Base URI is longer than full path",
			commonlogger.StatusWarning,
			fullPath,
		),
	)
}

// FailedToMapRESTEndpoint logs that the REST endpoint could not be mapped
func (l *Logger) FailedToMapRESTEndpoint(err error) {
	l.logger.LogError(commonlogger.NewLogError("Failed to map REST endpoint", err))
}

// MissingGRPCMethod logs a MissingGRPCMethodError
func (l *Logger) MissingGRPCMethod(fullPath string) {
	l.logger.LogMessage(commonlogger.NewLogMessage("Missing gRPC method", commonlogger.StatusWarning, fullPath))
}

// MissingMapper logs a MissingMapperError
func (l *Logger) MissingMapper() {
	l.logger.LogError(commonlogger.NewLogError("Missing mapper", NilMapperError))
}

// MissingGRPCInterceptions logs a MissingGRPCInterceptionsError
func (l *Logger) MissingGRPCInterceptions() {
	l.logger.LogError(commonlogger.NewLogError("Missing gRPC interceptions", NilGRPCInterceptionsError))
}
