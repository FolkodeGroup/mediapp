package main

import (
	"os"
	"testing"
)

// Extraemos la lógica de conexión a una función para poder testearla
func getDBURL() string {
	return os.Getenv("DATABASE_URL")
}

func TestGetDBURL_Empty(t *testing.T) {
	os.Unsetenv("DATABASE_URL")
	if url := getDBURL(); url != "" {
		t.Errorf("esperaba cadena vacía, obtuvo %v", url)
	}
}

func TestGetDBURL_Set(t *testing.T) {
	expected := "postgres://user:pass@localhost:5432/db"
	os.Setenv("DATABASE_URL", expected)
	if url := getDBURL(); url != expected {
		t.Errorf("esperaba %v, obtuvo %v", expected, url)
	}
}

