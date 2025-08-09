package main

import (
	"fmt"
	"mediapp/internal/logger"
	"go.uber.org/zap"
)

func main() {
	// Inicializamos el logger
	logger.Init()
	defer logger.Sync()

	// Ejemplo de logs
	logger.L().Info("Servidor iniciado",
		zap.String("version", "1.0.0"),
		zap.Int("puerto", 8080),
	)

	logger.L().Error("Error al conectar a la base de datos",
		zap.String("host", "localhost"),
		zap.Error(fakeError()),
	)
}

func fakeError() error {
	return fmt.Errorf("falló la conexión")
}
