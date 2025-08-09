package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

// Init inicializa el logger global
func Init() {
	var err error
	log, err = zap.NewProduction() // JSON por defecto
	if err != nil {
		panic("no se pudo inicializar el logger: " + err.Error())
	}
}

// L devuelve el logger global
func L() *zap.Logger {
	return log
}

// Sync para vaciar buffers al cerrar la app
func Sync() {
	_ = log.Sync()
}
