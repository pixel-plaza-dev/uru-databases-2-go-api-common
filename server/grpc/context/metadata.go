package context

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	jwtmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin/middleware/jwt"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc"
	clientauthinterceptor "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc/client/interceptor/auth"
	"google.golang.org/grpc/metadata"
)

// GetOutgoingCtx returns a context with the token string
func GetOutgoingCtx(ctx *gin.Context) (context.Context, error) {
	// Get the token string from the context
	token, err := jwtmiddleware.GetCtxTokenString(ctx)
	if err != nil {
		// Check if the token is missing
		if errors.Is(err, clientauthinterceptor.MissingTokenError) {
			return context.Background(), nil
		}
		return nil, err
	}

	// Append the token to the gRPC context
	grpcCtx := metadata.AppendToOutgoingContext(context.Background(), commongrpc.AuthorizationMetadataKey, token)

	return grpcCtx, nil
}
