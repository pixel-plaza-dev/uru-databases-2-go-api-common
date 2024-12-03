package types

type (
	// ErrorResponse struct
	ErrorResponse struct {
		Error string `json:"error"`
	}

	// BadRequest struct for bad request responses
	BadRequest struct {
		ErrorResponse
	}

	// InternalServerError struct for internal server error responses
	InternalServerError struct {
		ErrorResponse
	}
)

// NewBadRequest creates a new bad request response
func NewBadRequest(err error) *BadRequest {
	return &BadRequest{
		ErrorResponse: ErrorResponse{
			Error: err.Error(),
		},
	}
}

// NewInternalServerError creates a new internal server error response
func NewInternalServerError() *InternalServerError {
	return &InternalServerError{
		ErrorResponse: ErrorResponse{
			Error: "internal server error",
		},
	}
}
