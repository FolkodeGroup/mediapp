-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    nombre_rol VARCHAR(50) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE IF EXISTS roles;
