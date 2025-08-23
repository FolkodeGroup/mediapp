package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Definimos nuestras propias keys de contexto aquí
type contextKey string

const (
	requestIDKey contextKey = "request_id"
	loggerKey    contextKey = "logger"
)

// RequestIDMiddleware genera un ID único por request y lo añade al contexto y headers
func RequestIDMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Usar el header existente o generar uno nuevo
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Establecer en el header de respuesta
		c.Writer.Header().Set("X-Request-ID", requestID)
		
		// Crear logger con request_id
		requestLogger := logger.With(zap.String("request_id", requestID))
		
		// Crear contexto con request_id y logger
		ctx := context.WithValue(c.Request.Context(), requestIDKey, requestID)
		ctx = context.WithValue(ctx, loggerKey, requestLogger)
		
		// También ponerlo en el contexto de gin para fácil acceso
		c.Set("request_id", requestID)
		
		// Actualizar la request con el nuevo contexto
		c.Request = c.Request.WithContext(ctx)

		// Loggear el inicio de la request usando el logger con request_id
		requestLogger.Info("Request iniciada",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("ip", c.ClientIP()))

		c.Next()
	}
}