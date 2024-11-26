package response

import (
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/context"
)

type (
	// Handler interface
	Handler interface {
		Handle(response interface{}, err error) (interface{}, error)
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

// Handle handles the response
func (d DefaultHandler) Handle(response interface{}, err error) (interface{}, error) {
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(d.mode, err)
	}
	return response, nil
}
