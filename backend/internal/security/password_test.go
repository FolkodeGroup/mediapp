package security

import (
	"testing"
)

// TestHashPassword verifica que HashPassword genera un hash válido
func TestHashPassword(t *testing.T) {
	password := "mypassword123"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Se esperaba que HashPassword no devolviera error, pero obtuvo: %v", err)
	}

	if hashedPassword == "" {
		t.Error("Se esperaba que el hash no fuera vacío")
	}

	// Verificar que el hash generado sea válido comparándolo con la misma contraseña
	if !CheckPasswordHash(password, hashedPassword) {
		t.Error("Se esperaba que la contraseña original coincidiera con el hash generado")
	}
}

// TestCheckPasswordHash verifica que CheckPasswordHash funcione correctamente
func TestCheckPasswordHash(t *testing.T) {
	password := "mypassword123"
	wrongPassword := "wrongpass456"

	// Generar un hash para la contraseña correcta
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Error al generar hash: %v", err)
	}

	// Caso positivo: contraseña correcta
	if !CheckPasswordHash(password, hashedPassword) {
		t.Error("Se esperaba que la contraseña correcta pasara la verificación")
	}

	// Caso negativo: contraseña incorrecta
	if CheckPasswordHash(wrongPassword, hashedPassword) {
		t.Error("Se esperaba que la contraseña incorrecta NO pasara la verificación")
	}
}