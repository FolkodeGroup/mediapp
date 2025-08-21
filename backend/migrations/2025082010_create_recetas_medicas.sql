-- +goose Up
CREATE TABLE IF NOT EXISTS recetas_medicas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    paciente_id UUID NOT NULL REFERENCES pacientes(id) ON DELETE CASCADE,
    usuario_id UUID NOT NULL REFERENCES usuarios(id),
    contenido TEXT,
    fecha_emision TIMESTAMP NOT NULL DEFAULT NOW(),
    firma_digital BOOLEAN DEFAULT FALSE
);

-- +goose Down
DROP TABLE IF EXISTS recetas_medicas;
