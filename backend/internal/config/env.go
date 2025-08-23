// internal/config/env.go
package config

import (
	"os"
)

func LoadEnv() {
	// Esta función puede estar vacía si ya estás usando godotenv.Load()
	// o puedes poner aquí la lógica de carga de variables de entorno
	
	// Por ejemplo, podrías hacer validaciones básicas:
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8080")
	}
	
	if os.Getenv("ENV") == "" {
		os.Setenv("ENV", "development")
	}
}