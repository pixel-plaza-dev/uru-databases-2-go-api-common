package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin"
	commonauth "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin/middleware/jwt"
)

// SetCtxToken sets the token in the context
func SetCtxToken(ctx *gin.Context, token string) {
	ctx.Set(commonauth.AuthorizationHeaderKey, &token)
}

// GetToken tries to get the token from the context
func GetToken(ctx context.Context) (string, error) {
	// Get the token from the context
	value := ctx.Value(commonauth.AuthorizationHeaderKey)
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
