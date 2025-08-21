-- +goose Up
CREATE TABLE IF NOT EXISTS consultorios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    direccion TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS consultorios;
