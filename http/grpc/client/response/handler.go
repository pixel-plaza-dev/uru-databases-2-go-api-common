package response

import (
	"github.com/gin-gonic/gin"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
	commongintypes "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonclientstatus "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/status"
	"google.golang.org/grpc/codes"
	"net/http"
)

type (
	// Handler interface
	Handler interface {
		HandlePrepareCtxError(ctx *gin.Context, err error)
		HandleResponse(ctx *gin.Context, code int, response interface{}, err error)
	}

	// DefaultHandler struct
	DefaultHandler struct {
		mode *commonflag.ModeFlag
	}
)

// NewDefaultHandler creates a new default request handler
func NewDefaultHandler(mode *commonflag.ModeFlag) (*DefaultHandler, error) {
	// Check if the flag mode is nil
	if mode == nil {
		return nil, commonflag.NilModeFlagError
	}
	return &DefaultHandler{mode: mode}, nil
}

// HandlePrepareCtxError handles the prepare context error
func (d DefaultHandler) HandlePrepareCtxError(ctx *gin.Context, err error) {
	if d.mode.IsProd() {
		err = commongin.InternalServerError
	}

	ctx.JSON(http.StatusInternalServerError, commongintypes.NewErrorResponse(err))
}

// HandleResponse handles the response from the gRPC server
func (d DefaultHandler) HandleResponse(ctx *gin.Context, code int, response interface{}, err error) {
	// Check if the error is nil
	if err == nil {
		ctx.JSON(code, response)
		return
	}

	// Extract the gRPC code and error from the status
	extractedCode, extractedErr := commonclientstatus.ExtractErrorFromStatus(d.mode, err)

	// Check the extracted code and error
	switch extractedCode {
	case codes.AlreadyExists:
		ctx.JSON(http.StatusConflict, commongintypes.NewErrorResponse(extractedErr))
	case codes.NotFound:
		ctx.JSON(http.StatusNotFound, commongintypes.NewErrorResponse(extractedErr))
	case codes.InvalidArgument:
		ctx.JSON(http.StatusBadRequest, commongintypes.NewErrorResponse(extractedErr))
	case codes.PermissionDenied:
		ctx.JSON(http.StatusForbidden, commongintypes.NewErrorResponse(extractedErr))
	case codes.Unauthenticated:
		ctx.JSON(http.StatusUnauthorized, commongintypes.NewErrorResponse(extractedErr))
	case codes.Unimplemented:
		ctx.JSON(http.StatusNotImplemented, commongintypes.NewErrorResponse(extractedErr))
	case codes.Unavailable:
		ctx.JSON(http.StatusServiceUnavailable, commongintypes.NewErrorResponse(extractedErr))
	default:
		ctx.JSON(http.StatusInternalServerError, commongintypes.NewErrorResponse(extractedErr))
	}
}
