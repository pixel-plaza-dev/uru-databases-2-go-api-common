package metadata

import (
	"github.com/gin-gonic/gin"
	"github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin/middleware"
	commongrpcctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc/client/context"
	"google.golang.org/grpc/credentials/oauth"
)

type (
	// Authentication interface
	Authentication interface {
		Authenticate() gin.HandlerFunc
	}

	// Middleware struct
	Middleware struct {
		accessToken string
	}
)

// NewMiddleware creates a new authentication middleware
func NewMiddleware(tokenSource *oauth.TokenSource) (*Middleware, error) {
	// Get the access token from the token source
	token, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}

	return &Middleware{
		accessToken: token.AccessToken,
	}, nil
}

// Authenticate returns a Gin middleware that sets the authentication metadata to the gRPC request
func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the JWT token
		jwtToken, err := middleware.GetToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(401, gin.H{"error": MissingTokenError})
			return
		}

		// Create the context metadata
		ctxMetadata := commongrpcctx.NewAuthenticatedCtxMetadata(m.accessToken, jwtToken)

		// Get the gRPC client context
		grpcCtx := commongrpcctx.GetCtxWithMetadata(ctxMetadata, ctx)

		// Set the gRPC client context to the Gin context
		for _, metadataField := range ctxMetadata.MetadataFields {
			ctx.Set(metadataField.Key, grpcCtx.Value(metadataField.Key))
		}

		// Continue
		ctx.Next()
	}
}
