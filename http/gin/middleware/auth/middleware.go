package auth

import (
	"github.com/gin-gonic/gin"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
	commonginctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/context"
	commongintypes "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/types"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	commonlogger "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/utils/logger"
	pbtypesgrpc "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/grpc"
	pbtypesrest "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/types/rest"
	"strings"
)

// Middleware struct
type Middleware struct {
	validator commonjwtvalidator.Validator
	logger    *Logger
	mode      *commonflag.ModeFlag
}

// NewMiddleware creates a new authentication middleware
func NewMiddleware(
	validator commonjwtvalidator.Validator,
	logger *Logger,
	mode *commonflag.ModeFlag,
) (*Middleware, error) {
	// Check if either the validator, logger, or mode flag is nil
	if validator == nil {
		return nil, commonjwtvalidator.NilValidatorError
	}
	if logger == nil {
		return nil, commonlogger.NilLoggerError
	}
	if mode == nil {
		return nil, commonflag.NilModeFlagError
	}

	return &Middleware{
		validator: validator,
		logger:    logger,
		mode:      mode,
	}, nil
}

// Authenticate return the middleware function that authenticates the request
func (m Middleware) Authenticate(
	mapper *pbtypesrest.Mapper, grpcInterceptions *map[pbtypesgrpc.Method]pbtypesgrpc.Interception,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check if either the mapper or the gRPC interceptions is nil
		if mapper == nil || grpcInterceptions == nil {
			if mapper == nil {
				m.logger.MissingMapper()
			}
			if grpcInterceptions == nil {
				m.logger.MissingGRPCInterceptions()
			}
			ctx.JSON(500, commongintypes.NewErrorResponse(commongin.InternalServerError))
			ctx.Abort()
			return
		}

		// Get the request URI and method
		requestURI := ctx.Request.RequestURI

		// Get the gRPC method interception
		interception, ok := (*grpcInterceptions)[mapper.GRPCMethod]
		if !ok {
			if m.mode.IsDev() {
				m.logger.MissingGRPCMethod(requestURI)
			}
			ctx.JSON(500, commongintypes.NewErrorResponse(commongin.InternalServerError))
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
