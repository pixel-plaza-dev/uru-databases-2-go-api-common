package context

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	commongin "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/gin"
)

// SetCtxTokenString sets the token string in the context
func SetCtxTokenString(ctx *gin.Context, token string) {
	ctx.Set(commongin.AuthorizationHeaderKey, &token)
}

// SetCtxTokenClaims sets the token claims in the context
func SetCtxTokenClaims(ctx *gin.Context, claims *jwt.MapClaims) {
	ctx.Set(commongin.CtxTokenClaimsKey, claims)
}

// GetCtxTokenString tries to get the token string from the context
func GetCtxTokenString(ctx *gin.Context) (string, error) {
	// Get the token from the context
	value := ctx.Value(commongin.AuthorizationHeaderKey)
	if value == nil {
		return "", MissingTokenInContextError
	}

	// Check the type of the value
	token, ok := value.(string)
	if !ok {
		return "", UnexpectedTokenTypeInContextError
	}

	return token, nil
}

// GetCtxTokenClaims tries to get the token claims from the context
func GetCtxTokenClaims(ctx *gin.Context) (*jwt.MapClaims, error) {
	// Get the token claims from the context
	value := ctx.Value(commongin.CtxTokenClaimsKey)
	if value == nil {
		return nil, MissingTokenClaimsInContextError
	}

	// Check the type of the value
	claims, ok := value.(*jwt.MapClaims)
	if !ok {
		return nil, UnexpectedTokenClaimsTypeInContextError
	}

	return claims, nil
}
