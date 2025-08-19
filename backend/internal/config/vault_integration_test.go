//go:build integration
// +build integration

package config

import (
	"os"
	"testing"
)

func TestLoadSecretsFromVault(t *testing.T) {
	os.Setenv("VAULT_ADDR", "http://localhost:8200")
	os.Setenv("VAULT_TOKEN", "root")
	LoadSecretsFromVault()
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret != "supersecreto123" {
		t.Fatalf("Se esperaba JWT_SECRET_KEY=supersecreto123, pero se obtuvo: %s", secret)
	}
}
