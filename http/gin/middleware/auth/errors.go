package auth

import "errors"

var (
	InvalidAuthorizationHeaderError = errors.New("invalid authorization header")
)
