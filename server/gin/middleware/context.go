package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin"
	commonauth "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin/middleware/auth"
)

// SetCtxToken sets the token in the context
func SetCtxToken(ctx *gin.Context, token *jwt.Token) {
	ctx.Set(commonauth.AuthorizationHeaderKey, &token)
}

// GetToken tries to get the token from the context
func GetToken(ctx context.Context) (*jwt.Token, error) {
	// Get the token from the context
	value := ctx.Value(commonauth.AuthorizationHeaderKey)
	if value == nil {
		return nil, commongin.NoTokenInContextError
	}

	// Check the type of the value
	t, ok := value.(*jwt.Token)
	if !ok {
		return nil, commongin.UnexpectedTokenTypeInContextError
	}

	return t, nil
}
