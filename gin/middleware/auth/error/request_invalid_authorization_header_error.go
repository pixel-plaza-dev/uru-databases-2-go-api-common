package error

type RequestInvalidAuthorizationHeaderError struct{}

// Error returns the error message
func (e RequestInvalidAuthorizationHeaderError) Error() string {
	return "Request missing or invalid authorization header"
}
