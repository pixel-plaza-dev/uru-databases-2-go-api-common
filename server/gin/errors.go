package gin

import "errors"

var (
	NoTokenInContextError                   = errors.New("no token in context")
	UnexpectedTokenTypeInContextError       = errors.New("unexpected token type in context")
	UnexpectedTokenClaimsTypeInContextError = errors.New("unexpected token claims type in context")
)
