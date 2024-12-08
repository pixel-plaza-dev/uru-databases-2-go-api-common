package auth

import "errors"

var (
	InvalidAuthorizationHeaderError = errors.New("invalid authorization header")
	NilMapperError                  = errors.New("mapper cannot be nil")
	NilGRPCInterceptionsError       = errors.New("grpc interceptions cannot be nil")
)
