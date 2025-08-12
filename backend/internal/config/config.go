package config

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv carga las variables de entorno desde un archivo .env si existe
func LoadEnv() {
	_ = godotenv.Load()
}

// GetEnv obtiene una variable de entorno o un valor por defecto
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
