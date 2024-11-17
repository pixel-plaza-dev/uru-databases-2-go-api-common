package header

import (
	"github.com/gin-gonic/gin"
)

// SecurityHeaders adds security headers to the response
func SecurityHeaders() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("X-Frame-Options", "DENY")
		context.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		context.Header("X-XSS-Protection", "1; mode=block")
		context.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		context.Header("Referrer-Policy", "strict-origin")
		context.Header("X-Content-Type-Options", "nosniff")
		context.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		context.Next()
	}
}
