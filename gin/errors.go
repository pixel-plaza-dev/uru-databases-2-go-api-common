package gin

import "errors"

var NoTokenInContextError = errors.New("no token in context")
var UnexpectedTokenTypeInContextError = errors.New("unexpected token type in context")
