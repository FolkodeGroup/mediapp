-- +goose Up
CREATE TABLE IF NOT EXISTS historias_clinicas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    paciente_id UUID NOT NULL REFERENCES pacientes(id) ON DELETE CASCADE,
    usuario_id UUID NOT NULL REFERENCES usuarios(id),
    fecha_consulta TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS historias_clinicas;
