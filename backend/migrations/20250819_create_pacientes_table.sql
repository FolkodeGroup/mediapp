-- +goose Up
-- Normalización de la tabla pacientes
CREATE TABLE IF NOT EXISTS pacientes (
    id SERIAL PRIMARY KEY,
    nro_documento VARCHAR(32) NOT NULL,
    tipo_documento VARCHAR(16) NOT NULL,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    fecha_nacimiento DATE NOT NULL,
    email VARCHAR(255) NOT NULL,
    telefono VARCHAR(32),
    direccion VARCHAR(255),
    estado_civil VARCHAR(32),
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_actualizacion TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS pacientes;
