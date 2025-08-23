package logger

import (
	"context"

	"go.uber.org/zap"
)

// FromContext devuelve un logger con requestID si está disponible
func FromContext(ctx context.Context) *zap.Logger {
	// Intentar obtener el logger del contexto (si fue setteado por el middleware)
	if l, ok := ctx.Value("logger").(*zap.Logger); ok {
		return l
	}
	
	// Fallback al logger global
	return L()
}