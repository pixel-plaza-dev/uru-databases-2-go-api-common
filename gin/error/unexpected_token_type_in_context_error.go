package error

type UnexpectedTokenTypeInContextError struct{}

// Error returns the error message
func (e UnexpectedTokenTypeInContextError) Error() string {
	return "Unexpected token type in context"
}
