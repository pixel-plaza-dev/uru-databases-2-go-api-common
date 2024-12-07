package request

import (
	"errors"
)

var (
	NilHandlerError = errors.New("request handler cannot be nil")
)
