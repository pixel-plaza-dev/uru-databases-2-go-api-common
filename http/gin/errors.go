package gin

import (
	"errors"
)

var (
	InternalServerError = errors.New("internal server error")
	ServiceUnavailable  = errors.New("service unavailable")
	Unauthenticated     = errors.New("missing or invalid bearer token on authentication header")
	Unauthorized        = errors.New("not authorized to access the resource")
	InDevelopment       = errors.New("in development")
)
