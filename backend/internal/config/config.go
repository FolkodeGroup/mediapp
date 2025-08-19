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
	log.Printf("Conectando a Vault en: %s", config.Address)

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Error creando cliente de Vault: %v", err)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))
	log.Println("Token de Vault configurado correctamente.")

	// Lee el secreto de la ruta que definimos en Vault
	secret, err := client.Logical().Read("secret/data/mediapp")
	if err != nil {
		log.Fatalf("Error leyendo secreto de Vault: %v", err)
	}

	if secret == nil {
		log.Println("No se encontró ningún secreto en la ruta especificada.")
		return
	}

	if secret.Data["data"] == nil {
		log.Println("La clave 'data' no está presente en el secreto.")
		return
	}

	data := secret.Data["data"].(map[string]interface{})
	if jwtSecret, ok := data["JWT_SECRET_KEY"].(string); ok {
		os.Setenv("JWT_SECRET_KEY", jwtSecret)
		log.Println("Secreto JWT_SECRET_KEY cargado desde Vault.")
	} else {
		log.Println("JWT_SECRET_KEY no encontrado en los datos del secreto.")
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
