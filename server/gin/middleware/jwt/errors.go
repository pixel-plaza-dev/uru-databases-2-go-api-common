package jwt

import "errors"

var (
	InvalidAuthorizationHeaderError    = errors.New("invalid authorization header")
	MissingRefreshTokenError           = errors.New("missing refresh token")
	MissingAccessTokenError            = errors.New("missing access token")
	UnableToCreateTokenValidationError = errors.New("unable to create token validation error")
)
