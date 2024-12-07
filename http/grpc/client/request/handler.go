package request

import (
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonclienterror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/error"
)

type (
	// Handler interface
	Handler interface {
		HandlePrepareCtxError(err error) error
		HandleStatusError(err error) error
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

// HandlePrepareCtxError handles the prepare context error
func (d DefaultHandler) HandlePrepareCtxError(err error) error {
	if d.mode.IsDev() {
		return err
	}
	return commongin.InternalServerError
}

// HandleStatusError handles the status error
func (d DefaultHandler) HandleStatusError(err error) error {
	return commonclienterror.ExtractErrorFromStatus(d.mode, err)
}
