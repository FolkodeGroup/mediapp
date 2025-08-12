-- +goose Up
-- Seed inicial: médico de prueba
INSERT INTO users (username, email, password)
VALUES ('medico_prueba', 'medico@prueba.com', '$2b$10$QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm123456');
-- Contraseña: Medico123 (bcrypt hash de ejemplo)

-- +goose Down
DELETE FROM users WHERE email = 'medico@prueba.com';
