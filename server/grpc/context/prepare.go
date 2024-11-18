package context

import (
	"context"
	"github.com/gin-gonic/gin"
)

// PrepareCtx prepares the context for the gRPC request
func PrepareCtx(ctx *gin.Context, request *any) (context.Context, error) {
	// Bind the request
	if err := ctx.ShouldBindJSON(request); err != nil {
		return nil, err
	}

	// Get the outgoing context
	grpcCtx, err := GetOutgoingCtx(ctx)
	if err != nil {
		return nil, err
	}

	return grpcCtx, nil
}
