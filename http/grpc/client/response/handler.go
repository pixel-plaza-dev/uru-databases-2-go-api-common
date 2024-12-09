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
		HandleErrorResponse(ctx *gin.Context, err error)
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
func (d *DefaultHandler) HandlePrepareCtxError(ctx *gin.Context, err error) {
	if d.mode.IsProd() {
		err = commongin.InternalServerError
	}

	ctx.JSON(http.StatusInternalServerError, commongintypes.NewErrorResponse(err))
}

// HandleResponse handles the response from the gRPC server
func (d *DefaultHandler) HandleResponse(ctx *gin.Context, code int, response interface{}, err error) {
	// Check if the error is nil
	if err == nil {
		ctx.JSON(code, response)
		return
	}

	// Handle the error response
	d.HandleErrorResponse(ctx, err)
}

// HandleErrorResponse handles the error response from the gRPC server
func (d *DefaultHandler) HandleErrorResponse(ctx *gin.Context, err error) {
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
		if d.mode == nil || d.mode.IsProd() {
			ctx.JSON(http.StatusForbidden, commongintypes.NewErrorResponse(commongin.Unauthorized))
		}
		ctx.JSON(http.StatusForbidden, commongintypes.NewErrorResponse(extractedErr))
	case codes.Unauthenticated:
		if d.mode == nil || d.mode.IsProd() {
			ctx.JSON(http.StatusUnauthorized, commongintypes.NewErrorResponse(commongin.Unauthenticated))
		}
		ctx.JSON(http.StatusUnauthorized, commongintypes.NewErrorResponse(extractedErr))
	case codes.Unimplemented:
		if d.mode == nil || d.mode.IsProd() {
			ctx.JSON(http.StatusNotImplemented, commongintypes.NewErrorResponse(commongin.InDevelopment))
		}
		ctx.JSON(http.StatusNotImplemented, commongintypes.NewErrorResponse(extractedErr))
	case codes.Unavailable:
		if d.mode == nil || d.mode.IsProd() {
			ctx.JSON(http.StatusServiceUnavailable, commongintypes.NewErrorResponse(commongin.ServiceUnavailable))
		}
		ctx.JSON(http.StatusServiceUnavailable, commongintypes.NewErrorResponse(extractedErr))
	default:
		if d.mode == nil || d.mode.IsProd() {
			ctx.JSON(http.StatusInternalServerError, commongintypes.NewErrorResponse(commongin.InternalServerError))
		}
		ctx.JSON(http.StatusInternalServerError, commongintypes.NewErrorResponse(extractedErr))
	}
}
