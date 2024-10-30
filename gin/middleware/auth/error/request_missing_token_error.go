package error

import (
	commonjwt "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/jwt"
)

type RequestMissingTokenError struct {
	Token commonjwt.Token
}

// Error returns the error message
func (r RequestMissingTokenError) Error() string {
	return "Request missing " + r.Token.String() + "token"
}
