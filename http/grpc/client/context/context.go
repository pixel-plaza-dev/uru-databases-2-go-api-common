package context

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	commonginctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/context"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	"google.golang.org/grpc/metadata"
)

// GetOutgoingCtx returns a context with the token string
func GetOutgoingCtx(ctx *gin.Context) (context.Context, error) {
	// Get the token string from the context
	token, err := commonginctx.GetCtxTokenString(ctx)
	if err != nil {
		// Check if the token is missing
		if errors.Is(err, commonginctx.MissingTokenInContextError) {
			return context.Background(), nil
		}
		return nil, err
	}

	// Append the token to the gRPC context
	grpcCtx := metadata.AppendToOutgoingContext(
		context.Background(),
		commongrpc.AuthorizationMetadataKey,
		token,
	)

	return grpcCtx, nil
}
