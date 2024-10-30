package error

type NoTokenInContextError struct{}

// Error returns the error message
func (e NoTokenInContextError) Error() string {
	return "No token in context"
}
