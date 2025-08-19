//go:build ignore

// Ejecutable de prueba manual para Vault. No usar en producción.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/FolkodeGroup/mediapp/internal/config"
)

func main() {
	manual := flag.Bool("manual", false, "Ejecutar prueba manual de Vault")
	flag.Parse()
	if *manual {
		os.Setenv("VAULT_ADDR", "http://localhost:8200")
		os.Setenv("VAULT_TOKEN", "root")
		config.LoadSecretsFromVault()
		fmt.Println("JWT_SECRET_KEY:", os.Getenv("JWT_SECRET_KEY"))
	} else {
		fmt.Println("Usa la bandera -manual para ejecutar la prueba manual de Vault.")
	}
}
