package auth

import "errors"

var (
	MissingRESTMappingError         = errors.New("missing rest endpoint mapping")
	MissingGRPCMethodError          = errors.New("missing grpc method")
	InvalidAuthorizationHeaderError = errors.New("invalid authorization header")
	EmptyBaseUriError               = errors.New("empty base uri")
	RESTMapNilError                 = errors.New("rest map is nil")
	GRPCInterceptionsNilError       = errors.New(
		"grpc interceptions map is nil",
	)
)
