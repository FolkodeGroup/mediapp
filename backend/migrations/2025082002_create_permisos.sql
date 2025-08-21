-- +goose Up
CREATE TABLE IF NOT EXISTS permisos (
    id SERIAL PRIMARY KEY,
    nombre_permiso VARCHAR(100) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE IF EXISTS permisos;
