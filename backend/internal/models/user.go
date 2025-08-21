package models

import (
	"time"

	"github.com/google/uuid"
)

// Usuario representa la tabla 'usuarios'
type Usuario struct {
	ID             uuid.UUID  `json:"id" db:"id"`
	Nombre         string     `json:"nombre" db:"nombre"`
	Email          string     `json:"email" db:"email"`
	ContrasenaHash string     `json:"-" db:"contrasena_hash"` // No exponer en JSON
	RolID          int        `json:"rol_id" db:"rol_id"`
	ConsultorioID  *uuid.UUID `json:"consultorio_id,omitempty" db:"consultorio_id"`
	Activo         bool       `json:"activo" db:"activo"`
	CreadoEn       time.Time  `json:"creado_en" db:"creado_en"`
}

// Paciente representa la tabla 'pacientes'
type Paciente struct {
	ID               uuid.UUID  `json:"id" db:"id"`
	Nombre           string     `json:"nombre" db:"nombre"`
	Apellido         string     `json:"apellido" db:"apellido"`
	FechaNacimiento  time.Time  `json:"fecha_nacimiento" db:"fecha_nacimiento"`
	NroCredencial    *string    `json:"nro_credencial,omitempty" db:"nro_credencial"`
	ObraSocial       *string    `json:"obra_social,omitempty" db:"obra_social"`
	CondicionIVA     *string    `json:"condicion_iva,omitempty" db:"condicion_iva"`
	Plan             *string    `json:"plan,omitempty" db:"plan"`
	CreadoPorUsuario *uuid.UUID `json:"creado_por_usuario,omitempty" db:"creado_por_usuario"`
	ConsultorioID    *uuid.UUID `json:"consultorio_id,omitempty" db:"consultorio_id"`
	CreadoEn         time.Time  `json:"creado_en" db:"creado_en"`
}

// Consultorio representa la tabla 'consultorios'
type Consultorio struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Direccion string    `json:"direccion" db:"direccion"`
}

// Rol representa la tabla 'roles'
type Rol struct {
	ID        int    `json:"id" db:"id"`
	NombreRol string `json:"nombre_rol" db:"nombre_rol"`
}

// Permiso representa la tabla 'permisos'
type Permiso struct {
	ID            int    `json:"id" db:"id"`
	NombrePermiso string `json:"nombre_permiso" db:"nombre_permiso"`
}

// RolPermiso representa la tabla 'rol_permiso'
type RolPermiso struct {
	RolID     int `json:"rol_id" db:"rol_id"`
	PermisoID int `json:"permiso_id" db:"permiso_id"`
}

// HistoriaClinica representa la tabla 'historias_clinicas'
type HistoriaClinica struct {
	ID            uuid.UUID `json:"id" db:"id"`
	PacienteID    uuid.UUID `json:"paciente_id" db:"paciente_id"`
	UsuarioID     uuid.UUID `json:"usuario_id" db:"usuario_id"`
	FechaConsulta time.Time `json:"fecha_consulta" db:"fecha_consulta"`
}

// HistoriaClinicaVersion representa la tabla 'historia_clinica_version'
type HistoriaClinicaVersion struct {
	ID                uuid.UUID `json:"id" db:"id"`
	HistoriaClinicaID uuid.UUID `json:"historia_clinica_id" db:"historia_clinica_id"`
	MotivoConsulta    *string   `json:"motivo_consulta,omitempty" db:"motivo_consulta"`
	Antecedentes      *string   `json:"antecedentes,omitempty" db:"antecedentes"`
	ExamenFisico      *string   `json:"examen_fisico,omitempty" db:"examen_fisico"`
	Diagnostico       *string   `json:"diagnostico,omitempty" db:"diagnostico"`
	Tratamiento       *string   `json:"tratamiento,omitempty" db:"tratamiento"`
	UsuarioID         uuid.UUID `json:"usuario_id" db:"usuario_id"`
	ModificadoEn      time.Time `json:"modificado_en" db:"modificado_en"`
}

// DatosPersonales representa la tabla 'datos_personales'
type DatosPersonales struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	PacienteID         uuid.UUID `json:"paciente_id" db:"paciente_id"`
	TelefonoEncriptado *[]byte   `json:"telefono_encriptado,omitempty" db:"telefono_encriptado"`
	DNIEncriptado      *[]byte   `json:"dni_encriptado,omitempty" db:"dni_encriptado"`
	Direccion          *string   `json:"direccion,omitempty" db:"direccion"`
}

// RecetaMedica representa la tabla 'recetas_medicas'
type RecetaMedica struct {
	ID           uuid.UUID `json:"id" db:"id"`
	PacienteID   uuid.UUID `json:"paciente_id" db:"paciente_id"`
	UsuarioID    uuid.UUID `json:"usuario_id" db:"usuario_id"`
	Contenido    *string   `json:"contenido,omitempty" db:"contenido"`
	FechaEmision time.Time `json:"fecha_emision" db:"fecha_emision"`
	FirmaDigital bool      `json:"firma_digital" db:"firma_digital"`
}

// Turno representa la tabla 'turnos'
type Turno struct {
	ID         uuid.UUID `json:"id" db:"id"`
	PacienteID uuid.UUID `json:"paciente_id" db:"paciente_id"`
	UsuarioID  uuid.UUID `json:"usuario_id" db:"usuario_id"`
	Fecha      time.Time `json:"fecha" db:"fecha"`
	Motivo     *string   `json:"motivo,omitempty" db:"motivo"`
}

// Auditoria representa la tabla 'auditorias'
type Auditoria struct {
	ID            int       `json:"id" db:"id"`
	UsuarioID     uuid.UUID `json:"usuario_id" db:"usuario_id"`
	Accion        string    `json:"accion" db:"accion"`
	TablaAfectada string    `json:"tabla_afectada" db:"tabla_afectada"`
	Fecha         time.Time `json:"fecha" db:"fecha"`
}
