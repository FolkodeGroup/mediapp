// Script simple para recuperar JWT_SECRET_KEY desde Vault
package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
)

func main() {
	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		vaultAddr = "http://localhost:8200"
	}
	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		vaultToken = "root"
	}

	config := api.DefaultConfig()
	config.Address = vaultAddr
	client, err := api.NewClient(config)
	if err != nil {
		panic(fmt.Sprintf("Error creando cliente Vault: %v", err))
	}
	client.SetToken(vaultToken)

	secret, err := client.Logical().Read("secret/data/mediapp")
	if err != nil {
		panic(fmt.Sprintf("Error leyendo secreto: %v", err))
	}

	if secret != nil && secret.Data["data"] != nil {
		data := secret.Data["data"].(map[string]interface{})
		jwt, ok := data["JWT_SECRET_KEY"].(string)
		if ok {
			fmt.Println("JWT_SECRET_KEY:", jwt)
			if jwt == "supersecreto123" {
				fmt.Println("Secreto verificado correctamente.")
			} else {
				fmt.Println("El secreto recuperado NO es el esperado.")
			}
		} else {
			fmt.Println("No se encontró JWT_SECRET_KEY en Vault.")
		}
	} else {
		fmt.Println("No se encontró el secreto en Vault.")
	}
}
