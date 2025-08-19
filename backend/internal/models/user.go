// Suponiendo que usas PostgreSQL con pgx o database/sql

package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `json:"id"`
	PasswordHash  string    `json:"-"` // No exponer en JSON
	RolID         int       `json:"rol_id"`
	ConsultorioID uuid.UUID `json:"consultorio_id"`
	Activo        bool      `json:"activo"`
	CreadoEn      time.Time `json:"creado_en"`
}

// El campo PasswordHash tiene tag json:"-" para que no se exponga en respuestas HTTP.
