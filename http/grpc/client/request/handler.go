package request

import (
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonclienterror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/error"
)

type (
	// Handler interface
	Handler interface {
		HandleError(err error) error
	}

	// DefaultHandler struct
	DefaultHandler struct {
		mode *commonflag.ModeFlag
	}
)

// NewDefaultHandler creates a new default response handler
func NewDefaultHandler(mode *commonflag.ModeFlag) *DefaultHandler {
	return &DefaultHandler{mode: mode}
}

// HandleError handles the error
func (d DefaultHandler) HandleError(err error) error {
	return commonclienterror.ExtractErrorFromStatus(d.mode, err)
}
