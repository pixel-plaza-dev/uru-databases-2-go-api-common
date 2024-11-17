package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin"
	jwtmiddleware "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin/middleware/jwt"
)

// SetCtxTokenString sets the token string in the context
func SetCtxTokenString(ctx *gin.Context, token string) {
	ctx.Set(jwtmiddleware.AuthorizationHeaderKey, &token)
}

// GetCtxTokenString tries to get the token string from the context
func GetCtxTokenString(ctx context.Context) (string, error) {
	// Get the token from the context
	value := ctx.Value(jwtmiddleware.AuthorizationHeaderKey)
	if value == nil {
		return "", commongin.NoTokenInContextError
	}

	// Check the type of the value
	token, ok := value.(string)
	if !ok {
		return "", commongin.UnexpectedTokenTypeInContextError
	}

	return token, nil
}

// SetCtxToken sets the token in the context
func SetCtxToken(ctx *gin.Context, token *jwt.Token) {
	ctx.Set(jwtmiddleware.TokenKey, token)
}

// GetCtxToken tries to get the token from the context
func GetCtxToken(ctx context.Context) (*jwt.Token, error) {
	// Get the token from the context
	value := ctx.Value(jwtmiddleware.TokenKey)
	if value == nil {
		return nil, commongin.NoTokenInContextError
	}

	// Check the type of the value
	token, ok := value.(*jwt.Token)
	if !ok {
		return nil, commongin.UnexpectedTokenTypeInContextError
	}

	return token, nil
}
