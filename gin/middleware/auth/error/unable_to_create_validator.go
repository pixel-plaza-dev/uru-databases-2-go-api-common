package error

type UnableToCreateValidatorError struct{}

// Error returns the error message
func (e UnableToCreateValidatorError) Error() string {
	return "Unable to create auth"
}
