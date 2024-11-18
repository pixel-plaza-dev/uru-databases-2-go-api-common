package context

import (
	"context"
	"github.com/gin-gonic/gin"
)

// PrepareCtx prepares the context for the gRPC request
func PrepareCtx(ctx *gin.Context, request *any) (*gin.Context, context.Context, *any, error) {
	// Bind the request
	if err := ctx.ShouldBindJSON(request); err != nil {
		return nil, nil, nil, err
	}

	// Get the outgoing context
	grpcCtx, err := GetOutgoingCtx(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	return ctx, grpcCtx, request, nil
}
