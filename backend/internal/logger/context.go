package logger

import (
	"context"

	"go.uber.org/zap"
)

// FromContext devuelve un logger con requestID si está disponible
func FromContext(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value("logger").(*zap.Logger); ok {
		return l
	}
	return L() // fallback al logger global
}