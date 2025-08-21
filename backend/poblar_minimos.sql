-- Inserta un rol "admin" si no existe
INSERT INTO roles (id, nombre_rol)
VALUES (1, 'admin')
ON CONFLICT (id) DO NOTHING;

-- Inserta un consultorio con el UUID usado en el test_api.sh
INSERT INTO consultorios (id, direccion)
VALUES ('7623abc7-5197-4f52-a9da-68594dffcf77', 'Consultorio Central')
ON CONFLICT (id) DO NOTHING;
