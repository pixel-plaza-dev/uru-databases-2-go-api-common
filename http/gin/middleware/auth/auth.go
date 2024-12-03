package auth

import (
	"github.com/gin-gonic/gin"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
	commonginctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/context"
	commongintypes "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	pbconfigrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/config/rest"
	pbtypesgrpc "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"strings"
)

type (
	// Authentication interface
	Authentication interface {
		Authenticate(
			baseUri string,
			mapper pbconfigrest.Mapper,
			grpcInterceptions *map[pbtypesgrpc.Method]pbtypesgrpc.Interception,
		) gin.HandlerFunc
	}

	// Middleware struct
	Middleware struct {
		validator commonjwtvalidator.Validator
		logger    Logger
	}
)

// NewMiddleware creates a new authentication middleware
func NewMiddleware(
	validator commonjwtvalidator.Validator,
	logger Logger,
) (*Middleware, error) {
	return &Middleware{
		validator: validator,
		logger:    logger,
	}, nil
}

// Authenticate return the middleware function that authenticates the request
func (m Middleware) Authenticate(
	mapper pbconfigrest.Mapper, grpcInterceptions *map[pbtypesgrpc.Method]pbtypesgrpc.Interception,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the full path and method
		fullPath := ctx.FullPath()
		method := ctx.Request.Method
		restMethod := pbtypesrest.GetMethod(method)

		// Check if the method is not supported
		if restMethod == pbtypesrest.INVALID {
			m.logger.MethodNotSupported(fullPath)
			ctx.JSON(
				500,
				commongintypes.NewInternalServerError(),
			)
			ctx.Abort()
			return
		}

		// Get the gRPC method
		grpcMethod, err := mapper.Traverse(fullPath, restMethod)
		if err != nil {
			m.logger.FailedToMapRESTEndpoint(err)
			ctx.JSON(
				500,
				commongintypes.NewInternalServerError(),
			)
			ctx.Abort()
			return
		}

		// Get the gRPC method interception
		interception, ok := (*grpcInterceptions)[*grpcMethod]
		if !ok {
			m.logger.MissingGRPCMethod(fullPath)
			ctx.JSON(500, commongintypes.NewInternalServerError())
			ctx.Abort()
			return
		}

		// Check if there is None interception
		if interception == pbtypesgrpc.None {
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
