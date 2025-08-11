// backend/internal/logger/logger.go
package logger

import (
	"go.uber.org/zap"
)

var log *zap.Logger

// Init inicializa el logger global
func Init() {
	var err error
	log, err = zap.NewProduction() // Logger en formato JSON por defecto
	if err != nil {
		panic("no se pudo inicializar el logger: " + err.Error())
	}
}

// L devuelve el logger global.
// Si no está inicializado, se crea uno temporal.
func L() *zap.Logger {
	if log == nil {
		tmp, _ := zap.NewProduction()
		return tmp
	}
	return log
}

// Sync vacía buffers antes de cerrar la app.
// Es recomendable llamarlo con defer en main().
func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}

/*
Ejemplo de uso:

package main

import (
	"backend/internal/logger"
)

func main() {
	logger.Init()
	defer logger.Sync()

	logger.L().Info("Servidor iniciado", zap.String("modo", "producción"))
}
*/
