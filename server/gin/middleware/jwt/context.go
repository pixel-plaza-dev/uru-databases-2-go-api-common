package jwt

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/server/gin"
)

// SetCtxTokenString sets the token string in the context
func SetCtxTokenString(ctx *gin.Context, token string) {
	ctx.Set(AuthorizationHeaderKey, &token)
}

// SetCtxTokenClaims sets the token claims in the context
func SetCtxTokenClaims(ctx *gin.Context, claims *jwt.MapClaims) {
	ctx.Set(CtxTokenClaimsKey, claims)
}

// GetCtxTokenString tries to get the token string from the context
func GetCtxTokenString(ctx context.Context) (string, error) {
	// Get the token from the context
	value := ctx.Value(AuthorizationHeaderKey)
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

// GetCtxTokenClaims tries to get the token claims from the context
func GetCtxTokenClaims(ctx context.Context) (*jwt.MapClaims, error) {
	// Get the token claims from the context
	value := ctx.Value(CtxTokenClaimsKey)
	if value == nil {
		return nil, commongin.NoTokenInContextError
	}

	// Check the type of the value
	claims, ok := value.(*jwt.MapClaims)
	if !ok {
		return nil, commongin.UnexpectedTokenTypeInContextError
	}

	return claims, nil
}
