package middleware

import (
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware() gin.HandlerFunc {
	// Limitar a 100 solicitudes por minuto por IP
	lmt := tollbooth.NewLimiter(100, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Minute, // TTL de 1 minuto
	})
	
	lmt.SetMessage(`{"error": "Rate limit exceeded"}`)
	lmt.SetStatusCode(http.StatusTooManyRequests)
	lmt.SetMessageContentType("application/json")

	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
			return
		}
		c.Next()
	}
}