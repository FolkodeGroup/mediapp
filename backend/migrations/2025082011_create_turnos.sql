-- +goose Up
CREATE TABLE IF NOT EXISTS turnos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    paciente_id UUID NOT NULL REFERENCES pacientes(id) ON DELETE CASCADE,
    usuario_id UUID NOT NULL REFERENCES usuarios(id),
    fecha TIMESTAMP NOT NULL,
    motivo TEXT
);

-- +goose Down
DROP TABLE IF EXISTS turnos;
