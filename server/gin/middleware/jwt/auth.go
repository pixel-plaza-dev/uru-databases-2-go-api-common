package jwt

import (
	"github.com/gin-gonic/gin"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/crypto/jwt/validator"
	"strings"
)

type (
	// Authentication interface
	Authentication interface {
		Authenticate() gin.HandlerFunc
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
func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the authorization from the header
		authorization := ctx.GetHeader(AuthorizationHeaderKey)

		// Check if the authorization is a bearer token
		parts := strings.Split(authorization, " ")

		// Return an error if the authorization is missing or invalid
		if len(parts) < 2 || parts[0] != BearerPrefix {
			ctx.JSON(
				401, gin.H{"error": InvalidAuthorizationHeaderError.Error()},
			)
			ctx.Abort()
			return
		}

		// Get the token from the header
		tokenString := parts[1]

		// Validate the token
		token, err := m.validator.GetToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		// Set the token in the ctx
		SetCtxTokenString(ctx, tokenString)
		SetCtxToken(ctx, token)

		// Continue
		ctx.Next()
	}
}
