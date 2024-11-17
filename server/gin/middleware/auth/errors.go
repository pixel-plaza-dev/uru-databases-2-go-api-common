package auth

import "errors"

var InvalidAuthorizationHeaderError = errors.New("invalid authorization header")
var MissingRefreshTokenError = errors.New("missing refresh token")
var MissingAccessTokenError = errors.New("missing access token")
var UnableToCreateTokenValidationError = errors.New("unable to create token validation error")
