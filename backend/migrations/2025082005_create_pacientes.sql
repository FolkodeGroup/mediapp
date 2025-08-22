-- +goose Up

CREATE TABLE IF NOT EXISTS pacientes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    fecha_nacimiento DATE NOT NULL,
    nro_credencial VARCHAR(50),
    obra_social VARCHAR(100),
    condicion_iva VARCHAR(50),
    plan VARCHAR(100),
    creado_por_usuario UUID REFERENCES usuarios(id),
    consultorio_id UUID REFERENCES consultorios(id),
    creado_en TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS pacientes;
