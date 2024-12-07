package rate_limiter

import (
	"errors"
	"github.com/gin-gonic/gin"
	commonredisratelimiter "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/database/redis/rate-limiter"
	"net/http"
)

// Middleware struct
type Middleware struct {
	rateLimiter commonredisratelimiter.RateLimiter
}

// NewMiddleware creates a new rate limiter middleware
func NewMiddleware(rateLimiter commonredisratelimiter.RateLimiter) (*Middleware, error) {
	// Check if the rate limiter is nil
	if rateLimiter == nil {
		return nil, commonredisratelimiter.NilRateLimiterError
	}

	return &Middleware{
		rateLimiter: rateLimiter,
	}, nil
}

// Limit limits the number of requests per IP address
func (m *Middleware) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the client IP address
		ip := c.ClientIP()

		// Limit the number of requests per IP address
		if err := m.rateLimiter.Limit(ip); err != nil {
			// Check if the rate limit is exceeded
			if errors.Is(err, commonredisratelimiter.TooManyRequestsError) {
				c.AbortWithStatus(http.StatusTooManyRequests)
				return
			}
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Next()
	}
}
