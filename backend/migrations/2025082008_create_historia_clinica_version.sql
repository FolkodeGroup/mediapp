-- +goose Up
CREATE TABLE IF NOT EXISTS historia_clinica_version (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    historia_clinica_id UUID NOT NULL REFERENCES historias_clinicas(id) ON DELETE CASCADE,
    motivo_consulta TEXT,
    antecedentes TEXT,
    examen_fisico TEXT,
    diagnostico TEXT,
    tratamiento TEXT,
    usuario_id UUID NOT NULL REFERENCES usuarios(id),
    modificado_en TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS historia_clinica_version;
