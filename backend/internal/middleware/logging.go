package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/FolkodeGroup/mediapp/internal/logger"
	"go.uber.org/zap"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Log de entrada
		log := logger.FromContext(c.Request.Context())
		log.Info("Petición recibida",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("clientIP", c.ClientIP()),
			zap.String("userAgent", c.Request.UserAgent()),
		)

		c.Next()

		// Log de salida
		log.Info("Petición completada",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", time.Since(start)),
		)
	}
}