# Instrucciones para usar Vault en desarrollo con Docker Compose

## 1. Iniciar Vault (modo desarrollo, sin unseal manual)

```
docker compose up -d vault
```
- Vault estará disponible en: http://localhost:8200
- Token root: `root`

## 2. Acceder a la UI de Vault

- Abre tu navegador en: [http://localhost:8200](http://localhost:8200)
- Ingresa el token: `root`

## 3. Acceder a la CLI de Vault

- Desde el contenedor:
```
docker exec -it mediapp-vault sh
# Ya dentro del contenedor puedes usar:
vault status
```
- O desde tu máquina (si tienes Vault CLI instalado):
```
export VAULT_ADDR='http://localhost:8200'
export VAULT_TOKEN='root'
vault status
```

## 4. Crear el secreto JWT_SECRET_KEY en Vault

- Desde la UI:
  - Ve a "Secrets" > "Create secret"
  - Ruta: `secret/data/mediapp`
  - Clave: `JWT_SECRET_KEY`  Valor: `supersecreto123`

- Desde la CLI (dentro del contenedor):
```
docker exec -e VAULT_ADDR='http://127.0.0.1:8200' -e VAULT_TOKEN='root' mediapp-vault vault kv put secret/mediapp JWT_SECRET_KEY=supersecreto123
```

## 5. Verificar el secreto desde la CLI

```
docker exec -e VAULT_ADDR='http://127.0.0.1:8200' -e VAULT_TOKEN='root' mediapp-vault vault kv get secret/mediapp
```

## 6. Test automático en Go

Crea el archivo `backend/internal/config/vault_integration_test.go` con este contenido:

```go
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
```

Ejecuta el test así:
```
cd backend
go test -tags=integration ./internal/config
```

---

Con esto tienes todo el flujo documentado y automatizado para desarrollo local con Vault.
