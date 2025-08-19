package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type contextKey string

const RequestIDKey contextKey = "requestID"

// RequestIDMiddleware genera un ID único por request y lo añade al contexto y logger
func RequestIDMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		ctx := context.WithValue(c.Request.Context(), RequestIDKey, requestID)
		c.Request = c.Request.WithContext(ctx)

		// Añadir requestID al logger
		c.Request = c.Request.WithContext(context.WithValue(ctx, "logger", logger.With(zap.String("requestID", requestID))))

		c.Next()
	}
}
