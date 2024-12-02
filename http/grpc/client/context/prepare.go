package context

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
)

// PrepareCtx prepares the context for the gRPC request
func PrepareCtx(ctx *gin.Context, request interface{}) (
	grpcCtx context.Context,
	err error,
) {
	// Bind the request
	if request != nil {
		err = ctx.ShouldBindJSON(request)
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}
	}

	// Get the outgoing context
	grpcCtx, err = GetOutgoingCtx(ctx)
	if err != nil {
		return nil, err
	}

	return grpcCtx, nil
}
