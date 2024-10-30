package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pixel-plaza-dev/uru-databases-2-api-common/gin/middleware"
	autherror "github.com/pixel-plaza-dev/uru-databases-2-api-common/gin/middleware/auth/error"
	commonjwterror "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/jwt/error"
	commonjwtvalidator "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/jwt/validator"
	"strings"
)

func Auth(validator *commonjwtvalidator.Validator, validateClaims func(*jwt.Token) (*jwt.Token, error)) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get the authorization from the header
		authorization := context.GetHeader("Authorization")

		// Check if the authorization is a bearer token
		parts := strings.Split(authorization, " ")
		if len(parts) < 2 || parts[0] != "Bearer" {
			// Return an error if the authorization is missing or invalid
			err := autherror.RequestInvalidAuthorizationHeaderError{}
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		// Get the token from the header
		tokenString := parts[1]

		// Validate the token
		token, err := validator.GetToken(tokenString, validateClaims)
		if err != nil {
			context.JSON(401, gin.H{"error": commonjwterror.InvalidTokenError{}.Error()})
			context.Abort()
			return
		}

		// Set the token in the context
		middleware.SetToken(context, token)

		context.Next()
	}
}
