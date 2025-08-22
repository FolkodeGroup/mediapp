-- +goose Up
CREATE TABLE IF NOT EXISTS datos_personales (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    paciente_id UUID NOT NULL REFERENCES pacientes(id) ON DELETE CASCADE,
    telefono_encriptado BYTEA,
    dni_encriptado BYTEA,
    direccion TEXT
);

-- +goose Down
DROP TABLE IF EXISTS datos_personales;
