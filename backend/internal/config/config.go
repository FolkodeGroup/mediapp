package config

import (
	"log"
	"os"

	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
)

// LoadSecretsFromVault carga los secretos desde HashiCorp Vault
func LoadSecretsFromVault() {
	// Obtiene la dirección y el token de Vault de las variables de entorno
	// configuradas en docker-compose.yml
	config := api.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Error creando cliente de Vault: %v", err)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	// Lee el secreto de la ruta que definimos en Vault
	secret, err := client.Logical().Read("secret/data/mediapp")
	if err != nil {
		log.Fatalf("Error leyendo secreto de Vault: %v", err)
	}

	if secret != nil && secret.Data["data"] != nil {
		data := secret.Data["data"].(map[string]interface{})
		if jwtSecret, ok := data["JWT_SECRET_KEY"].(string); ok {
			os.Setenv("JWT_SECRET_KEY", jwtSecret)
			log.Println("Secreto JWT_SECRET_KEY cargado desde Vault.")
		}
	}
}

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
