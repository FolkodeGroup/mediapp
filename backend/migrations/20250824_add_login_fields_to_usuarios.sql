-- +goose Up
ALTER TABLE usuarios
ADD COLUMN ultimo_login TIMESTAMP,
ADD COLUMN intentos_fallidos INT NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE usuarios
DROP COLUMN IF EXISTS ultimo_login,
DROP COLUMN IF EXISTS intentos_fallidos;
