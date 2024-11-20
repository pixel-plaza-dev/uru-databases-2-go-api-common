package auth

import "errors"

var (
	MissingRESTMappingError         = errors.New("missing rest endpoint mapping")
	MissingGRPCMethodError          = errors.New("missing grpc method")
	InvalidAuthorizationHeaderError = errors.New("invalid authorization header")
)
