-- +goose Up
CREATE TABLE IF NOT EXISTS rol_permiso (
    rol_id INT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    permiso_id INT NOT NULL REFERENCES permisos(id) ON DELETE CASCADE,
    PRIMARY KEY (rol_id, permiso_id)
);

-- +goose Down
DROP TABLE IF EXISTS rol_permiso;
