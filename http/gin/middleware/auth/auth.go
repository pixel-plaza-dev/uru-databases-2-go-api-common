package auth

import (
	"github.com/gin-gonic/gin"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
	commonginctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/context"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	commongrpc "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc"
	pbtypes "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/details/types"
	"strings"
)

type (
	// Authentication interface
	Authentication interface {
		Authenticate() gin.HandlerFunc
	}

	// Middleware struct
	Middleware struct {
		baseUri   string
		validator commonjwtvalidator.Validator
		restMap   *map[string]map[pbtypes.
				RESTMethod]pbtypes.GRPCMethod
		grpcInterceptions *map[pbtypes.GRPCMethod]pbtypes.Interception
		logger            Logger
	}
)

// NewMiddleware creates a new authentication middleware
func NewMiddleware(
	baseUri string,
	validator commonjwtvalidator.Validator,
	restMap *map[string]map[pbtypes.RESTMethod]pbtypes.
		GRPCMethod,
	grpcInterceptions *map[pbtypes.GRPCMethod]pbtypes.Interception,
	logger Logger,
) (*Middleware, error) {
	// Check if the base URI is empty
	if baseUri == "" {
		return nil, EmptyBaseUriError
	}

	// Check if the map is empty
	if restMap == nil {
		return nil, RESTMapNilError
	}

	// Check if the gRPC interceptions map is empty
	if grpcInterceptions == nil {
		return nil, GRPCInterceptionsNilError
	}

	return &Middleware{
		baseUri:           baseUri,
		validator:         validator,
		restMap:           restMap,
		grpcInterceptions: grpcInterceptions,
		logger:            logger,
	}, nil
}

// Authenticate return the middleware function that authenticates the request
func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the full endpoint and method
		fullRestEndpoint := ctx.FullPath()
		method := ctx.Request.Method
		restMethod := pbtypes.GetRESTMethod(method)

		// Check if the method is not supported
		if restMethod == pbtypes.INVALID {
			m.logger.MethodNotSupported(fullRestEndpoint)
			ctx.JSON(
				500,
				gin.H{"error": commongrpc.InternalServerError.Error()},
			)
			ctx.Abort()
			return
		}

		// Check if the base URI is longer than the full path
		if len(m.baseUri) > len(fullRestEndpoint) {
			m.logger.BaseUriIsLongerThanFullPath(fullRestEndpoint)
			ctx.JSON(500, gin.H{"error": commongin.InternalServerError.Error()})
			ctx.Abort()
			return
		}

		// Remove the base URI from the full REST endpoint
		restEndpoint := fullRestEndpoint[len(m.baseUri):]

		// Get the gRPC method
		grpcMethod, ok := (*m.restMap)[restEndpoint][restMethod]
		if !ok {
			m.logger.MissingRESTMapping(fullRestEndpoint)
			ctx.JSON(
				500,
				gin.H{"error": commongin.InternalServerError.Error()},
			)
			ctx.Abort()
			return
		}

		// Get the gRPC method interception
		interception, ok := (*m.grpcInterceptions)[grpcMethod]
		if !ok {
			m.logger.MissingGRPCMethod(fullRestEndpoint)
			ctx.JSON(500, gin.H{"error": commongin.InternalServerError.Error()})
			ctx.Abort()
			return
		}

		// Check if there is None interception
		if interception == pbtypes.None {
			ctx.Next()
			return
		}

		// Get the authorization from the header
		authorization := ctx.GetHeader(commongin.AuthorizationHeaderKey)

		// Check if the authorization is a bearer token
		parts := strings.Split(authorization, " ")

		// Return an error if the authorization is missing or invalid
		if len(parts) < 2 || parts[0] != commongin.BearerPrefix {
			ctx.JSON(
				401, gin.H{"error": InvalidAuthorizationHeaderError.Error()},
			)
			ctx.Abort()
			return
		}

		// Get the token from the header
		tokenString := parts[1]

		// Validate the token and get the validated claims
		claims, err := m.validator.GetValidatedClaims(tokenString, interception)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		// Set the token string and token claims to the context
		commonginctx.SetCtxTokenString(ctx, tokenString)
		commonginctx.SetCtxTokenClaims(ctx, claims)

		// Continue
		ctx.Next()
	}
}
