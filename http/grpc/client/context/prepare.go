package context

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	commonclientrequest "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/request"
	"io"
)

// PrepareCtx prepares the context for the gRPC request
func PrepareCtx(ctx *gin.Context, request interface{}, handler commonclientrequest.Handler) (
	grpcCtx context.Context,
	err error,
) {
	// Bind the request
	if request != nil {
		err = ctx.ShouldBindJSON(request)
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, handler.HandlePrepareCtxError(err)
		}
	}

	// Get the outgoing context
	grpcCtx, err = GetOutgoingCtx(ctx)
	if err != nil {
		return nil,
			handler.HandlePrepareCtxError(err)
	}

	return grpcCtx, nil
}
