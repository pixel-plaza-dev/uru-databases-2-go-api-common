package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pixel-plaza-dev/uru-databases-2-go-api-common/gin/middleware"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/jwt/validator"
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

// Authenticate return the middleware function
func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get the authorization from the header
		authorization := context.GetHeader("Authorization")

		// Check if the authorization is a bearer token
		parts := strings.Split(authorization, " ")

		// Return an error if the authorization is missing or invalid
		if len(parts) < 2 || parts[0] != "Bearer" {
			context.JSON(401, gin.H{"error": InvalidAuthorizationHeaderError})
			context.Abort()
			return
		}

		// Get the token from the header
		tokenString := parts[1]

		// Validate the token
		token, err := m.validator.GetToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err})
			context.Abort()
			return
		}

		// Set the token in the context
		middleware.SetCtxToken(context, token)

		context.Next()
	}
}
