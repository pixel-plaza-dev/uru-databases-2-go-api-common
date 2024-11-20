package auth

import (
	"github.com/gin-gonic/gin"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
	commonginctx "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin/context"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	"strings"
)

type (
	// Authentication interface
	Authentication interface {
		Authenticate(mustBeRefreshToken bool) gin.HandlerFunc
		AuthenticateAccessToken() gin.HandlerFunc
		AuthenticateRefreshToken() gin.HandlerFunc
	}

	// Middleware struct
	Middleware struct {
		validator commonjwtvalidator.Validator
	}
)

// NewMiddleware creates a new authentication middleware
func NewMiddleware(validator commonjwtvalidator.Validator) *Middleware {
	return &Middleware{
		validator: validator,
	}
}

// Authenticate return the middleware function that authenticates the request
func (m *Middleware) Authenticate(mustBeRefreshToken bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
		claims, err := m.validator.GetValidatedClaims(tokenString, mustBeRefreshToken)
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

// AuthenticateAccessToken return the middleware function that authenticates the request with an access token
func (m *Middleware) AuthenticateAccessToken() gin.HandlerFunc {
	return m.Authenticate(false)
}

// AuthenticateRefreshToken return the middleware function that authenticates the request with a refresh token
func (m *Middleware) AuthenticateRefreshToken() gin.HandlerFunc {
	return m.Authenticate(true)
}
